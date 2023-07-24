//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2CookieManagerVtbl struct {
	_IUnknownVtbl
	CreateCookie                   ComProc
	CopyCookie                     ComProc
	GetCookies                     ComProc
	AddOrUpdateCookie              ComProc
	DeleteCookie                   ComProc
	DeleteCookies                  ComProc
	DeleteCookiesWithDomainAndPath ComProc
	DeleteAllCookies               ComProc
}

type ICoreWebView2CookieManager struct {
	vtbl *_ICoreWebView2CookieManagerVtbl
}

func (i *ICoreWebView2CookieManager) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2CookieManager) CreateCookie(name string, value string, domain string, path string) (*ICoreWebView2Cookie, error) {
	var err error

	// Convert string 'name' to *uint16
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return nil, err
	}

	// Convert string 'value' to *uint16
	_value, err := windows.UTF16PtrFromString(value)
	if err != nil {
		return nil, err
	}

	// Convert string 'domain' to *uint16
	_domain, err := windows.UTF16PtrFromString(domain)
	if err != nil {
		return nil, err
	}

	// Convert string 'path' to *uint16
	_path, err := windows.UTF16PtrFromString(path)
	if err != nil {
		return nil, err
	}

	var cookie *ICoreWebView2Cookie

	_, _, err = i.vtbl.CreateCookie.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(_value)),
		uintptr(unsafe.Pointer(_domain)),
		uintptr(unsafe.Pointer(_path)),
		uintptr(unsafe.Pointer(&cookie)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return cookie, nil
}

func (i *ICoreWebView2CookieManager) CopyCookie(cookieParam *ICoreWebView2Cookie) (*ICoreWebView2Cookie, error) {
	var err error

	var cookie *ICoreWebView2Cookie

	_, _, err = i.vtbl.CopyCookie.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(cookieParam)),
		uintptr(unsafe.Pointer(&cookie)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return cookie, nil
}

func (i *ICoreWebView2CookieManager) GetCookies(uri string, handler *ICoreWebView2GetCookiesCompletedHandler) error {
	var err error

	// Convert string 'uri' to *uint16
	_uri, err := windows.UTF16PtrFromString(uri)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.GetCookies.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2CookieManager) AddOrUpdateCookie(cookie *ICoreWebView2Cookie) error {
	var err error

	_, _, err = i.vtbl.AddOrUpdateCookie.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(cookie)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2CookieManager) DeleteCookie(cookie *ICoreWebView2Cookie) error {
	var err error

	_, _, err = i.vtbl.DeleteCookie.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(cookie)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2CookieManager) DeleteCookies(name string, uri string) error {
	var err error

	// Convert string 'name' to *uint16
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return err
	}

	// Convert string 'uri' to *uint16
	_uri, err := windows.UTF16PtrFromString(uri)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.DeleteCookies.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(_uri)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2CookieManager) DeleteCookiesWithDomainAndPath(name string, domain string, path string) error {
	var err error

	// Convert string 'name' to *uint16
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return err
	}

	// Convert string 'domain' to *uint16
	_domain, err := windows.UTF16PtrFromString(domain)
	if err != nil {
		return err
	}

	// Convert string 'path' to *uint16
	_path, err := windows.UTF16PtrFromString(path)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.DeleteCookiesWithDomainAndPath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(_domain)),
		uintptr(unsafe.Pointer(_path)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2CookieManager) DeleteAllCookies() error {
	var err error

	_, _, err = i.vtbl.DeleteAllCookies.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
