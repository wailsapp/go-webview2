//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2ProcessExtendedInfoCollectionVtbl struct {
	IUnknownVtbl
	GetCount        ComProc
	GetValueAtIndex ComProc
}

type ICoreWebView2ProcessExtendedInfoCollection struct {
	Vtbl *ICoreWebView2ProcessExtendedInfoCollectionVtbl
}

func (i *ICoreWebView2ProcessExtendedInfoCollection) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2ProcessExtendedInfoCollection) GetCount() (uint, error) {

	var count uint

	hr, _, err := i.Vtbl.GetCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&count)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return count, err
}

func (i *ICoreWebView2ProcessExtendedInfoCollection) GetValueAtIndex(index uint32) (*ICoreWebView2ProcessExtendedInfo, error) {

	var processInfo *ICoreWebView2ProcessExtendedInfo

	hr, _, err := i.Vtbl.GetValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&index)),
		uintptr(unsafe.Pointer(&processInfo)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return processInfo, err
}
