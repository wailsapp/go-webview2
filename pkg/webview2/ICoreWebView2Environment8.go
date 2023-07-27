//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2Environment8Vtbl struct {
	IUnknownVtbl
	AddProcessInfosChanged    ComProc
	RemoveProcessInfosChanged ComProc
	GetProcessInfos           ComProc
}

type ICoreWebView2Environment8 struct {
	Vtbl *ICoreWebView2Environment8Vtbl
}

func (i *ICoreWebView2Environment8) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2Environment8() *ICoreWebView2Environment8 {
	var result *ICoreWebView2Environment8

	iidICoreWebView2Environment8 := NewGUID("{D6EB91DD-C3D2-45E5-BD29-6DC2BC4DE9CF}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2Environment8)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2Environment8) AddProcessInfosChanged(eventHandler *ICoreWebView2ProcessInfosChangedEventHandler) (EventRegistrationToken, error) {

	var token EventRegistrationToken

	hr, _, err := i.Vtbl.AddProcessInfosChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return EventRegistrationToken{}, syscall.Errno(hr)
	}
	return token, err
}

func (i *ICoreWebView2Environment8) RemoveProcessInfosChanged(token EventRegistrationToken) error {

	hr, _, err := i.Vtbl.RemoveProcessInfosChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2Environment8) GetProcessInfos() (*ICoreWebView2ProcessInfoCollection, error) {

	var value *ICoreWebView2ProcessInfoCollection

	hr, _, err := i.Vtbl.GetProcessInfos.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return value, err
}
