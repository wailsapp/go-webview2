//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2Environment10Vtbl struct {
	IUnknownVtbl
	CreateCoreWebView2ControllerOptions                ComProc
	CreateCoreWebView2ControllerWithOptions            ComProc
	CreateCoreWebView2CompositionControllerWithOptions ComProc
}

type ICoreWebView2Environment10 struct {
	Vtbl *ICoreWebView2Environment10Vtbl
}

func (i *ICoreWebView2Environment10) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2Environment10() *ICoreWebView2Environment10 {
	var result *ICoreWebView2Environment10

	iidICoreWebView2Environment10 := NewGUID("{ee0eb9df-6f12-46ce-b53f-3f47b9c928e0}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2Environment10)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2Environment10) CreateCoreWebView2ControllerOptions() (*ICoreWebView2ControllerOptions, error) {

	var options *ICoreWebView2ControllerOptions

	hr, _, err := i.Vtbl.CreateCoreWebView2ControllerOptions.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&options)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return options, err
}

func (i *ICoreWebView2Environment10) CreateCoreWebView2ControllerWithOptions(parentWindow HWND, options *ICoreWebView2ControllerOptions, handler *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler) error {

	hr, _, err := i.Vtbl.CreateCoreWebView2ControllerWithOptions.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&parentWindow)),
		uintptr(unsafe.Pointer(options)),
		uintptr(unsafe.Pointer(handler)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2Environment10) CreateCoreWebView2CompositionControllerWithOptions(parentWindow HWND, options *ICoreWebView2ControllerOptions, handler *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler) error {

	hr, _, err := i.Vtbl.CreateCoreWebView2CompositionControllerWithOptions.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&parentWindow)),
		uintptr(unsafe.Pointer(options)),
		uintptr(unsafe.Pointer(handler)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
