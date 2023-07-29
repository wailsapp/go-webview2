//go:build windows

package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type _ICoreWebView2FileVtbl struct {
	_IUnknownVtbl
	GetPath ComProc
}

type ICoreWebView2File struct {
	vtbl *_ICoreWebView2FileVtbl
}

func (i *ICoreWebView2File) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2File) Release() error {
	return i.vtbl.CallRelease(unsafe.Pointer(i))
}

func (i *ICoreWebView2File) GetPath() (string, error) {
	var err error
	var _path *uint16
	_, _, err = i.vtbl.GetPath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_path)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	}

	path := windows.UTF16PtrToString(_path)
	windows.CoTaskMemFree(unsafe.Pointer(_path))
	return path, nil
}
