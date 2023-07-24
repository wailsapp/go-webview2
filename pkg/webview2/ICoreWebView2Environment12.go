//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Environment12Vtbl struct {
	_IUnknownVtbl
	CreateSharedBuffer ComProc
}

type ICoreWebView2Environment12 struct {
	vtbl *_ICoreWebView2Environment12Vtbl
}

func (i *ICoreWebView2Environment12) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Environment12) CreateSharedBuffer(size uint64) (*ICoreWebView2SharedBuffer, error) {
	var err error

	var shared_buffer *ICoreWebView2SharedBuffer

	_, _, err = i.vtbl.CreateSharedBuffer.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&size)),
		uintptr(unsafe.Pointer(&shared_buffer)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return shared_buffer, nil
}
