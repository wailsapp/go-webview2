//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Environment9Vtbl struct {
	_IUnknownVtbl
	CreateContextMenuItem ComProc
}

type ICoreWebView2Environment9 struct {
	vtbl *_ICoreWebView2Environment9Vtbl
}

func (i *ICoreWebView2Environment9) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Environment9) CreateContextMenuItem(label string, iconStream *IStream, kind COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND) (*ICoreWebView2ContextMenuItem, error) {
	var err error

	// Convert string 'label' to *uint16
	_label, err := windows.UTF16PtrFromString(label)
	if err != nil {
		return nil, err
	}

	var item *ICoreWebView2ContextMenuItem

	_, _, err = i.vtbl.CreateContextMenuItem.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_label)),
		uintptr(unsafe.Pointer(iconStream)),
		uintptr(kind),
		uintptr(unsafe.Pointer(&item)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return item, nil
}
