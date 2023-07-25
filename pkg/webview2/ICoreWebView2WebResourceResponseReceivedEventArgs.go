//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2WebResourceResponseReceivedEventArgsVtbl struct {
	IUnknownVtbl
	GetRequest  ComProc
	GetResponse ComProc
}

type ICoreWebView2WebResourceResponseReceivedEventArgs struct {
	Vtbl *ICoreWebView2WebResourceResponseReceivedEventArgsVtbl
}

func (i *ICoreWebView2WebResourceResponseReceivedEventArgs) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2WebResourceResponseReceivedEventArgs) GetRequest() (*ICoreWebView2WebResourceRequest, error) {

	var request ICoreWebView2WebResourceRequest

	hr, _, err := i.Vtbl.GetRequest.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&request)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return &request, err
}

func (i *ICoreWebView2WebResourceResponseReceivedEventArgs) GetResponse() (*ICoreWebView2WebResourceResponseView, error) {

	var response ICoreWebView2WebResourceResponseView

	hr, _, err := i.Vtbl.GetResponse.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&response)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return &response, err
}
