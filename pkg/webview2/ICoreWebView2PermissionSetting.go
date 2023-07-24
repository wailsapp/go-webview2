//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2PermissionSettingVtbl struct {
	_IUnknownVtbl
	GetPermissionKind   ComProc
	GetPermissionOrigin ComProc
	GetPermissionState  ComProc
}

type ICoreWebView2PermissionSetting struct {
	vtbl *_ICoreWebView2PermissionSettingVtbl
}

func (i *ICoreWebView2PermissionSetting) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2PermissionSetting) GetPermissionKind() (*COREWEBVIEW2_PERMISSION_KIND, error) {
	var err error

	var value *COREWEBVIEW2_PERMISSION_KIND

	_, _, err = i.vtbl.GetPermissionKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2PermissionSetting) GetPermissionOrigin() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetPermissionOrigin.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	value := windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2PermissionSetting) GetPermissionState() (*COREWEBVIEW2_PERMISSION_STATE, error) {
	var err error

	var value *COREWEBVIEW2_PERMISSION_STATE

	_, _, err = i.vtbl.GetPermissionState.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}
