//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2ControllerOptions2Vtbl struct {
	_IUnknownVtbl
	GetScriptLocale ComProc
	PutScriptLocale ComProc
}

type ICoreWebView2ControllerOptions2 struct {
	vtbl *_ICoreWebView2ControllerOptions2Vtbl
}

func (i *ICoreWebView2ControllerOptions2) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2ControllerOptions2) GetScriptLocale() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _locale *uint16

	_, _, err = i.vtbl.GetScriptLocale.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_locale)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	locale := windows.UTF16PtrToString(_locale)
	windows.CoTaskMemFree(unsafe.Pointer(_locale))
	return locale, nil
}

func (i *ICoreWebView2ControllerOptions2) PutScriptLocale(locale string) error {
	var err error

	// Convert string 'locale' to *uint16
	_locale, err := windows.UTF16PtrFromString(locale)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PutScriptLocale.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_locale)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
