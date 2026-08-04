package types

import (
	"bytes"
	"fmt"
	"github.com/leaanthony/slicer"
	"io"
	"log"
	"strings"
	"text/template"
)

type InterfaceDeclaration struct {
	Header    *InterfaceHeader   `parser:"'[' @@ ']'"`
	Name      string             `parser:"'interface' @Ident"`
	BaseClass string             `parser:" ':' @Ident '{' "`
	Methods   []*InterfaceMethod `parser:"@@+ '}'"`

	// private
	decl         *Declaration
	InvokeMethod *InterfaceMethod
	includes     slicer.StringSlicer
}

func (d *InterfaceDeclaration) Process(decl *Declaration) error {
	d.decl = decl

	// Find Invoke method
	for _, method := range d.Methods {
		err := method.Process(d)
		if err != nil {
			return err
		}
		if string(method.Name) == "Invoke" {
			// No break: a method declared AFTER Invoke would never be processed, and would then be
			// emitted with an empty parameter list and no call arguments. Every shipped handler has
			// Invoke alone, so this only ever saved a few iterations -- and since the generator now
			// gofmts its output, such a file fails generation rather than being written out broken.
			d.InvokeMethod = method
		}
	}
	d.includes.AddUnique(`"unsafe"`)
	if len(d.Methods) == 1 && d.Methods[0] == d.InvokeMethod {
		return nil
	}
	d.includes.AddUnique(`"syscall"`)
	d.includes.AddUnique(`"golang.org/x/sys/windows"`)
	return nil
}

func (d *InterfaceDeclaration) Generate(packageName string, w io.Writer) error {
	err := d.generateVtbl(packageName, w)
	if err != nil {
		return err
	}

	err = d.generateInvoke(w)
	if err != nil {
		return err
	}

	err = d.generateInterfaceMethods(w)
	if err != nil {
		return err
	}

	return nil
}

func (d *InterfaceDeclaration) generateVtbl(packageName string, w io.Writer) error {
	rootInterface, err := d.RootInterface()
	if err != nil {
		return err
	}
	data := struct {
		PackageName     string
		Name            string
		Methods         []*InterfaceMethod
		HasInvokeMethod bool
		Includes        []string
		BaseClass       string
		RootInterface   string
		Header          *InterfaceHeader
	}{
		PackageName: packageName,
		BaseClass:   d.BaseClass,
		// The vtable embeds the IMMEDIATE base, because that is what the memory layout is; the
		// QueryInterface accessor hangs off the chain ROOT, because that is which object can
		// answer it. Two different questions, so two fields.
		RootInterface:   rootInterface,
		Header:          d.Header,
		Name:            d.Name,
		Methods:         d.Methods,
		HasInvokeMethod: d.HasInvokeMethod(),
		Includes:        d.includes.AsSlice(),
	}
	if d.BaseClass == "IUnknown" {
		data.BaseClass = ""
	}
	mustTemplate("Interface Vtbl", "interfacevtbl.tmpl", &data, w)
	return nil
}

func (d *InterfaceDeclaration) GetBaseClass() string {
	if d.BaseClass == "IUnknown" {
		return ""
	}
	return d.BaseClass
}

// RootInterface walks the declared inheritance chain to its first member -- the ancestor whose own
// base is IUnknown -- and returns "" if this interface IS that ancestor.
//
// This, not the immediate base, is where a QueryInterface accessor belongs, and the distinction is
// the whole of what makes those accessors either useful or useless. QueryInterface asks an OBJECT
// for another of its interfaces, so any interface on the same object can answer for any other. A
// chain like ICoreWebView2 -> _2 -> ... -> _27 is all one object, so ICoreWebView2 can hand out
// ICoreWebView2_14 directly and a caller need not walk thirteen accessors to reach it.
//
// What the chain root identifies is WHICH object. ICoreWebView2Controller2's root is
// ICoreWebView2Controller, a different object from the webview -- so the accessor emitted on
// ICoreWebView2 asked the wrong object and could only ever fail, while the controller, which can
// answer, had no accessor at all. Rooting on the chain start fixes exactly those and leaves the
// ICoreWebView2_N family where it already is, which is both correct and not an API break.
//
// An interface whose base is IUnknown returns "" here and gets no accessor at all -- see
// interfacevtbl.tmpl, whose guard is BaseClass rather than RootInterface.
func (d *InterfaceDeclaration) RootInterface() (string, error) {
	root := ""
	seen := map[string]bool{d.Name: true}
	for cur := d; cur != nil && cur.BaseClass != "IUnknown"; {
		root = cur.BaseClass
		next := cur.decl.library.interfaces[cur.BaseClass]
		if next == nil {
			// A base this library does not declare ends the chain, and is deliberately NOT an
			// error. com.tmpl hand-writes IUnknown, IStream and IDataObject precisely so that
			// interfaces can derive from types no IDL declares; the IDL also carries forward
			// declarations; and this generator's own test fixtures are single-interface fragments
			// whose base is absent by construction. Rejecting the case breaks all three.
			//
			// The cost is that a base which really is missing surfaces as `undefined: <Base>Vtbl`
			// when the CONSUMING package is built, rather than here where the cause is known.
			// Telling "hand-written elsewhere" from "absent" would need a registry of what com.tmpl
			// defines, which is a worse coupling than the deferred error.
			break
		}
		if seen[next.Name] {
			// A cyclic "A : B, B : A" would otherwise spin forever. The generator runs on whatever
			// IDL Microsoft publishes next, and a hang is a worse failure than a wrong answer
			// because there is nothing to read.
			return "", fmt.Errorf("inheritance cycle reached %s while resolving the chain root "+
				"of %s", next.Name, d.Name)
		}
		seen[next.Name] = true
		cur = next
	}
	return root, nil
}

