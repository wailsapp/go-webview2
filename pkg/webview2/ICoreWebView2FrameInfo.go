//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2FrameInfoVtbl struct {
	_IUnknownVtbl
	GetName   ComProc
	GetSource ComProc
}

type ICoreWebView2FrameInfo struct {
	vtbl *_ICoreWebView2FrameInfoVtbl
}

func (i *ICoreWebView2FrameInfo) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2FrameInfo) GetName() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _name *uint16

	_, _, err = i.vtbl.GetName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	name := windows.UTF16PtrToString(_name)
	windows.CoTaskMemFree(unsafe.Pointer(_name))
	return name, nil
}

func (i *ICoreWebView2FrameInfo) GetSource() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _source *uint16

	_, _, err = i.vtbl.GetSource.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_source)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	source := windows.UTF16PtrToString(_source)
	windows.CoTaskMemFree(unsafe.Pointer(_source))
	return source, nil
}
