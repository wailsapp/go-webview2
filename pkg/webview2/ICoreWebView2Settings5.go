//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2Settings5Vtbl struct {
	IUnknownVtbl
	GetIsPinchZoomEnabled ComProc
	PutIsPinchZoomEnabled ComProc
}

type ICoreWebView2Settings5 struct {
	Vtbl *ICoreWebView2Settings5Vtbl
}

func (i *ICoreWebView2Settings5) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2Settings5() *ICoreWebView2Settings5 {
	var result *ICoreWebView2Settings5

	iidICoreWebView2Settings5 := NewGUID("{183e7052-1d03-43a0-ab99-98e043b66b39}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2Settings5)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2Settings5) GetIsPinchZoomEnabled() (*bool, error) {
	// Create int32 to hold bool result
	var _enabled int32

	hr, _, err := i.Vtbl.GetIsPinchZoomEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_enabled)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	// Get result and cleanup
	enabled := ptr(_enabled != 0)
	return enabled, err
}

func (i *ICoreWebView2Settings5) PutIsPinchZoomEnabled(enabled bool) error {

	hr, _, err := i.Vtbl.PutIsPinchZoomEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
