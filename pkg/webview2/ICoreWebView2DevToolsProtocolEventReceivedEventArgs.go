//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2DevToolsProtocolEventReceivedEventArgsVtbl struct {
	IUnknownVtbl
	GetParameterObjectAsJson ComProc
}

type ICoreWebView2DevToolsProtocolEventReceivedEventArgs struct {
	Vtbl *ICoreWebView2DevToolsProtocolEventReceivedEventArgsVtbl
}

func (i *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) GetParameterObjectAsJson() (string, error) {
	// Create *uint16 to hold result
	var _parameterObjectAsJson *uint16

	hr, _, err := i.Vtbl.GetParameterObjectAsJson.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_parameterObjectAsJson)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return "", syscall.Errno(hr)
	}
	// Get result and cleanup
	parameterObjectAsJson := UTF16PtrToString(_parameterObjectAsJson)
	CoTaskMemFree(unsafe.Pointer(_parameterObjectAsJson))
	return parameterObjectAsJson, err
}
