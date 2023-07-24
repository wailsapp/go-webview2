//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2_6Vtbl struct {
	_IUnknownVtbl
	OpenTaskManagerWindow ComProc
}

type ICoreWebView2_6 struct {
	vtbl *_ICoreWebView2_6Vtbl
}

func (i *ICoreWebView2_6) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2_6) OpenTaskManagerWindow() error {
	var err error

	_, _, err = i.vtbl.OpenTaskManagerWindow.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
