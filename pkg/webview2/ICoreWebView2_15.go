//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2_15Vtbl struct {
	_IUnknownVtbl
	AddFaviconChanged    ComProc
	RemoveFaviconChanged ComProc
	GetFaviconUri        ComProc
	GetFavicon           ComProc
}

type ICoreWebView2_15 struct {
	vtbl *_ICoreWebView2_15Vtbl
}

func (i *ICoreWebView2_15) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2_15) AddFaviconChanged(eventHandler *ICoreWebView2FaviconChangedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddFaviconChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2_15) RemoveFaviconChanged(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveFaviconChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_15) GetFaviconUri() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetFaviconUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	value := windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2_15) GetFavicon(format COREWEBVIEW2_FAVICON_IMAGE_FORMAT, completedHandler *ICoreWebView2GetFaviconCompletedHandler) error {
	var err error

	_, _, err = i.vtbl.GetFavicon.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(format),
		uintptr(unsafe.Pointer(completedHandler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
