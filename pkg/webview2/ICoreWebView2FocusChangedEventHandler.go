//go:build windows

package webview2

type _ICoreWebView2FocusChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2FocusChangedEventHandler struct {
	vtbl *_ICoreWebView2FocusChangedEventHandlerVtbl
	impl _ICoreWebView2FocusChangedEventHandlerImpl
}

func (i *ICoreWebView2FocusChangedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2FocusChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FocusChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FocusChangedEventHandlerIUnknownAddRef(this *ICoreWebView2FocusChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FocusChangedEventHandlerIUnknownRelease(this *ICoreWebView2FocusChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FocusChangedEventHandlerInvoke(this *ICoreWebView2FocusChangedEventHandler, sender *ICoreWebView2Controller, args *_IUnknown) uintptr {
	return this.impl.FocusChanged(sender, args)
}

type _ICoreWebView2FocusChangedEventHandlerImpl interface {
	_IUnknownImpl
	FocusChanged(sender *ICoreWebView2Controller, args *_IUnknown) uintptr
}

var _ICoreWebView2FocusChangedEventHandlerFn = _ICoreWebView2FocusChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2FocusChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2FocusChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2FocusChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2FocusChangedEventHandlerInvoke),
}

func NewICoreWebView2FocusChangedEventHandler(impl _ICoreWebView2FocusChangedEventHandlerImpl) *ICoreWebView2FocusChangedEventHandler {
	return &ICoreWebView2FocusChangedEventHandler{
		vtbl: &_ICoreWebView2FocusChangedEventHandlerFn,
		impl: impl,
	}
}
