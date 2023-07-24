//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2PermissionSettingCollectionViewVtbl struct {
	_IUnknownVtbl
	GetValueAtIndex ComProc
	GetCount        ComProc
}

type ICoreWebView2PermissionSettingCollectionView struct {
	vtbl *_ICoreWebView2PermissionSettingCollectionViewVtbl
}

func (i *ICoreWebView2PermissionSettingCollectionView) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2PermissionSettingCollectionView) GetValueAtIndex(index uint32) (*ICoreWebView2PermissionSetting, error) {
	var err error

	var permissionSetting *ICoreWebView2PermissionSetting

	_, _, err = i.vtbl.GetValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&index)),
		uintptr(unsafe.Pointer(&permissionSetting)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return permissionSetting, nil
}

func (i *ICoreWebView2PermissionSettingCollectionView) GetCount() (*uint32, error) {
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
