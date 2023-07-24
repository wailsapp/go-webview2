//go:build windows

package webview2

type _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2DevToolsProtocolEventReceivedEventHandler struct {
	vtbl *_ICoreWebView2DevToolsProtocolEventReceivedEventHandlerVtbl
	impl _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerImpl
}

func (i *ICoreWebView2DevToolsProtocolEventReceivedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerIUnknownQueryInterface(this *ICoreWebView2DevToolsProtocolEventReceivedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerIUnknownAddRef(this *ICoreWebView2DevToolsProtocolEventReceivedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerIUnknownRelease(this *ICoreWebView2DevToolsProtocolEventReceivedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerInvoke(this *ICoreWebView2DevToolsProtocolEventReceivedEventHandler, sender *ICoreWebView2, args *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) uintptr {
	return this.impl.DevToolsProtocolEventReceived(sender, args)
}

type _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerImpl interface {
	_IUnknownImpl
	DevToolsProtocolEventReceived(sender *ICoreWebView2, args *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) uintptr
}

var _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerFn = _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2DevToolsProtocolEventReceivedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2DevToolsProtocolEventReceivedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2DevToolsProtocolEventReceivedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2DevToolsProtocolEventReceivedEventHandlerInvoke),
}

func NewICoreWebView2DevToolsProtocolEventReceivedEventHandler(impl _ICoreWebView2DevToolsProtocolEventReceivedEventHandlerImpl) *ICoreWebView2DevToolsProtocolEventReceivedEventHandler {
	return &ICoreWebView2DevToolsProtocolEventReceivedEventHandler{
		vtbl: &_ICoreWebView2DevToolsProtocolEventReceivedEventHandlerFn,
		impl: impl,
	}
}
