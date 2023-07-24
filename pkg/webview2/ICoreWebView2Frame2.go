//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Frame2Vtbl struct {
	_IUnknownVtbl
	AddNavigationStarting     ComProc
	RemoveNavigationStarting  ComProc
	AddContentLoading         ComProc
	RemoveContentLoading      ComProc
	AddNavigationCompleted    ComProc
	RemoveNavigationCompleted ComProc
	AddDOMContentLoaded       ComProc
	RemoveDOMContentLoaded    ComProc
	ExecuteScript             ComProc
	PostWebMessageAsJson      ComProc
	PostWebMessageAsString    ComProc
	AddWebMessageReceived     ComProc
	RemoveWebMessageReceived  ComProc
}

type ICoreWebView2Frame2 struct {
	vtbl *_ICoreWebView2Frame2Vtbl
}

func (i *ICoreWebView2Frame2) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Frame2) AddNavigationStarting(eventHandler *ICoreWebView2FrameNavigationStartingEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddNavigationStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2Frame2) RemoveNavigationStarting(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveNavigationStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Frame2) AddContentLoading(eventHandler *ICoreWebView2FrameContentLoadingEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddContentLoading.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2Frame2) RemoveContentLoading(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveContentLoading.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Frame2) AddNavigationCompleted(eventHandler *ICoreWebView2FrameNavigationCompletedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddNavigationCompleted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2Frame2) RemoveNavigationCompleted(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveNavigationCompleted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Frame2) AddDOMContentLoaded(eventHandler *ICoreWebView2FrameDOMContentLoadedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddDOMContentLoaded.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2Frame2) RemoveDOMContentLoaded(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveDOMContentLoaded.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Frame2) ExecuteScript(javaScript string, handler *ICoreWebView2ExecuteScriptCompletedHandler) error {
	var err error

	// Convert string 'javaScript' to *uint16
	_javaScript, err := windows.UTF16PtrFromString(javaScript)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.ExecuteScript.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_javaScript)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Frame2) PostWebMessageAsJson(webMessageAsJson string) error {
	var err error

	// Convert string 'webMessageAsJson' to *uint16
	_webMessageAsJson, err := windows.UTF16PtrFromString(webMessageAsJson)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PostWebMessageAsJson.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_webMessageAsJson)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Frame2) PostWebMessageAsString(webMessageAsString string) error {
	var err error

	// Convert string 'webMessageAsString' to *uint16
	_webMessageAsString, err := windows.UTF16PtrFromString(webMessageAsString)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PostWebMessageAsString.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_webMessageAsString)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Frame2) AddWebMessageReceived(handler *ICoreWebView2FrameWebMessageReceivedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddWebMessageReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2Frame2) RemoveWebMessageReceived(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveWebMessageReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
