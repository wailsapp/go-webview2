//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2CustomSchemeRegistrationVtbl struct {
	_IUnknownVtbl
	GetSchemeName            ComProc
	GetTreatAsSecure         ComProc
	PutTreatAsSecure         ComProc
	GetAllowedOrigins        ComProc
	SetAllowedOrigins        ComProc
	GetHasAuthorityComponent ComProc
	PutHasAuthorityComponent ComProc
}

type ICoreWebView2CustomSchemeRegistration struct {
	vtbl *_ICoreWebView2CustomSchemeRegistrationVtbl
}

func (i *ICoreWebView2CustomSchemeRegistration) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2CustomSchemeRegistration) GetSchemeName() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _schemeName *uint16

	_, _, err = i.vtbl.GetSchemeName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_schemeName)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	schemeName := windows.UTF16PtrToString(_schemeName)
	windows.CoTaskMemFree(unsafe.Pointer(_schemeName))
	return schemeName, nil
}

func (i *ICoreWebView2CustomSchemeRegistration) GetTreatAsSecure() (bool, error) {
	var err error

	var treatAsSecure bool

	_, _, err = i.vtbl.GetTreatAsSecure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&treatAsSecure)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return treatAsSecure, nil
}

func (i *ICoreWebView2CustomSchemeRegistration) PutTreatAsSecure(value bool) error {
	var err error

	_, _, err = i.vtbl.PutTreatAsSecure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2CustomSchemeRegistration) GetAllowedOrigins() (*uint32, string, error) {
	var err error

	var allowedOriginsCount *uint32 // Create *uint16 to hold result
	var _allowedOrigins *uint16

	_, _, err = i.vtbl.GetAllowedOrigins.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&allowedOriginsCount)),
		uintptr(unsafe.Pointer(_allowedOrigins)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, "", err
	} // Get result and cleanup
	allowedOrigins := windows.UTF16PtrToString(_allowedOrigins)
	windows.CoTaskMemFree(unsafe.Pointer(_allowedOrigins))
	return allowedOriginsCount, allowedOrigins, nil
}

func (i *ICoreWebView2CustomSchemeRegistration) SetAllowedOrigins(allowedOriginsCount uint32, allowedOrigins string) error {
	var err error

	// Convert string 'allowedOrigins' to *uint16
	_allowedOrigins, err := windows.UTF16PtrFromString(allowedOrigins)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.SetAllowedOrigins.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&allowedOriginsCount)),
		uintptr(unsafe.Pointer(_allowedOrigins)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2CustomSchemeRegistration) GetHasAuthorityComponent() (bool, error) {
	var err error

	var hasAuthorityComponent bool

	_, _, err = i.vtbl.GetHasAuthorityComponent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&hasAuthorityComponent)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return hasAuthorityComponent, nil
}

func (i *ICoreWebView2CustomSchemeRegistration) PutHasAuthorityComponent(hasAuthorityComponent bool) error {
	var err error

	_, _, err = i.vtbl.PutHasAuthorityComponent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&hasAuthorityComponent)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
