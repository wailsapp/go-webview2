//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2ProcessFailedEventArgsVtbl struct {
	_IUnknownVtbl
	GetProcessFailedKind ComProc
}

type ICoreWebView2ProcessFailedEventArgs struct {
	vtbl *_ICoreWebView2ProcessFailedEventArgsVtbl
}

func (i *ICoreWebView2ProcessFailedEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2ProcessFailedEventArgs) GetProcessFailedKind() (*COREWEBVIEW2_PROCESS_FAILED_KIND, error) {
	var err error

	var processFailedKind *COREWEBVIEW2_PROCESS_FAILED_KIND

	_, _, err = i.vtbl.GetProcessFailedKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&processFailedKind)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return processFailedKind, nil
}
