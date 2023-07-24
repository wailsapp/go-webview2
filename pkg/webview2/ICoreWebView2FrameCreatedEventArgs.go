//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2FrameCreatedEventArgsVtbl struct {
	_IUnknownVtbl
	GetFrame ComProc
}

type ICoreWebView2FrameCreatedEventArgs struct {
	vtbl *_ICoreWebView2FrameCreatedEventArgsVtbl
}

func (i *ICoreWebView2FrameCreatedEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2FrameCreatedEventArgs) GetFrame() (*ICoreWebView2Frame, error) {
	var err error

	var frame *ICoreWebView2Frame

	_, _, err = i.vtbl.GetFrame.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&frame)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return frame, nil
}
