//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Settings5Vtbl struct {
	_IUnknownVtbl
	GetIsPinchZoomEnabled ComProc
	PutIsPinchZoomEnabled ComProc
}

type ICoreWebView2Settings5 struct {
	vtbl *_ICoreWebView2Settings5Vtbl
}

func (i *ICoreWebView2Settings5) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Settings5) GetIsPinchZoomEnabled() (bool, error) {
	var err error

	var enabled bool

	_, _, err = i.vtbl.GetIsPinchZoomEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return enabled, nil
}

func (i *ICoreWebView2Settings5) PutIsPinchZoomEnabled(enabled bool) error {
	var err error

	_, _, err = i.vtbl.PutIsPinchZoomEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
