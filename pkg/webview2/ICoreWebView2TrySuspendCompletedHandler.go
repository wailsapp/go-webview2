//go:build windows

package webview2

type _ICoreWebView2TrySuspendCompletedHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2TrySuspendCompletedHandler struct {
	vtbl *_ICoreWebView2TrySuspendCompletedHandlerVtbl
	impl _ICoreWebView2TrySuspendCompletedHandlerImpl
}

func (i *ICoreWebView2TrySuspendCompletedHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2TrySuspendCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2TrySuspendCompletedHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2TrySuspendCompletedHandlerIUnknownAddRef(this *ICoreWebView2TrySuspendCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2TrySuspendCompletedHandlerIUnknownRelease(this *ICoreWebView2TrySuspendCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2TrySuspendCompletedHandlerInvoke(this *ICoreWebView2TrySuspendCompletedHandler, errorCode uintptr, isSuccessful bool) uintptr {
	return this.impl.TrySuspendCompleted(errorCode, isSuccessful)
}

type _ICoreWebView2TrySuspendCompletedHandlerImpl interface {
	_IUnknownImpl
	TrySuspendCompleted(errorCode uintptr, isSuccessful bool) uintptr
}

var _ICoreWebView2TrySuspendCompletedHandlerFn = _ICoreWebView2TrySuspendCompletedHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2TrySuspendCompletedHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2TrySuspendCompletedHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2TrySuspendCompletedHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2TrySuspendCompletedHandlerInvoke),
}

func NewICoreWebView2TrySuspendCompletedHandler(impl _ICoreWebView2TrySuspendCompletedHandlerImpl) *ICoreWebView2TrySuspendCompletedHandler {
	return &ICoreWebView2TrySuspendCompletedHandler{
		vtbl: &_ICoreWebView2TrySuspendCompletedHandlerFn,
		impl: impl,
	}
}
