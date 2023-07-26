//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2FrameInfoCollectionIteratorVtbl struct {
	IUnknownVtbl
	GetHasCurrent ComProc
	GetCurrent    ComProc
	MoveNext      ComProc
}

type ICoreWebView2FrameInfoCollectionIterator struct {
	Vtbl *ICoreWebView2FrameInfoCollectionIteratorVtbl
}

func (i *ICoreWebView2FrameInfoCollectionIterator) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2FrameInfoCollectionIterator) GetHasCurrent() (*bool, error) {
	// Create int32 to hold bool result
	var _hasCurrent int32

	hr, _, err := i.Vtbl.GetHasCurrent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_hasCurrent)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	// Get result and cleanup
	hasCurrent := ptr(_hasCurrent != 0)
	return hasCurrent, err
}

func (i *ICoreWebView2FrameInfoCollectionIterator) GetCurrent() (*ICoreWebView2FrameInfo, error) {

	var frameInfo *ICoreWebView2FrameInfo

	hr, _, err := i.Vtbl.GetCurrent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&frameInfo)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return frameInfo, err
}

func (i *ICoreWebView2FrameInfoCollectionIterator) MoveNext() (*bool, error) {
	// Create int32 to hold bool result
	var _hasNext int32

	hr, _, err := i.Vtbl.MoveNext.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_hasNext)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	// Get result and cleanup
	hasNext := ptr(_hasNext != 0)
	return hasNext, err
}
