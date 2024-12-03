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

func (i *ICoreWebView2ObjectCollectionView) AddRef() uint32 {
	ret, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))

	return uint32(ret)
}

func (i *ICoreWebView2ObjectCollectionView) Release() uint32 {
	ret, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))

	return uint32(ret)
}

func (i *ICoreWebView2ObjectCollectionView) GetCount() (uint32, error) {
	
	var value uint32
	hr, _, _ := i.vtbl.GetCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, windows.Errno(hr)
	}
	return value, nil
}

func (i *ICoreWebView2ObjectCollectionView) GetValueAtIndex(index uint32) (*_IUnknownVtbl, error) {
	
	var value *_IUnknownVtbl
	hr, _, _ := i.vtbl.GetValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(index),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, windows.Errno(hr)
	}
	return value, nil
}
