//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2DevToolsProtocolEventReceivedEventArgs2Vtbl struct {
	_IUnknownVtbl
	GetSessionId ComProc
}

type ICoreWebView2DevToolsProtocolEventReceivedEventArgs2 struct {
	vtbl *_ICoreWebView2DevToolsProtocolEventReceivedEventArgs2Vtbl
}

func (i *ICoreWebView2DevToolsProtocolEventReceivedEventArgs2) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2DevToolsProtocolEventReceivedEventArgs2) GetSessionId() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _sessionId *uint16

	_, _, err = i.vtbl.GetSessionId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_sessionId)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	sessionId := windows.UTF16PtrToString(_sessionId)
	windows.CoTaskMemFree(unsafe.Pointer(_sessionId))
	return sessionId, nil
}
