//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2DOMContentLoadedEventArgsVtbl struct {
	_IUnknownVtbl
	GetNavigationId ComProc
}

type ICoreWebView2DOMContentLoadedEventArgs struct {
	vtbl *_ICoreWebView2DOMContentLoadedEventArgsVtbl
}

func (i *ICoreWebView2DOMContentLoadedEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2DOMContentLoadedEventArgs) GetNavigationId() (*uint64, error) {
	var err error

	var navigationId *uint64

	_, _, err = i.vtbl.GetNavigationId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&navigationId)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return navigationId, nil
}
