//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2ContentLoadingEventArgsVtbl struct {
	IUnknownVtbl
	GetIsErrorPage  ComProc
	GetNavigationId ComProc
}

type ICoreWebView2ContentLoadingEventArgs struct {
	Vtbl *ICoreWebView2ContentLoadingEventArgsVtbl
}

func (i *ICoreWebView2ContentLoadingEventArgs) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2ContentLoadingEventArgs) GetIsErrorPage() (bool, error) {
	// Create int32 to hold bool result
	var _isErrorPage int32

	hr, _, err := i.Vtbl.GetIsErrorPage.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_isErrorPage)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return false, syscall.Errno(hr)
	}
	// Get result and cleanup
	isErrorPage := _isErrorPage != 0
	return isErrorPage, err
}

func (i *ICoreWebView2ContentLoadingEventArgs) GetNavigationId() (uint64, error) {

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
