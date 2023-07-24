//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2DownloadOperationVtbl struct {
	_IUnknownVtbl
	AddBytesReceivedChanged       ComProc
	RemoveBytesReceivedChanged    ComProc
	AddEstimatedEndTimeChanged    ComProc
	RemoveEstimatedEndTimeChanged ComProc
	AddStateChanged               ComProc
	RemoveStateChanged            ComProc
	GetUri                        ComProc
	GetContentDisposition         ComProc
	GetMimeType                   ComProc
	GetTotalBytesToReceive        ComProc
	GetBytesReceived              ComProc
	GetEstimatedEndTime           ComProc
	GetResultFilePath             ComProc
	GetState                      ComProc
	GetInterruptReason            ComProc
	Cancel                        ComProc
	Pause                         ComProc
	Resume                        ComProc
	GetCanResume                  ComProc
}

type ICoreWebView2DownloadOperation struct {
	vtbl *_ICoreWebView2DownloadOperationVtbl
}

func (i *ICoreWebView2DownloadOperation) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2DownloadOperation) AddBytesReceivedChanged(eventHandler *ICoreWebView2BytesReceivedChangedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddBytesReceivedChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2DownloadOperation) RemoveBytesReceivedChanged(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveBytesReceivedChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadOperation) AddEstimatedEndTimeChanged(eventHandler *ICoreWebView2EstimatedEndTimeChangedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddEstimatedEndTimeChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2DownloadOperation) RemoveEstimatedEndTimeChanged(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveEstimatedEndTimeChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadOperation) AddStateChanged(eventHandler *ICoreWebView2StateChangedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddStateChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2DownloadOperation) RemoveStateChanged(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveStateChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadOperation) GetUri() (string, error) {
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

func (i *ICoreWebView2DownloadOperation) GetContentDisposition() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _contentDisposition *uint16

	_, _, err = i.vtbl.GetContentDisposition.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_contentDisposition)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	contentDisposition := windows.UTF16PtrToString(_contentDisposition)
	windows.CoTaskMemFree(unsafe.Pointer(_contentDisposition))
	return contentDisposition, nil
}

func (i *ICoreWebView2DownloadOperation) GetMimeType() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _mimeType *uint16

	_, _, err = i.vtbl.GetMimeType.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_mimeType)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	mimeType := windows.UTF16PtrToString(_mimeType)
	windows.CoTaskMemFree(unsafe.Pointer(_mimeType))
	return mimeType, nil
}

func (i *ICoreWebView2DownloadOperation) GetTotalBytesToReceive() (*INT64, error) {
	var err error

	var totalBytesToReceive *INT64

	_, _, err = i.vtbl.GetTotalBytesToReceive.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&totalBytesToReceive)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return totalBytesToReceive, nil
}

func (i *ICoreWebView2DownloadOperation) GetBytesReceived() (*INT64, error) {
	var err error

	var bytesReceived *INT64

	_, _, err = i.vtbl.GetBytesReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&bytesReceived)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return bytesReceived, nil
}

func (i *ICoreWebView2DownloadOperation) GetEstimatedEndTime() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _estimatedEndTime *uint16

	_, _, err = i.vtbl.GetEstimatedEndTime.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_estimatedEndTime)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	estimatedEndTime := windows.UTF16PtrToString(_estimatedEndTime)
	windows.CoTaskMemFree(unsafe.Pointer(_estimatedEndTime))
	return estimatedEndTime, nil
}

func (i *ICoreWebView2DownloadOperation) GetResultFilePath() (string, error) {
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

func (i *ICoreWebView2DownloadOperation) GetState() (*COREWEBVIEW2_DOWNLOAD_STATE, error) {
	var err error

	var downloadState *COREWEBVIEW2_DOWNLOAD_STATE

	_, _, err = i.vtbl.GetState.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&downloadState)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return downloadState, nil
}

func (i *ICoreWebView2DownloadOperation) GetInterruptReason() (*COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON, error) {
	var err error

	var interruptReason *COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON

	_, _, err = i.vtbl.GetInterruptReason.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&interruptReason)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return interruptReason, nil
}

func (i *ICoreWebView2DownloadOperation) Cancel() error {
	var err error

	_, _, err = i.vtbl.Cancel.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadOperation) Pause() error {
	var err error

	_, _, err = i.vtbl.Pause.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadOperation) Resume() error {
	var err error

	_, _, err = i.vtbl.Resume.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2DownloadOperation) GetCanResume() (bool, error) {
	var err error

	var canResume bool

	_, _, err = i.vtbl.GetCanResume.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&canResume)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return canResume, nil
}
