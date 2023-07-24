//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Environment8Vtbl struct {
	_IUnknownVtbl
	AddProcessInfosChanged    ComProc
	RemoveProcessInfosChanged ComProc
	GetProcessInfos           ComProc
}

type ICoreWebView2Environment8 struct {
	vtbl *_ICoreWebView2Environment8Vtbl
}

func (i *ICoreWebView2Environment8) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Environment8) AddProcessInfosChanged(eventHandler *ICoreWebView2ProcessInfosChangedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddProcessInfosChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2Environment8) RemoveProcessInfosChanged(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveProcessInfosChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Environment8) GetProcessInfos() (*ICoreWebView2ProcessInfoCollection, error) {
	var err error

	var value *ICoreWebView2ProcessInfoCollection

	_, _, err = i.vtbl.GetProcessInfos.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}
