//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2HttpResponseHeadersVtbl struct {
	_IUnknownVtbl
	AppendHeader ComProc
	Contains     ComProc
	GetHeader    ComProc
	GetHeaders   ComProc
	GetIterator  ComProc
}

type ICoreWebView2HttpResponseHeaders struct {
	vtbl *_ICoreWebView2HttpResponseHeadersVtbl
}

func (i *ICoreWebView2HttpResponseHeaders) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2HttpResponseHeaders) AppendHeader(name string, value string) error {
	var err error

	// Convert string 'name' to *uint16
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return err
	}

	// Convert string 'value' to *uint16
	_value, err := windows.UTF16PtrFromString(value)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.AppendHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2HttpResponseHeaders) Contains(name string) (bool, error) {
	var err error

	// Convert string 'name' to *uint16
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return false, err
	}

	var contains bool

	_, _, err = i.vtbl.Contains.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(&contains)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return contains, nil
}

func (i *ICoreWebView2HttpResponseHeaders) GetHeader(name string) (string, error) {
	var err error

	// Convert string 'name' to *uint16
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return "", err
	}
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	value := windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2HttpResponseHeaders) GetHeaders(name string) (*ICoreWebView2HttpHeadersCollectionIterator, error) {
	var err error

	// Convert string 'name' to *uint16
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return nil, err
	}

	var iterator *ICoreWebView2HttpHeadersCollectionIterator

	_, _, err = i.vtbl.GetHeaders.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(&iterator)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return iterator, nil
}

func (i *ICoreWebView2HttpResponseHeaders) GetIterator() (*ICoreWebView2HttpHeadersCollectionIterator, error) {
	var err error

	var iterator *ICoreWebView2HttpHeadersCollectionIterator

	_, _, err = i.vtbl.GetIterator.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&iterator)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return iterator, nil
}