func (d *InterfaceDeclaration) generateInvoke(w io.Writer) error {
	if d.InvokeMethod == nil {
		return nil
	}
	data := struct {
		Name         string
		InvokeMethod *InterfaceMethod
		Declaration  *InterfaceDeclaration
	}{
		Declaration:  d,
		Name:         d.Name,
		InvokeMethod: d.InvokeMethod,
	}
	mustTemplate("Interface Invoke", "interfaceInvoke.tmpl", &data, w)
	return nil
}

func (d *InterfaceDeclaration) HasInvokeMethod() bool {
	return d.InvokeMethod != nil
}

func mustTemplate(templateName string, filename string, data interface{}, w io.Writer) {
	templateData, err := templates.ReadFile("templates/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err := template.New(templateName).Parse(string(templateData))
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func (d *InterfaceDeclaration) generateInterfaceMethods(w io.Writer) error {
	if len(d.Methods) == 1 && d.Methods[0] == d.InvokeMethod {
		return nil
	}
	for _, method := range d.Methods {
		data := struct {
			Name   string
			Method *InterfaceMethod
		}{
			Name:   d.Name,
			Method: method,
		}
		mustTemplate("Interface Methods", "interfaceMethod.tmpl", &data, w)
	}
	return nil
}

type InterfaceMethod struct {
	Prop       *Prop               `parser:"('[' @('propget'|'propput') ']')?"`
	ReturnType string              `parser:"@Ident"`
	Name       InterfaceMethodName `parser:"@Ident '('"`
	Params     []*Param            `parser:" @@* ')' ';'"`

	// private
	GoMethodName string

	GoInputs        string
	InputParamNames string

	GoReturnTypes string

	ProcessedName    string
	inputParams      []*Param
	outputParams     []*Param
	OutputParamNames string
	GoOutputs        string
	decl             *InterfaceDeclaration
}

func (m *InterfaceMethod) Process(decl *InterfaceDeclaration) error {
	m.decl = decl
	// Generate Go Method name
	goMethodName := strings.TrimPrefix(decl.Name, "ICoreWebView2")
	goMethodName = strings.TrimSuffix(goMethodName, "Handler")
	goMethodName = strings.TrimSuffix(goMethodName, "Event")
	m.GoMethodName = goMethodName

	m.ProcessedName = string(m.Name)
	if m.Prop != nil {
		m.ProcessedName = string(*m.Prop) + m.ProcessedName
	}
	return m.processParams()
}

func (m *InterfaceMethod) processParams() error {
	for _, param := range m.Params {
		param.Process(m)
		if param.IsOutputParam() {
			m.outputParams = append(m.outputParams, param)
		} else {
			m.inputParams = append(m.inputParams, param)
		}
	}

	if err := m.processInputParams(); err != nil {
		return err
	}
	return m.processOutputParams()
}

func (m *InterfaceMethod) processInputParams() error {
	var inputs slicer.StringSlicer
	var inputParamNames slicer.StringSlicer
	for _, param := range m.inputParams {
		inputs.Add(param.Name + " " + param.AsInputType())
		inputParamNames.Add(param.Name)
		if err := param.processSetup(); err != nil {
			return err
		}
	}
	m.GoInputs = inputs.Join(", ")
	m.InputParamNames = inputParamNames.Join(", ")
	return nil
}

func (m *InterfaceMethod) processOutputParams() error {
	var outputs slicer.StringSlicer
	var outputParamNames slicer.StringSlicer
	var outputParamTypes slicer.StringSlicer
	for _, param := range m.outputParams {
		outputs.Add(param.Name + " " + param.GoType)
		outputParamNames.Add(param.Name)
		outputParamTypes.Add(param.GoType)
		if err := param.processSetup(); err != nil {
			return err
		}
	}
	// Add the mandatory error
	outputs.Add("err error")
	outputParamNames.Add("err")
	outputParamTypes.Add("error")

	m.GoOutputs = outputs.Join(", ")
	m.OutputParamNames = outputParamNames.Join(", ")
	m.GoReturnTypes = outputParamTypes.Join(", ")
	if outputParamTypes.Length() > 1 {
		m.GoReturnTypes = "(" + m.GoReturnTypes + ")"
	}
	return nil
}

// CallbackInputs is GoInputs for a callback: the same parameters, declared in the shape the
// Windows callback ABI actually delivers them in. Only interfaceInvoke.tmpl uses it. See
// Param.AsCallbackType.
func (m *InterfaceMethod) CallbackInputs() string {
	var inputs slicer.StringSlicer
	for _, param := range m.inputParams {
		inputs.Add(param.Name + " " + param.AsCallbackType())
	}
	return inputs.Join(", ")
}

func (m *InterfaceMethod) SetupCode() string {
	var buffer bytes.Buffer
	for _, param := range m.Params {
		param.SetupCode(&buffer)
	}
	return buffer.String()
}

func (m *InterfaceMethod) CleanupCode() string {
	var buffer bytes.Buffer
	for _, param := range m.Params {
		param.CleanupCode(&buffer)
	}
	return buffer.String()
}

func (m *InterfaceMethod) VtableCallInputs() string {
	var buffer bytes.Buffer
	for _, input := range m.Params {
		buffer.WriteString("\t\t" + input.VtableCallInput + ",\n")
	}
	return buffer.String()
}

func (m *InterfaceMethod) ReturnsHRESULT() bool {
	return m.ReturnType == "HRESULT"
}

// ErrorValues is the early return taken when converting a string input fails. It keeps
// "err" -- unlike the paths below -- because there the error is UTF16PtrFromString's, which
// is a real Go error, and it is in scope (inputStringSetup.tmpl binds it).
func (m *InterfaceMethod) ErrorValues() string {
	var errorValues slicer.StringSlicer
	for _, outputParam := range m.outputParams {
		errorValues.Add(outputParam.defaultErrorValue())
	}
	errorValues.Add("err")
	return errorValues.Join(", ")
}
func (m *InterfaceMethod) ErrorValuesHRESULT() string {
	var errorValues slicer.StringSlicer
	for _, outputParam := range m.outputParams {
		errorValues.Add(outputParam.defaultErrorValue())
	}
	if m.ReturnsHRESULT() {
		errorValues.Add("syscall.Errno(hr)")
	} else {
		// Not "err": the Call's Errno is not bound any more (see interfaceMethod.tmpl), and
		// a method with no HRESULT has no status to report.
		errorValues.Add("nil")
	}
	return errorValues.Join(", ")
}

func (m *InterfaceMethod) GetHResultVariable() string {
	if m.ReturnsHRESULT() {
		return "hr"
	}
	return "_"
}

// SuccessValues is reached only after the HRESULT check passed, so the error is nil by
// construction. It used to be "err" -- the Call's Errno -- which is non-nil on success.
func (m *InterfaceMethod) SuccessValues() string {
	var successValues slicer.StringSlicer
	for _, outputParam := range m.outputParams {
		successValues.Add(outputParam.GetReturnVariableName())
	}
	successValues.Add("nil")
	return successValues.Join(", ")
}

type InterfaceHeader struct {
	UUID *UUID `parser:"'uuid' '(' @UUID ')' ',' 'object' ',' 'pointer_default' '(' 'unique' ')'"`
}

func (h *InterfaceHeader) AsString() string {
	uuid := *h.UUID
	return string(`"{` + uuid + `}"`)
}

type InterfaceMethodName string

func (m *InterfaceMethodName) Capture(values []string) error {
	if len(values) == 0 {
		return nil
	}
	result := values[0]
	if strings.HasPrefix(values[0], "add_") {
		result = "Add" + result[4:]
	}
	if strings.HasPrefix(values[0], "remove_") {
		result = "Remove" + result[7:]
	}
	*m = InterfaceMethodName(result)
	return nil
}

type Prop string

func (p *Prop) Capture(values []string) error {
	if len(values) == 0 {
		return nil
	}
	result := strings.Title(values[0][4:])
	*p = Prop(result)
	return nil
}
