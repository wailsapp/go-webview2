//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2PermissionRequestedEventArgs2Vtbl struct {
	_IUnknownVtbl
	GetHandled ComProc
	PutHandled ComProc
}

type ICoreWebView2PermissionRequestedEventArgs2 struct {
	vtbl *_ICoreWebView2PermissionRequestedEventArgs2Vtbl
}

func (i *ICoreWebView2PermissionRequestedEventArgs2) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2PermissionRequestedEventArgs2) GetHandled() (bool, error) {
	var err error

	var handled bool

	_, _, err = i.vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handled)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return handled, nil
}

func (i *ICoreWebView2PermissionRequestedEventArgs2) PutHandled(handled bool) error {
	var err error

	_, _, err = i.vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handled)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
