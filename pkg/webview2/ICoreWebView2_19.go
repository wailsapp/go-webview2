//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2_19Vtbl struct {
	_IUnknownVtbl
	GetMemoryUsageTargetLevel ComProc
	PutMemoryUsageTargetLevel ComProc
}

type ICoreWebView2_19 struct {
	vtbl *_ICoreWebView2_19Vtbl
}

func (i *ICoreWebView2_19) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2_19) GetMemoryUsageTargetLevel() (*COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL, error) {
	var err error

	var level *COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL

	_, _, err = i.vtbl.GetMemoryUsageTargetLevel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&level)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return level, nil
}

func (i *ICoreWebView2_19) PutMemoryUsageTargetLevel(level COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL) error {
	var err error

	_, _, err = i.vtbl.PutMemoryUsageTargetLevel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(level),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
