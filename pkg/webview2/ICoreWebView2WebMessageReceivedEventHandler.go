//go:build windows

package webview2

type _ICoreWebView2WebMessageReceivedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2WebMessageReceivedEventHandler struct {
	vtbl *_ICoreWebView2WebMessageReceivedEventHandlerVtbl
	impl _ICoreWebView2WebMessageReceivedEventHandlerImpl
}

func (i *ICoreWebView2WebMessageReceivedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2WebMessageReceivedEventHandlerIUnknownQueryInterface(this *ICoreWebView2WebMessageReceivedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2WebMessageReceivedEventHandlerIUnknownAddRef(this *ICoreWebView2WebMessageReceivedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2WebMessageReceivedEventHandlerIUnknownRelease(this *ICoreWebView2WebMessageReceivedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2WebMessageReceivedEventHandlerInvoke(this *ICoreWebView2WebMessageReceivedEventHandler, sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr {
	return this.impl.WebMessageReceived(sender, args)
}

type _ICoreWebView2WebMessageReceivedEventHandlerImpl interface {
	_IUnknownImpl
	WebMessageReceived(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr
}

var _ICoreWebView2WebMessageReceivedEventHandlerFn = _ICoreWebView2WebMessageReceivedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2WebMessageReceivedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2WebMessageReceivedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2WebMessageReceivedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2WebMessageReceivedEventHandlerInvoke),
}

func NewICoreWebView2WebMessageReceivedEventHandler(impl _ICoreWebView2WebMessageReceivedEventHandlerImpl) *ICoreWebView2WebMessageReceivedEventHandler {
	return &ICoreWebView2WebMessageReceivedEventHandler{
		vtbl: &_ICoreWebView2WebMessageReceivedEventHandlerFn,
		impl: impl,
	}
}
