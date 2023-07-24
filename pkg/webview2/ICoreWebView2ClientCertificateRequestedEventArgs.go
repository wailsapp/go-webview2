//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2ClientCertificateRequestedEventArgsVtbl struct {
	_IUnknownVtbl
	GetHost                          ComProc
	GetPort                          ComProc
	GetIsProxy                       ComProc
	GetAllowedCertificateAuthorities ComProc
	GetMutuallyTrustedCertificates   ComProc
	GetSelectedCertificate           ComProc
	PutSelectedCertificate           ComProc
	GetCancel                        ComProc
	PutCancel                        ComProc
	GetHandled                       ComProc
	PutHandled                       ComProc
	GetDeferral                      ComProc
}

type ICoreWebView2ClientCertificateRequestedEventArgs struct {
	vtbl *_ICoreWebView2ClientCertificateRequestedEventArgsVtbl
}

func (i *ICoreWebView2ClientCertificateRequestedEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetHost() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetHost.Call(
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

func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetPort() (int, error) {
	var err error

	var value int

	_, _, err = i.vtbl.GetPort.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return value, nil
}

func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetIsProxy() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetIsProxy.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetAllowedCertificateAuthorities() (*ICoreWebView2StringCollection, error) {
	var err error

	var value *ICoreWebView2StringCollection

	_, _, err = i.vtbl.GetAllowedCertificateAuthorities.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetMutuallyTrustedCertificates() (*ICoreWebView2ClientCertificateCollection, error) {
	var err error

	var value *ICoreWebView2ClientCertificateCollection

	_, _, err = i.vtbl.GetMutuallyTrustedCertificates.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetSelectedCertificate() (*ICoreWebView2ClientCertificate, error) {
	var err error

	var value *ICoreWebView2ClientCertificate

	_, _, err = i.vtbl.GetSelectedCertificate.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ClientCertificateRequestedEventArgs) PutSelectedCertificate(value *ICoreWebView2ClientCertificate) error {
	var err error

	_, _, err = i.vtbl.PutSelectedCertificate.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetCancel() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2ClientCertificateRequestedEventArgs) PutCancel(value bool) error {
	var err error

	_, _, err = i.vtbl.PutCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetHandled() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2ClientCertificateRequestedEventArgs) PutHandled(value bool) error {
	var err error

	_, _, err = i.vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
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
