//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2_17Vtbl struct {
	_IUnknownVtbl
	PostSharedBufferToScript ComProc
}

type ICoreWebView2_17 struct {
	vtbl *_ICoreWebView2_17Vtbl
}

func (i *ICoreWebView2_17) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2_17) PostSharedBufferToScript(sharedBuffer *ICoreWebView2SharedBuffer, access COREWEBVIEW2_SHARED_BUFFER_ACCESS, additionalDataAsJson string) error {
	var err error

	// Convert string 'additionalDataAsJson' to *uint16
	_additionalDataAsJson, err := windows.UTF16PtrFromString(additionalDataAsJson)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PostSharedBufferToScript.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(sharedBuffer)),
		uintptr(access),
		uintptr(unsafe.Pointer(_additionalDataAsJson)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
