//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2_9Vtbl struct {
	IUnknownVtbl
	AddIsDefaultDownloadDialogOpenChanged    ComProc
	RemoveIsDefaultDownloadDialogOpenChanged ComProc
	GetIsDefaultDownloadDialogOpen           ComProc
	OpenDefaultDownloadDialog                ComProc
	CloseDefaultDownloadDialog               ComProc
	GetDefaultDownloadDialogCornerAlignment  ComProc
	PutDefaultDownloadDialogCornerAlignment  ComProc
	GetDefaultDownloadDialogMargin           ComProc
	PutDefaultDownloadDialogMargin           ComProc
}

type ICoreWebView2_9 struct {
	Vtbl *ICoreWebView2_9Vtbl
}

func (i *ICoreWebView2_9) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2_9() *ICoreWebView2_9 {
	var result *ICoreWebView2_9

	iidICoreWebView2_9 := NewGUID("{4d7b2eab-9fdc-468d-b998-a9260b5ed651}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_9)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2_9) AddIsDefaultDownloadDialogOpenChanged(handler *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler) (*EventRegistrationToken, error) {

	var token *EventRegistrationToken

	hr, _, err := i.Vtbl.AddIsDefaultDownloadDialogOpenChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return token, err
}

func (i *ICoreWebView2_9) RemoveIsDefaultDownloadDialogOpenChanged(token EventRegistrationToken) error {

	hr, _, err := i.Vtbl.RemoveIsDefaultDownloadDialogOpenChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2_9) GetIsDefaultDownloadDialogOpen() (*bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.Vtbl.GetIsDefaultDownloadDialogOpen.Call(
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

func (i *ICoreWebView2_9) OpenDefaultDownloadDialog() error {

	hr, _, err := i.Vtbl.OpenDefaultDownloadDialog.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2_9) CloseDefaultDownloadDialog() error {

	hr, _, err := i.Vtbl.CloseDefaultDownloadDialog.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2_9) GetDefaultDownloadDialogCornerAlignment() (*COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT, error) {

	var value *COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT

	hr, _, err := i.Vtbl.GetDefaultDownloadDialogCornerAlignment.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return value, err
}

func (i *ICoreWebView2_9) PutDefaultDownloadDialogCornerAlignment(value COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT) error {

	hr, _, err := i.Vtbl.PutDefaultDownloadDialogCornerAlignment.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2_9) GetDefaultDownloadDialogMargin() (*POINT, error) {

	var value *POINT

	hr, _, err := i.Vtbl.GetDefaultDownloadDialogMargin.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return value, err
}

func (i *ICoreWebView2_9) PutDefaultDownloadDialogMargin(value POINT) error {

	hr, _, err := i.Vtbl.PutDefaultDownloadDialogMargin.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
