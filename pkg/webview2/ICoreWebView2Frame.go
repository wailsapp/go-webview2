//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2FrameVtbl struct {
	_IUnknownVtbl
	GetName                          ComProc
	AddNameChanged                   ComProc
	RemoveNameChanged                ComProc
	AddHostObjectToScriptWithOrigins ComProc
	RemoveHostObjectFromScript       ComProc
	AddDestroyed                     ComProc
	RemoveDestroyed                  ComProc
	IsDestroyed                      ComProc
}

type ICoreWebView2Frame struct {
	vtbl *_ICoreWebView2FrameVtbl
}

func (i *ICoreWebView2Frame) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Frame) GetName() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _name *uint16

	_, _, err = i.vtbl.GetName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	name := windows.UTF16PtrToString(_name)
	windows.CoTaskMemFree(unsafe.Pointer(_name))
	return name, nil
}

func (i *ICoreWebView2Frame) AddNameChanged(eventHandler *ICoreWebView2FrameNameChangedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddNameChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2Frame) RemoveNameChanged(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveNameChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Frame) AddHostObjectToScriptWithOrigins(name string, object *VARIANT, originsCount uint32, origins string) error {
	var err error

	// Convert string 'name' to *uint16
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return err
	}

	// Convert string 'origins' to *uint16
	_origins, err := windows.UTF16PtrFromString(origins)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.AddHostObjectToScriptWithOrigins.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(object)),
		uintptr(unsafe.Pointer(&originsCount)),
		uintptr(unsafe.Pointer(_origins)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Frame) RemoveHostObjectFromScript(name string) error {
	var err error

	// Convert string 'name' to *uint16
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.RemoveHostObjectFromScript.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Frame) AddDestroyed(eventHandler *ICoreWebView2FrameDestroyedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddDestroyed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2Frame) RemoveDestroyed(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveDestroyed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Frame) IsDestroyed() (bool, error) {
	var err error

	var destroyed bool

	_, _, err = i.vtbl.IsDestroyed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&destroyed)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return destroyed, nil
}
