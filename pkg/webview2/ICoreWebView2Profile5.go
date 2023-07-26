//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2Profile5Vtbl struct {
	IUnknownVtbl
	GetCookieManager ComProc
}

type ICoreWebView2Profile5 struct {
	Vtbl *ICoreWebView2Profile5Vtbl
}

func (i *ICoreWebView2Profile5) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2Profile5() *ICoreWebView2Profile5 {
	var result *ICoreWebView2Profile5

	iidICoreWebView2Profile5 := NewGUID("{2EE5B76E-6E80-4DF2-BCD3-D4EC3340A01B}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2Profile5)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2Profile5) GetCookieManager() (*ICoreWebView2CookieManager, error) {

	var cookieManager *ICoreWebView2CookieManager

	hr, _, err := i.Vtbl.GetCookieManager.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cookieManager)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return cookieManager, err
}
