//go:build windows

package webview2

type _ICoreWebView2PrintToPdfCompletedHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2PrintToPdfCompletedHandler struct {
	vtbl *_ICoreWebView2PrintToPdfCompletedHandlerVtbl
	impl _ICoreWebView2PrintToPdfCompletedHandlerImpl
}

func (i *ICoreWebView2PrintToPdfCompletedHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2PrintToPdfCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2PrintToPdfCompletedHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2PrintToPdfCompletedHandlerIUnknownAddRef(this *ICoreWebView2PrintToPdfCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2PrintToPdfCompletedHandlerIUnknownRelease(this *ICoreWebView2PrintToPdfCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2PrintToPdfCompletedHandlerInvoke(this *ICoreWebView2PrintToPdfCompletedHandler, errorCode uintptr, isSuccessful bool) uintptr {
	return this.impl.PrintToPdfCompleted(errorCode, isSuccessful)
}

type _ICoreWebView2PrintToPdfCompletedHandlerImpl interface {
	_IUnknownImpl
	PrintToPdfCompleted(errorCode uintptr, isSuccessful bool) uintptr
}

var _ICoreWebView2PrintToPdfCompletedHandlerFn = _ICoreWebView2PrintToPdfCompletedHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2PrintToPdfCompletedHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2PrintToPdfCompletedHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2PrintToPdfCompletedHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2PrintToPdfCompletedHandlerInvoke),
}

func NewICoreWebView2PrintToPdfCompletedHandler(impl _ICoreWebView2PrintToPdfCompletedHandlerImpl) *ICoreWebView2PrintToPdfCompletedHandler {
	return &ICoreWebView2PrintToPdfCompletedHandler{
		vtbl: &_ICoreWebView2PrintToPdfCompletedHandlerFn,
		impl: impl,
	}
}
