//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Frame3Vtbl struct {
	_IUnknownVtbl
	AddPermissionRequested    ComProc
	RemovePermissionRequested ComProc
}

type ICoreWebView2Frame3 struct {
	vtbl *_ICoreWebView2Frame3Vtbl
}

func (i *ICoreWebView2Frame3) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Frame3) AddPermissionRequested(handler *ICoreWebView2FramePermissionRequestedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddPermissionRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2Frame3) RemovePermissionRequested(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemovePermissionRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
