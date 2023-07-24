//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2ContextMenuRequestedEventArgsVtbl struct {
	_IUnknownVtbl
	GetMenuItems         ComProc
	GetContextMenuTarget ComProc
	GetLocation          ComProc
	PutSelectedCommandId ComProc
	GetSelectedCommandId ComProc
	PutHandled           ComProc
	GetHandled           ComProc
	GetDeferral          ComProc
}

type ICoreWebView2ContextMenuRequestedEventArgs struct {
	vtbl *_ICoreWebView2ContextMenuRequestedEventArgsVtbl
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetMenuItems() (*ICoreWebView2ContextMenuItemCollection, error) {
	var err error

	var value *ICoreWebView2ContextMenuItemCollection

	_, _, err = i.vtbl.GetMenuItems.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetContextMenuTarget() (*ICoreWebView2ContextMenuTarget, error) {
	var err error

	var value *ICoreWebView2ContextMenuTarget

	_, _, err = i.vtbl.GetContextMenuTarget.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetLocation() (*POINT, error) {
	var err error

	var value *POINT

	_, _, err = i.vtbl.GetLocation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) PutSelectedCommandId(value INT32) error {
	var err error

	_, _, err = i.vtbl.PutSelectedCommandId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetSelectedCommandId() (*INT32, error) {
	var err error

	var value *INT32

	_, _, err = i.vtbl.GetSelectedCommandId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) PutHandled(value bool) error {
	var err error

	_, _, err = i.vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetHandled() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
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
