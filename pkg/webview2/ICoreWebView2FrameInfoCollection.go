//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2FrameInfoCollectionVtbl struct {
	_IUnknownVtbl
	GetIterator ComProc
}

type ICoreWebView2FrameInfoCollection struct {
	vtbl *_ICoreWebView2FrameInfoCollectionVtbl
}

func (i *ICoreWebView2FrameInfoCollection) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2FrameInfoCollection) GetIterator() (*ICoreWebView2FrameInfoCollectionIterator, error) {
	var err error

	var iterator *ICoreWebView2FrameInfoCollectionIterator

	_, _, err = i.vtbl.GetIterator.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&iterator)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return iterator, nil
}
