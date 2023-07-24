//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2BrowserProcessExitedEventArgsVtbl struct {
	_IUnknownVtbl
	GetBrowserProcessExitKind ComProc
	GetBrowserProcessId       ComProc
}

type ICoreWebView2BrowserProcessExitedEventArgs struct {
	vtbl *_ICoreWebView2BrowserProcessExitedEventArgsVtbl
}

func (i *ICoreWebView2BrowserProcessExitedEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2BrowserProcessExitedEventArgs) GetBrowserProcessExitKind() (*COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND, error) {
	var err error

	var browserProcessExitKind *COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND

	_, _, err = i.vtbl.GetBrowserProcessExitKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&browserProcessExitKind)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return browserProcessExitKind, nil
}

func (i *ICoreWebView2BrowserProcessExitedEventArgs) GetBrowserProcessId() (*uint32, error) {
	var err error

	var value *uint32

	_, _, err = i.vtbl.GetBrowserProcessId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}
