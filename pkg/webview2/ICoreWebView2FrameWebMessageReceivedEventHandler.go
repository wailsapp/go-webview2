//go:build windows

package webview2

type _ICoreWebView2FrameWebMessageReceivedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2FrameWebMessageReceivedEventHandler struct {
	vtbl *_ICoreWebView2FrameWebMessageReceivedEventHandlerVtbl
	impl _ICoreWebView2FrameWebMessageReceivedEventHandlerImpl
}

func (i *ICoreWebView2FrameWebMessageReceivedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2FrameWebMessageReceivedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FrameWebMessageReceivedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FrameWebMessageReceivedEventHandlerIUnknownAddRef(this *ICoreWebView2FrameWebMessageReceivedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FrameWebMessageReceivedEventHandlerIUnknownRelease(this *ICoreWebView2FrameWebMessageReceivedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FrameWebMessageReceivedEventHandlerInvoke(this *ICoreWebView2FrameWebMessageReceivedEventHandler, sender *ICoreWebView2Frame, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr {
	return this.impl.FrameWebMessageReceived(sender, args)
}

type _ICoreWebView2FrameWebMessageReceivedEventHandlerImpl interface {
	_IUnknownImpl
	FrameWebMessageReceived(sender *ICoreWebView2Frame, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr
}

var _ICoreWebView2FrameWebMessageReceivedEventHandlerFn = _ICoreWebView2FrameWebMessageReceivedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2FrameWebMessageReceivedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2FrameWebMessageReceivedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2FrameWebMessageReceivedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2FrameWebMessageReceivedEventHandlerInvoke),
}

func NewICoreWebView2FrameWebMessageReceivedEventHandler(impl _ICoreWebView2FrameWebMessageReceivedEventHandlerImpl) *ICoreWebView2FrameWebMessageReceivedEventHandler {
	return &ICoreWebView2FrameWebMessageReceivedEventHandler{
		vtbl: &_ICoreWebView2FrameWebMessageReceivedEventHandlerFn,
		impl: impl,
	}
}
