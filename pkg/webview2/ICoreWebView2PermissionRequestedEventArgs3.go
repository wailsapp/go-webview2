//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2PermissionRequestedEventArgs3Vtbl struct {
	_IUnknownVtbl
	GetSavesInProfile ComProc
	PutSavesInProfile ComProc
}

type ICoreWebView2PermissionRequestedEventArgs3 struct {
	vtbl *_ICoreWebView2PermissionRequestedEventArgs3Vtbl
}

func (i *ICoreWebView2PermissionRequestedEventArgs3) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2PermissionRequestedEventArgs3) GetSavesInProfile() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetSavesInProfile.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2PermissionRequestedEventArgs3) PutSavesInProfile(value bool) error {
	var err error

	_, _, err = i.vtbl.PutSavesInProfile.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
