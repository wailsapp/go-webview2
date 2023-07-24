//go:build windows

package webview2

type _ICoreWebView2FaviconChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2FaviconChangedEventHandler struct {
	vtbl *_ICoreWebView2FaviconChangedEventHandlerVtbl
	impl _ICoreWebView2FaviconChangedEventHandlerImpl
}

func (i *ICoreWebView2FaviconChangedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2FaviconChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FaviconChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FaviconChangedEventHandlerIUnknownAddRef(this *ICoreWebView2FaviconChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FaviconChangedEventHandlerIUnknownRelease(this *ICoreWebView2FaviconChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FaviconChangedEventHandlerInvoke(this *ICoreWebView2FaviconChangedEventHandler, sender *ICoreWebView2, args *_IUnknown) uintptr {
	return this.impl.FaviconChanged(sender, args)
}

type _ICoreWebView2FaviconChangedEventHandlerImpl interface {
	_IUnknownImpl
	FaviconChanged(sender *ICoreWebView2, args *_IUnknown) uintptr
}

var _ICoreWebView2FaviconChangedEventHandlerFn = _ICoreWebView2FaviconChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2FaviconChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2FaviconChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2FaviconChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2FaviconChangedEventHandlerInvoke),
}

func NewICoreWebView2FaviconChangedEventHandler(impl _ICoreWebView2FaviconChangedEventHandlerImpl) *ICoreWebView2FaviconChangedEventHandler {
	return &ICoreWebView2FaviconChangedEventHandler{
		vtbl: &_ICoreWebView2FaviconChangedEventHandlerFn,
		impl: impl,
	}
}
