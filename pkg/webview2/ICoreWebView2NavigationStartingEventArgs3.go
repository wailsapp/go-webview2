//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2NavigationStartingEventArgs3Vtbl struct {
	IUnknownVtbl
	GetNavigationKind ComProc
}

type ICoreWebView2NavigationStartingEventArgs3 struct {
	Vtbl *ICoreWebView2NavigationStartingEventArgs3Vtbl
}

func (i *ICoreWebView2NavigationStartingEventArgs3) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2NavigationStartingEventArgs3() *ICoreWebView2NavigationStartingEventArgs3 {
	var result *ICoreWebView2NavigationStartingEventArgs3

	iidICoreWebView2NavigationStartingEventArgs3 := NewGUID("{DDFFE494-4942-4BD2-AB73-35B8FF40E19F}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2NavigationStartingEventArgs3)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2NavigationStartingEventArgs3) GetNavigationKind() (COREWEBVIEW2_NAVIGATION_KIND, error) {

	var navigation_kind COREWEBVIEW2_NAVIGATION_KIND

	hr, _, err := i.Vtbl.GetNavigationKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&navigation_kind)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return COREWEBVIEW2_NAVIGATION_KIND{}, syscall.Errno(hr)
	}
	return navigation_kind, err
}
