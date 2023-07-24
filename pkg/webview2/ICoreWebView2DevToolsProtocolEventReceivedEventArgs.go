//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2DevToolsProtocolEventReceivedEventArgsVtbl struct {
	_IUnknownVtbl
	GetParameterObjectAsJson ComProc
}

type ICoreWebView2DevToolsProtocolEventReceivedEventArgs struct {
	vtbl *_ICoreWebView2DevToolsProtocolEventReceivedEventArgsVtbl
}

func (i *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) GetParameterObjectAsJson() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _parameterObjectAsJson *uint16

	_, _, err = i.vtbl.GetParameterObjectAsJson.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_parameterObjectAsJson)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	parameterObjectAsJson := windows.UTF16PtrToString(_parameterObjectAsJson)
	windows.CoTaskMemFree(unsafe.Pointer(_parameterObjectAsJson))
	return parameterObjectAsJson, nil
}
