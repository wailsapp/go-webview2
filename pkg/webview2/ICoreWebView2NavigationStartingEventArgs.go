//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2NavigationStartingEventArgsVtbl struct {
	_IUnknownVtbl
	GetUri             ComProc
	GetIsUserInitiated ComProc
	GetIsRedirected    ComProc
	GetRequestHeaders  ComProc
	GetCancel          ComProc
	PutCancel          ComProc
	GetNavigationId    ComProc
}

type ICoreWebView2NavigationStartingEventArgs struct {
	vtbl *_ICoreWebView2NavigationStartingEventArgsVtbl
}

func (i *ICoreWebView2NavigationStartingEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2NavigationStartingEventArgs) GetUri() (string, error) {
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

func (i *ICoreWebView2NavigationStartingEventArgs) GetIsUserInitiated() (bool, error) {
	var err error

	var isUserInitiated bool

	_, _, err = i.vtbl.GetIsUserInitiated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isUserInitiated)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return isUserInitiated, nil
}

func (i *ICoreWebView2NavigationStartingEventArgs) GetIsRedirected() (bool, error) {
	var err error

	var isRedirected bool

	_, _, err = i.vtbl.GetIsRedirected.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isRedirected)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return isRedirected, nil
}

func (i *ICoreWebView2NavigationStartingEventArgs) GetRequestHeaders() (*ICoreWebView2HttpRequestHeaders, error) {
	var err error

	var requestHeaders *ICoreWebView2HttpRequestHeaders

	_, _, err = i.vtbl.GetRequestHeaders.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&requestHeaders)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return requestHeaders, nil
}

func (i *ICoreWebView2NavigationStartingEventArgs) GetCancel() (bool, error) {
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

func (i *ICoreWebView2NavigationStartingEventArgs) PutCancel(cancel bool) error {
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

func (i *ICoreWebView2NavigationStartingEventArgs) GetNavigationId() (*uint64, error) {
	var err error

	var navigationId *uint64

	_, _, err = i.vtbl.GetNavigationId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&navigationId)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return navigationId, nil
}
