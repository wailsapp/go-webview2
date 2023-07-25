//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2_17Vtbl struct {
	IUnknownVtbl
	PostSharedBufferToScript ComProc
}

type ICoreWebView2_17 struct {
	Vtbl *ICoreWebView2_17Vtbl
}

func (i *ICoreWebView2_17) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2_17() *ICoreWebView2_17 {
	var result *ICoreWebView2_17

	iidICoreWebView2_17 := NewGUID("{702E75D4-FD44-434D-9D70-1A68A6B1192A}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_17)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2_17) PostSharedBufferToScript(sharedBuffer *ICoreWebView2SharedBuffer, access COREWEBVIEW2_SHARED_BUFFER_ACCESS, additionalDataAsJson string) error {

	// Convert string 'additionalDataAsJson' to *uint16
	_additionalDataAsJson, err := UTF16PtrFromString(additionalDataAsJson)
	if err != nil {
		return err
	}

	hr, _, err := i.Vtbl.PostSharedBufferToScript.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(sharedBuffer)),
		uintptr(access),
		uintptr(unsafe.Pointer(_additionalDataAsJson)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
