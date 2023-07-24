//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2CompositionControllerVtbl struct {
	_IUnknownVtbl
	GetRootVisualTarget ComProc
	PutRootVisualTarget ComProc
	SendMouseInput      ComProc
	SendPointerInput    ComProc
	GetCursor           ComProc
	GetSystemCursorId   ComProc
	AddCursorChanged    ComProc
	RemoveCursorChanged ComProc
}

type ICoreWebView2CompositionController struct {
	vtbl *_ICoreWebView2CompositionControllerVtbl
}

func (i *ICoreWebView2CompositionController) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2CompositionController) GetRootVisualTarget() (*_IUnknown, error) {
	var err error

	var target *_IUnknown

	_, _, err = i.vtbl.GetRootVisualTarget.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&target)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return target, nil
}

func (i *ICoreWebView2CompositionController) PutRootVisualTarget(target *_IUnknown) error {
	var err error

	_, _, err = i.vtbl.PutRootVisualTarget.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(target)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2CompositionController) SendMouseInput(eventKind COREWEBVIEW2_MOUSE_EVENT_KIND, virtualKeys COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS, mouseData uint32, point POINT) error {
	var err error

	_, _, err = i.vtbl.SendMouseInput.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(eventKind),
		uintptr(virtualKeys),
		uintptr(unsafe.Pointer(&mouseData)),
		uintptr(unsafe.Pointer(&point)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2CompositionController) SendPointerInput(eventKind COREWEBVIEW2_POINTER_EVENT_KIND, pointerInfo *ICoreWebView2PointerInfo) error {
	var err error

	_, _, err = i.vtbl.SendPointerInput.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(eventKind),
		uintptr(unsafe.Pointer(pointerInfo)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2CompositionController) GetCursor() (*HCURSOR, error) {
	var err error

	var cursor *HCURSOR

	_, _, err = i.vtbl.GetCursor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cursor)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return cursor, nil
}

func (i *ICoreWebView2CompositionController) GetSystemCursorId() (*uint32, error) {
	var err error

	var systemCursorId *uint32

	_, _, err = i.vtbl.GetSystemCursorId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&systemCursorId)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return systemCursorId, nil
}

func (i *ICoreWebView2CompositionController) AddCursorChanged(eventHandler *ICoreWebView2CursorChangedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddCursorChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2CompositionController) RemoveCursorChanged(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveCursorChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
