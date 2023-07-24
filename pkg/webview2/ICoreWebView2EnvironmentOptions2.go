//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2EnvironmentOptions2Vtbl struct {
	_IUnknownVtbl
	GetExclusiveUserDataFolderAccess ComProc
	PutExclusiveUserDataFolderAccess ComProc
}

type ICoreWebView2EnvironmentOptions2 struct {
	vtbl *_ICoreWebView2EnvironmentOptions2Vtbl
}

func (i *ICoreWebView2EnvironmentOptions2) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2EnvironmentOptions2) GetExclusiveUserDataFolderAccess() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetExclusiveUserDataFolderAccess.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2EnvironmentOptions2) PutExclusiveUserDataFolderAccess(value bool) error {
	var err error

	_, _, err = i.vtbl.PutExclusiveUserDataFolderAccess.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
