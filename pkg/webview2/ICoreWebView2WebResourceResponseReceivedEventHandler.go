//go:build windows

package webview2

type _ICoreWebView2WebResourceResponseReceivedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2WebResourceResponseReceivedEventHandler struct {
	vtbl *_ICoreWebView2WebResourceResponseReceivedEventHandlerVtbl
	impl _ICoreWebView2WebResourceResponseReceivedEventHandlerImpl
}

func (i *ICoreWebView2WebResourceResponseReceivedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2WebResourceResponseReceivedEventHandlerIUnknownQueryInterface(this *ICoreWebView2WebResourceResponseReceivedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2WebResourceResponseReceivedEventHandlerIUnknownAddRef(this *ICoreWebView2WebResourceResponseReceivedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2WebResourceResponseReceivedEventHandlerIUnknownRelease(this *ICoreWebView2WebResourceResponseReceivedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2WebResourceResponseReceivedEventHandlerInvoke(this *ICoreWebView2WebResourceResponseReceivedEventHandler, sender *ICoreWebView2, args *ICoreWebView2WebResourceResponseReceivedEventArgs) uintptr {
	return this.impl.WebResourceResponseReceived(sender, args)
}

type _ICoreWebView2WebResourceResponseReceivedEventHandlerImpl interface {
	_IUnknownImpl
	WebResourceResponseReceived(sender *ICoreWebView2, args *ICoreWebView2WebResourceResponseReceivedEventArgs) uintptr
}

var _ICoreWebView2WebResourceResponseReceivedEventHandlerFn = _ICoreWebView2WebResourceResponseReceivedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2WebResourceResponseReceivedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2WebResourceResponseReceivedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2WebResourceResponseReceivedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2WebResourceResponseReceivedEventHandlerInvoke),
}

func NewICoreWebView2WebResourceResponseReceivedEventHandler(impl _ICoreWebView2WebResourceResponseReceivedEventHandlerImpl) *ICoreWebView2WebResourceResponseReceivedEventHandler {
	return &ICoreWebView2WebResourceResponseReceivedEventHandler{
		vtbl: &_ICoreWebView2WebResourceResponseReceivedEventHandlerFn,
		impl: impl,
	}
}
