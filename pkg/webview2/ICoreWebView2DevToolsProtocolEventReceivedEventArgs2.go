//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2DevToolsProtocolEventReceivedEventArgs2Vtbl struct {
	IUnknownVtbl
	GetSessionId ComProc
}

type ICoreWebView2DevToolsProtocolEventReceivedEventArgs2 struct {
	Vtbl *ICoreWebView2DevToolsProtocolEventReceivedEventArgs2Vtbl
}

func (i *ICoreWebView2DevToolsProtocolEventReceivedEventArgs2) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2DevToolsProtocolEventReceivedEventArgs2() *ICoreWebView2DevToolsProtocolEventReceivedEventArgs2 {
	var result *ICoreWebView2DevToolsProtocolEventReceivedEventArgs2

	iidICoreWebView2DevToolsProtocolEventReceivedEventArgs2 := NewGUID("{2DC4959D-1494-4393-95BA-BEA4CB9EBD1B}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2DevToolsProtocolEventReceivedEventArgs2)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2DevToolsProtocolEventReceivedEventArgs2) GetSessionId() (string, error) {
	// Create *uint16 to hold result
	var _sessionId *uint16

	hr, _, err := i.Vtbl.GetSessionId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_sessionId)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return "", syscall.Errno(hr)
	}
	// Get result and cleanup
	sessionId := UTF16PtrToString(_sessionId)
	CoTaskMemFree(unsafe.Pointer(_sessionId))
	return sessionId, err
}
