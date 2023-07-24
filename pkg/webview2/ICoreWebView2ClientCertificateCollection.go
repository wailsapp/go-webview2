//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2ClientCertificateCollectionVtbl struct {
	_IUnknownVtbl
	GetCount        ComProc
	GetValueAtIndex ComProc
}

type ICoreWebView2ClientCertificateCollection struct {
	vtbl *_ICoreWebView2ClientCertificateCollectionVtbl
}

func (i *ICoreWebView2ClientCertificateCollection) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2ClientCertificateCollection) GetCount() (uint, error) {
	var err error

	var value uint

	_, _, err = i.vtbl.GetCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return value, nil
}

func (i *ICoreWebView2ClientCertificateCollection) GetValueAtIndex(index uint) (*ICoreWebView2ClientCertificate, error) {
	var err error

	var certificate *ICoreWebView2ClientCertificate

	_, _, err = i.vtbl.GetValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&index)),
		uintptr(unsafe.Pointer(&certificate)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return certificate, nil
}
