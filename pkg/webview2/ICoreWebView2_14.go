//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2_14Vtbl struct {
	_IUnknownVtbl
	AddServerCertificateErrorDetected    ComProc
	RemoveServerCertificateErrorDetected ComProc
	ClearServerCertificateErrorActions   ComProc
}

type ICoreWebView2_14 struct {
	vtbl *_ICoreWebView2_14Vtbl
}

func (i *ICoreWebView2_14) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2_14) AddServerCertificateErrorDetected(eventHandler *ICoreWebView2ServerCertificateErrorDetectedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddServerCertificateErrorDetected.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2_14) RemoveServerCertificateErrorDetected(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveServerCertificateErrorDetected.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_14) ClearServerCertificateErrorActions(handler *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler) error {
	var err error

	_, _, err = i.vtbl.ClearServerCertificateErrorActions.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
