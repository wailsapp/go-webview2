//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2BrowserExtensionListVtbl struct {
	IUnknownVtbl
	GetCount        ComProc
	GetValueAtIndex ComProc
}

type ICoreWebView2BrowserExtensionList struct {
	Vtbl *ICoreWebView2BrowserExtensionListVtbl
}

func (i *ICoreWebView2BrowserExtensionList) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2BrowserExtensionList) GetCount() (uint, error) {

	var count uint

	hr, _, err := i.Vtbl.GetCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&count)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return count, err
}

func (i *ICoreWebView2BrowserExtensionList) GetValueAtIndex(index uint) (*ICoreWebView2BrowserExtension, error) {

	var extension *ICoreWebView2BrowserExtension

	hr, _, err := i.Vtbl.GetValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&index)),
		uintptr(unsafe.Pointer(&extension)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return extension, err
}
