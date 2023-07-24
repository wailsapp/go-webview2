//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2WebMessageReceivedEventArgsVtbl struct {
	_IUnknownVtbl
	GetSource                ComProc
	GetWebMessageAsJson      ComProc
	TryGetWebMessageAsString ComProc
}

type ICoreWebView2WebMessageReceivedEventArgs struct {
	vtbl *_ICoreWebView2WebMessageReceivedEventArgsVtbl
}

func (i *ICoreWebView2WebMessageReceivedEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2WebMessageReceivedEventArgs) GetSource() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _source *uint16

	_, _, err = i.vtbl.GetSource.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_source)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	source := windows.UTF16PtrToString(_source)
	windows.CoTaskMemFree(unsafe.Pointer(_source))
	return source, nil
}

func (i *ICoreWebView2WebMessageReceivedEventArgs) GetWebMessageAsJson() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _webMessageAsJson *uint16

	_, _, err = i.vtbl.GetWebMessageAsJson.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_webMessageAsJson)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	webMessageAsJson := windows.UTF16PtrToString(_webMessageAsJson)
	windows.CoTaskMemFree(unsafe.Pointer(_webMessageAsJson))
	return webMessageAsJson, nil
}

func (i *ICoreWebView2WebMessageReceivedEventArgs) TryGetWebMessageAsString() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _webMessageAsString *uint16

	_, _, err = i.vtbl.TryGetWebMessageAsString.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_webMessageAsString)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	webMessageAsString := windows.UTF16PtrToString(_webMessageAsString)
	windows.CoTaskMemFree(unsafe.Pointer(_webMessageAsString))
	return webMessageAsString, nil
}
