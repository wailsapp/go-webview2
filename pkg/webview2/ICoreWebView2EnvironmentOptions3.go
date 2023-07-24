//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2EnvironmentOptions3Vtbl struct {
	_IUnknownVtbl
	GetIsCustomCrashReportingEnabled ComProc
	PutIsCustomCrashReportingEnabled ComProc
}

type ICoreWebView2EnvironmentOptions3 struct {
	vtbl *_ICoreWebView2EnvironmentOptions3Vtbl
}

func (i *ICoreWebView2EnvironmentOptions3) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2EnvironmentOptions3) GetIsCustomCrashReportingEnabled() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetIsCustomCrashReportingEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2EnvironmentOptions3) PutIsCustomCrashReportingEnabled(value bool) error {
	var err error

	_, _, err = i.vtbl.PutIsCustomCrashReportingEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
