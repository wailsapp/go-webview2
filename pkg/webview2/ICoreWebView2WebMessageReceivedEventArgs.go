//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2WebMessageReceivedEventArgsVtbl struct {
	IUnknownVtbl
	GetSource                ComProc
	GetWebMessageAsJson      ComProc
	TryGetWebMessageAsString ComProc
}

type ICoreWebView2WebMessageReceivedEventArgs struct {
	Vtbl *ICoreWebView2WebMessageReceivedEventArgsVtbl
}

func (i *ICoreWebView2WebMessageReceivedEventArgs) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2WebMessageReceivedEventArgs) GetSource() (*string, error) {
	// Create *uint16 to hold result
	var _source *uint16

	hr, _, err := i.Vtbl.GetSource.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_source)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	// Get result and cleanup
	source := ptr(UTF16PtrToString(_source))
	CoTaskMemFree(unsafe.Pointer(_source))
	return source, err
}

func (i *ICoreWebView2WebMessageReceivedEventArgs) GetWebMessageAsJson() (*string, error) {
	// Create *uint16 to hold result
	var _webMessageAsJson *uint16

	hr, _, err := i.Vtbl.GetWebMessageAsJson.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_webMessageAsJson)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	// Get result and cleanup
	webMessageAsJson := ptr(UTF16PtrToString(_webMessageAsJson))
	CoTaskMemFree(unsafe.Pointer(_webMessageAsJson))
	return webMessageAsJson, err
}

func (i *ICoreWebView2WebMessageReceivedEventArgs) TryGetWebMessageAsString() (*string, error) {
	// Create *uint16 to hold result
	var _webMessageAsString *uint16

	hr, _, err := i.Vtbl.TryGetWebMessageAsString.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_webMessageAsString)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	// Get result and cleanup
	webMessageAsString := ptr(UTF16PtrToString(_webMessageAsString))
	CoTaskMemFree(unsafe.Pointer(_webMessageAsString))
	return webMessageAsString, err
}
