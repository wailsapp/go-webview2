//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Profile5Vtbl struct {
	_IUnknownVtbl
	GetCookieManager ComProc
}

type ICoreWebView2Profile5 struct {
	vtbl *_ICoreWebView2Profile5Vtbl
}

func (i *ICoreWebView2Profile5) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Profile5) GetCookieManager() (*ICoreWebView2CookieManager, error) {
	var err error

	var cookieManager *ICoreWebView2CookieManager

	_, _, err = i.vtbl.GetCookieManager.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cookieManager)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return cookieManager, nil
}
