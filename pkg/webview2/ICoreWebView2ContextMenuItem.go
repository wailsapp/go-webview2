//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2ContextMenuItemVtbl struct {
	_IUnknownVtbl
	GetName                   ComProc
	GetLabel                  ComProc
	GetCommandId              ComProc
	GetShortcutKeyDescription ComProc
	GetIcon                   ComProc
	GetKind                   ComProc
	PutIsEnabled              ComProc
	GetIsEnabled              ComProc
	PutIsChecked              ComProc
	GetIsChecked              ComProc
	GetChildren               ComProc
	AddCustomItemSelected     ComProc
	RemoveCustomItemSelected  ComProc
}

type ICoreWebView2ContextMenuItem struct {
	vtbl *_ICoreWebView2ContextMenuItemVtbl
}

func (i *ICoreWebView2ContextMenuItem) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2ContextMenuItem) GetName() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	value := windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2ContextMenuItem) GetLabel() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetLabel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	value := windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2ContextMenuItem) GetCommandId() (*INT32, error) {
	var err error

	var value *INT32

	_, _, err = i.vtbl.GetCommandId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuItem) GetShortcutKeyDescription() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetShortcutKeyDescription.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	value := windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2ContextMenuItem) GetIcon() (*IStream, error) {
	var err error

	var value *IStream

	_, _, err = i.vtbl.GetIcon.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuItem) GetKind() (*COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND, error) {
	var err error

	var value *COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND

	_, _, err = i.vtbl.GetKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuItem) PutIsEnabled(value bool) error {
	var err error

	_, _, err = i.vtbl.PutIsEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2ContextMenuItem) GetIsEnabled() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetIsEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuItem) PutIsChecked(value bool) error {
	var err error

	_, _, err = i.vtbl.PutIsChecked.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2ContextMenuItem) GetIsChecked() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetIsChecked.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuItem) GetChildren() (*ICoreWebView2ContextMenuItemCollection, error) {
	var err error

	var value *ICoreWebView2ContextMenuItemCollection

	_, _, err = i.vtbl.GetChildren.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2ContextMenuItem) AddCustomItemSelected(eventHandler *ICoreWebView2CustomItemSelectedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddCustomItemSelected.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2ContextMenuItem) RemoveCustomItemSelected(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveCustomItemSelected.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
