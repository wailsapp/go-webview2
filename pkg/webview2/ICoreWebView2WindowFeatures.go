//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2WindowFeaturesVtbl struct {
	_IUnknownVtbl
	GetHasPosition             ComProc
	GetHasSize                 ComProc
	GetLeft                    ComProc
	GetTop                     ComProc
	GetHeight                  ComProc
	GetWidth                   ComProc
	GetShouldDisplayMenuBar    ComProc
	GetShouldDisplayStatus     ComProc
	GetShouldDisplayToolbar    ComProc
	GetShouldDisplayScrollBars ComProc
}

type ICoreWebView2WindowFeatures struct {
	vtbl *_ICoreWebView2WindowFeaturesVtbl
}

func (i *ICoreWebView2WindowFeatures) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2WindowFeatures) GetHasPosition() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetHasPosition.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2WindowFeatures) GetHasSize() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetHasSize.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2WindowFeatures) GetLeft() (*uint32, error) {
	var err error

	var value *uint32

	_, _, err = i.vtbl.GetLeft.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2WindowFeatures) GetTop() (*uint32, error) {
	var err error

	var value *uint32

	_, _, err = i.vtbl.GetTop.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2WindowFeatures) GetHeight() (*uint32, error) {
	var err error

	var value *uint32

	_, _, err = i.vtbl.GetHeight.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2WindowFeatures) GetWidth() (*uint32, error) {
	var err error

	var value *uint32

	_, _, err = i.vtbl.GetWidth.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2WindowFeatures) GetShouldDisplayMenuBar() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetShouldDisplayMenuBar.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2WindowFeatures) GetShouldDisplayStatus() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetShouldDisplayStatus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2WindowFeatures) GetShouldDisplayToolbar() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetShouldDisplayToolbar.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2WindowFeatures) GetShouldDisplayScrollBars() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetShouldDisplayScrollBars.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}
