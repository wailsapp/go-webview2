//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Environment4Vtbl struct {
	_IUnknownVtbl
	GetAutomationProviderForWindow ComProc
}

type ICoreWebView2Environment4 struct {
	vtbl *_ICoreWebView2Environment4Vtbl
}

func (i *ICoreWebView2Environment4) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Environment4) GetAutomationProviderForWindow(hwnd HWND) (*_IUnknown, error) {
	var err error

	var provider *_IUnknown

	_, _, err = i.vtbl.GetAutomationProviderForWindow.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&hwnd)),
		uintptr(unsafe.Pointer(&provider)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return provider, nil
}
