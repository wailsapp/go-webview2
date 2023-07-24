//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2ContentLoadingEventArgsVtbl struct {
	_IUnknownVtbl
	GetIsErrorPage  ComProc
	GetNavigationId ComProc
}

type ICoreWebView2ContentLoadingEventArgs struct {
	vtbl *_ICoreWebView2ContentLoadingEventArgsVtbl
}

func (i *ICoreWebView2ContentLoadingEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2ContentLoadingEventArgs) GetIsErrorPage() (bool, error) {
	var err error

	var isErrorPage bool

	_, _, err = i.vtbl.GetIsErrorPage.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isErrorPage)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return isErrorPage, nil
}

func (i *ICoreWebView2ContentLoadingEventArgs) GetNavigationId() (*uint64, error) {
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
