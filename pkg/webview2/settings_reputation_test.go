//go:build windows

package webview2

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"testing"
	"unsafe"
)

type completionProbe struct {
	IUnknownImpl
	complete func(uintptr, string) uintptr
}

func (p *completionProbe) AddScriptToExecuteOnDocumentCreatedCompleted(code uintptr, result string) uintptr {
	return p.complete(code, result)
}
func (p *completionProbe) CallDevToolsProtocolMethodCompleted(code uintptr, result string) uintptr {
	return p.complete(code, result)
}
func (p *completionProbe) ExecuteScriptCompleted(code uintptr, result string) uintptr {
	return p.complete(code, result)
}

func TestScriptCompletionCallbacksUseNativeStringArguments(t *testing.T) {
	value, err := windows.UTF16PtrFromString("local \u2603")
	privacyCheck(t, err == nil, err)
	calls := 0
	probe := &completionProbe{complete: func(code uintptr, result string) uintptr {
		privacyCheck(t, code == 0x80004005 && result == "local \u2603", code, result)
		calls++
		return 17
	}}
	for _, tc := range []struct {
		object unsafe.Pointer
		invoke ComProc
	}{
		{object: unsafe.Pointer(NewICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler(probe)), invoke: ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerFn.Invoke},
		{object: unsafe.Pointer(NewICoreWebView2CallDevToolsProtocolMethodCompletedHandler(probe)), invoke: ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerFn.Invoke},
		{object: unsafe.Pointer(NewICoreWebView2ExecuteScriptCompletedHandler(probe)), invoke: ICoreWebView2ExecuteScriptCompletedHandlerFn.Invoke},
	} {
		calls = 0
		hr, _, _ := tc.invoke.Call(uintptr(tc.object), 0x80004005, uintptr(unsafe.Pointer(value)))
		privacyCheck(t, hr == 17 && calls == 1, hr, calls)
	}
}

func TestReputationCheckingRoundTrip(t *testing.T) {
	// Independent ABI oracle: IUnknown (3), base Settings (18), Settings2-7 (14).
	const getReputation, putReputation = 35, 36
	var table [37]ComProc
	for slot := range table {
		table[slot] = NewComProc(func(this, value uintptr) uintptr { return 0x80004005 })
	}
	var value int32
	var hr uintptr
	table[getReputation] = NewComProc(func(this uintptr, out *int32) uintptr { *out = value; return hr })
	table[putReputation] = NewComProc(func(this, v uintptr) uintptr { value = int32(v); return hr })
	raw := struct{ table *[37]ComProc }{&table}
	settings := (*ICoreWebView2Settings8)(unsafe.Pointer(&raw))
	for _, want := range []bool{true, false} {
		privacyCheck(t, settings.PutIsReputationCheckingRequired(want) == nil, "NoError")
		got, err := settings.GetIsReputationCheckingRequired()
		privacyCheck(t, err == nil && want == got, want, got, err)
	}
	hr = 0x80004005
	_, err := settings.GetIsReputationCheckingRequired()
	privacyCheck(t, errors.Is(err, syscall.Errno(hr)) && errors.Is(settings.PutIsReputationCheckingRequired(false), syscall.Errno(hr)), err)
}

func TestReputationCheckingRefusesMissingInterface(t *testing.T) {
	for _, hr := range []uintptr{0, 0x80004002} {
		settings := &ICoreWebView2Settings{Vtbl: &ICoreWebView2SettingsVtbl{IUnknownVtbl: IUnknownVtbl{QueryInterface: NewComProc(func(this, iid, out uintptr) uintptr { return hr })}}}
		got, err := settings.GetICoreWebView2Settings8()
		privacyCheck(t, got == nil && err != nil && (hr == 0 || errors.Is(err, syscall.Errno(hr))), got, err)
	}
}

func privacyCheck(t *testing.T, ok bool, values ...any) {
	t.Helper()
	if !ok {
		t.Fatal(values...)
	}
}
