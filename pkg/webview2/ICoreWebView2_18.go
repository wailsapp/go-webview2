//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2_18Vtbl struct {
	_IUnknownVtbl
	AddLaunchingExternalUriScheme    ComProc
	RemoveLaunchingExternalUriScheme ComProc
}

type ICoreWebView2_18 struct {
	vtbl *_ICoreWebView2_18Vtbl
}

func (i *ICoreWebView2_18) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2_18) AddLaunchingExternalUriScheme(eventHandler *ICoreWebView2LaunchingExternalUriSchemeEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddLaunchingExternalUriScheme.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2_18) RemoveLaunchingExternalUriScheme(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveLaunchingExternalUriScheme.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
