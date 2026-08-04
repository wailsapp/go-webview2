package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// These are properties of the WHOLE generated binding, checked against the pinned IDL, rather than
// goldens for one interface. Each corresponds to a defect family that shipped for months, and the
// reason none was caught by review is the same in every case: the wrong output is valid Go that
// compiles, links and returns S_OK. A golden file records what one interface looked like on the day
// it was written; only a property says what must be true of all 300 of them.
//
// They also fail loudly if a template change is right for the interface a golden covers and wrong
// elsewhere -- which is the shape of every hand-patched fix in this package's history.

// pinnedIDL reads the version from latest_version.txt rather than repeating it, because every
// older IDL is still in the tree: a hand-copied constant would keep passing against 2903.40 long
// after update_version_mapping.go moved the pin, and the tests would be asserting about output
// nobody ships.
func pinnedIDL(t *testing.T) string {
	t.Helper()
	version, err := os.ReadFile(filepath.Join("..", "latest_version.txt"))
	require.NoError(t, err)
	return "WebView2." + strings.TrimSpace(string(version)) + ".idl"
}

func generateFromPinnedIDL(t *testing.T) map[string]string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("..", pinnedIDL(t)))
	require.NoError(t, err)
	files, err := ParseIDL(data)
	require.NoError(t, err)
	require.NotEmpty(t, files)
	out := make(map[string]string, len(files))
	for _, f := range files {
		out[f.FileName] = f.Content.String()
	}
	return out
}

// interfaceBases returns each interface's declared base, straight from the IDL. The IDL is the only
// authority on this: the version-suffixed names invite inferring the chain from the numbering, and
// ICoreWebView2EnvironmentOptions2 through 8 are exactly where that inference is wrong -- each of
// those derives from IUnknown, not from its predecessor.
func interfaceBases(t *testing.T) map[string]string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("..", pinnedIDL(t)))
	require.NoError(t, err)
	idl, err := Parser.ParseBytes("", data)
	require.NoError(t, err)
	require.NoError(t, idl.Process())

	bases := map[string]string{}
	for _, lib := range idl.Libraries {
		for _, d := range lib.Declarations {
			if d.Interface != nil {
				bases[d.Interface.Name] = d.Interface.BaseClass
			}
		}
	}
	require.NotEmpty(t, bases)
	return bases
}

// TestVtableEmbedsDeclaredBase is the regression test for the worst defect in this package's
// history, and it is a property rather than a slot-offset assertion on purpose.
//
// A COM vtable is flat: a derived interface's vtable begins with its ENTIRE base chain, then its own
// methods. Every derived vtable was generated as IUnknownVtbl plus its own methods, so each method
// sat too early by however many methods the chain above it declares, and a call landed on whichever
// unrelated function occupies that offset. ICoreWebView2_14's AddServerCertificateErrorDetected
// dispatched at slot 4 -- ICoreWebView2::get_Settings, which has one out-parameter and therefore
// wrote the Settings pointer over the caller's event-handler struct -- instead of slot 107. It
// returned S_OK, so registration "succeeded" and the event never fired. 88 interfaces were affected.
//
// Asserting embedding rather than computed offsets is deliberate: embedding the immediate base is
// sufficient by induction, and a test that recomputes the offsets would be asserting its own
// arithmetic against itself.
func TestVtableEmbedsDeclaredBase(t *testing.T) {
	files := generateFromPinnedIDL(t)
	bases := interfaceBases(t)

	checked := 0
	for name, base := range bases {
		content, ok := files[name+".go"]
		require.True(t, ok, "%s is declared but generated no file", name)
		want := base + "Vtbl"
		if base == "IUnknown" {
			want = "IUnknownVtbl"
		}
		// The embedded field is the first line of the vtable struct.
		decl := name + "Vtbl struct {\n\t"
		idx := strings.Index(content, decl)
		require.NotEqual(t, -1, idx, "%s has no vtable struct", name)
		rest := content[idx+len(decl):]
		got := rest[:strings.IndexByte(rest, '\n')]
		require.Equal(t, want, got,
			"%sVtbl must embed %s: the IDL declares %s : %s, and a vtable that omits an "+
				"intermediate interface's slots shifts every later method to the wrong offset",
			name, want, name, base)
		checked++
	}
	// Exact, not a floor. A floor with two units of headroom would let one interface silently stop
	// generating a file: the inner assertions are skipped for it and the count still clears the bar.
	require.Equal(t, len(bases), checked, "every declared interface must have been checked")
	require.NotZero(t, checked, "no interfaces were checked at all")
}

