//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Controller4Vtbl struct {
	_IUnknownVtbl
	GetAllowExternalDrop ComProc
	PutAllowExternalDrop ComProc
}

type ICoreWebView2Controller4 struct {
	vtbl *_ICoreWebView2Controller4Vtbl
}

func (i *ICoreWebView2Controller4) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Controller4) GetAllowExternalDrop() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetAllowExternalDrop.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2Controller4) PutAllowExternalDrop(value bool) error {
	var err error

	_, _, err = i.vtbl.PutAllowExternalDrop.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
