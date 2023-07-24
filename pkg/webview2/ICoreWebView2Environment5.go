//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Environment5Vtbl struct {
	_IUnknownVtbl
	AddBrowserProcessExited    ComProc
	RemoveBrowserProcessExited ComProc
}

type ICoreWebView2Environment5 struct {
	vtbl *_ICoreWebView2Environment5Vtbl
}

func (i *ICoreWebView2Environment5) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Environment5) AddBrowserProcessExited(eventHandler *ICoreWebView2BrowserProcessExitedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddBrowserProcessExited.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2Environment5) RemoveBrowserProcessExited(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveBrowserProcessExited.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
