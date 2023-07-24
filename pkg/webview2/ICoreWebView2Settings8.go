//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Settings8Vtbl struct {
	_IUnknownVtbl
	GetIsReputationCheckingRequired ComProc
	PutIsReputationCheckingRequired ComProc
}

type ICoreWebView2Settings8 struct {
	vtbl *_ICoreWebView2Settings8Vtbl
}

func (i *ICoreWebView2Settings8) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Settings8) GetIsReputationCheckingRequired() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetIsReputationCheckingRequired.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2Settings8) PutIsReputationCheckingRequired(value bool) error {
	var err error

	_, _, err = i.vtbl.PutIsReputationCheckingRequired.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
