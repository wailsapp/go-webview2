//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2Profile4Vtbl struct {
	IUnknownVtbl
	SetPermissionState              ComProc
	GetNonDefaultPermissionSettings ComProc
}

type ICoreWebView2Profile4 struct {
	Vtbl *ICoreWebView2Profile4Vtbl
}

func (i *ICoreWebView2Profile4) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2Profile4() *ICoreWebView2Profile4 {
	var result *ICoreWebView2Profile4

	iidICoreWebView2Profile4 := NewGUID("{8F4ae680-192e-4eC8-833a-21cfadaef628}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2Profile4)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2Profile4) SetPermissionState(permissionKind COREWEBVIEW2_PERMISSION_KIND, origin string, state COREWEBVIEW2_PERMISSION_STATE, completedHandler *ICoreWebView2SetPermissionStateCompletedHandler) error {

	// Convert string 'origin' to *uint16
	_origin, err := UTF16PtrFromString(origin)
	if err != nil {
		return err
	}

	hr, _, err := i.Vtbl.SetPermissionState.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(permissionKind),
		uintptr(unsafe.Pointer(_origin)),
		uintptr(state),
		uintptr(unsafe.Pointer(completedHandler)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2Profile4) GetNonDefaultPermissionSettings(completedHandler *ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler) error {

	hr, _, err := i.Vtbl.GetNonDefaultPermissionSettings.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(completedHandler)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
