//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2_16Vtbl struct {
	_IUnknownVtbl
	Print            ComProc
	ShowPrintUI      ComProc
	PrintToPdfStream ComProc
}

type ICoreWebView2_16 struct {
	vtbl *_ICoreWebView2_16Vtbl
}

func (i *ICoreWebView2_16) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2_16) Print(printSettings *ICoreWebView2PrintSettings, handler *ICoreWebView2PrintCompletedHandler) error {
	var err error

	_, _, err = i.vtbl.Print.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(printSettings)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_16) ShowPrintUI(printDialogKind COREWEBVIEW2_PRINT_DIALOG_KIND) error {
	var err error

	_, _, err = i.vtbl.ShowPrintUI.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(printDialogKind),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_16) PrintToPdfStream(printSettings *ICoreWebView2PrintSettings, handler *ICoreWebView2PrintToPdfStreamCompletedHandler) error {
	var err error

	_, _, err = i.vtbl.PrintToPdfStream.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(printSettings)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
