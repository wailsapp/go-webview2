//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2_2Vtbl struct {
	_IUnknownVtbl
	AddWebResourceResponseReceived    ComProc
	RemoveWebResourceResponseReceived ComProc
	NavigateWithWebResourceRequest    ComProc
	AddDOMContentLoaded               ComProc
	RemoveDOMContentLoaded            ComProc
	GetCookieManager                  ComProc
	GetEnvironment                    ComProc
}

type ICoreWebView2_2 struct {
	vtbl *_ICoreWebView2_2Vtbl
}

func (i *ICoreWebView2_2) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2_2) AddWebResourceResponseReceived(eventHandler *ICoreWebView2WebResourceResponseReceivedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddWebResourceResponseReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2_2) RemoveWebResourceResponseReceived(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveWebResourceResponseReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_2) NavigateWithWebResourceRequest(request *ICoreWebView2WebResourceRequest) error {
	var err error

	_, _, err = i.vtbl.NavigateWithWebResourceRequest.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(request)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_2) AddDOMContentLoaded(eventHandler *ICoreWebView2DOMContentLoadedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddDOMContentLoaded.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2_2) RemoveDOMContentLoaded(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveDOMContentLoaded.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_2) GetCookieManager() (*ICoreWebView2CookieManager, error) {
	var err error

	var cookieManager *ICoreWebView2CookieManager

	_, _, err = i.vtbl.GetCookieManager.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cookieManager)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return cookieManager, nil
}

func (i *ICoreWebView2_2) GetEnvironment() (*ICoreWebView2Environment, error) {
	var err error

	var environment *ICoreWebView2Environment

	_, _, err = i.vtbl.GetEnvironment.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&environment)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return environment, nil
}
