//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Profile3Vtbl struct {
	_IUnknownVtbl
	GetPreferredTrackingPreventionLevel ComProc
	PutPreferredTrackingPreventionLevel ComProc
}

type ICoreWebView2Profile3 struct {
	vtbl *_ICoreWebView2Profile3Vtbl
}

func (i *ICoreWebView2Profile3) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Profile3) GetPreferredTrackingPreventionLevel() (*COREWEBVIEW2_TRACKING_PREVENTION_LEVEL, error) {
	var err error

	var value *COREWEBVIEW2_TRACKING_PREVENTION_LEVEL

	_, _, err = i.vtbl.GetPreferredTrackingPreventionLevel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2Profile3) PutPreferredTrackingPreventionLevel(value COREWEBVIEW2_TRACKING_PREVENTION_LEVEL) error {
	var err error

	_, _, err = i.vtbl.PutPreferredTrackingPreventionLevel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
