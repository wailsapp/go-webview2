//go:build windows

package webview2

// Every correctness claim about this package's argument marshalling has, until now, been an argument
// about the Windows x64 calling convention rather than an observation. That is how eight families of
// marshalling bug survived here: the wrong code compiles, links, runs and returns S_OK, so reading it
// is the only check there was, and reading it is what failed.
//
// These tests execute it instead. The trick is that a COM object is nothing but a pointer to a table
// of function pointers, so one can be built entirely out of Go: fill a generated Vtbl struct with
// NewComProc(someGoFunc) and hand its address to the generated wrapper. The wrapper marshals exactly
// as it would for a real WebView2, the fake callee receives whatever actually arrived, and the test
// asserts on it.
//
// The consequence worth noticing: this needs no WebView2 Runtime, no display, no Edge install, no
// network. It needs Windows and nothing else, so it runs on a stock windows-latest CI runner -- which
// is the difference between "we reasoned carefully" and "we know".
//
// Each test below names the family it pins and what the pre-fix code did instead.

import (
	"math"
	"testing"
	"unsafe"
)

const sOK = 0

// nativePhysicalKeyStatus is COREWEBVIEW2_PHYSICAL_KEY_STATUS as the C header declares it: a BOOL is
// a 4-byte int, so this is 24 bytes. The generated Go struct must match it exactly, because the callee
// writes this layout through the pointer the caller supplies.
type nativePhysicalKeyStatus struct {
	RepeatCount, ScanCode                                   uint32
	IsExtendedKey, IsMenuKeyDown, WasKeyDown, IsKeyReleased int32
}

// --- by-value BOOL -----------------------------------------------------------------------------

// A COM BOOL in-parameter is a 4-byte integer passed by value. The generator used to hand over the
// ADDRESS of a Go bool, so the callee read a pointer as an integer: nonzero, therefore always true,
// and a different value on every call.
func TestBoolInParamArrivesByValue(t *testing.T) {
	for _, want := range []bool{true, false} {
		var got uintptr = 0xDEAD
		vtbl := ICoreWebView2SettingsVtbl{}
		vtbl.PutIsScriptEnabled = NewComProc(func(this *ICoreWebView2Settings, v uintptr) uintptr {
			got = v
			return sOK
		})
		obj := &ICoreWebView2Settings{Vtbl: &vtbl}

		if err := obj.PutIsScriptEnabled(want); err != nil {
			t.Fatal(err)
		}
		wantN := uintptr(0)
		if want {
			wantN = 1
		}
		if got != wantN {
			t.Errorf("PutIsScriptEnabled(%v): callee saw %#x, want %#x "+
				"(a large value means it received an address)", want, got, wantN)
		}
	}
}

// --- by-value 8-byte aggregate ----------------------------------------------------------------

// EventRegistrationToken is struct{ int64 }, so it goes in a register as that int64. Every one of the
// 61 remove_* methods used to pass its address, which is why no event handler could be removed -- and
// remove_ still returned S_OK, so nothing looked wrong.
func TestEventRegistrationTokenArrivesByValue(t *testing.T) {
	var got uintptr
	vtbl := ICoreWebView2Vtbl{}
	vtbl.RemoveNavigationCompleted = NewComProc(func(this *ICoreWebView2, v uintptr) uintptr {
		got = v
		return sOK
	})
	obj := &ICoreWebView2{Vtbl: &vtbl}

	token := EventRegistrationToken{value: 0x0123456789ABCDEF}
	if err := obj.RemoveNavigationCompleted(token); err != nil {
		t.Fatal(err)
	}
	if got != uintptr(0x0123456789ABCDEF) {
		t.Errorf("RemoveNavigationCompleted: callee saw %#x, want %#x", got, uintptr(0x0123456789ABCDEF))
	}
}

// POINT is two int32s -- 8 bytes, so also register-passed, with X in the low half. RECT (16 bytes) is
// the opposite case and is passed by address; the two are the reason a single rule cannot cover
// aggregates.
func TestPointArrivesByValueWithFieldOrderIntact(t *testing.T) {
	var got uintptr
	vtbl := ICoreWebView2PointerInfoVtbl{}
	vtbl.PutPixelLocation = NewComProc(func(this *ICoreWebView2PointerInfo, v uintptr) uintptr {
		got = v
		return sOK
	})
	obj := &ICoreWebView2PointerInfo{Vtbl: &vtbl}

	if err := obj.PutPixelLocation(POINT{X: 0x11112222, Y: 0x33334444}); err != nil {
		t.Fatal(err)
	}
	if lo, hi := uint32(got), uint32(got>>32); lo != 0x11112222 || hi != 0x33334444 {
		t.Errorf("PutPixelLocation: callee saw X=%#x Y=%#x, want X=0x11112222 Y=0x33334444", lo, hi)
	}
}

// --- by-value double ---------------------------------------------------------------------------

