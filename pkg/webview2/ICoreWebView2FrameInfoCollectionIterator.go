//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2FrameInfoCollectionIteratorVtbl struct {
	_IUnknownVtbl
	GetHasCurrent ComProc
	GetCurrent    ComProc
	MoveNext      ComProc
}

type ICoreWebView2FrameInfoCollectionIterator struct {
	vtbl *_ICoreWebView2FrameInfoCollectionIteratorVtbl
}

func (i *ICoreWebView2FrameInfoCollectionIterator) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2FrameInfoCollectionIterator) GetHasCurrent() (bool, error) {
	var err error

	var hasCurrent bool

	_, _, err = i.vtbl.GetHasCurrent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&hasCurrent)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return hasCurrent, nil
}

func (i *ICoreWebView2FrameInfoCollectionIterator) GetCurrent() (*ICoreWebView2FrameInfo, error) {
	var err error

	var frameInfo *ICoreWebView2FrameInfo

	_, _, err = i.vtbl.GetCurrent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&frameInfo)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return frameInfo, nil
}

func (i *ICoreWebView2FrameInfoCollectionIterator) MoveNext() (bool, error) {
	var err error

	var hasNext bool

	_, _, err = i.vtbl.MoveNext.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&hasNext)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return hasNext, nil
}
