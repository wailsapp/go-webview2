//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2DeferralVtbl struct {
	_IUnknownVtbl
	Complete ComProc
}

type ICoreWebView2Deferral struct {
	vtbl *_ICoreWebView2DeferralVtbl
}

func (i *ICoreWebView2Deferral) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Deferral) Complete() error {
	var err error

	_, _, err = i.vtbl.Complete.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
