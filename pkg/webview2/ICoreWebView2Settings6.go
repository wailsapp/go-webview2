//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Settings6Vtbl struct {
	_IUnknownVtbl
	GetIsSwipeNavigationEnabled ComProc
	PutIsSwipeNavigationEnabled ComProc
}

type ICoreWebView2Settings6 struct {
	vtbl *_ICoreWebView2Settings6Vtbl
}

func (i *ICoreWebView2Settings6) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Settings6) GetIsSwipeNavigationEnabled() (bool, error) {
	var err error

	var enabled bool

	_, _, err = i.vtbl.GetIsSwipeNavigationEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return enabled, nil
}

func (i *ICoreWebView2Settings6) PutIsSwipeNavigationEnabled(enabled bool) error {
	var err error

	_, _, err = i.vtbl.PutIsSwipeNavigationEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
