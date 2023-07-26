//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2Environment2Vtbl struct {
	IUnknownVtbl
	CreateWebResourceRequest ComProc
}

type ICoreWebView2Environment2 struct {
	Vtbl *ICoreWebView2Environment2Vtbl
}

func (i *ICoreWebView2Environment2) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2Environment2() *ICoreWebView2Environment2 {
	var result *ICoreWebView2Environment2

	iidICoreWebView2Environment2 := NewGUID("{41F3632B-5EF4-404F-AD82-2D606C5A9A21}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2Environment2)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2Environment2) CreateWebResourceRequest(uri string, method string, postData *IStream, headers string) (*ICoreWebView2WebResourceRequest, error) {

	// Convert string 'uri' to *uint16
	_uri, err := UTF16PtrFromString(uri)
	if err != nil {
		return nil, err
	}
	// Convert string 'method' to *uint16
	_method, err := UTF16PtrFromString(method)
	if err != nil {
		return nil, err
	}
	// Convert string 'headers' to *uint16
	_headers, err := UTF16PtrFromString(headers)
	if err != nil {
		return nil, err
	}
	var request *ICoreWebView2WebResourceRequest

	hr, _, err := i.Vtbl.CreateWebResourceRequest.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
		uintptr(unsafe.Pointer(_method)),
		uintptr(unsafe.Pointer(postData)),
		uintptr(unsafe.Pointer(_headers)),
		uintptr(unsafe.Pointer(&request)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return request, err
}
