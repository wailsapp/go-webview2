package types

import (
	"fmt"
	"io"
	"strings"
)

type Param struct {
	Direction *Direction `parser:"@@?"`
	Type      string     `parser:"@Ident"`
	Const     string     `parser:"@('const')?"`
	Pointer   string     `parser:"@('*')*"`
	Name      string     `parser:"@Ident ','?"`

	// Processed
	GoType       string
	InputGoType  string
	OutputGoType string

	// This is used to generate setup code for the Go inputs
	setupTemplate   string
	cleanupTemplate string
	LocalName       string
	decl            *InterfaceMethod
	VtableCallInput string
}

func (p *Param) IsOutputParam() bool {
	if p.Direction == nil {
		return false
	}
	return p.Direction.Dir == "out"
}

func (p *Param) LocalVariableType() string {
	//if p.IsOutputParam() && p.isDoublePointer() {
	//	return p.GoType[1:]
	//}
	return p.GoType
}

func (p *Param) Process(decl *InterfaceMethod) {
	p.decl = decl
	p.GoType = IdlTypeToGoType(p.Type)
	if p.isDoublePointer() {
		p.GoType = "*" + p.GoType
	}
	p.OutputGoType = p.GoType
	if p.IsOutputParam() && strings.HasPrefix(p.OutputGoType, "**") {
		p.OutputGoType = p.GoType[1:]
	}
	p.InputGoType = p.GoType
}

func (p *Param) isPointer() bool {
	return p.Pointer != ""
}

func (p *Param) isSinglePointer() bool {
	return p.Pointer == "*"
}

func (p *Param) isDoublePointer() bool {
	return p.Pointer == "**"
}

func (p *Param) AsInputType() string {
	if p.isPointer() && p.GoType != "string" {
		return "*" + p.GoType
	}
	return p.GoType
}

// AsCallbackType is the type an INBOUND parameter must be declared as, and is deliberately not
// AsInputType. An outbound call is free to take a Go string, because the generated method
// converts it with UTF16PtrFromString before it reaches the vtable. A callback has no such
// step: syscall.NewCallback hands the Go function the raw machine words WebView2 passed, so the
// declared type IS the ABI. LPCWSTR arrives as a pointer, hence *uint16.
//
// Declaring it "string" does not merely mis-marshal; it fails to load at all. Every handler
// vtable is built in a package-level var initialiser, so NewComProc -> syscall.NewCallback runs
// during package init, and NewCallback rejects any argument wider than a uintptr. A string
// header is 16 bytes on amd64, so importing pkg/webview2 panicked before main() -- whether or
// not the program ever used these three handlers:
//
//	panic: compileCallback: argument size is larger than uintptr
//
// BOOL is left as Go's bool (ICoreWebView2PrintToPdf/TrySuspendCompletedHandler) although it is
// the same family. bool is 1 byte against a 4-byte BOOL, so the callee reads the low byte of the
// word -- correct for the 0/1 Win32 BOOL actually carries, and NewCallback accepts it because
// 1 <= sizeof(uintptr). Widening it to int32 would change those two Impl interfaces for every
// caller that implements them, to fix nothing observable.
func (p *Param) AsCallbackType() string {
	if p.GoType == "string" {
		// One star per level of IDL indirection, so LPCWSTR is *uint16 and LPCWSTR* is **uint16.
		// Every string callback parameter in the pinned IDL is a plain LPCWSTR, but returning
		// *uint16 regardless of depth would be its own silently-wrong answer if that changes.
		return strings.Repeat("*", len(p.Pointer)) + "*uint16"
	}
	return p.AsInputType()
}

func (p *Param) processSetup() error {
	p.processSetupInputs()
	p.processSetupOutputs()
	return p.processVtableCallInput()
}

func (p *Param) SetupCode(w io.Writer) {
	if p.setupTemplate == "" {
		return
	}
	data := struct {
		Param       *Param
		ErrorValues string
	}{
		Param:       p,
		ErrorValues: p.decl.ErrorValues(),
	}
	mustTemplate("Param Setup: "+p.setupTemplate, p.setupTemplate, &data, w)
}
func (p *Param) CleanupCode(w io.Writer) {
	if p.cleanupTemplate == "" {
		return
	}
	mustTemplate("Param Cleanup: "+p.cleanupTemplate, p.cleanupTemplate, p, w)
}

func (p *Param) IsInputParam() bool {
	return !p.IsOutputParam()
}

