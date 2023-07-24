//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2CompositionController3Vtbl struct {
	_IUnknownVtbl
	DragEnter ComProc
	DragLeave ComProc
	DragOver  ComProc
	Drop      ComProc
}

type ICoreWebView2CompositionController3 struct {
	vtbl *_ICoreWebView2CompositionController3Vtbl
}

func (i *ICoreWebView2CompositionController3) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2CompositionController3) DragEnter(dataObject *IDataObject, keyState DWORD, point POINT) (*DWORD, error) {
	var err error

	var effect *DWORD

	_, _, err = i.vtbl.DragEnter.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(dataObject)),
		uintptr(unsafe.Pointer(&keyState)),
		uintptr(unsafe.Pointer(&point)),
		uintptr(unsafe.Pointer(&effect)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return effect, nil
}

func (i *ICoreWebView2CompositionController3) DragLeave() error {
	var err error

	_, _, err = i.vtbl.DragLeave.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2CompositionController3) DragOver(keyState DWORD, point POINT) (*DWORD, error) {
	var err error

	var effect *DWORD

	_, _, err = i.vtbl.DragOver.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&keyState)),
		uintptr(unsafe.Pointer(&point)),
		uintptr(unsafe.Pointer(&effect)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return effect, nil
}

func (i *ICoreWebView2CompositionController3) Drop(dataObject *IDataObject, keyState DWORD, point POINT) (*DWORD, error) {
	var err error

	var effect *DWORD

	_, _, err = i.vtbl.Drop.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(dataObject)),
		uintptr(unsafe.Pointer(&keyState)),
		uintptr(unsafe.Pointer(&point)),
		uintptr(unsafe.Pointer(&effect)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return effect, nil
}
