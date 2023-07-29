//go:build windows

package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type _ICoreWebView2ObjectCollectionViewVtbl struct {
	_IUnknownVtbl
	GetCount        ComProc
	GetValueAtIndex ComProc
}

type ICoreWebView2ObjectCollectionView struct {
	vtbl *_ICoreWebView2ObjectCollectionViewVtbl
}

func (i *ICoreWebView2ObjectCollectionView) Release() error {
	return i.vtbl.CallRelease(unsafe.Pointer(i))
}

func (i *ICoreWebView2ObjectCollectionView) GetCount() (uint32, error) {
	var err error
	var value uint32
	_, _, err = i.vtbl.GetCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return value, nil
}

func (i *ICoreWebView2ObjectCollectionView) GetValueAtIndex(index uint32) (*_IUnknownVtbl, error) {
	var err error
	var value *_IUnknownVtbl
	_, _, err = i.vtbl.GetValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(index),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}
