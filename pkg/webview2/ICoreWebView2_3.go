//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2_3Vtbl struct {
	_IUnknownVtbl
	TrySuspend                          ComProc
	Resume                              ComProc
	GetIsSuspended                      ComProc
	SetVirtualHostNameToFolderMapping   ComProc
	ClearVirtualHostNameToFolderMapping ComProc
}

type ICoreWebView2_3 struct {
	vtbl *_ICoreWebView2_3Vtbl
}

func (i *ICoreWebView2_3) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2_3) TrySuspend(handler *ICoreWebView2TrySuspendCompletedHandler) error {
	var err error

	_, _, err = i.vtbl.TrySuspend.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_3) Resume() error {
	var err error

	_, _, err = i.vtbl.Resume.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_3) GetIsSuspended() (bool, error) {
	var err error

	var isSuspended bool

	_, _, err = i.vtbl.GetIsSuspended.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isSuspended)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return isSuspended, nil
}

func (i *ICoreWebView2_3) SetVirtualHostNameToFolderMapping(hostName string, folderPath string, accessKind COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND) error {
	var err error

	// Convert string 'hostName' to *uint16
	_hostName, err := windows.UTF16PtrFromString(hostName)
	if err != nil {
		return err
	}

	// Convert string 'folderPath' to *uint16
	_folderPath, err := windows.UTF16PtrFromString(folderPath)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.SetVirtualHostNameToFolderMapping.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_hostName)),
		uintptr(unsafe.Pointer(_folderPath)),
		uintptr(accessKind),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_3) ClearVirtualHostNameToFolderMapping(hostName string) error {
	var err error

	// Convert string 'hostName' to *uint16
	_hostName, err := windows.UTF16PtrFromString(hostName)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.ClearVirtualHostNameToFolderMapping.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_hostName)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
