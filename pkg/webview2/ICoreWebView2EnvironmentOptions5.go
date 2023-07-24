//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2EnvironmentOptions5Vtbl struct {
	_IUnknownVtbl
	GetEnableTrackingPrevention ComProc
	PutEnableTrackingPrevention ComProc
}

type ICoreWebView2EnvironmentOptions5 struct {
	vtbl *_ICoreWebView2EnvironmentOptions5Vtbl
}

func (i *ICoreWebView2EnvironmentOptions5) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2EnvironmentOptions5) GetEnableTrackingPrevention() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetEnableTrackingPrevention.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2EnvironmentOptions5) PutEnableTrackingPrevention(value bool) error {
	var err error

	_, _, err = i.vtbl.PutEnableTrackingPrevention.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
