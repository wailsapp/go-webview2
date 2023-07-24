//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Settings7Vtbl struct {
	_IUnknownVtbl
	GetHiddenPdfToolbarItems ComProc
	PutHiddenPdfToolbarItems ComProc
}

type ICoreWebView2Settings7 struct {
	vtbl *_ICoreWebView2Settings7Vtbl
}

func (i *ICoreWebView2Settings7) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Settings7) GetHiddenPdfToolbarItems() (*COREWEBVIEW2_PDF_TOOLBAR_ITEMS, error) {
	var err error

	var hidden_pdf_toolbar_items *COREWEBVIEW2_PDF_TOOLBAR_ITEMS

	_, _, err = i.vtbl.GetHiddenPdfToolbarItems.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&hidden_pdf_toolbar_items)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return hidden_pdf_toolbar_items, nil
}

func (i *ICoreWebView2Settings7) PutHiddenPdfToolbarItems(hidden_pdf_toolbar_items COREWEBVIEW2_PDF_TOOLBAR_ITEMS) error {
	var err error

	_, _, err = i.vtbl.PutHiddenPdfToolbarItems.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(hidden_pdf_toolbar_items),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
