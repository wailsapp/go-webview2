//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2SourceChangedEventArgsVtbl struct {
	IUnknownVtbl
	GetIsNewDocument ComProc
}

type ICoreWebView2SourceChangedEventArgs struct {
	Vtbl *ICoreWebView2SourceChangedEventArgsVtbl
}

func (i *ICoreWebView2SourceChangedEventArgs) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2SourceChangedEventArgs) GetIsNewDocument() (*bool, error) {
	// Create int32 to hold bool result
	var _isNewDocument int32

	hr, _, err := i.Vtbl.GetIsNewDocument.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_isNewDocument)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	} // Get result and cleanup
	isNewDocument := _isNewDocument != 0
	return &isNewDocument, err
}
