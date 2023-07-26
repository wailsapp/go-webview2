//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2FrameCreatedEventArgsVtbl struct {
	IUnknownVtbl
	GetFrame ComProc
}

type ICoreWebView2FrameCreatedEventArgs struct {
	Vtbl *ICoreWebView2FrameCreatedEventArgsVtbl
}

func (i *ICoreWebView2FrameCreatedEventArgs) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2FrameCreatedEventArgs) GetFrame() (*ICoreWebView2Frame, error) {

	var frame *ICoreWebView2Frame

	hr, _, err := i.Vtbl.GetFrame.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&frame)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return frame, err
}
