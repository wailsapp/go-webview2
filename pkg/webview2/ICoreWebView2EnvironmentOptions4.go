//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2EnvironmentOptions4Vtbl struct {
	_IUnknownVtbl
	GetCustomSchemeRegistrations ComProc
	SetCustomSchemeRegistrations ComProc
}

type ICoreWebView2EnvironmentOptions4 struct {
	vtbl *_ICoreWebView2EnvironmentOptions4Vtbl
}

func (i *ICoreWebView2EnvironmentOptions4) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2EnvironmentOptions4) GetCustomSchemeRegistrations() (*uint32, *ICoreWebView2CustomSchemeRegistration, error) {
	var err error

	var count *uint32
	var schemeRegistrations *ICoreWebView2CustomSchemeRegistration

	_, _, err = i.vtbl.GetCustomSchemeRegistrations.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&count)),
		uintptr(unsafe.Pointer(&schemeRegistrations)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, nil, err
	}
	return count, schemeRegistrations, nil
}

func (i *ICoreWebView2EnvironmentOptions4) SetCustomSchemeRegistrations(count uint32, schemeRegistrations *ICoreWebView2CustomSchemeRegistration) error {
	var err error

	_, _, err = i.vtbl.SetCustomSchemeRegistrations.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&count)),
		uintptr(unsafe.Pointer(&schemeRegistrations)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
