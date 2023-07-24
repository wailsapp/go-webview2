//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Profile4Vtbl struct {
	_IUnknownVtbl
	SetPermissionState              ComProc
	GetNonDefaultPermissionSettings ComProc
}

type ICoreWebView2Profile4 struct {
	vtbl *_ICoreWebView2Profile4Vtbl
}

func (i *ICoreWebView2Profile4) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Profile4) SetPermissionState(permissionKind COREWEBVIEW2_PERMISSION_KIND, origin string, state COREWEBVIEW2_PERMISSION_STATE, completedHandler *ICoreWebView2SetPermissionStateCompletedHandler) error {
	var err error

	// Convert string 'origin' to *uint16
	_origin, err := windows.UTF16PtrFromString(origin)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.SetPermissionState.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(permissionKind),
		uintptr(unsafe.Pointer(_origin)),
		uintptr(state),
		uintptr(unsafe.Pointer(completedHandler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Profile4) GetNonDefaultPermissionSettings(completedHandler *ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler) error {
	var err error

	_, _, err = i.vtbl.GetNonDefaultPermissionSettings.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(completedHandler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