func (p *Param) processVtableCallInput() error {
	variableName := p.GetVariableName()

	// This block used to test p.Type -- the IDL type, so "BOOL", "INT32", "double" -- against
	// Go type names in lower case. It therefore never matched, and every by-value input
	// parameter fell through to the &address catch-all at the bottom of this function, which
	// makes the callee read a pointer as an integer. p.GoType is the mapped Go type and is
	// what the comparisons meant. One wrong field name, and no scalar in-parameter in the
	// whole binding was passed correctly.
	if !p.isPointer() {
		switch {
		case p.GoType == "bool":
			// A COM BOOL in-param is a 4-byte integer passed by value. uintptr(someBool) is
			// not a legal Go conversion, hence the helper in com.tmpl.
			p.VtableCallInput = "boolToUintptr(" + variableName + ")"
			return nil
		case strings.HasPrefix(p.GoType, "int"), strings.HasPrefix(p.GoType, "uint"):
			p.VtableCallInput = "uintptr(" + variableName + ")"
			return nil
		}
		// Handle typedefs and aggregates. See maps.go for the ABI rule and for why no single rule
		// covers them: EventRegistrationToken alone is 61 call sites, and every remove_ method in
		// the binding handed the callee the ADDRESS of the 8-byte token it was supposed to match, so
		// no event handler could ever be removed -- and remove_ still returned S_OK.
		if uintptrTypedef[p.Type] {
			p.VtableCallInput = "uintptr(" + variableName + ")"
			return nil
		}
		if expr, ok := byValueArgument[p.Type]; ok {
			p.VtableCallInput = fmt.Sprintf(expr, variableName)
			return nil
		}
		if _, ok := byRefAggregate[p.Type]; ok {
			// Wider than a register, so here the address genuinely is the argument.
			p.VtableCallInput = "uintptr(unsafe.Pointer(&" + variableName + "))"
			return nil
		}
		// A double goes in XMM0-XMM3 for the first four arguments, and Go's amd64 syscall
		// assembly copies each of those four argument slots into the matching XMM register
		// specifically so that floats work (runtime/sys_windows_amd64.s, "Load first 4 args into
		// correspondent registers ... in case any of the arguments are floating point values").
		// So the bit pattern IS the argument, and Float64bits produces it.
		//
		// An earlier version of this comment claimed no marshalling answer existed. That was
		// wrong, and it was wrong about the scope too: 12 methods take a double in-parameter --
		// PutZoomFactor, PutRasterizationScale, PutScaleFactor, PutExpires, PutPageWidth/Height,
		// the four PutMargin*, SetBoundsAndZoomFactor and ClearBrowsingDataInTimeRange -- and all
		// of them were handing the callee the ADDRESS of a Go float reinterpreted as a double.
		//
		// windows/arm64 remains unsolved: sys_windows_arm64.s loads only R0-R7 and never V0-V7,
		// carrying a TODO to do what amd64 does. Passing the bits in an integer register is no
		// worse there than passing a pointer was, so this is a strict improvement on both.
		if p.GoType == "float64" {
			p.VtableCallInput = "uintptr(math.Float64bits(" + variableName + "))"
			p.decl.decl.includes.AddUnique(`"math"`)
			return nil
		}
		// float32 is left unclassified on purpose: no IDL declares one, and it would need
		// Float32bits in the low half of the register rather than Float64bits.
		// Strings and enums are classified further down (the LPCWSTR case and IsEnum), so those two
		// are exempt rather than unclassified.
		if p.GoType != "string" && !p.IsEnum() {
			// Refuse to guess. Every defect in this family was the &address default landing on a
			// type nobody had classified, and each one returned S_OK while corrupting exactly one
			// argument. A new IDL type should cost one line in maps.go, not a silent wire bug.
			//
			// An error rather than log.Fatalf: log.Fatalf calls os.Exit, and this runs inside the
			// generator's own tests, so a reintroduced bug killed the test binary mid-run -- no
			// attributable failure, and every test after it silently never ran.
			return fmt.Errorf("by-value in-parameter %q of type %q (%s.%s) is in neither "+
				"byValueArgument nor byRefAggregate in maps.go. Classify it: an aggregate of "+
				"1, 2, 4 or 8 bytes is passed in a register as an integer of that width; "+
				"anything else is passed by address",
				variableName, p.Type, p.decl.decl.Name, p.decl.ProcessedName)
		}
	}
	switch p.Type {
	case "LPCWSTR", "LPWSTR":
		// Direction matters here and was not being distinguished. An OUT parameter is
		// declared LPWSTR* : the callee writes a string pointer into storage we own, so it
		// needs the ADDRESS of our local *uint16. Passing the local's (nil) value instead
		// gave the callee a null to write through, so every string getter returned empty --
		// silently, since the HRESULT was S_OK. An IN parameter is already the *uint16 that
		// UTF16PtrFromString produced and is passed as-is.
		if p.IsOutputParam() {
			p.VtableCallInput = "uintptr(unsafe.Pointer(&" + variableName + "))"
		} else {
			p.VtableCallInput = "uintptr(unsafe.Pointer(" + variableName + "))"
		}
		return nil
	}
	if p.Pointer == "**" {
		// Direction matters here for the same reason it does for "*" below, and this branch was
		// the one left without the check. An IN parameter is already the T** the caller built, so
		// taking its address hands the callee a T*** -- it then reads our local's own value as the
		// first element and calls through it. ICoreWebView2Environment14.CreateObjectCollection is
		// the live example, and it is a wild call rather than a wrong value.
		if p.IsOutputParam() {
			p.VtableCallInput = "uintptr(unsafe.Pointer(&" + variableName + "))"
		} else {
			p.VtableCallInput = "uintptr(unsafe.Pointer(" + variableName + "))"
		}
		return nil
	}
	if p.Pointer == "*" {
		if p.IsOutputParam() {
			p.VtableCallInput = "uintptr(unsafe.Pointer(&" + variableName + "))"
		} else {
			p.VtableCallInput = "uintptr(unsafe.Pointer(" + variableName + "))"
		}
		return nil
	}
	if p.IsEnum() {
		p.VtableCallInput = "uintptr(" + variableName + ")"
		return nil
	}
	p.VtableCallInput = "uintptr(unsafe.Pointer(&" + variableName + "))"
	return nil
}

