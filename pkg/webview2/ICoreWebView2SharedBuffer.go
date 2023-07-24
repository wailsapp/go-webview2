//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2SharedBufferVtbl struct {
	_IUnknownVtbl
	GetSize              ComProc
	GetBuffer            ComProc
	OpenStream           ComProc
	GetFileMappingHandle ComProc
	Close                ComProc
}

type ICoreWebView2SharedBuffer struct {
	vtbl *_ICoreWebView2SharedBufferVtbl
}

func (i *ICoreWebView2SharedBuffer) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2SharedBuffer) GetSize() (*uint64, error) {
	var err error

	var value *uint64

	_, _, err = i.vtbl.GetSize.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2SharedBuffer) GetBuffer() (*uint8, error) {
	var err error

	var value *uint8

	_, _, err = i.vtbl.GetBuffer.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2SharedBuffer) OpenStream() (*IStream, error) {
	var err error

	var value *IStream

	_, _, err = i.vtbl.OpenStream.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2SharedBuffer) GetFileMappingHandle() (*HANDLE, error) {
	var err error

	var value *HANDLE

	_, _, err = i.vtbl.GetFileMappingHandle.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2SharedBuffer) Close() error {
	var err error

	_, _, err = i.vtbl.Close.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
