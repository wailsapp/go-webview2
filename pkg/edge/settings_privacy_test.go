//go:build windows

package edge

import (
	"errors"
	"github.com/wailsapp/go-webview2/pkg/webview2"
	"golang.org/x/sys/windows"
	"syscall"
	"testing"
	"unsafe"
)

func TestReputationCheckingDelegatesToSettings8(t *testing.T) {
	for _, tc := range []struct {
		name                     string
		value, missing           bool
		getHR, queryHR, setterHR uintptr
	}{
		{name: "enable", value: true},
		{name: "disable"},
		{name: "get settings failure", getHR: uintptr(windows.E_FAIL)},
		{name: "query failure", queryHR: uintptr(windows.E_NOINTERFACE)},
		{name: "missing interface", missing: true},
		{name: "setter failure", setterHR: uintptr(windows.E_FAIL)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var queries, writes, baseReleased, eightReleased int
			eight := &webview2.ICoreWebView2Settings8{Vtbl: &webview2.ICoreWebView2Settings8Vtbl{
				ICoreWebView2SettingsVtbl: webview2.ICoreWebView2SettingsVtbl{IUnknownVtbl: webview2.IUnknownVtbl{
					Release: webview2.NewComProc(func(uintptr) uintptr { eightReleased++; return 0 }),
				}},
				PutIsReputationCheckingRequired: webview2.NewComProc(func(this, value uintptr) uintptr {
					writes++
					privacyCheck(t, value == uintptr(boolToInt(tc.value)), "forwarded value", value, tc.value)
					return tc.setterHR
				}),
			}}
			settings := &ICoreWebViewSettings{vtbl: &_ICoreWebViewSettingsVtbl{_IUnknownVtbl: _IUnknownVtbl{
				Release: NewComProc(func(uintptr) uintptr { baseReleased++; return 0 }),
				QueryInterface: NewComProc(func(this, iid uintptr, out *unsafe.Pointer) uintptr {
					queries++
					if tc.queryHR == 0 && !tc.missing {
						*out = unsafe.Pointer(eight)
					}
					return tc.queryHR
				}),
			}}}
			chromium := &Chromium{webview: &ICoreWebView2{vtbl: &iCoreWebView2Vtbl{GetSettings: NewComProc(func(this uintptr, out *unsafe.Pointer) uintptr {
				if tc.getHR == 0 {
					*out = unsafe.Pointer(settings)
				}
				return tc.getHR
			})}}}
			want := tc.getHR
			if want == 0 {
				want = tc.queryHR
			}
			if want == 0 && tc.missing {
				want = uintptr(windows.E_NOINTERFACE)
			}
			if want == 0 {
				want = tc.setterHR
			}
			err := chromium.PutIsReputationCheckingRequired(tc.value)
			privacyCheck(t, (want == 0 && err == nil) || (want != 0 && errors.Is(err, syscall.Errno(want))), "forwarded error", want, err)
			wantBase, wantEight := 0, 0
			if tc.getHR == 0 {
				wantBase = 1
			}
			if tc.getHR == 0 && tc.queryHR == 0 && !tc.missing {
				wantEight = 1
			}
			privacyCheck(t, queries == wantBase && baseReleased == wantBase && writes == wantEight && eightReleased == wantEight, queries, baseReleased, writes, eightReleased)
		})
	}
}

func TestSettings4RefusesQueryInterfaceFailure(t *testing.T) {
	for _, hr := range []uintptr{0x80004002, 0} {
		settings := &ICoreWebViewSettings{vtbl: &_ICoreWebViewSettingsVtbl{_IUnknownVtbl: _IUnknownVtbl{
			QueryInterface: NewComProc(func(this, iid, result uintptr) uintptr { return hr }),
		}}}
		got, err := settings.GetICoreWebView2Settings4()
		privacyCheck(t, got == nil && err != nil && (hr == 0 || errors.Is(err, syscall.Errno(hr))), got, err)
	}
}

func TestFormStorageSettersPropagateCapabilityAndHRESULT(t *testing.T) {
	for _, set := range []func(*Chromium, bool) error{(*Chromium).PutIsGeneralAutofillEnabled, (*Chromium).PutIsPasswordAutosaveEnabled} {
		privacyCheck(t, errors.Is(set(&Chromium{}, false), UnsupportedCapabilityError), "expected HRESULT/capability refusal")
		var value int32
		var hr uintptr
		put := NewComProc(func(this, v uintptr) uintptr { value = int32(v); return hr })
		get := NewComProc(func(this uintptr, v *int32) uintptr { *v = value; return hr })
		release := NewComProc(func(uintptr) uintptr { return 1 })
		four := &ICoreWebView2Settings4{Vtbl: &ICoreWebView2Settings4Vtbl{_IUnknownVtbl: _IUnknownVtbl{Release: release}, PutIsGeneralAutofillEnabled: put, PutIsPasswordAutosaveEnabled: put, GetIsGeneralAutofillEnabled: get, GetIsPasswordAutosaveEnabled: get}}
		settings := &ICoreWebViewSettings{vtbl: &_ICoreWebViewSettingsVtbl{_IUnknownVtbl: _IUnknownVtbl{Release: release, QueryInterface: NewComProc(func(this uintptr, iid *GUID, out *unsafe.Pointer) uintptr {
			if hr == 0 {
				*out = unsafe.Pointer(four)
			}
			return hr
		})}}}
		view := &ICoreWebView2{vtbl: &iCoreWebView2Vtbl{GetSettings: NewComProc(func(this uintptr, out *unsafe.Pointer) uintptr {
			*out = unsafe.Pointer(settings)
			return 0
		})}}
		chromium := &Chromium{webview: view, webview2RuntimeVersion: "155.0.4283.62"}
		for _, want := range []bool{true, false} {
			privacyCheck(t, set(chromium, want) == nil, "NoError")
			for _, read := range []func() (bool, error){four.GetIsGeneralAutofillEnabled, four.GetIsPasswordAutosaveEnabled} {
				got, err := read()
				privacyCheck(t, err == nil && want == got, want, got, err)
			}
		}
		hr = 0x80004005
		privacyCheck(t, errors.Is(set(chromium, false), syscall.Errno(hr)), "expected HRESULT/capability refusal")
		_, err := four.GetIsPasswordAutosaveEnabled()
		privacyCheck(t, errors.Is(err, syscall.Errno(hr)) && errors.Is(four.PutIsGeneralAutofillEnabled(false), syscall.Errno(hr)), err)
		privacyCheck(t, errors.Is(chromium.PutIsReputationCheckingRequired(false), syscall.Errno(hr)), "expected HRESULT/capability refusal")
		var table [37]ComProc
		table[2], table[36], hr = release, put, 0
		raw := struct{ Vtbl *[37]ComProc }{&table}
		settings.vtbl.QueryInterface = NewComProc(func(this uintptr, iid *GUID, out *unsafe.Pointer) uintptr {
			*out = unsafe.Pointer(&raw)
			return 0
		})
		privacyCheck(t, chromium.PutIsReputationCheckingRequired(false) == nil, "NoError")
		privacyCheck(t, value == 0, "Zero")
		hr = 0x80004005
		privacyCheck(t, errors.Is(chromium.PutIsReputationCheckingRequired(false), syscall.Errno(hr)), "expected HRESULT/capability refusal")
	}
}

func privacyCheck(t *testing.T, ok bool, values ...any) {
	t.Helper()
	if !ok {
		t.Fatal(values...)
	}
}
