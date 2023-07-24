//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2ProcessInfoVtbl struct {
	_IUnknownVtbl
	GetProcessId ComProc
	GetKind      ComProc
}

type ICoreWebView2ProcessInfo struct {
	vtbl *_ICoreWebView2ProcessInfoVtbl
}

func (i *ICoreWebView2ProcessInfo) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2ProcessInfo) GetProcessId() (*INT32, error) {
	var err error

	var value *INT32

	_, _, err = i.vtbl.GetProcessId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ProcessInfo) GetKind() (*COREWEBVIEW2_PROCESS_KIND, error) {
	var err error

	var kind *COREWEBVIEW2_PROCESS_KIND

	_, _, err = i.vtbl.GetKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&kind)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return kind, nil
}
