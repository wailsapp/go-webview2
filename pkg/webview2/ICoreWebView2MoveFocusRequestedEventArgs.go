//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2MoveFocusRequestedEventArgsVtbl struct {
	_IUnknownVtbl
	GetReason  ComProc
	GetHandled ComProc
	PutHandled ComProc
}

type ICoreWebView2MoveFocusRequestedEventArgs struct {
	vtbl *_ICoreWebView2MoveFocusRequestedEventArgsVtbl
}

func (i *ICoreWebView2MoveFocusRequestedEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2MoveFocusRequestedEventArgs) GetReason() (*COREWEBVIEW2_MOVE_FOCUS_REASON, error) {
	var err error

	var reason *COREWEBVIEW2_MOVE_FOCUS_REASON

	_, _, err = i.vtbl.GetReason.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&reason)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return reason, nil
}

func (i *ICoreWebView2MoveFocusRequestedEventArgs) GetHandled() (bool, error) {
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

func (i *ICoreWebView2MoveFocusRequestedEventArgs) PutHandled(value bool) error {
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
