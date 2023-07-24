//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Environment6Vtbl struct {
	_IUnknownVtbl
	CreatePrintSettings ComProc
}

type ICoreWebView2Environment6 struct {
	vtbl *_ICoreWebView2Environment6Vtbl
}

func (i *ICoreWebView2Environment6) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Environment6) CreatePrintSettings() (*ICoreWebView2PrintSettings, error) {
	var err error

	var printSettings *ICoreWebView2PrintSettings

	_, _, err = i.vtbl.CreatePrintSettings.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&printSettings)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return printSettings, nil
}
