//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2_13Vtbl struct {
	_IUnknownVtbl
	GetProfile ComProc
}

type ICoreWebView2_13 struct {
	vtbl *_ICoreWebView2_13Vtbl
}

func (i *ICoreWebView2_13) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2_13) GetProfile() (*ICoreWebView2Profile, error) {
	var err error

	var value *ICoreWebView2Profile

	_, _, err = i.vtbl.GetProfile.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}
