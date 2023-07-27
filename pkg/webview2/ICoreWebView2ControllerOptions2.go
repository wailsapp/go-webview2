//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2ControllerOptions2Vtbl struct {
	IUnknownVtbl
	GetScriptLocale ComProc
	PutScriptLocale ComProc
}

type ICoreWebView2ControllerOptions2 struct {
	Vtbl *ICoreWebView2ControllerOptions2Vtbl
}

func (i *ICoreWebView2ControllerOptions2) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2ControllerOptions2() *ICoreWebView2ControllerOptions2 {
	var result *ICoreWebView2ControllerOptions2

	iidICoreWebView2ControllerOptions2 := NewGUID("{06c991d8-9e7e-11ed-a8fc-0242ac120002}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2ControllerOptions2)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2ControllerOptions2) GetScriptLocale() (string, error) {
	// Create *uint16 to hold result
	var _locale *uint16

	hr, _, err := i.Vtbl.GetScriptLocale.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_locale)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return "", syscall.Errno(hr)
	}
	// Get result and cleanup
	locale := UTF16PtrToString(_locale)
	CoTaskMemFree(unsafe.Pointer(_locale))
	return locale, err
}

func (i *ICoreWebView2ControllerOptions2) PutScriptLocale(locale string) error {

	// Convert string 'locale' to *uint16
	_locale, err := UTF16PtrFromString(locale)
	if err != nil {
		return err
	}

	hr, _, err := i.Vtbl.PutScriptLocale.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_locale)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
