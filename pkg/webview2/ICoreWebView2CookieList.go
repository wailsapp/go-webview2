//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2CookieListVtbl struct {
	_IUnknownVtbl
	GetCount        ComProc
	GetValueAtIndex ComProc
}

type ICoreWebView2CookieList struct {
	vtbl *_ICoreWebView2CookieListVtbl
}

func (i *ICoreWebView2CookieList) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2CookieList) GetCount() (uint, error) {
	var err error

	var count uint

	_, _, err = i.vtbl.GetCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&count)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return count, nil
}

func (i *ICoreWebView2CookieList) GetValueAtIndex(index uint) (*ICoreWebView2Cookie, error) {
	var err error

	var cookie *ICoreWebView2Cookie

	_, _, err = i.vtbl.GetValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&index)),
		uintptr(unsafe.Pointer(&cookie)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return cookie, nil
}
