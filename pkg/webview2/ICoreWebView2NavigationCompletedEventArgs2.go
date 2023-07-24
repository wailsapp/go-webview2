//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2NavigationCompletedEventArgs2Vtbl struct {
	_IUnknownVtbl
	GetHttpStatusCode ComProc
}

type ICoreWebView2NavigationCompletedEventArgs2 struct {
	vtbl *_ICoreWebView2NavigationCompletedEventArgs2Vtbl
}

func (i *ICoreWebView2NavigationCompletedEventArgs2) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2NavigationCompletedEventArgs2) GetHttpStatusCode() (int, error) {
	var err error

	var http_status_code int

	_, _, err = i.vtbl.GetHttpStatusCode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(http_status_code),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return http_status_code, nil
}
