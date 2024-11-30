//go:build windows

package webview2

import 
	"syscall"
rg/x/sys/windows"
	"syscall"

	"golang.org/x/sys/windows"
)

	AddSaveFileSecurityCheckStarting    ComProc
	IUnknownVtbl
	AddSaveFileSecurityCheckStarting ComProc
	RemoveSaveFileSecurityCheckStarting ComProc
}

type ICoreWebView2_26 struct {
	Vtbl *ICoreWebView2_26Vtbl
}

func (i *ICoreWebView2_26) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter


func (i *ICoreWebView2) GetICoreWebView2_26() *ICoreWebView2_26 {
	var result *ICoreWebView2_26

	iidICoreWebView2_26 := NewGUID("{806268b8-f897-5685-88e5-c45fca0b1a48}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_26)),
		uintptr(unsafe.Pointer(&result)))

	return result


func (i *ICoreWebView2_26) AddSaveFileSecurityCheckStarting(eventHandler *ICoreWebView2SaveFileSecurityCheckStartingEventHandler) (EventRegistrationToken, error) {

	var token EventRegistrationToken

	hr, _, err := i.Vtbl.AddSaveFileSecurityCheckStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return EventRegistrationToken{}, syscall.Errno(hr)
	}
	return token, err
}



	hr, _, err := i.Vtbl.RemoveSaveFileSecurityCheckStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
