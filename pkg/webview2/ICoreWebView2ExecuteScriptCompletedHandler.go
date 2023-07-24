//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
)

type _ICoreWebView2ExecuteScriptCompletedHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2ExecuteScriptCompletedHandler struct {
	vtbl *_ICoreWebView2ExecuteScriptCompletedHandlerVtbl
	impl _ICoreWebView2ExecuteScriptCompletedHandlerImpl
}

func (i *ICoreWebView2ExecuteScriptCompletedHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2ExecuteScriptCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2ExecuteScriptCompletedHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ExecuteScriptCompletedHandlerIUnknownAddRef(this *ICoreWebView2ExecuteScriptCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ExecuteScriptCompletedHandlerIUnknownRelease(this *ICoreWebView2ExecuteScriptCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ExecuteScriptCompletedHandlerInvoke(this *ICoreWebView2ExecuteScriptCompletedHandler, errorCode uintptr, resultObjectAsJson string) uintptr {
	return this.impl.ExecuteScriptCompleted(errorCode, resultObjectAsJson)
}

type _ICoreWebView2ExecuteScriptCompletedHandlerImpl interface {
	_IUnknownImpl
	ExecuteScriptCompleted(errorCode uintptr, resultObjectAsJson string) uintptr
}

var _ICoreWebView2ExecuteScriptCompletedHandlerFn = _ICoreWebView2ExecuteScriptCompletedHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2ExecuteScriptCompletedHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2ExecuteScriptCompletedHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2ExecuteScriptCompletedHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2ExecuteScriptCompletedHandlerInvoke),
}

func NewICoreWebView2ExecuteScriptCompletedHandler(impl _ICoreWebView2ExecuteScriptCompletedHandlerImpl) *ICoreWebView2ExecuteScriptCompletedHandler {
	return &ICoreWebView2ExecuteScriptCompletedHandler{
		vtbl: &_ICoreWebView2ExecuteScriptCompletedHandlerFn,
		impl: impl,
	}
}
