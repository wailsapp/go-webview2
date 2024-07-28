//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2WebResourceRequestedEventArgs2Vtbl struct {
	IUnknownVtbl
	GetRequestedSourceKind ComProc
}

type ICoreWebView2WebResourceRequestedEventArgs2 struct {
	Vtbl *ICoreWebView2WebResourceRequestedEventArgs2Vtbl
}

func (i *ICoreWebView2WebResourceRequestedEventArgs2) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2WebResourceRequestedEventArgs2() *ICoreWebView2WebResourceRequestedEventArgs2 {
	var result *ICoreWebView2WebResourceRequestedEventArgs2

	iidICoreWebView2WebResourceRequestedEventArgs2 := NewGUID("{9C562C24-B219-4D7F-92F6-B187FBBADD56}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2WebResourceRequestedEventArgs2)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2WebResourceRequestedEventArgs2) GetRequestedSourceKind() (COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS, error) {

	var requestedSourceKind COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS

	hr, _, err := i.Vtbl.GetRequestedSourceKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&requestedSourceKind)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return requestedSourceKind, err
}
