//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Environment10Vtbl struct {
	_IUnknownVtbl
	CreateCoreWebView2ControllerOptions                ComProc
	CreateCoreWebView2ControllerWithOptions            ComProc
	CreateCoreWebView2CompositionControllerWithOptions ComProc
}

type ICoreWebView2Environment10 struct {
	vtbl *_ICoreWebView2Environment10Vtbl
}

func (i *ICoreWebView2Environment10) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Environment10) CreateCoreWebView2ControllerOptions() (*ICoreWebView2ControllerOptions, error) {
	var err error

	var options *ICoreWebView2ControllerOptions

	_, _, err = i.vtbl.CreateCoreWebView2ControllerOptions.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&options)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return options, nil
}

func (i *ICoreWebView2Environment10) CreateCoreWebView2ControllerWithOptions(parentWindow HWND, options *ICoreWebView2ControllerOptions, handler *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler) error {
	var err error

	_, _, err = i.vtbl.CreateCoreWebView2ControllerWithOptions.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&parentWindow)),
		uintptr(unsafe.Pointer(options)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Environment10) CreateCoreWebView2CompositionControllerWithOptions(parentWindow HWND, options *ICoreWebView2ControllerOptions, handler *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler) error {
	var err error

	_, _, err = i.vtbl.CreateCoreWebView2CompositionControllerWithOptions.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&parentWindow)),
		uintptr(unsafe.Pointer(options)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