func (p *Param) ClearLocalName() string {
	p.LocalName = ""
	return ""
}

func (p *Param) GetVariableName() string {
	result := p.LocalName
	if result == "" {
		result = p.Name
	}
	return result
}

func (p *Param) GetReturnVariableName() string {
	result := p.LocalName
	if result == "" {
		result = p.Name
	}
	return result
}

func (p *Param) IsEnum() bool {
	return p.decl.decl.decl.library.enums.Contains(p.Type)
}

func (p *Param) processSetupInputs() {
	if !p.IsInputParam() {
		return
	}
	switch p.GoType {
	case "string":
		// We need to convert to *uint16
		p.setupTemplate = "inputStringSetup.tmpl"
		p.LocalName = "_" + p.Name
	}
}

func (p *Param) processSetupOutputs() {
	if !p.IsOutputParam() {
		return
	}
	switch p.GoType {
	case "string":
		p.LocalName = "_" + p.Name
		p.setupTemplate = "outputStringSetup.tmpl"
		p.cleanupTemplate = "outputStringCleanup.tmpl"
	case "bool":
		p.LocalName = "_" + p.Name
		p.setupTemplate = "outputBoolSetup.tmpl"
		p.cleanupTemplate = "outputBoolCleanup.tmpl"
	default:
		p.setupTemplate = "outputDefaultSetup.tmpl"
	}
	if p.Pointer != "" {
		p.decl.decl.includes.AddUnique(`"unsafe"`)
	}
}

func (p *Param) defaultErrorValue() string {

	switch true {
	// A pointer's zero value is nil, and this has to be tested FIRST: every case below asks what
	// kind of thing p is, and for *T the answer that matters is only that it is a pointer.
	case p.OutputGoType[0] == '*':
		return "nil"
	// uintptrTypedef rather than the three of the eight that happen to appear as out-parameters
	// today (HANDLE, HWND, HCURSOR). The default branch below returns GoType{} as the zero value,
	// which for a uintptr typedef does not compile -- so an HICON out-parameter in some later IDL
	// would have broken the build with nothing to point at. Same knowledge, one place.
	//
	// The p.GoType guard matters because uintptrTypedef is keyed on the IDL type, which carries no
	// indirection: without it [out] HANDLE** returned the integer 0 for a *HANDLE, which does not
	// compile. That was a regression against the p.GoType == "HANDLE" test this replaced.
	case p.IsEnum(), strings.HasPrefix(p.GoType, "uint"), strings.HasPrefix(p.GoType, "int"),
		uintptrTypedef[p.Type] && p.GoType == p.Type:
		return "0"
	case strings.HasPrefix(p.GoType, "float"):
		return "0.0"
	case p.GoType == "bool":
		return "false"
	case p.GoType == "string":
		return `""`
	default:
		return p.GoType + "{}"
	}
}

type Direction struct {
	Dir    string `parser:"'[' @('out'|'in')"`
	Retval string `parser:"(',' @('retval'|'size_is' '(' Ident ')') )? ']'"`
}
