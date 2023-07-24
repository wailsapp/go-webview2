//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2ProcessInfoCollectionVtbl struct {
	_IUnknownVtbl
	GetCount        ComProc
	GetValueAtIndex ComProc
}

type ICoreWebView2ProcessInfoCollection struct {
	vtbl *_ICoreWebView2ProcessInfoCollectionVtbl
}

func (i *ICoreWebView2ProcessInfoCollection) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2ProcessInfoCollection) GetCount() (uint, error) {
	var err error

	var count uint

	_, _, err = i.vtbl.GetCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&count)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return count, nil
}

func (i *ICoreWebView2ProcessInfoCollection) GetValueAtIndex(index uint32) (*ICoreWebView2ProcessInfo, error) {
	var err error

	var processInfo *ICoreWebView2ProcessInfo

	_, _, err = i.vtbl.GetValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&index)),
		uintptr(unsafe.Pointer(&processInfo)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return processInfo, nil
}
