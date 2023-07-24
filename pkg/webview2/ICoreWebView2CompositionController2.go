//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2CompositionController2Vtbl struct {
	_IUnknownVtbl
	GetAutomationProvider ComProc
}

type ICoreWebView2CompositionController2 struct {
	vtbl *_ICoreWebView2CompositionController2Vtbl
}

func (i *ICoreWebView2CompositionController2) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2CompositionController2) GetAutomationProvider() (*_IUnknown, error) {
	var err error

	var provider *_IUnknown

	_, _, err = i.vtbl.GetAutomationProvider.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&provider)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return provider, nil
}
