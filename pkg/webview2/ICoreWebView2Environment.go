//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2EnvironmentVtbl struct {
	_IUnknownVtbl
	CreateCoreWebView2Controller     ComProc
	CreateWebResourceResponse        ComProc
	GetBrowserVersionString          ComProc
	AddNewBrowserVersionAvailable    ComProc
	RemoveNewBrowserVersionAvailable ComProc
}

type ICoreWebView2Environment struct {
	vtbl *_ICoreWebView2EnvironmentVtbl
}

func (i *ICoreWebView2Environment) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Environment) CreateCoreWebView2Controller(parentWindow HWND, handler *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler) error {
	var err error

	_, _, err = i.vtbl.CreateCoreWebView2Controller.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&parentWindow)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Environment) CreateWebResourceResponse(content *IStream, statusCode int, reasonPhrase string, headers string) (*ICoreWebView2WebResourceResponse, error) {
	var err error

	// Convert string 'reasonPhrase' to *uint16
	_reasonPhrase, err := windows.UTF16PtrFromString(reasonPhrase)
	if err != nil {
		return nil, err
	}

	// Convert string 'headers' to *uint16
	_headers, err := windows.UTF16PtrFromString(headers)
	if err != nil {
		return nil, err
	}

	var response *ICoreWebView2WebResourceResponse

	_, _, err = i.vtbl.CreateWebResourceResponse.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(content)),
		uintptr(statusCode),
		uintptr(unsafe.Pointer(_reasonPhrase)),
		uintptr(unsafe.Pointer(_headers)),
		uintptr(unsafe.Pointer(&response)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return response, nil
}

func (i *ICoreWebView2Environment) GetBrowserVersionString() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _versionInfo *uint16

	_, _, err = i.vtbl.GetBrowserVersionString.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_versionInfo)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	versionInfo := windows.UTF16PtrToString(_versionInfo)
	windows.CoTaskMemFree(unsafe.Pointer(_versionInfo))
	return versionInfo, nil
}

func (i *ICoreWebView2Environment) AddNewBrowserVersionAvailable(eventHandler *ICoreWebView2NewBrowserVersionAvailableEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddNewBrowserVersionAvailable.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2Environment) RemoveNewBrowserVersionAvailable(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveNewBrowserVersionAvailable.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
