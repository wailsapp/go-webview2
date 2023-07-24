//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2WebResourceResponseViewVtbl struct {
	_IUnknownVtbl
	GetHeaders      ComProc
	GetStatusCode   ComProc
	GetReasonPhrase ComProc
	GetContent      ComProc
}

type ICoreWebView2WebResourceResponseView struct {
	vtbl *_ICoreWebView2WebResourceResponseViewVtbl
}

func (i *ICoreWebView2WebResourceResponseView) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2WebResourceResponseView) GetHeaders() (*ICoreWebView2HttpResponseHeaders, error) {
	var err error

	var headers *ICoreWebView2HttpResponseHeaders

	_, _, err = i.vtbl.GetHeaders.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&headers)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return headers, nil
}

func (i *ICoreWebView2WebResourceResponseView) GetStatusCode() (int, error) {
	var err error

	var statusCode int

	_, _, err = i.vtbl.GetStatusCode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(statusCode),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return statusCode, nil
}

func (i *ICoreWebView2WebResourceResponseView) GetReasonPhrase() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _reasonPhrase *uint16

	_, _, err = i.vtbl.GetReasonPhrase.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_reasonPhrase)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	reasonPhrase := windows.UTF16PtrToString(_reasonPhrase)
	windows.CoTaskMemFree(unsafe.Pointer(_reasonPhrase))
	return reasonPhrase, nil
}

func (i *ICoreWebView2WebResourceResponseView) GetContent(handler *ICoreWebView2WebResourceResponseViewGetContentCompletedHandler) error {
	var err error

	_, _, err = i.vtbl.GetContent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
