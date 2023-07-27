//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2FileVtbl struct {
	IUnknownVtbl
	GetPath ComProc
}

type ICoreWebView2File struct {
	Vtbl *ICoreWebView2FileVtbl
}

func (i *ICoreWebView2File) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2File) GetPath() (string, error) {
	// Create *uint16 to hold result
	var _path *uint16

	hr, _, err := i.Vtbl.GetPath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_path)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return "", syscall.Errno(hr)
	}
	// Get result and cleanup
	path := UTF16PtrToString(_path)
	CoTaskMemFree(unsafe.Pointer(_path))
	return path, err
}
