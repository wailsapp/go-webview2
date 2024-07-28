//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2BrowserExtensionVtbl struct {
	IUnknownVtbl
	GetId        ComProc
	GetName      ComProc
	Remove       ComProc
	GetIsEnabled ComProc
	Enable       ComProc
}

type ICoreWebView2BrowserExtension struct {
	Vtbl *ICoreWebView2BrowserExtensionVtbl
}

func (i *ICoreWebView2BrowserExtension) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2BrowserExtension) GetId() (string, error) {
	// Create *uint16 to hold result
	var _value *uint16

	hr, _, err := i.Vtbl.GetId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return "", syscall.Errno(hr)
	}
	// Get result and cleanup
	value := UTF16PtrToString(_value)
	CoTaskMemFree(unsafe.Pointer(_value))
	return value, err
}

func (i *ICoreWebView2BrowserExtension) GetName() (string, error) {
	// Create *uint16 to hold result
	var _value *uint16

	hr, _, err := i.Vtbl.GetName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return "", syscall.Errno(hr)
	}
	// Get result and cleanup
	value := UTF16PtrToString(_value)
	CoTaskMemFree(unsafe.Pointer(_value))
	return value, err
}

func (i *ICoreWebView2BrowserExtension) Remove(handler *ICoreWebView2BrowserExtensionRemoveCompletedHandler) error {

	hr, _, err := i.Vtbl.Remove.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2BrowserExtension) GetIsEnabled() (bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.Vtbl.GetIsEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return false, syscall.Errno(hr)
	}
	// Get result and cleanup
	value := _value != 0
	return value, err
}

func (i *ICoreWebView2BrowserExtension) Enable(isEnabled bool, handler *ICoreWebView2BrowserExtensionEnableCompletedHandler) error {

	hr, _, err := i.Vtbl.Enable.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isEnabled)),
		uintptr(unsafe.Pointer(handler)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
