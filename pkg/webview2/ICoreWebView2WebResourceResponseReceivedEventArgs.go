//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2WebResourceResponseReceivedEventArgsVtbl struct {
	_IUnknownVtbl
	GetRequest  ComProc
	GetResponse ComProc
}

type ICoreWebView2WebResourceResponseReceivedEventArgs struct {
	vtbl *_ICoreWebView2WebResourceResponseReceivedEventArgsVtbl
}

func (i *ICoreWebView2WebResourceResponseReceivedEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2WebResourceResponseReceivedEventArgs) GetRequest() (*ICoreWebView2WebResourceRequest, error) {
	var err error

	var request *ICoreWebView2WebResourceRequest

	_, _, err = i.vtbl.GetRequest.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&request)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return request, nil
}

func (i *ICoreWebView2WebResourceResponseReceivedEventArgs) GetResponse() (*ICoreWebView2WebResourceResponseView, error) {
	var err error

	var response *ICoreWebView2WebResourceResponseView

	_, _, err = i.vtbl.GetResponse.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&response)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return response, nil
}