// TestQueryInterfaceAccessorReceiver checks the other half of the same mistake, with the rule that
// actually applies: QueryInterface asks an OBJECT for another of its interfaces, so the accessor
// belongs on an interface of the object that can answer -- which the declared chain's ROOT names,
// not the immediate base.
//
// Every accessor was emitted on ICoreWebView2. For the ICoreWebView2_N chain that is correct, since
// it is all one object. For every other chain it is useless: GetICoreWebView2Controller2 hung off
// ICoreWebView2, a different object, so it could only fail, while ICoreWebView2Controller had no
// accessor at all. 56 accessors were misrooted; the 26 on ICoreWebView2 stay put, so no existing
// caller breaks.
//
// Interfaces whose base is IUnknown keep the ICoreWebView2 receiver they ship with -- there is no
// sibling interface to reach them from, and re-rooting them onto IUnknown would move ~170 methods
// onto it and delete accessors that already ship. That is an API break, not a bug fix.
func TestQueryInterfaceAccessorReceiver(t *testing.T) {
	files := generateFromPinnedIDL(t)
	bases := interfaceBases(t)

	// root walks the declared chain to the ancestor whose own base is IUnknown. The seen-set is not
	// pedantry: production code grew the same guard, and without it a cyclic IDL would hang this
	// test instead of failing it.
	root := func(name string) string {
		out := ""
		seen := map[string]bool{name: true}
		for {
			base, ok := bases[name]
			if !ok || base == "IUnknown" {
				return out
			}
			require.False(t, seen[base], "inheritance cycle at %s", base)
			seen[base] = true
			out, name = base, base
		}
	}

	// An accessor is generated exactly when the IDL base is not IUnknown -- interfacevtbl.tmpl
	// guards on BaseClass, which generateVtbl blanks for IUnknown. Deriving the expected count from
	// that rule, rather than from a magic floor, is what makes a silently-missing accessor fail.
	wantAccessors := 0
	for _, base := range bases {
		if base != "IUnknown" {
			wantAccessors++
		}
	}

	checked, rerooted := 0, 0
	for name, base := range bases {
		content, ok := files[name+".go"]
		require.True(t, ok, "%s is declared but generated no file", name)
		if base == "IUnknown" {
			require.NotContains(t, content, fmt.Sprintf(") Get%s() *%s {", name, name),
				"%s derives from IUnknown, so no accessor should be generated for it", name)
			continue
		}
		receiver := root(name)
		if receiver == "" {
			receiver = "ICoreWebView2" // base is IUnknown; see above
		}
		require.Contains(t, content,
			fmt.Sprintf("func (i *%s) Get%s() *%s {", receiver, name, name),
			"Get%s must be a method on %s, the object a caller can QueryInterface from", name, receiver)
		if receiver != "ICoreWebView2" {
			rerooted++
		}
		checked++
	}
	require.Equal(t, wantAccessors, checked,
		"every interface with a non-IUnknown base must have an accessor")
	require.NotZero(t, rerooted,
		"expected the Controller/Environment/Profile/Settings/Frame chains to root on their own object")
}

var invokeSignature = regexp.MustCompile(`(?m)^func \w+Invoke\(this \*\w+(?:, ([^)]*))?\) uintptr \{`)

// TestCallbackParamsFitInAUintptr is the regression test for the family PR #36 patched in the
// output. A callback reached through syscall.NewCallback may not declare a parameter wider than a
// uintptr, and NewCallback checks this at CALLBACK CONSTRUCTION time -- which happens in a
// package-level var initialiser, so a violation panics during package init for any program that
// merely imports pkg/webview2, used or not:
//
//	panic: compileCallback: argument size is larger than uintptr
//
// A Go string is a 16-byte header, and three CompletedHandlers declared their LPCWSTR result as one.
// That makes this property the difference between the package being importable and not, which is
// also why it is worth a test rather than trusting three goldens.
func TestCallbackParamsFitInAUintptr(t *testing.T) {
	files := generateFromPinnedIDL(t)

	checked := 0
	for fileName, content := range files {
		for _, m := range invokeSignature.FindAllStringSubmatch(content, -1) {
			if m[1] == "" {
				continue
			}
			for _, param := range strings.Split(m[1], ",") {
				fields := strings.Fields(strings.TrimSpace(param))
				require.Len(t, fields, 2, "unexpected parameter %q in %s", param, fileName)
				typ := fields[1]
				require.NotEqual(t, "string", typ,
					"%s: callback parameter %q is a Go string (16 bytes); LPCWSTR arrives as a "+
						"pointer, so it must be declared *uint16 or syscall.NewCallback rejects "+
						"the whole vtable at init", fileName, fields[0])
				// NewCallback rejects floats outright, with its own panic
				// ("compileCallback: float arguments not supported"), so they fail at init exactly
				// as an oversized argument does.
				require.NotContains(t, []string{"float32", "float64"}, typ,
					"%s: callback parameter %q is a float; syscall.NewCallback refuses to build "+
						"the callback at all", fileName, fields[0])
				// Anything wider than a register fails the same way, not just a string. These two
				// are the aggregates maps.go classifies as too wide to pass in one -- keep the
				// lists together if either changes.
				for _, wide := range []string{"RECT", "COREWEBVIEW2_PHYSICAL_KEY_STATUS"} {
					require.NotEqual(t, wide, typ,
						"%s: callback parameter %q is a %s, which exceeds a register; "+
							"syscall.NewCallback rejects the whole vtable at init",
						fileName, fields[0], wide)
				}
				require.False(t, strings.HasPrefix(typ, "[]") || strings.Contains(typ, "interface{"),
					"%s: callback parameter %q has type %s, which is wider than a uintptr",
					fileName, fields[0], typ)
			}
			checked++
		}
	}
	require.Greater(t, checked, 50, "expected the pinned IDL to yield many handler Invoke functions")
}

