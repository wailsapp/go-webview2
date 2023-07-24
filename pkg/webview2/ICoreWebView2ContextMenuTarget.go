//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2ContextMenuTargetVtbl struct {
	_IUnknownVtbl
	GetKind                    ComProc
	GetIsEditable              ComProc
	GetIsRequestedForMainFrame ComProc
	GetPageUri                 ComProc
	GetFrameUri                ComProc
	GetHasLinkUri              ComProc
	GetLinkUri                 ComProc
	GetHasLinkText             ComProc
	GetLinkText                ComProc
	GetHasSourceUri            ComProc
	GetSourceUri               ComProc
	GetHasSelection            ComProc
	GetSelectionText           ComProc
}

type ICoreWebView2ContextMenuTarget struct {
	vtbl *_ICoreWebView2ContextMenuTargetVtbl
}

func (i *ICoreWebView2ContextMenuTarget) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2ContextMenuTarget) GetKind() (*COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND, error) {
	var err error

	var value *COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND

	_, _, err = i.vtbl.GetKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetIsEditable() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetIsEditable.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetIsRequestedForMainFrame() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetIsRequestedForMainFrame.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetPageUri() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetPageUri.Call(
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

func (i *ICoreWebView2ContextMenuTarget) GetFrameUri() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetFrameUri.Call(
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

func (i *ICoreWebView2ContextMenuTarget) GetHasLinkUri() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetHasLinkUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetLinkUri() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetLinkUri.Call(
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

func (i *ICoreWebView2ContextMenuTarget) GetHasLinkText() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetHasLinkText.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetLinkText() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetLinkText.Call(
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

func (i *ICoreWebView2ContextMenuTarget) GetHasSourceUri() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetHasSourceUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetSourceUri() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetSourceUri.Call(
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

func (i *ICoreWebView2ContextMenuTarget) GetHasSelection() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetHasSelection.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetSelectionText() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetSelectionText.Call(
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
