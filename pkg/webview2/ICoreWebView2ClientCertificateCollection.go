//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2ClientCertificateCollectionVtbl struct {
	IUnknownVtbl
	GetCount        ComProc
	GetValueAtIndex ComProc
}

type ICoreWebView2ClientCertificateCollection struct {
	Vtbl *ICoreWebView2ClientCertificateCollectionVtbl
}

func (i *ICoreWebView2ClientCertificateCollection) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2ClientCertificateCollection) GetCount() (uint, error) {

	var value uint

	hr, _, err := i.Vtbl.GetCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return value, err
}

func (i *ICoreWebView2ClientCertificateCollection) GetValueAtIndex(index uint) (*ICoreWebView2ClientCertificate, error) {

	var certificate *ICoreWebView2ClientCertificate

	hr, _, err := i.Vtbl.GetValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&index)),
		uintptr(unsafe.Pointer(&certificate)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return certificate, err
}
