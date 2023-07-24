//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2HttpHeadersCollectionIteratorVtbl struct {
	_IUnknownVtbl
	GetCurrentHeader    ComProc
	GetHasCurrentHeader ComProc
	MoveNext            ComProc
}

type ICoreWebView2HttpHeadersCollectionIterator struct {
	vtbl *_ICoreWebView2HttpHeadersCollectionIteratorVtbl
}

func (i *ICoreWebView2HttpHeadersCollectionIterator) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2HttpHeadersCollectionIterator) GetCurrentHeader() (string, string, error) {
	var err error
	// Create *uint16 to hold result
	var _name *uint16 // Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetCurrentHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", "", err
	} // Get result and cleanup
	name := windows.UTF16PtrToString(_name)
	windows.CoTaskMemFree(unsafe.Pointer(_name)) // Get result and cleanup
	value := windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return name, value, nil
}

func (i *ICoreWebView2HttpHeadersCollectionIterator) GetHasCurrentHeader() (bool, error) {
	var err error

	var hasCurrent bool

	_, _, err = i.vtbl.GetHasCurrentHeader.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&hasCurrent)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return hasCurrent, nil
}

func (i *ICoreWebView2HttpHeadersCollectionIterator) MoveNext() (bool, error) {
	var err error

	var hasNext bool

	_, _, err = i.vtbl.MoveNext.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&hasNext)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return hasNext, nil
}