// A double is passed in XMM0-XMM3 for the first four arguments, and Go's syscall assembly copies each
// of those argument slots into both the integer register and the matching XMM register -- which is
// what makes math.Float64bits work here. The callback observes the integer register, so it sees the
// same bits the callee's XMM register receives.
//
// Twelve methods take a double, and each used to pass the ADDRESS of a Go float64, so the callee read
// a heap address as an IEEE-754 double: a denormal, or roughly 1e-300.
func TestDoubleInParamArrivesAsItsBits(t *testing.T) {
	var got uintptr
	vtbl := ICoreWebView2ControllerVtbl{}
	vtbl.PutZoomFactor = NewComProc(func(this *ICoreWebView2Controller, v uintptr) uintptr {
		got = v
		return sOK
	})
	obj := &ICoreWebView2Controller{Vtbl: &vtbl}

	const zoom = 1.75
	if err := obj.PutZoomFactor(zoom); err != nil {
		t.Fatal(err)
	}
	if got != uintptr(math.Float64bits(zoom)) {
		t.Errorf("PutZoomFactor(%v): callee saw %#x, want %#x (%v as observed by the callee)",
			zoom, got, math.Float64bits(zoom), math.Float64frombits(uint64(got)))
	}
}

// --- 32-bit out-parameter ----------------------------------------------------------------------

// The callee writes four bytes through the pointer we supply. When the local was Go's 64-bit int, the
// high half stayed zero and the sign never extended, so a negative NTSTATUS came back as a large
// positive: GetExitCode returned 3221225477 for STATUS_ACCESS_VIOLATION.
func TestInt32OutParamKeepsItsSign(t *testing.T) {
	const wantExit = int32(-1073741819) // 0xC0000005, STATUS_ACCESS_VIOLATION

	vtbl := ICoreWebView2ProcessFailedEventArgs2Vtbl{}
	// The callback declares its out-parameter as *int32 rather than uintptr. A callback may take
	// pointer arguments directly, so nothing has to be converted back from an integer -- which keeps
	// this test out of the unsafe.Pointer rules it is meant to be checking.
	vtbl.GetExitCode = NewComProc(func(this *ICoreWebView2ProcessFailedEventArgs2, out *int32) uintptr {
		*out = wantExit
		return sOK
	})
	obj := &ICoreWebView2ProcessFailedEventArgs2{Vtbl: &vtbl}

	got, err := obj.GetExitCode()
	if err != nil {
		t.Fatal(err)
	}
	if got != wantExit {
		t.Errorf("GetExitCode: got %d, want %d", got, wantExit)
	}
}

// --- string out-parameter ----------------------------------------------------------------------

// An out-parameter string is declared LPWSTR*: the callee writes a string POINTER into storage we
// own, so it needs the address of our local *uint16. Passing the local's nil VALUE instead gave the
// callee a null to write through, so all 109 string getters returned "" -- with S_OK.
//
// The assertion that matters is therefore that the callee is given somewhere to write. This writes nil
// rather than a real string on purpose: a real one would have to come from CoTaskMemAlloc, because the
// generated cleanup frees it with CoTaskMemFree and handing that a Go pointer corrupts the COM heap --
// and reading a syscall's returned address back into an unsafe.Pointer is the one pattern `go vet`
// cannot verify, which would make this file fail the vet step it is meant to pass. CoTaskMemFree(nil)
// and UTF16PtrToString(nil) are both defined no-ops, so the path runs end to end with no allocation.
// The conversion itself is library code rather than generated code, and the in-parameter test below
// covers a real round-trip in the other direction.
func TestStringOutParamIsGivenSomewhereToWrite(t *testing.T) {
	received := false

	vtbl := ICoreWebView2Vtbl{}
	vtbl.GetSource = NewComProc(func(this *ICoreWebView2, out **uint16) uintptr {
		received = out != nil
		*out = nil
		return sOK
	})
	obj := &ICoreWebView2{Vtbl: &vtbl}

	got, err := obj.GetSource()
	if err != nil {
		t.Fatal(err)
	}
	if !received {
		t.Error("GetSource: callee received a NULL out-parameter, so it had nowhere to write -- " +
			"the local's value was passed instead of its address")
	}
	if got != "" {
		t.Errorf("GetSource: got %q from a nil write, want the empty string", got)
	}
}

// A string IN-parameter is already the *uint16 that UTF16PtrFromString produced and is passed as-is.
func TestStringInParamArrivesAsAUTF16Pointer(t *testing.T) {
	const want = "https://example.invalid/"

	var got string
	vtbl := ICoreWebView2Vtbl{}
	vtbl.Navigate = NewComProc(func(this *ICoreWebView2, uri *uint16) uintptr {
		got = UTF16PtrToString(uri)
		return sOK
	})
	obj := &ICoreWebView2{Vtbl: &vtbl}

	if err := obj.Navigate(want); err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Navigate: callee saw %q, want %q", got, want)
	}
}

