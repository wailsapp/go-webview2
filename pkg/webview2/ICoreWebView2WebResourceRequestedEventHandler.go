//go:build windows

package webview2

type _ICoreWebView2WebResourceRequestedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2WebResourceRequestedEventHandler struct {
	vtbl *_ICoreWebView2WebResourceRequestedEventHandlerVtbl
	impl _ICoreWebView2WebResourceRequestedEventHandlerImpl
}

func (i *ICoreWebView2WebResourceRequestedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2WebResourceRequestedEventHandlerIUnknownQueryInterface(this *ICoreWebView2WebResourceRequestedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2WebResourceRequestedEventHandlerIUnknownAddRef(this *ICoreWebView2WebResourceRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2WebResourceRequestedEventHandlerIUnknownRelease(this *ICoreWebView2WebResourceRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2WebResourceRequestedEventHandlerInvoke(this *ICoreWebView2WebResourceRequestedEventHandler, sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr {
	return this.impl.WebResourceRequested(sender, args)
}

type _ICoreWebView2WebResourceRequestedEventHandlerImpl interface {
	_IUnknownImpl
	WebResourceRequested(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr
}

var _ICoreWebView2WebResourceRequestedEventHandlerFn = _ICoreWebView2WebResourceRequestedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2WebResourceRequestedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2WebResourceRequestedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2WebResourceRequestedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2WebResourceRequestedEventHandlerInvoke),
}

func NewICoreWebView2WebResourceRequestedEventHandler(impl _ICoreWebView2WebResourceRequestedEventHandlerImpl) *ICoreWebView2WebResourceRequestedEventHandler {
	return &ICoreWebView2WebResourceRequestedEventHandler{
		vtbl: &_ICoreWebView2WebResourceRequestedEventHandlerFn,
		impl: impl,
	}
}
