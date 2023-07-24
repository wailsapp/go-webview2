//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2_8Vtbl struct {
	_IUnknownVtbl
	AddIsMutedChanged                   ComProc
	RemoveIsMutedChanged                ComProc
	GetIsMuted                          ComProc
	PutIsMuted                          ComProc
	AddIsDocumentPlayingAudioChanged    ComProc
	RemoveIsDocumentPlayingAudioChanged ComProc
	GetIsDocumentPlayingAudio           ComProc
}

type ICoreWebView2_8 struct {
	vtbl *_ICoreWebView2_8Vtbl
}

func (i *ICoreWebView2_8) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2_8) AddIsMutedChanged(eventHandler *ICoreWebView2IsMutedChangedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddIsMutedChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2_8) RemoveIsMutedChanged(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveIsMutedChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_8) GetIsMuted() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetIsMuted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2_8) PutIsMuted(value bool) error {
	var err error

	_, _, err = i.vtbl.PutIsMuted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_8) AddIsDocumentPlayingAudioChanged(eventHandler *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddIsDocumentPlayingAudioChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2_8) RemoveIsDocumentPlayingAudioChanged(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveIsDocumentPlayingAudioChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2_8) GetIsDocumentPlayingAudio() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetIsDocumentPlayingAudio.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}
