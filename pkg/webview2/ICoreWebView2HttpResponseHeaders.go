//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2HttpResponseHeadersVtbl struct {
	IUnknownVtbl
	AppendHeader ComProc
	Contains     ComProc
	GetHeader    ComProc
	GetHeaders   ComProc
	GetIterator  ComProc
}

type ICoreWebView2HttpResponseHeaders struct {
	Vtbl *ICoreWebView2HttpResponseHeadersVtbl
}

func (i *ICoreWebView2HttpResponseHeaders) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2HttpResponseHeaders) AppendHeader(name string, value string) error {

	// Convert string 'name' to *uint16
	_name, err := UTF16PtrFromString(name)
	if err != nil {
		return err
	}
	// Convert string 'value' to *uint16
	_value, err := UTF16PtrFromString(value)
	if err != nil {
		return err
	}

	hr, _, err := i.Vtbl.AppendHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2HttpResponseHeaders) Contains(name string) (bool, error) {

	// Convert string 'name' to *uint16
	_name, err := UTF16PtrFromString(name)
	if err != nil {
		return false, err
	} // Create int32 to hold bool result
	var _contains int32

	hr, _, err := i.Vtbl.Contains.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(&_contains)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return false, syscall.Errno(hr)
	}
	// Get result and cleanup
	contains := _contains != 0
	return contains, err
}

func (i *ICoreWebView2HttpResponseHeaders) GetHeader(name string) (string, error) {

	// Convert string 'name' to *uint16
	_name, err := UTF16PtrFromString(name)
	if err != nil {
		return "", err
	} // Create *uint16 to hold result
	var _value *uint16

	hr, _, err := i.Vtbl.GetHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
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

func (i *ICoreWebView2HttpResponseHeaders) GetHeaders(name string) (*ICoreWebView2HttpHeadersCollectionIterator, error) {

	// Convert string 'name' to *uint16
	_name, err := UTF16PtrFromString(name)
	if err != nil {
		return nil, err
	}
	var iterator *ICoreWebView2HttpHeadersCollectionIterator

	hr, _, err := i.Vtbl.GetHeaders.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(&iterator)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return iterator, err
}

func (i *ICoreWebView2HttpResponseHeaders) GetIterator() (*ICoreWebView2HttpHeadersCollectionIterator, error) {

	var iterator *ICoreWebView2HttpHeadersCollectionIterator

	hr, _, err := i.Vtbl.GetIterator.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&iterator)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return iterator, err
}