// --- aggregate out-parameter, and the width of a struct field ---------------------------------

// COREWEBVIEW2_PHYSICAL_KEY_STATUS is 24 bytes: two UINT32 and four BOOL, where a BOOL is a 4-byte
// int. Generating those flags as Go's 1-byte bool made the struct 12 bytes, so the callee's 24-byte
// write ran 12 bytes past the end of the local, and the last three flags were read from padding and
// so were permanently false.
//
// The write below is deliberately a full 24 bytes at the offsets a real callee uses, so this test
// fails -- by corrupting memory or by returning false flags -- if the field widths regress.
func TestAggregateOutParamMatchesTheNativeLayout(t *testing.T) {
	if got, want := unsafe.Sizeof(COREWEBVIEW2_PHYSICAL_KEY_STATUS{}), uintptr(24); got != want {
		t.Fatalf("COREWEBVIEW2_PHYSICAL_KEY_STATUS is %d bytes, want %d: a BOOL field is 4 bytes, "+
			"and the callee writes the native layout regardless of what Go declares", got, want)
	}

	vtbl := ICoreWebView2AcceleratorKeyPressedEventArgsVtbl{}
	vtbl.GetPhysicalKeyStatus = NewComProc(func(this *ICoreWebView2AcceleratorKeyPressedEventArgs, native *nativePhysicalKeyStatus) uintptr {
		native.RepeatCount = 7
		native.ScanCode = 0x1E
		native.IsExtendedKey = 1
		native.IsMenuKeyDown = 0
		native.WasKeyDown = 1
		native.IsKeyReleased = 1
		return sOK
	})
	obj := &ICoreWebView2AcceleratorKeyPressedEventArgs{Vtbl: &vtbl}

	got, err := obj.GetPhysicalKeyStatus()
	if err != nil {
		t.Fatal(err)
	}
	switch {
	case got.RepeatCount != 7 || got.ScanCode != 0x1E:
		t.Errorf("GetPhysicalKeyStatus: RepeatCount=%d ScanCode=%#x, want 7 and 0x1e",
			got.RepeatCount, got.ScanCode)
	case got.IsExtendedKey == 0 || got.WasKeyDown == 0 || got.IsKeyReleased == 0:
		t.Errorf("GetPhysicalKeyStatus: flags read from the wrong offsets: %+v", got)
	case got.IsMenuKeyDown != 0:
		t.Errorf("GetPhysicalKeyStatus: IsMenuKeyDown should be 0, got %d", got.IsMenuKeyDown)
	}
}

// --- uintptr typedef ---------------------------------------------------------------------------

// HWND is `type HWND uintptr` -- an integer, so the value IS the argument. Passing its address gave
// WebView2 a garbage window handle.
func TestHandleTypedefArrivesByValue(t *testing.T) {
	const want = uintptr(0x00CAFE00)

	var got uintptr
	vtbl := ICoreWebView2EnvironmentVtbl{}
	vtbl.CreateCoreWebView2Controller = NewComProc(
		func(this *ICoreWebView2Environment, hwnd uintptr, handler uintptr) uintptr {
			got = hwnd
			return sOK
		})
	obj := &ICoreWebView2Environment{Vtbl: &vtbl}

	if err := obj.CreateCoreWebView2Controller(HWND(want), nil); err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("CreateCoreWebView2Controller: callee saw hwnd=%#x, want %#x", got, want)
	}
}

// --- the HRESULT contract ----------------------------------------------------------------------

// ComProc.Call's third result is a syscall.Errno that is non-nil on success, so returning it made
// every successful call look like a failure. The HRESULT is the status, and a failing one must come
// back as the error while the out-parameters stay at their zero values.
func TestHResultDecidesTheError(t *testing.T) {
	const eFail = 0x80004005 // E_FAIL

	vtbl := ICoreWebView2ProcessFailedEventArgs2Vtbl{}
	vtbl.GetExitCode = NewComProc(func(this *ICoreWebView2ProcessFailedEventArgs2, out *int32) uintptr {
		*out = 42 // written, but the call fails
		return eFail
	})
	obj := &ICoreWebView2ProcessFailedEventArgs2{Vtbl: &vtbl}

	got, err := obj.GetExitCode()
	if err == nil {
		t.Fatal("GetExitCode returned nil error for E_FAIL")
	}
	if got != 0 {
		t.Errorf("GetExitCode returned %d alongside an error; want the zero value", got)
	}

	// And the success path must produce a nil error, not Errno(0) boxed into a non-nil interface.
	vtbl.GetExitCode = NewComProc(func(this *ICoreWebView2ProcessFailedEventArgs2, out *int32) uintptr {
		*out = 42
		return sOK
	})
	got, err = obj.GetExitCode()
	if err != nil {
		t.Fatalf("GetExitCode returned a non-nil error on S_OK: %v", err)
	}
	if got != 42 {
		t.Errorf("GetExitCode: got %d, want 42", got)
	}
}