// TestCallErrnoIsNeverReturnedAsError guards the fix that upstream applied to the generated output
// twice by hand and never to the template. ComProc.Call's third result is a syscall.Errno, which is
// non-nil on SUCCESS ("The operation completed successfully"), so binding it to err and returning it
// made every successful call look like a failure. HRESULT is the real status.
func TestCallErrnoIsNeverReturnedAsError(t *testing.T) {
	files := generateFromPinnedIDL(t)

	for fileName, content := range files {
		if fileName == "com.go" {
			// com.tmpl's hand-written IStream.Read does bind err, and correctly: it compares
			// against windows.ERROR_SUCCESS, which IS Errno(0), rather than against nil.
			continue
		}
		require.NotContains(t, content, ", _, err := i.Vtbl.",
			"%s binds ComProc.Call's Errno, which is non-nil on success", fileName)
	}
}

// TestByValueArgumentsAreNotPassedByAddress is table-driven over synthetic IDL rather than a
// property over the real one, because the wrong and right forms are distinguishable only if you know
// the parameter's type -- which the generated text alone does not tell you.
//
// The Windows x64 rule being pinned: an aggregate of exactly 1, 2, 4 or 8 bytes is passed IN A
// REGISTER as an integer of that width, anything else by address. So POINT (8 bytes) and RECT (16)
// take opposite forms, which is how a single &address default came to look plausible.
func TestByValueArgumentsAreNotPassedByAddress(t *testing.T) {
	cases := []struct {
		name  string
		param string
		want  string
	}{
		{"BOOL in-param", "[in] BOOL value", "boolToUintptr(value),"},
		{"UINT32 in-param", "[in] UINT32 value", "uintptr(value),"},
		{"HWND is a uintptr typedef", "[in] HWND value", "uintptr(value),"},
		{"8-byte token goes in a register", "[in] EventRegistrationToken value",
			"uintptr(*(*uint64)(unsafe.Pointer(&value))),"},
		{"8-byte POINT goes in a register", "[in] POINT value",
			"uintptr(*(*uint64)(unsafe.Pointer(&value))),"},
		{"4-byte COLOR goes in a register", "[in] COREWEBVIEW2_COLOR value",
			"uintptr(*(*uint32)(unsafe.Pointer(&value))),"},
		{"16-byte RECT is too wide, so by address", "[in] RECT value",
			"uintptr(unsafe.Pointer(&value)),"},
		{"in-param pointer is already the address", "[in] ICoreWebView2Settings* value",
			"uintptr(unsafe.Pointer(value)),"},
		{"out-param needs the address of our storage", "[out, retval] UINT32* value",
			"uintptr(unsafe.Pointer(&value)),"},
		{"out-param string needs the address of our pointer", "[out, retval] LPWSTR* value",
			"uintptr(unsafe.Pointer(&_value)),"},
		{"in-param string is already a *uint16", "[in] LPCWSTR value",
			"uintptr(unsafe.Pointer(_value)),"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			idl := fmt.Sprintf(`
[uuid(26d34152-879f-4065-bea2-3daa2cfadfb8), version(1.0)]
library WebView2 {
[uuid(A0D6DF20-3B92-416D-AA0C-437A9C727857), object, pointer_default(unique)]
interface ICoreWebView2Probe : IUnknown {
  HRESULT Probe(%s);
}
}`, c.param)

			files, err := ParseIDL([]byte(idl))
			require.NoError(t, err)

			var content string
			for _, f := range files {
				if f.FileName == "ICoreWebView2Probe.go" {
					content = f.Content.String()
				}
			}
			require.NotEmpty(t, content, "probe interface was not generated")
			require.Contains(t, content, c.want,
				"wrong marshalling for %q.\ngenerated:\n%s", c.param, content)
		})
	}
}

