//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2Settings8Vtbl struct {
	ICoreWebView2SettingsVtbl
	GetUserAgent                        ComProc
	PutUserAgent                        ComProc
	GetAreBrowserAcceleratorKeysEnabled ComProc
	PutAreBrowserAcceleratorKeysEnabled ComProc
	GetIsPasswordAutosaveEnabled        ComProc
	PutIsPasswordAutosaveEnabled        ComProc
	GetIsGeneralAutofillEnabled         ComProc
	PutIsGeneralAutofillEnabled         ComProc
	GetIsPinchZoomEnabled               ComProc
	PutIsPinchZoomEnabled               ComProc
	GetIsSwipeNavigationEnabled         ComProc
	PutIsSwipeNavigationEnabled         ComProc
	GetHiddenPdfToolbarItems            ComProc
	PutHiddenPdfToolbarItems            ComProc
	GetIsReputationCheckingRequired     ComProc
	PutIsReputationCheckingRequired     ComProc
}

type ICoreWebView2Settings8 struct {
	Vtbl *ICoreWebView2Settings8Vtbl
}

func (i *ICoreWebView2Settings8) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2Settings) GetICoreWebView2Settings8() (*ICoreWebView2Settings8, error) {
	var result *ICoreWebView2Settings8

	iidICoreWebView2Settings8 := NewGUID("{9e6b0e8f-86ad-4e81-8147-a9b5edb68650}")
	hr, _, _ := i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2Settings8)),
		uintptr(unsafe.Pointer(&result)))

	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	if result == nil {
		return nil, syscall.Errno(windows.E_NOINTERFACE)
	}
	return result, nil
}

func (i *ICoreWebView2Settings8) GetIsReputationCheckingRequired() (bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, _ := i.Vtbl.GetIsReputationCheckingRequired.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return false, syscall.Errno(hr)
	}
	// Get result and cleanup
	value := _value != 0
	return value, nil
}

func (i *ICoreWebView2Settings8) PutIsReputationCheckingRequired(value bool) error {
	var enabled uintptr
	if value {
		enabled = 1
	}
	hr, _, _ := i.Vtbl.PutIsReputationCheckingRequired.Call(
		uintptr(unsafe.Pointer(i)),
		enabled,
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return nil
}
