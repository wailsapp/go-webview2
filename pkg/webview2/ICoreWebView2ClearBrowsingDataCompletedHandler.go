//go:build windows

package webview2

type _ICoreWebView2ClearBrowsingDataCompletedHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2ClearBrowsingDataCompletedHandler struct {
	vtbl *_ICoreWebView2ClearBrowsingDataCompletedHandlerVtbl
	impl _ICoreWebView2ClearBrowsingDataCompletedHandlerImpl
}

func (i *ICoreWebView2ClearBrowsingDataCompletedHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2ClearBrowsingDataCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2ClearBrowsingDataCompletedHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ClearBrowsingDataCompletedHandlerIUnknownAddRef(this *ICoreWebView2ClearBrowsingDataCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ClearBrowsingDataCompletedHandlerIUnknownRelease(this *ICoreWebView2ClearBrowsingDataCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ClearBrowsingDataCompletedHandlerInvoke(this *ICoreWebView2ClearBrowsingDataCompletedHandler, errorCode uintptr) uintptr {
	return this.impl.ClearBrowsingDataCompleted(errorCode)
}

type _ICoreWebView2ClearBrowsingDataCompletedHandlerImpl interface {
	_IUnknownImpl
	ClearBrowsingDataCompleted(errorCode uintptr) uintptr
}

var _ICoreWebView2ClearBrowsingDataCompletedHandlerFn = _ICoreWebView2ClearBrowsingDataCompletedHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2ClearBrowsingDataCompletedHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2ClearBrowsingDataCompletedHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2ClearBrowsingDataCompletedHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2ClearBrowsingDataCompletedHandlerInvoke),
}

func NewICoreWebView2ClearBrowsingDataCompletedHandler(impl _ICoreWebView2ClearBrowsingDataCompletedHandlerImpl) *ICoreWebView2ClearBrowsingDataCompletedHandler {
	return &ICoreWebView2ClearBrowsingDataCompletedHandler{
		vtbl: &_ICoreWebView2ClearBrowsingDataCompletedHandlerFn,
		impl: impl,
	}
}
