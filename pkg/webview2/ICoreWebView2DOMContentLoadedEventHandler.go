//go:build windows

package webview2

type _ICoreWebView2DOMContentLoadedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2DOMContentLoadedEventHandler struct {
	vtbl *_ICoreWebView2DOMContentLoadedEventHandlerVtbl
	impl _ICoreWebView2DOMContentLoadedEventHandlerImpl
}

func (i *ICoreWebView2DOMContentLoadedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2DOMContentLoadedEventHandlerIUnknownQueryInterface(this *ICoreWebView2DOMContentLoadedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2DOMContentLoadedEventHandlerIUnknownAddRef(this *ICoreWebView2DOMContentLoadedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2DOMContentLoadedEventHandlerIUnknownRelease(this *ICoreWebView2DOMContentLoadedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2DOMContentLoadedEventHandlerInvoke(this *ICoreWebView2DOMContentLoadedEventHandler, sender *ICoreWebView2, args *ICoreWebView2DOMContentLoadedEventArgs) uintptr {
	return this.impl.DOMContentLoaded(sender, args)
}

type _ICoreWebView2DOMContentLoadedEventHandlerImpl interface {
	_IUnknownImpl
	DOMContentLoaded(sender *ICoreWebView2, args *ICoreWebView2DOMContentLoadedEventArgs) uintptr
}

var _ICoreWebView2DOMContentLoadedEventHandlerFn = _ICoreWebView2DOMContentLoadedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2DOMContentLoadedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2DOMContentLoadedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2DOMContentLoadedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2DOMContentLoadedEventHandlerInvoke),
}

func NewICoreWebView2DOMContentLoadedEventHandler(impl _ICoreWebView2DOMContentLoadedEventHandlerImpl) *ICoreWebView2DOMContentLoadedEventHandler {
	return &ICoreWebView2DOMContentLoadedEventHandler{
		vtbl: &_ICoreWebView2DOMContentLoadedEventHandlerFn,
		impl: impl,
	}
}
