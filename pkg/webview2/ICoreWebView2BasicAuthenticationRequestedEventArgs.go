//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2BasicAuthenticationRequestedEventArgsVtbl struct {
	_IUnknownVtbl
	GetUri       ComProc
	GetChallenge ComProc
	GetResponse  ComProc
	GetCancel    ComProc
	PutCancel    ComProc
	GetDeferral  ComProc
}

type ICoreWebView2BasicAuthenticationRequestedEventArgs struct {
	vtbl *_ICoreWebView2BasicAuthenticationRequestedEventArgsVtbl
}

func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetUri() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetUri.Call(
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

func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetChallenge() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _challenge *uint16

	_, _, err = i.vtbl.GetChallenge.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_challenge)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	challenge := windows.UTF16PtrToString(_challenge)
	windows.CoTaskMemFree(unsafe.Pointer(_challenge))
	return challenge, nil
}

func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetResponse() (*ICoreWebView2BasicAuthenticationResponse, error) {
	var err error

	var response *ICoreWebView2BasicAuthenticationResponse

	_, _, err = i.vtbl.GetResponse.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&response)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return response, nil
}

func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetCancel() (bool, error) {
	var err error

	var cancel bool

	_, _, err = i.vtbl.GetCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cancel)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return cancel, nil
}

func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) PutCancel(cancel bool) error {
	var err error

	_, _, err = i.vtbl.PutCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cancel)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var err error

	var deferral *ICoreWebView2Deferral

	_, _, err = i.vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&deferral)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return deferral, nil
}
