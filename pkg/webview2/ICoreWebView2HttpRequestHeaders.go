//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2HttpRequestHeadersVtbl struct {
	_IUnknownVtbl
	GetHeader    ComProc
	GetHeaders   ComProc
	Contains     ComProc
	SetHeader    ComProc
	RemoveHeader ComProc
	GetIterator  ComProc
}

type ICoreWebView2HttpRequestHeaders struct {
	vtbl *_ICoreWebView2HttpRequestHeadersVtbl
}

func (i *ICoreWebView2HttpRequestHeaders) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2HttpRequestHeaders) GetHeader(name string) (string, error) {
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

func (i *ICoreWebView2HttpRequestHeaders) GetHeaders(name string) (*ICoreWebView2HttpHeadersCollectionIterator, error) {
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

func (i *ICoreWebView2HttpRequestHeaders) Contains(name string) (bool, error) {
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

func (i *ICoreWebView2HttpRequestHeaders) SetHeader(name string, value string) error {
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

	_, _, err = i.vtbl.SetHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2HttpRequestHeaders) RemoveHeader(name string) error {
	var err error

	// Convert string 'name' to *uint16
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.RemoveHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2HttpRequestHeaders) GetIterator() (*ICoreWebView2HttpHeadersCollectionIterator, error) {
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
