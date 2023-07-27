//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2DOMContentLoadedEventArgsVtbl struct {
	IUnknownVtbl
	GetNavigationId ComProc
}

type ICoreWebView2DOMContentLoadedEventArgs struct {
	Vtbl *ICoreWebView2DOMContentLoadedEventArgsVtbl
}

func (i *ICoreWebView2DOMContentLoadedEventArgs) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2DOMContentLoadedEventArgs) GetNavigationId() (uint64, error) {

	var navigationId uint64

	hr, _, err := i.Vtbl.GetNavigationId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&navigationId)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return navigationId, err
}
