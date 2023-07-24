//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2EnvironmentOptionsVtbl struct {
	_IUnknownVtbl
	GetAdditionalBrowserArguments             ComProc
	PutAdditionalBrowserArguments             ComProc
	GetLanguage                               ComProc
	PutLanguage                               ComProc
	GetTargetCompatibleBrowserVersion         ComProc
	PutTargetCompatibleBrowserVersion         ComProc
	GetAllowSingleSignOnUsingOSPrimaryAccount ComProc
	PutAllowSingleSignOnUsingOSPrimaryAccount ComProc
}

type ICoreWebView2EnvironmentOptions struct {
	vtbl *_ICoreWebView2EnvironmentOptionsVtbl
}

func (i *ICoreWebView2EnvironmentOptions) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2EnvironmentOptions) GetAdditionalBrowserArguments() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetAdditionalBrowserArguments.Call(
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

func (i *ICoreWebView2EnvironmentOptions) PutAdditionalBrowserArguments(value string) error {
	var err error

	// Convert string 'value' to *uint16
	_value, err := windows.UTF16PtrFromString(value)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PutAdditionalBrowserArguments.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2EnvironmentOptions) GetLanguage() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetLanguage.Call(
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

func (i *ICoreWebView2EnvironmentOptions) PutLanguage(value string) error {
	var err error

	// Convert string 'value' to *uint16
	_value, err := windows.UTF16PtrFromString(value)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PutLanguage.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2EnvironmentOptions) GetTargetCompatibleBrowserVersion() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetTargetCompatibleBrowserVersion.Call(
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

func (i *ICoreWebView2EnvironmentOptions) PutTargetCompatibleBrowserVersion(value string) error {
	var err error

	// Convert string 'value' to *uint16
	_value, err := windows.UTF16PtrFromString(value)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PutTargetCompatibleBrowserVersion.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2EnvironmentOptions) GetAllowSingleSignOnUsingOSPrimaryAccount() (bool, error) {
	var err error

	var allow bool

	_, _, err = i.vtbl.GetAllowSingleSignOnUsingOSPrimaryAccount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&allow)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return allow, nil
}

func (i *ICoreWebView2EnvironmentOptions) PutAllowSingleSignOnUsingOSPrimaryAccount(allow bool) error {
	var err error

	_, _, err = i.vtbl.PutAllowSingleSignOnUsingOSPrimaryAccount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&allow)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
