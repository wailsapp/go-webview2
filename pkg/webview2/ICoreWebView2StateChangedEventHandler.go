//go:build windows

package webview2

type _ICoreWebView2StateChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2StateChangedEventHandler struct {
	vtbl *_ICoreWebView2StateChangedEventHandlerVtbl
	impl _ICoreWebView2StateChangedEventHandlerImpl
}

func (i *ICoreWebView2StateChangedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2StateChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2StateChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2StateChangedEventHandlerIUnknownAddRef(this *ICoreWebView2StateChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2StateChangedEventHandlerIUnknownRelease(this *ICoreWebView2StateChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2StateChangedEventHandlerInvoke(this *ICoreWebView2StateChangedEventHandler, sender *ICoreWebView2DownloadOperation, args *_IUnknown) uintptr {
	return this.impl.StateChanged(sender, args)
}

type _ICoreWebView2StateChangedEventHandlerImpl interface {
	_IUnknownImpl
	StateChanged(sender *ICoreWebView2DownloadOperation, args *_IUnknown) uintptr
}

var _ICoreWebView2StateChangedEventHandlerFn = _ICoreWebView2StateChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2StateChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2StateChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2StateChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2StateChangedEventHandlerInvoke),
}

func NewICoreWebView2StateChangedEventHandler(impl _ICoreWebView2StateChangedEventHandlerImpl) *ICoreWebView2StateChangedEventHandler {
	return &ICoreWebView2StateChangedEventHandler{
		vtbl: &_ICoreWebView2StateChangedEventHandlerFn,
		impl: impl,
	}
}
