//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2CookieVtbl struct {
	_IUnknownVtbl
	GetName       ComProc
	GetValue      ComProc
	PutValue      ComProc
	GetDomain     ComProc
	GetPath       ComProc
	GetExpires    ComProc
	PutExpires    ComProc
	GetIsHttpOnly ComProc
	PutIsHttpOnly ComProc
	GetSameSite   ComProc
	PutSameSite   ComProc
	GetIsSecure   ComProc
	PutIsSecure   ComProc
	GetIsSession  ComProc
}

type ICoreWebView2Cookie struct {
	vtbl *_ICoreWebView2CookieVtbl
}

func (i *ICoreWebView2Cookie) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Cookie) GetName() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _name *uint16

	_, _, err = i.vtbl.GetName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	name := windows.UTF16PtrToString(_name)
	windows.CoTaskMemFree(unsafe.Pointer(_name))
	return name, nil
}

func (i *ICoreWebView2Cookie) GetValue() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetValue.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	value := windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2Cookie) PutValue(value string) error {
	var err error

	// Convert string 'value' to *uint16
	_value, err := windows.UTF16PtrFromString(value)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PutValue.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Cookie) GetDomain() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _domain *uint16

	_, _, err = i.vtbl.GetDomain.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_domain)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	domain := windows.UTF16PtrToString(_domain)
	windows.CoTaskMemFree(unsafe.Pointer(_domain))
	return domain, nil
}

func (i *ICoreWebView2Cookie) GetPath() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _path *uint16

	_, _, err = i.vtbl.GetPath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_path)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	path := windows.UTF16PtrToString(_path)
	windows.CoTaskMemFree(unsafe.Pointer(_path))
	return path, nil
}

func (i *ICoreWebView2Cookie) GetExpires() (float64, error) {
	var err error

	var expires float64

	_, _, err = i.vtbl.GetExpires.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&expires)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0.0, err
	}
	return expires, nil
}

func (i *ICoreWebView2Cookie) PutExpires(expires float64) error {
	var err error

	_, _, err = i.vtbl.PutExpires.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&expires)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Cookie) GetIsHttpOnly() (bool, error) {
	var err error

	var isHttpOnly bool

	_, _, err = i.vtbl.GetIsHttpOnly.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isHttpOnly)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return isHttpOnly, nil
}

func (i *ICoreWebView2Cookie) PutIsHttpOnly(isHttpOnly bool) error {
	var err error

	_, _, err = i.vtbl.PutIsHttpOnly.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isHttpOnly)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Cookie) GetSameSite() (*COREWEBVIEW2_COOKIE_SAME_SITE_KIND, error) {
	var err error

	var sameSite *COREWEBVIEW2_COOKIE_SAME_SITE_KIND

	_, _, err = i.vtbl.GetSameSite.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&sameSite)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return sameSite, nil
}

func (i *ICoreWebView2Cookie) PutSameSite(sameSite COREWEBVIEW2_COOKIE_SAME_SITE_KIND) error {
	var err error

	_, _, err = i.vtbl.PutSameSite.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(sameSite),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Cookie) GetIsSecure() (bool, error) {
	var err error

	var isSecure bool

	_, _, err = i.vtbl.GetIsSecure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isSecure)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return isSecure, nil
}

func (i *ICoreWebView2Cookie) PutIsSecure(isSecure bool) error {
	var err error

	_, _, err = i.vtbl.PutIsSecure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isSecure)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Cookie) GetIsSession() (bool, error) {
	var err error

	var isSession bool

	_, _, err = i.vtbl.GetIsSession.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isSession)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return isSession, nil
}
