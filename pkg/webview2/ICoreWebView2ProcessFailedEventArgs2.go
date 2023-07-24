//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2ProcessFailedEventArgs2Vtbl struct {
	_IUnknownVtbl
	GetReason                     ComProc
	GetExitCode                   ComProc
	GetProcessDescription         ComProc
	GetFrameInfosForFailedProcess ComProc
}

type ICoreWebView2ProcessFailedEventArgs2 struct {
	vtbl *_ICoreWebView2ProcessFailedEventArgs2Vtbl
}

func (i *ICoreWebView2ProcessFailedEventArgs2) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2ProcessFailedEventArgs2) GetReason() (*COREWEBVIEW2_PROCESS_FAILED_REASON, error) {
	var err error

	var reason *COREWEBVIEW2_PROCESS_FAILED_REASON

	_, _, err = i.vtbl.GetReason.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&reason)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return reason, nil
}

func (i *ICoreWebView2ProcessFailedEventArgs2) GetExitCode() (int, error) {
	var err error

	var exitCode int

	_, _, err = i.vtbl.GetExitCode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(exitCode),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return exitCode, nil
}

func (i *ICoreWebView2ProcessFailedEventArgs2) GetProcessDescription() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _processDescription *uint16

	_, _, err = i.vtbl.GetProcessDescription.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_processDescription)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	processDescription := windows.UTF16PtrToString(_processDescription)
	windows.CoTaskMemFree(unsafe.Pointer(_processDescription))
	return processDescription, nil
}

func (i *ICoreWebView2ProcessFailedEventArgs2) GetFrameInfosForFailedProcess() (*ICoreWebView2FrameInfoCollection, error) {
	var err error

	var frames *ICoreWebView2FrameInfoCollection

	_, _, err = i.vtbl.GetFrameInfosForFailedProcess.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&frames)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return frames, nil
}
