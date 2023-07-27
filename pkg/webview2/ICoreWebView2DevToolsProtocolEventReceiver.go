//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2DevToolsProtocolEventReceiverVtbl struct {
	IUnknownVtbl
	AddDevToolsProtocolEventReceived    ComProc
	RemoveDevToolsProtocolEventReceived ComProc
}

type ICoreWebView2DevToolsProtocolEventReceiver struct {
	Vtbl *ICoreWebView2DevToolsProtocolEventReceiverVtbl
}

func (i *ICoreWebView2DevToolsProtocolEventReceiver) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2DevToolsProtocolEventReceiver) AddDevToolsProtocolEventReceived(handler *ICoreWebView2DevToolsProtocolEventReceivedEventHandler) (EventRegistrationToken, error) {

	var token EventRegistrationToken

	hr, _, err := i.Vtbl.AddDevToolsProtocolEventReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return EventRegistrationToken{}, syscall.Errno(hr)
	}
	return token, err
}

func (i *ICoreWebView2DevToolsProtocolEventReceiver) RemoveDevToolsProtocolEventReceived(token EventRegistrationToken) error {

	hr, _, err := i.Vtbl.RemoveDevToolsProtocolEventReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
