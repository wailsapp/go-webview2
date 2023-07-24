//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2NavigationStartingEventArgs2Vtbl struct {
	_IUnknownVtbl
	GetAdditionalAllowedFrameAncestors ComProc
	PutAdditionalAllowedFrameAncestors ComProc
}

type ICoreWebView2NavigationStartingEventArgs2 struct {
	vtbl *_ICoreWebView2NavigationStartingEventArgs2Vtbl
}

func (i *ICoreWebView2NavigationStartingEventArgs2) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2NavigationStartingEventArgs2) GetAdditionalAllowedFrameAncestors() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetAdditionalAllowedFrameAncestors.Call(
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

func (i *ICoreWebView2NavigationStartingEventArgs2) PutAdditionalAllowedFrameAncestors(value string) error {
	var err error

	// Convert string 'value' to *uint16
	_value, err := windows.UTF16PtrFromString(value)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PutAdditionalAllowedFrameAncestors.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
