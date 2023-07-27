//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2Settings6Vtbl struct {
	IUnknownVtbl
	GetIsSwipeNavigationEnabled ComProc
	PutIsSwipeNavigationEnabled ComProc
}

type ICoreWebView2Settings6 struct {
	Vtbl *ICoreWebView2Settings6Vtbl
}

func (i *ICoreWebView2Settings6) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2Settings6() *ICoreWebView2Settings6 {
	var result *ICoreWebView2Settings6

	iidICoreWebView2Settings6 := NewGUID("{11cb3acd-9bc8-43b8-83bf-f40753714f87}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2Settings6)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2Settings6) GetIsSwipeNavigationEnabled() (*bool, error) {
	// Create int32 to hold bool result
	var _enabled int32

	hr, _, err := i.Vtbl.GetIsSwipeNavigationEnabled.Call(
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

func (i *ICoreWebView2Settings6) PutIsSwipeNavigationEnabled(enabled bool) error {

	hr, _, err := i.Vtbl.PutIsSwipeNavigationEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&enabled)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
