//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Environment3Vtbl struct {
	_IUnknownVtbl
	CreateCoreWebView2CompositionController ComProc
	CreateCoreWebView2PointerInfo           ComProc
}

type ICoreWebView2Environment3 struct {
	vtbl *_ICoreWebView2Environment3Vtbl
}

func (i *ICoreWebView2Environment3) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Environment3) CreateCoreWebView2CompositionController(parentWindow HWND, handler *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler) error {
	var err error

	_, _, err = i.vtbl.CreateCoreWebView2CompositionController.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&parentWindow)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Environment3) CreateCoreWebView2PointerInfo() (*ICoreWebView2PointerInfo, error) {
	var err error

	var pointerInfo *ICoreWebView2PointerInfo

	_, _, err = i.vtbl.CreateCoreWebView2PointerInfo.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pointerInfo)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return pointerInfo, nil
}
