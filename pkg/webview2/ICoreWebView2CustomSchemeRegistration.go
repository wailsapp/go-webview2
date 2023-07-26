//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2CustomSchemeRegistrationVtbl struct {
	IUnknownVtbl
	GetSchemeName            ComProc
	GetTreatAsSecure         ComProc
	PutTreatAsSecure         ComProc
	GetAllowedOrigins        ComProc
	SetAllowedOrigins        ComProc
	GetHasAuthorityComponent ComProc
	PutHasAuthorityComponent ComProc
}

type ICoreWebView2CustomSchemeRegistration struct {
	Vtbl *ICoreWebView2CustomSchemeRegistrationVtbl
}

func (i *ICoreWebView2CustomSchemeRegistration) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2CustomSchemeRegistration) GetSchemeName() (*string, error) {
	// Create *uint16 to hold result
	var _schemeName *uint16

	hr, _, err := i.Vtbl.GetSchemeName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_schemeName)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	// Get result and cleanup
	schemeName := ptr(UTF16PtrToString(_schemeName))
	CoTaskMemFree(unsafe.Pointer(_schemeName))
	return schemeName, err
}

func (i *ICoreWebView2CustomSchemeRegistration) GetTreatAsSecure() (*bool, error) {
	// Create int32 to hold bool result
	var _treatAsSecure int32

	hr, _, err := i.Vtbl.GetTreatAsSecure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_treatAsSecure)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	// Get result and cleanup
	treatAsSecure := ptr(_treatAsSecure != 0)
	return treatAsSecure, err
}

func (i *ICoreWebView2CustomSchemeRegistration) PutTreatAsSecure(value bool) error {

	hr, _, err := i.Vtbl.PutTreatAsSecure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2CustomSchemeRegistration) GetAllowedOrigins() (*uint32, *string, error) {

	var allowedOriginsCount *uint32 // Create *uint16 to hold result
	var _allowedOrigins *uint16

	hr, _, err := i.Vtbl.GetAllowedOrigins.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&allowedOriginsCount)),
		uintptr(unsafe.Pointer(_allowedOrigins)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, nil, syscall.Errno(hr)
	}
	// Get result and cleanup
	allowedOrigins := ptr(UTF16PtrToString(_allowedOrigins))
	CoTaskMemFree(unsafe.Pointer(_allowedOrigins))
	return allowedOriginsCount, allowedOrigins, err
}

func (i *ICoreWebView2CustomSchemeRegistration) SetAllowedOrigins(allowedOriginsCount uint32, allowedOrigins *string) error {

	hr, _, err := i.Vtbl.SetAllowedOrigins.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&allowedOriginsCount)),
		uintptr(unsafe.Pointer(allowedOrigins)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2CustomSchemeRegistration) GetHasAuthorityComponent() (*bool, error) {
	// Create int32 to hold bool result
	var _hasAuthorityComponent int32

	hr, _, err := i.Vtbl.GetHasAuthorityComponent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_hasAuthorityComponent)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	// Get result and cleanup
	hasAuthorityComponent := ptr(_hasAuthorityComponent != 0)
	return hasAuthorityComponent, err
}

func (i *ICoreWebView2CustomSchemeRegistration) PutHasAuthorityComponent(hasAuthorityComponent bool) error {

	hr, _, err := i.Vtbl.PutHasAuthorityComponent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&hasAuthorityComponent)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
