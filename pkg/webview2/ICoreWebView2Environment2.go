//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Environment2Vtbl struct {
	_IUnknownVtbl
	CreateWebResourceRequest ComProc
}

type ICoreWebView2Environment2 struct {
	vtbl *_ICoreWebView2Environment2Vtbl
}

func (i *ICoreWebView2Environment2) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Environment2) CreateWebResourceRequest(uri string, method string, postData *IStream, headers string) (*ICoreWebView2WebResourceRequest, error) {
	var err error

	// Convert string 'uri' to *uint16
	_uri, err := windows.UTF16PtrFromString(uri)
	if err != nil {
		return nil, err
	}

	// Convert string 'method' to *uint16
	_method, err := windows.UTF16PtrFromString(method)
	if err != nil {
		return nil, err
	}

	// Convert string 'headers' to *uint16
	_headers, err := windows.UTF16PtrFromString(headers)
	if err != nil {
		return nil, err
	}

	var request *ICoreWebView2WebResourceRequest

	_, _, err = i.vtbl.CreateWebResourceRequest.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
		uintptr(unsafe.Pointer(_method)),
		uintptr(unsafe.Pointer(postData)),
		uintptr(unsafe.Pointer(_headers)),
		uintptr(unsafe.Pointer(&request)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return request, nil
}
