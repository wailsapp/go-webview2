//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Environment11Vtbl struct {
	_IUnknownVtbl
	GetFailureReportFolderPath ComProc
}

type ICoreWebView2Environment11 struct {
	vtbl *_ICoreWebView2Environment11Vtbl
}

func (i *ICoreWebView2Environment11) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Environment11) GetFailureReportFolderPath() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetFailureReportFolderPath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	value := windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}
