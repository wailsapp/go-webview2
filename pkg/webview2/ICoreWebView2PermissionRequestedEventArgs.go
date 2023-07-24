//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2PermissionRequestedEventArgsVtbl struct {
	_IUnknownVtbl
	GetUri             ComProc
	GetPermissionKind  ComProc
	GetIsUserInitiated ComProc
	GetState           ComProc
	PutState           ComProc
	GetDeferral        ComProc
}

type ICoreWebView2PermissionRequestedEventArgs struct {
	vtbl *_ICoreWebView2PermissionRequestedEventArgsVtbl
}

func (i *ICoreWebView2PermissionRequestedEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2PermissionRequestedEventArgs) GetUri() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _uri *uint16

	_, _, err = i.vtbl.GetUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	uri := windows.UTF16PtrToString(_uri)
	windows.CoTaskMemFree(unsafe.Pointer(_uri))
	return uri, nil
}

func (i *ICoreWebView2PermissionRequestedEventArgs) GetPermissionKind() (*COREWEBVIEW2_PERMISSION_KIND, error) {
	var err error

	var permissionKind *COREWEBVIEW2_PERMISSION_KIND

	_, _, err = i.vtbl.GetPermissionKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&permissionKind)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return permissionKind, nil
}

func (i *ICoreWebView2PermissionRequestedEventArgs) GetIsUserInitiated() (bool, error) {
	var err error

	var isUserInitiated bool

	_, _, err = i.vtbl.GetIsUserInitiated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isUserInitiated)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return isUserInitiated, nil
}

func (i *ICoreWebView2PermissionRequestedEventArgs) GetState() (*COREWEBVIEW2_PERMISSION_STATE, error) {
	var err error

	var state *COREWEBVIEW2_PERMISSION_STATE

	_, _, err = i.vtbl.GetState.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&state)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return state, nil
}

func (i *ICoreWebView2PermissionRequestedEventArgs) PutState(state COREWEBVIEW2_PERMISSION_STATE) error {
	var err error

	_, _, err = i.vtbl.PutState.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(state),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PermissionRequestedEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var err error

	var deferral *ICoreWebView2Deferral

	_, _, err = i.vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&deferral)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return deferral, nil
}
