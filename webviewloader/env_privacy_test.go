//go:build windows && !native_webview2loader

package webviewloader

import (
	"errors"
	"github.com/wailsapp/go-webview2/pkg/combridge"
	"golang.org/x/sys/windows"
	"syscall"
	"testing"
	"unsafe"
)

func TestEnvironmentOptionsHandOffCustomCrashReporting(t *testing.T) {
	// Independent SDK identity table: scripts/WebView2.1.0.2903.40.idl.
	sdk := struct{ options3, options4 string }{
		options3: "{4a5c436e-a9e3-4a2e-89c3-910d3513f5cc}",
		options4: "{ac52d13f-0d38-475a-9dca-876580d6793e}",
	}
	iid, err := windows.GUIDFromString(sdk.options3)
	privacyCheck(t, err == nil, err)
	fourth, err := windows.GUIDFromString(sdk.options4)
	privacyCheck(t, err == nil, err)
	for _, enabled := range []bool{false, true} {
		opts := &environmentOptions{}
		WithCustomCrashReportingEnabled(enabled)(opts)
		bridge := newEnvironmentOptionsCom(opts)
		unknown := combridge.IUnknownFromUintptr(bridge.Ref())
		var third *combridge.IUnknownImpl
		err := unknown.QueryInterface(&iid, &third)
		privacyCheck(t, err == nil, "SDK Options3 QueryInterface", err)
		native := (*struct {
			Vtbl *struct{ Query, Add, Release, Get, Put uintptr }
		})(unsafe.Pointer(third))
		var got int32
		hr, _, _ := syscall.SyscallN(native.Vtbl.Get, uintptr(unsafe.Pointer(third)), uintptr(unsafe.Pointer(&got)))
		third.Release()
		var unsupported *combridge.IUnknownImpl
		err = unknown.QueryInterface(&fourth, &unsupported)
		privacyCheck(t, errors.Is(err, syscall.Errno(windows.E_NOINTERFACE)) && unsupported == nil, "Options4 must not answer for Options3", err, unsupported)
		bridge.Close()
		privacyCheck(t, hr == 0 && enabled == (got != 0), hr, enabled, got)
	}
}

func privacyCheck(t *testing.T, ok bool, values ...any) {
	t.Helper()
	if !ok {
		t.Fatal(values...)
	}
}
