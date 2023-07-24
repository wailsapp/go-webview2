//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2_9Vtbl struct {
	_IUnknownVtbl
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
	vtbl *_ICoreWebView2_9Vtbl
}

func (i *ICoreWebView2_9) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2_9) AddIsDefaultDownloadDialogOpenChanged(handler *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddIsDefaultDownloadDialogOpenChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2_9) RemoveIsDefaultDownloadDialogOpenChanged(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveIsDefaultDownloadDialogOpenChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_9) GetIsDefaultDownloadDialogOpen() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetIsDefaultDownloadDialogOpen.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2_9) OpenDefaultDownloadDialog() error {
	var err error

	_, _, err = i.vtbl.OpenDefaultDownloadDialog.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_9) CloseDefaultDownloadDialog() error {
	var err error

	_, _, err = i.vtbl.CloseDefaultDownloadDialog.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_9) GetDefaultDownloadDialogCornerAlignment() (*COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT, error) {
	var err error

	var value *COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT

	_, _, err = i.vtbl.GetDefaultDownloadDialogCornerAlignment.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2_9) PutDefaultDownloadDialogCornerAlignment(value COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT) error {
	var err error

	_, _, err = i.vtbl.PutDefaultDownloadDialogCornerAlignment.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_9) GetDefaultDownloadDialogMargin() (*POINT, error) {
	var err error

	var value *POINT

	_, _, err = i.vtbl.GetDefaultDownloadDialogMargin.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2_9) PutDefaultDownloadDialogMargin(value POINT) error {
	var err error

	_, _, err = i.vtbl.PutDefaultDownloadDialogMargin.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
