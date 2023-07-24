//go:build windows

package webview2

type _ICoreWebView2ContextMenuRequestedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2ContextMenuRequestedEventHandler struct {
	vtbl *_ICoreWebView2ContextMenuRequestedEventHandlerVtbl
	impl _ICoreWebView2ContextMenuRequestedEventHandlerImpl
}

func (i *ICoreWebView2ContextMenuRequestedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2ContextMenuRequestedEventHandlerIUnknownQueryInterface(this *ICoreWebView2ContextMenuRequestedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ContextMenuRequestedEventHandlerIUnknownAddRef(this *ICoreWebView2ContextMenuRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ContextMenuRequestedEventHandlerIUnknownRelease(this *ICoreWebView2ContextMenuRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ContextMenuRequestedEventHandlerInvoke(this *ICoreWebView2ContextMenuRequestedEventHandler, sender *ICoreWebView2, args *ICoreWebView2ContextMenuRequestedEventArgs) uintptr {
	return this.impl.ContextMenuRequested(sender, args)
}

type _ICoreWebView2ContextMenuRequestedEventHandlerImpl interface {
	_IUnknownImpl
	ContextMenuRequested(sender *ICoreWebView2, args *ICoreWebView2ContextMenuRequestedEventArgs) uintptr
}

var _ICoreWebView2ContextMenuRequestedEventHandlerFn = _ICoreWebView2ContextMenuRequestedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2ContextMenuRequestedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2ContextMenuRequestedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2ContextMenuRequestedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2ContextMenuRequestedEventHandlerInvoke),
}

func NewICoreWebView2ContextMenuRequestedEventHandler(impl _ICoreWebView2ContextMenuRequestedEventHandlerImpl) *ICoreWebView2ContextMenuRequestedEventHandler {
	return &ICoreWebView2ContextMenuRequestedEventHandler{
		vtbl: &_ICoreWebView2ContextMenuRequestedEventHandlerFn,
		impl: impl,
	}
}
