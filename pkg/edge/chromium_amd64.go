//go:build windows
// +build windows

package edge

import (
	"unsafe"

	"github.com/AlpineAIO/go-webview2/internal/w32"
)

func (e *Chromium) SetSize(bounds w32.Rect) {
	if e.controller == nil {
		return
	}

	e.controller.vtbl.PutBounds.Call(
		uintptr(unsafe.Pointer(e.controller)),
		uintptr(unsafe.Pointer(&bounds)),
	)
}
