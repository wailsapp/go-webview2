//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2ServerCertificateErrorDetectedEventArgsVtbl struct {
	_IUnknownVtbl
	GetErrorStatus       ComProc
	GetRequestUri        ComProc
	GetServerCertificate ComProc
	GetAction            ComProc
	PutAction            ComProc
	GetDeferral          ComProc
}

type ICoreWebView2ServerCertificateErrorDetectedEventArgs struct {
	vtbl *_ICoreWebView2ServerCertificateErrorDetectedEventArgsVtbl
}

func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) GetErrorStatus() (*COREWEBVIEW2_WEB_ERROR_STATUS, error) {
	var err error

	var value *COREWEBVIEW2_WEB_ERROR_STATUS

	_, _, err = i.vtbl.GetErrorStatus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) GetRequestUri() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetRequestUri.Call(
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

func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) GetServerCertificate() (*ICoreWebView2Certificate, error) {
	var err error

	var value *ICoreWebView2Certificate

	_, _, err = i.vtbl.GetServerCertificate.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) GetAction() (*COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION, error) {
	var err error

	var value *COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION

	_, _, err = i.vtbl.GetAction.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) PutAction(value COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION) error {
	var err error

	_, _, err = i.vtbl.PutAction.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2ServerCertificateErrorDetectedEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
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
