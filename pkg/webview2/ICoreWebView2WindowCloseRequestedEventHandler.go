//go:build windows

package webview2

type _ICoreWebView2WindowCloseRequestedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2WindowCloseRequestedEventHandler struct {
	vtbl *_ICoreWebView2WindowCloseRequestedEventHandlerVtbl
	impl _ICoreWebView2WindowCloseRequestedEventHandlerImpl
}

func (i *ICoreWebView2WindowCloseRequestedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2WindowCloseRequestedEventHandlerIUnknownQueryInterface(this *ICoreWebView2WindowCloseRequestedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2WindowCloseRequestedEventHandlerIUnknownAddRef(this *ICoreWebView2WindowCloseRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2WindowCloseRequestedEventHandlerIUnknownRelease(this *ICoreWebView2WindowCloseRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2WindowCloseRequestedEventHandlerInvoke(this *ICoreWebView2WindowCloseRequestedEventHandler, sender *ICoreWebView2, args *_IUnknown) uintptr {
	return this.impl.WindowCloseRequested(sender, args)
}

type _ICoreWebView2WindowCloseRequestedEventHandlerImpl interface {
	_IUnknownImpl
	WindowCloseRequested(sender *ICoreWebView2, args *_IUnknown) uintptr
}

var _ICoreWebView2WindowCloseRequestedEventHandlerFn = _ICoreWebView2WindowCloseRequestedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2WindowCloseRequestedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2WindowCloseRequestedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2WindowCloseRequestedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2WindowCloseRequestedEventHandlerInvoke),
}

func NewICoreWebView2WindowCloseRequestedEventHandler(impl _ICoreWebView2WindowCloseRequestedEventHandlerImpl) *ICoreWebView2WindowCloseRequestedEventHandler {
	return &ICoreWebView2WindowCloseRequestedEventHandler{
		vtbl: &_ICoreWebView2WindowCloseRequestedEventHandlerFn,
		impl: impl,
	}
}
