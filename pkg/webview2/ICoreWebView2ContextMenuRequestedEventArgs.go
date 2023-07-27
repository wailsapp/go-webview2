//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2ContextMenuRequestedEventArgsVtbl struct {
	IUnknownVtbl
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
	Vtbl *ICoreWebView2ContextMenuRequestedEventArgsVtbl
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetMenuItems() (*ICoreWebView2ContextMenuItemCollection, error) {

	var value *ICoreWebView2ContextMenuItemCollection

	hr, _, err := i.Vtbl.GetMenuItems.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return value, err
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetContextMenuTarget() (*ICoreWebView2ContextMenuTarget, error) {

	var value *ICoreWebView2ContextMenuTarget

	hr, _, err := i.Vtbl.GetContextMenuTarget.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return value, err
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetLocation() (*POINT, error) {

	var value *POINT

	hr, _, err := i.Vtbl.GetLocation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return value, err
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) PutSelectedCommandId(value int32) error {

	hr, _, err := i.Vtbl.PutSelectedCommandId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetSelectedCommandId() (*int32, error) {

	var value *int32

	hr, _, err := i.Vtbl.GetSelectedCommandId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return value, err
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) PutHandled(value bool) error {

	hr, _, err := i.Vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetHandled() (*bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.Vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	// Get result and cleanup
	value := ptr(_value != 0)
	return value, err
}

func (i *ICoreWebView2ContextMenuRequestedEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {

	var deferral *ICoreWebView2Deferral

	hr, _, err := i.Vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&deferral)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return deferral, err
}
