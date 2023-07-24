//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2ContextMenuItemCollectionVtbl struct {
	_IUnknownVtbl
	GetCount           ComProc
	GetValueAtIndex    ComProc
	RemoveValueAtIndex ComProc
	InsertValueAtIndex ComProc
}

type ICoreWebView2ContextMenuItemCollection struct {
	vtbl *_ICoreWebView2ContextMenuItemCollectionVtbl
}

func (i *ICoreWebView2ContextMenuItemCollection) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2ContextMenuItemCollection) GetCount() (*uint32, error) {
	var err error

	var value *uint32

	_, _, err = i.vtbl.GetCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuItemCollection) GetValueAtIndex(index uint32) (*ICoreWebView2ContextMenuItem, error) {
	var err error

	var value *ICoreWebView2ContextMenuItem

	_, _, err = i.vtbl.GetValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&index)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuItemCollection) RemoveValueAtIndex(index uint32) error {
	var err error

	_, _, err = i.vtbl.RemoveValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&index)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2ContextMenuItemCollection) InsertValueAtIndex(index uint32, value *ICoreWebView2ContextMenuItem) error {
	var err error

	_, _, err = i.vtbl.InsertValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&index)),
		uintptr(unsafe.Pointer(value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
