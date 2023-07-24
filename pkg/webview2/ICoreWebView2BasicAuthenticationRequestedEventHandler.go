//go:build windows

package webview2

type _ICoreWebView2BasicAuthenticationRequestedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2BasicAuthenticationRequestedEventHandler struct {
	vtbl *_ICoreWebView2BasicAuthenticationRequestedEventHandlerVtbl
	impl _ICoreWebView2BasicAuthenticationRequestedEventHandlerImpl
}

func (i *ICoreWebView2BasicAuthenticationRequestedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2BasicAuthenticationRequestedEventHandlerIUnknownQueryInterface(this *ICoreWebView2BasicAuthenticationRequestedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2BasicAuthenticationRequestedEventHandlerIUnknownAddRef(this *ICoreWebView2BasicAuthenticationRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2BasicAuthenticationRequestedEventHandlerIUnknownRelease(this *ICoreWebView2BasicAuthenticationRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2BasicAuthenticationRequestedEventHandlerInvoke(this *ICoreWebView2BasicAuthenticationRequestedEventHandler, sender *ICoreWebView2, args *ICoreWebView2BasicAuthenticationRequestedEventArgs) uintptr {
	return this.impl.BasicAuthenticationRequested(sender, args)
}

type _ICoreWebView2BasicAuthenticationRequestedEventHandlerImpl interface {
	_IUnknownImpl
	BasicAuthenticationRequested(sender *ICoreWebView2, args *ICoreWebView2BasicAuthenticationRequestedEventArgs) uintptr
}

var _ICoreWebView2BasicAuthenticationRequestedEventHandlerFn = _ICoreWebView2BasicAuthenticationRequestedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2BasicAuthenticationRequestedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2BasicAuthenticationRequestedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2BasicAuthenticationRequestedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2BasicAuthenticationRequestedEventHandlerInvoke),
}

func NewICoreWebView2BasicAuthenticationRequestedEventHandler(impl _ICoreWebView2BasicAuthenticationRequestedEventHandlerImpl) *ICoreWebView2BasicAuthenticationRequestedEventHandler {
	return &ICoreWebView2BasicAuthenticationRequestedEventHandler{
		vtbl: &_ICoreWebView2BasicAuthenticationRequestedEventHandlerFn,
		impl: impl,
	}
}
