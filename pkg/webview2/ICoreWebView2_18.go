//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2_18Vtbl struct {
	IUnknownVtbl
	AddLaunchingExternalUriScheme    ComProc
	RemoveLaunchingExternalUriScheme ComProc
}

type ICoreWebView2_18 struct {
	Vtbl *ICoreWebView2_18Vtbl
}

func (i *ICoreWebView2_18) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2_18() *ICoreWebView2_18 {
	var result *ICoreWebView2_18

	iidICoreWebView2_18 := NewGUID("{7A626017-28BE-49B2-B865-3BA2B3522D90}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_18)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2_18) AddLaunchingExternalUriScheme(eventHandler *ICoreWebView2LaunchingExternalUriSchemeEventHandler) (EventRegistrationToken, error) {

	var token EventRegistrationToken

	hr, _, err := i.Vtbl.AddLaunchingExternalUriScheme.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return EventRegistrationToken{}, syscall.Errno(hr)
	}
	return token, err
}

func (i *ICoreWebView2_18) RemoveLaunchingExternalUriScheme(token EventRegistrationToken) error {

	hr, _, err := i.Vtbl.RemoveLaunchingExternalUriScheme.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
