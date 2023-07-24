//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2_5Vtbl struct {
	_IUnknownVtbl
	AddClientCertificateRequested    ComProc
	RemoveClientCertificateRequested ComProc
}

type ICoreWebView2_5 struct {
	vtbl *_ICoreWebView2_5Vtbl
}

func (i *ICoreWebView2_5) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2_5) AddClientCertificateRequested(eventHandler *ICoreWebView2ClientCertificateRequestedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddClientCertificateRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2_5) RemoveClientCertificateRequested(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveClientCertificateRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
