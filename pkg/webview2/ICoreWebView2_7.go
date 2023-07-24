//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2_7Vtbl struct {
	_IUnknownVtbl
	PrintToPdf ComProc
}

type ICoreWebView2_7 struct {
	vtbl *_ICoreWebView2_7Vtbl
}

func (i *ICoreWebView2_7) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2_7) PrintToPdf(resultFilePath string, printSettings *ICoreWebView2PrintSettings, handler *ICoreWebView2PrintToPdfCompletedHandler) error {
	var err error

	// Convert string 'resultFilePath' to *uint16
	_resultFilePath, err := windows.UTF16PtrFromString(resultFilePath)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PrintToPdf.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_resultFilePath)),
		uintptr(unsafe.Pointer(printSettings)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
