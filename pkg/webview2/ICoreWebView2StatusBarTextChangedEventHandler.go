//go:build windows

package webview2

type _ICoreWebView2StatusBarTextChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2StatusBarTextChangedEventHandler struct {
	vtbl *_ICoreWebView2StatusBarTextChangedEventHandlerVtbl
	impl _ICoreWebView2StatusBarTextChangedEventHandlerImpl
}

func (i *ICoreWebView2StatusBarTextChangedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2StatusBarTextChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2StatusBarTextChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2StatusBarTextChangedEventHandlerIUnknownAddRef(this *ICoreWebView2StatusBarTextChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2StatusBarTextChangedEventHandlerIUnknownRelease(this *ICoreWebView2StatusBarTextChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2StatusBarTextChangedEventHandlerInvoke(this *ICoreWebView2StatusBarTextChangedEventHandler, sender *ICoreWebView2, args *_IUnknown) uintptr {
	return this.impl.StatusBarTextChanged(sender, args)
}

type _ICoreWebView2StatusBarTextChangedEventHandlerImpl interface {
	_IUnknownImpl
	StatusBarTextChanged(sender *ICoreWebView2, args *_IUnknown) uintptr
}

var _ICoreWebView2StatusBarTextChangedEventHandlerFn = _ICoreWebView2StatusBarTextChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2StatusBarTextChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2StatusBarTextChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2StatusBarTextChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2StatusBarTextChangedEventHandlerInvoke),
}

func NewICoreWebView2StatusBarTextChangedEventHandler(impl _ICoreWebView2StatusBarTextChangedEventHandlerImpl) *ICoreWebView2StatusBarTextChangedEventHandler {
	return &ICoreWebView2StatusBarTextChangedEventHandler{
		vtbl: &_ICoreWebView2StatusBarTextChangedEventHandlerFn,
		impl: impl,
	}
}
