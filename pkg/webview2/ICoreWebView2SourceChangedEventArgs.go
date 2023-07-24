//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2SourceChangedEventArgsVtbl struct {
	_IUnknownVtbl
	GetIsNewDocument ComProc
}

type ICoreWebView2SourceChangedEventArgs struct {
	vtbl *_ICoreWebView2SourceChangedEventArgsVtbl
}

func (i *ICoreWebView2SourceChangedEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2SourceChangedEventArgs) GetIsNewDocument() (bool, error) {
	var err error

	var isNewDocument bool

	_, _, err = i.vtbl.GetIsNewDocument.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isNewDocument)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return isNewDocument, nil
}
