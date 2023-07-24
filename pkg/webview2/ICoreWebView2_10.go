//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2_10Vtbl struct {
	_IUnknownVtbl
	AddBasicAuthenticationRequested    ComProc
	RemoveBasicAuthenticationRequested ComProc
}

type ICoreWebView2_10 struct {
	vtbl *_ICoreWebView2_10Vtbl
}

func (i *ICoreWebView2_10) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2_10) AddBasicAuthenticationRequested(eventHandler *ICoreWebView2BasicAuthenticationRequestedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddBasicAuthenticationRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2_10) RemoveBasicAuthenticationRequested(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveBasicAuthenticationRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
