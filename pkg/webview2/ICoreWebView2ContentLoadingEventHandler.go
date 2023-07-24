//go:build windows

package webview2

type _ICoreWebView2ContentLoadingEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2ContentLoadingEventHandler struct {
	vtbl *_ICoreWebView2ContentLoadingEventHandlerVtbl
	impl _ICoreWebView2ContentLoadingEventHandlerImpl
}

func (i *ICoreWebView2ContentLoadingEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2ContentLoadingEventHandlerIUnknownQueryInterface(this *ICoreWebView2ContentLoadingEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ContentLoadingEventHandlerIUnknownAddRef(this *ICoreWebView2ContentLoadingEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ContentLoadingEventHandlerIUnknownRelease(this *ICoreWebView2ContentLoadingEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ContentLoadingEventHandlerInvoke(this *ICoreWebView2ContentLoadingEventHandler, sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr {
	return this.impl.ContentLoading(sender, args)
}

type _ICoreWebView2ContentLoadingEventHandlerImpl interface {
	_IUnknownImpl
	ContentLoading(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr
}

var _ICoreWebView2ContentLoadingEventHandlerFn = _ICoreWebView2ContentLoadingEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2ContentLoadingEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2ContentLoadingEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2ContentLoadingEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2ContentLoadingEventHandlerInvoke),
}

func NewICoreWebView2ContentLoadingEventHandler(impl _ICoreWebView2ContentLoadingEventHandlerImpl) *ICoreWebView2ContentLoadingEventHandler {
	return &ICoreWebView2ContentLoadingEventHandler{
		vtbl: &_ICoreWebView2ContentLoadingEventHandlerFn,
		impl: impl,
	}
}
