//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2BasicAuthenticationResponseVtbl struct {
	_IUnknownVtbl
	GetUserName ComProc
	PutUserName ComProc
	GetPassword ComProc
	PutPassword ComProc
}

type ICoreWebView2BasicAuthenticationResponse struct {
	vtbl *_ICoreWebView2BasicAuthenticationResponseVtbl
}

func (i *ICoreWebView2BasicAuthenticationResponse) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2BasicAuthenticationResponse) GetUserName() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _userName *uint16

	_, _, err = i.vtbl.GetUserName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_userName)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	userName := windows.UTF16PtrToString(_userName)
	windows.CoTaskMemFree(unsafe.Pointer(_userName))
	return userName, nil
}

func (i *ICoreWebView2BasicAuthenticationResponse) PutUserName(userName string) error {
	var err error

	// Convert string 'userName' to *uint16
	_userName, err := windows.UTF16PtrFromString(userName)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PutUserName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_userName)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2BasicAuthenticationResponse) GetPassword() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _password *uint16

	_, _, err = i.vtbl.GetPassword.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_password)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	password := windows.UTF16PtrToString(_password)
	windows.CoTaskMemFree(unsafe.Pointer(_password))
	return password, nil
}

func (i *ICoreWebView2BasicAuthenticationResponse) PutPassword(password string) error {
	var err error

	// Convert string 'password' to *uint16
	_password, err := windows.UTF16PtrFromString(password)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PutPassword.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_password)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
