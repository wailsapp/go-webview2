//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2DownloadStartingEventArgsVtbl struct {
	_IUnknownVtbl
	GetDownloadOperation ComProc
	GetCancel            ComProc
	PutCancel            ComProc
	GetResultFilePath    ComProc
	PutResultFilePath    ComProc
	GetHandled           ComProc
	PutHandled           ComProc
	GetDeferral          ComProc
}

type ICoreWebView2DownloadStartingEventArgs struct {
	vtbl *_ICoreWebView2DownloadStartingEventArgsVtbl
}

func (i *ICoreWebView2DownloadStartingEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2DownloadStartingEventArgs) GetDownloadOperation() (*ICoreWebView2DownloadOperation, error) {
	var err error

	var downloadOperation *ICoreWebView2DownloadOperation

	_, _, err = i.vtbl.GetDownloadOperation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&downloadOperation)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return downloadOperation, nil
}

func (i *ICoreWebView2DownloadStartingEventArgs) GetCancel() (bool, error) {
	var err error

	var cancel bool

	_, _, err = i.vtbl.GetCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cancel)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return cancel, nil
}

func (i *ICoreWebView2DownloadStartingEventArgs) PutCancel(cancel bool) error {
	var err error

	_, _, err = i.vtbl.PutCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cancel)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadStartingEventArgs) GetResultFilePath() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _resultFilePath *uint16

	_, _, err = i.vtbl.GetResultFilePath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_resultFilePath)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	resultFilePath := windows.UTF16PtrToString(_resultFilePath)
	windows.CoTaskMemFree(unsafe.Pointer(_resultFilePath))
	return resultFilePath, nil
}

func (i *ICoreWebView2DownloadStartingEventArgs) PutResultFilePath(resultFilePath string) error {
	var err error

	// Convert string 'resultFilePath' to *uint16
	_resultFilePath, err := windows.UTF16PtrFromString(resultFilePath)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PutResultFilePath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_resultFilePath)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadStartingEventArgs) GetHandled() (bool, error) {
	var err error

	var handled bool

	_, _, err = i.vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handled)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return handled, nil
}

func (i *ICoreWebView2DownloadStartingEventArgs) PutHandled(handled bool) error {
	var err error

	_, _, err = i.vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handled)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadStartingEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
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
