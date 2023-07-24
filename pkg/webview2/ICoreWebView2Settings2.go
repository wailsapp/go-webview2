//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Settings2Vtbl struct {
	_IUnknownVtbl
	GetUserAgent ComProc
	PutUserAgent ComProc
}

type ICoreWebView2Settings2 struct {
	vtbl *_ICoreWebView2Settings2Vtbl
}

func (i *ICoreWebView2Settings2) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Settings2) GetUserAgent() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _userAgent *uint16

	_, _, err = i.vtbl.GetUserAgent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_userAgent)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	userAgent := windows.UTF16PtrToString(_userAgent)
	windows.CoTaskMemFree(unsafe.Pointer(_userAgent))
	return userAgent, nil
}

func (i *ICoreWebView2Settings2) PutUserAgent(userAgent string) error {
	var err error

	// Convert string 'userAgent' to *uint16
	_userAgent, err := windows.UTF16PtrFromString(userAgent)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PutUserAgent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_userAgent)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
