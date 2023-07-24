//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Settings4Vtbl struct {
	_IUnknownVtbl
	GetIsPasswordAutosaveEnabled ComProc
	PutIsPasswordAutosaveEnabled ComProc
	GetIsGeneralAutofillEnabled  ComProc
	PutIsGeneralAutofillEnabled  ComProc
}

type ICoreWebView2Settings4 struct {
	vtbl *_ICoreWebView2Settings4Vtbl
}

func (i *ICoreWebView2Settings4) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Settings4) GetIsPasswordAutosaveEnabled() (bool, error) {
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

func (i *ICoreWebView2Settings4) PutIsPasswordAutosaveEnabled(value bool) error {
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

func (i *ICoreWebView2Settings4) GetIsGeneralAutofillEnabled() (bool, error) {
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

func (i *ICoreWebView2Settings4) PutIsGeneralAutofillEnabled(value bool) error {
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
