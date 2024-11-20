package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2CookieList vtable
type iCoreWebView2CookieListVtbl struct {
	_IUnknownVtbl
	GetCount ComProc
	GetItem  ComProc
}

// ICoreWebView2CookieList represents a list of cookies
type ICoreWebView2CookieList struct {
	vtbl *iCoreWebView2CookieListVtbl
}

// GetCount gets the number of cookies in the list
func (i *ICoreWebView2CookieList) GetCount() (uint32, error) {
	var count uint32
	hr, _, _ := i.vtbl.GetCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&count)),
	)
	if hr != 0 {
		return 0, syscall.Errno(hr)
	}
	return count, nil
}

// GetItem gets the cookie at the specified index
func (i *ICoreWebView2CookieList) GetItem(index uint32) (*ICoreWebView2Cookie, error) {
	var cookie *ICoreWebView2Cookie
	hr, _, _ := i.vtbl.GetItem.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(index),
		uintptr(unsafe.Pointer(&cookie)),
	)
	if hr != 0 {
		return nil, syscall.Errno(hr)
	}
	return cookie, nil
}

// Release releases the ICoreWebView2CookieList interface
func (i *ICoreWebView2CookieList) Release() error {
	return i.vtbl.CallRelease(unsafe.Pointer(i))
}
