//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2_12Vtbl struct {
	_IUnknownVtbl
	AddStatusBarTextChanged    ComProc
	RemoveStatusBarTextChanged ComProc
	GetStatusBarText           ComProc
}

type ICoreWebView2_12 struct {
	vtbl *_ICoreWebView2_12Vtbl
}

func (i *ICoreWebView2_12) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2_12) AddStatusBarTextChanged(eventHandler *ICoreWebView2StatusBarTextChangedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddStatusBarTextChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2_12) RemoveStatusBarTextChanged(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveStatusBarTextChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_12) GetStatusBarText() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetStatusBarText.Call(
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
