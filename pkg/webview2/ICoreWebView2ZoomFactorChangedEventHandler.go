//go:build windows

package webview2

type _ICoreWebView2ZoomFactorChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2ZoomFactorChangedEventHandler struct {
	vtbl *_ICoreWebView2ZoomFactorChangedEventHandlerVtbl
	impl _ICoreWebView2ZoomFactorChangedEventHandlerImpl
}

func (i *ICoreWebView2ZoomFactorChangedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2ZoomFactorChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2ZoomFactorChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ZoomFactorChangedEventHandlerIUnknownAddRef(this *ICoreWebView2ZoomFactorChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ZoomFactorChangedEventHandlerIUnknownRelease(this *ICoreWebView2ZoomFactorChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ZoomFactorChangedEventHandlerInvoke(this *ICoreWebView2ZoomFactorChangedEventHandler, sender *ICoreWebView2Controller, args *_IUnknown) uintptr {
	return this.impl.ZoomFactorChanged(sender, args)
}

type _ICoreWebView2ZoomFactorChangedEventHandlerImpl interface {
	_IUnknownImpl
	ZoomFactorChanged(sender *ICoreWebView2Controller, args *_IUnknown) uintptr
}

var _ICoreWebView2ZoomFactorChangedEventHandlerFn = _ICoreWebView2ZoomFactorChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2ZoomFactorChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2ZoomFactorChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2ZoomFactorChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2ZoomFactorChangedEventHandlerInvoke),
}

func NewICoreWebView2ZoomFactorChangedEventHandler(impl _ICoreWebView2ZoomFactorChangedEventHandlerImpl) *ICoreWebView2ZoomFactorChangedEventHandler {
	return &ICoreWebView2ZoomFactorChangedEventHandler{
		vtbl: &_ICoreWebView2ZoomFactorChangedEventHandlerFn,
		impl: impl,
	}
}