// TestCommittedOutputMatchesGenerator is the invariant this whole generator-first arrangement
// exists to establish, and until now it was the one thing left to a human running diff -r.
//
// The committed pkg/webview2 was NOT the output of the committed generator: fixes had been applied
// to the 306 output files and never to the templates, so each regeneration silently reverted them.
// Nothing detected that, because nothing compared the two. This does.
//
// It replaces a test that asserted the output was gofmt-clean. That one could not fail for the
// reason it claimed: its input had already been through gofmtAll, so it was checking that
// formatting formatted content is a no-op. Deliberately mangling a template still left it green.
// gofmtAll's own error return is the real guard against a template emitting invalid Go.
func TestCommittedOutputMatchesGenerator(t *testing.T) {
	const committed = "../../pkg/webview2"
	if _, err := os.Stat(committed); err != nil {
		t.Skipf("%s is absent; the scripts module stays independently testable", committed)
	}

	generated := generateFromPinnedIDL(t)

	entries, err := os.ReadDir(committed)
	require.NoError(t, err)
	onDisk := map[string]bool{}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") || strings.HasSuffix(e.Name(), "_test.go") {
			continue
		}
		onDisk[e.Name()] = true
	}

	for name, want := range generated {
		require.True(t, onDisk[name],
			"the generator produces %s but it is not committed; run the regeneration", name)
		got, err := os.ReadFile(filepath.Join(committed, name))
		require.NoError(t, err)
		require.Equal(t, want, string(got),
			"committed %s differs from generator output -- it has been hand-edited, or the "+
				"regeneration was not run after a template change. Never fix generated output "+
				"directly: the next regeneration reverts it.", name)
	}
	for name := range onDisk {
		require.Contains(t, generated, name,
			"%s is committed but the generator does not produce it", name)
	}
}

// TestVtableDeclaresEveryDeclaredMethod is the other half of the slot-offset argument, and the half
// that was previously only asserted in a commit message.
//
// Embedding the base vtable is sufficient for correct offsets ONLY IF each interface also contributes
// exactly its own methods, in the IDL's order. COM guarantees declaration order is vtable order, so
// the remaining risk is a method going missing: if the parser ever dropped one -- a form it does not
// recognise, a grammar change -- every slot after the gap would shift, and the failure would look
// exactly like the bug this series fixed while every embedding assertion still passed.
//
// So count them. One embedded base, and one ComProc per declared method, for all 252 interfaces.
func TestVtableDeclaresEveryDeclaredMethod(t *testing.T) {
	files := generateFromPinnedIDL(t)

	data, err := os.ReadFile(filepath.Join("..", pinnedIDL(t)))
	require.NoError(t, err)
	idl, err := Parser.ParseBytes("", data)
	require.NoError(t, err)
	require.NoError(t, idl.Process())

	checked := 0
	for _, lib := range idl.Libraries {
		for _, d := range lib.Declarations {
			if d.Interface == nil {
				continue
			}
			name := d.Interface.Name
			content, ok := files[name+".go"]
			require.True(t, ok, "%s is declared but generated no file", name)

			body := vtableBody(t, content, name)
			procs, embeds := 0, 0
			for _, line := range strings.Split(body, "\n") {
				switch line = strings.TrimSpace(line); {
				case line == "":
				case strings.HasSuffix(line, "ComProc"):
					procs++
				case strings.HasSuffix(line, "Vtbl"):
					embeds++
				}
			}
			require.Equal(t, len(d.Interface.Methods), procs,
				"%s declares %d methods in the IDL but its vtable has %d slots; every slot after "+
					"the gap dispatches to the wrong function",
				name, len(d.Interface.Methods), procs)
			require.Equal(t, 1, embeds, "%s must embed exactly one base vtable", name)
			checked++
		}
	}
	require.NotZero(t, checked)
}

// vtableBody returns the field block of <name>Vtbl.
func vtableBody(t *testing.T, content, name string) string {
	t.Helper()
	open := name + "Vtbl struct {"
	i := strings.Index(content, open)
	require.NotEqual(t, -1, i, "%s has no vtable struct", name)
	rest := content[i+len(open):]
	j := strings.Index(rest, "\n}")
	require.NotEqual(t, -1, j, "%s's vtable struct is unterminated", name)
	return rest[:j]
}
