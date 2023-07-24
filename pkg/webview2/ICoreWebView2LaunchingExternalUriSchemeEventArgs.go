//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2LaunchingExternalUriSchemeEventArgsVtbl struct {
	_IUnknownVtbl
	GetUri              ComProc
	GetInitiatingOrigin ComProc
	GetIsUserInitiated  ComProc
	GetCancel           ComProc
	PutCancel           ComProc
	GetDeferral         ComProc
}

type ICoreWebView2LaunchingExternalUriSchemeEventArgs struct {
	vtbl *_ICoreWebView2LaunchingExternalUriSchemeEventArgsVtbl
}

func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) GetUri() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetUri.Call(
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

func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) GetInitiatingOrigin() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetInitiatingOrigin.Call(
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

func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) GetIsUserInitiated() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetIsUserInitiated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) GetCancel() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) PutCancel(value bool) error {
	var err error

	_, _, err = i.vtbl.PutCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2LaunchingExternalUriSchemeEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var err error

	var value *ICoreWebView2Deferral

	_, _, err = i.vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}
