//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2WebResourceResponseViewVtbl struct {
	IUnknownVtbl
	GetHeaders      ComProc
	GetStatusCode   ComProc
	GetReasonPhrase ComProc
	GetContent      ComProc
}

type ICoreWebView2WebResourceResponseView struct {
	Vtbl *ICoreWebView2WebResourceResponseViewVtbl
}

func (i *ICoreWebView2WebResourceResponseView) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2WebResourceResponseView) GetHeaders() (*ICoreWebView2HttpResponseHeaders, error) {

	var headers ICoreWebView2HttpResponseHeaders

	hr, _, err := i.Vtbl.GetHeaders.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&headers)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return &headers, err
}

func (i *ICoreWebView2WebResourceResponseView) GetStatusCode() (*int, error) {

	var statusCode int

	hr, _, err := i.Vtbl.GetStatusCode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(statusCode),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return &statusCode, err
}

func (i *ICoreWebView2WebResourceResponseView) GetReasonPhrase() (*string, error) {
	// Create *uint16 to hold result
	var _reasonPhrase *uint16

	hr, _, err := i.Vtbl.GetReasonPhrase.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_reasonPhrase)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	} // Get result and cleanup
	reasonPhrase := UTF16PtrToString(_reasonPhrase)
	CoTaskMemFree(unsafe.Pointer(_reasonPhrase))
	return &reasonPhrase, err
}

func (i *ICoreWebView2WebResourceResponseView) GetContent(handler *ICoreWebView2WebResourceResponseViewGetContentCompletedHandler) error {

	hr, _, err := i.Vtbl.GetContent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
