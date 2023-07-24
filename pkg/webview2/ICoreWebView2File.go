//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2FileVtbl struct {
	_IUnknownVtbl
	GetPath ComProc
}

type ICoreWebView2File struct {
	vtbl *_ICoreWebView2FileVtbl
}

func (i *ICoreWebView2File) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2File) GetPath() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _path *uint16

	_, _, err = i.vtbl.GetPath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_path)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	path := windows.UTF16PtrToString(_path)
	windows.CoTaskMemFree(unsafe.Pointer(_path))
	return path, nil
}
