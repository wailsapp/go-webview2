//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2WebMessageReceivedEventArgs2Vtbl struct {
	_IUnknownVtbl
	GetAdditionalObjects ComProc
}

type ICoreWebView2WebMessageReceivedEventArgs2 struct {
	vtbl *_ICoreWebView2WebMessageReceivedEventArgs2Vtbl
}

func (i *ICoreWebView2WebMessageReceivedEventArgs2) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2WebMessageReceivedEventArgs2) GetAdditionalObjects() (*ICoreWebView2ObjectCollectionView, error) {
	var err error

	var value *ICoreWebView2ObjectCollectionView

	_, _, err = i.vtbl.GetAdditionalObjects.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}
