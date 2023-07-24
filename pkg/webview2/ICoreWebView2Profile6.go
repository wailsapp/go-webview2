//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Profile6Vtbl struct {
	_IUnknownVtbl
	GetIsPasswordAutosaveEnabled ComProc
	PutIsPasswordAutosaveEnabled ComProc
	GetIsGeneralAutofillEnabled  ComProc
	PutIsGeneralAutofillEnabled  ComProc
}

type ICoreWebView2Profile6 struct {
	vtbl *_ICoreWebView2Profile6Vtbl
}

func (i *ICoreWebView2Profile6) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Profile6) GetIsPasswordAutosaveEnabled() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetIsPasswordAutosaveEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2Profile6) PutIsPasswordAutosaveEnabled(value bool) error {
	var err error

	_, _, err = i.vtbl.PutIsPasswordAutosaveEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Profile6) GetIsGeneralAutofillEnabled() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetIsGeneralAutofillEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2Profile6) PutIsGeneralAutofillEnabled(value bool) error {
	var err error

	_, _, err = i.vtbl.PutIsGeneralAutofillEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
