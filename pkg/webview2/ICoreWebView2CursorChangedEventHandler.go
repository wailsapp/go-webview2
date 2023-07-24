//go:build windows

package webview2

type _ICoreWebView2CursorChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2CursorChangedEventHandler struct {
	vtbl *_ICoreWebView2CursorChangedEventHandlerVtbl
	impl _ICoreWebView2CursorChangedEventHandlerImpl
}

func (i *ICoreWebView2CursorChangedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2CursorChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2CursorChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2CursorChangedEventHandlerIUnknownAddRef(this *ICoreWebView2CursorChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2CursorChangedEventHandlerIUnknownRelease(this *ICoreWebView2CursorChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2CursorChangedEventHandlerInvoke(this *ICoreWebView2CursorChangedEventHandler, sender *ICoreWebView2CompositionController, args *_IUnknown) uintptr {
	return this.impl.CursorChanged(sender, args)
}

type _ICoreWebView2CursorChangedEventHandlerImpl interface {
	_IUnknownImpl
	CursorChanged(sender *ICoreWebView2CompositionController, args *_IUnknown) uintptr
}

var _ICoreWebView2CursorChangedEventHandlerFn = _ICoreWebView2CursorChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2CursorChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2CursorChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2CursorChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2CursorChangedEventHandlerInvoke),
}

func NewICoreWebView2CursorChangedEventHandler(impl _ICoreWebView2CursorChangedEventHandlerImpl) *ICoreWebView2CursorChangedEventHandler {
	return &ICoreWebView2CursorChangedEventHandler{
		vtbl: &_ICoreWebView2CursorChangedEventHandlerFn,
		impl: impl,
	}
}
