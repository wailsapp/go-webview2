//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2_11Vtbl struct {
	_IUnknownVtbl
	CallDevToolsProtocolMethodForSession ComProc
	AddContextMenuRequested              ComProc
	RemoveContextMenuRequested           ComProc
}

type ICoreWebView2_11 struct {
	vtbl *_ICoreWebView2_11Vtbl
}

func (i *ICoreWebView2_11) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2_11) CallDevToolsProtocolMethodForSession(sessionId string, methodName string, parametersAsJson string, handler *ICoreWebView2CallDevToolsProtocolMethodCompletedHandler) error {
	var err error

	// Convert string 'sessionId' to *uint16
	_sessionId, err := windows.UTF16PtrFromString(sessionId)
	if err != nil {
		return err
	}

	// Convert string 'methodName' to *uint16
	_methodName, err := windows.UTF16PtrFromString(methodName)
	if err != nil {
		return err
	}

	// Convert string 'parametersAsJson' to *uint16
	_parametersAsJson, err := windows.UTF16PtrFromString(parametersAsJson)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.CallDevToolsProtocolMethodForSession.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_sessionId)),
		uintptr(unsafe.Pointer(_methodName)),
		uintptr(unsafe.Pointer(_parametersAsJson)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_11) AddContextMenuRequested(eventHandler *ICoreWebView2ContextMenuRequestedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddContextMenuRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2_11) RemoveContextMenuRequested(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveContextMenuRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
