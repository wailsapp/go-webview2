//go:build windows

package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type _ICoreWebView2ContainsFullScreenElementChangedEventArgsVtbl struct {
	_IUnknownVtbl
	GetContainsFullScreenElement ComProc
}

type ICoreWebView2ContainsFullScreenElementChangedEventArgs struct {
	vtbl *_ICoreWebView2ContainsFullScreenElementChangedEventArgsVtbl
}

func (i *ICoreWebView2ContainsFullScreenElementChangedEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2ContainsFullScreenElementChangedEventArgs) GetContainsFullScreenElement() (bool, error) {
	var err error
	var result bool
	_, _, err = i.vtbl.GetContainsFullScreenElement.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&result)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return result, nil
}
