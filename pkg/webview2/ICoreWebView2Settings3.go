//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Settings3Vtbl struct {
	_IUnknownVtbl
	GetAreBrowserAcceleratorKeysEnabled ComProc
	PutAreBrowserAcceleratorKeysEnabled ComProc
}

type ICoreWebView2Settings3 struct {
	vtbl *_ICoreWebView2Settings3Vtbl
}

func (i *ICoreWebView2Settings3) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Settings3) GetAreBrowserAcceleratorKeysEnabled() (bool, error) {
	var err error

	var areBrowserAcceleratorKeysEnabled bool

	_, _, err = i.vtbl.GetAreBrowserAcceleratorKeysEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&areBrowserAcceleratorKeysEnabled)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return areBrowserAcceleratorKeysEnabled, nil
}

func (i *ICoreWebView2Settings3) PutAreBrowserAcceleratorKeysEnabled(areBrowserAcceleratorKeysEnabled bool) error {
	var err error

	_, _, err = i.vtbl.PutAreBrowserAcceleratorKeysEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&areBrowserAcceleratorKeysEnabled)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
