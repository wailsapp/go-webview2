//go:build windows

package webview2

type _ICoreWebView2NewWindowRequestedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2NewWindowRequestedEventHandler struct {
	vtbl *_ICoreWebView2NewWindowRequestedEventHandlerVtbl
	impl _ICoreWebView2NewWindowRequestedEventHandlerImpl
}

func (i *ICoreWebView2NewWindowRequestedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2NewWindowRequestedEventHandlerIUnknownQueryInterface(this *ICoreWebView2NewWindowRequestedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2NewWindowRequestedEventHandlerIUnknownAddRef(this *ICoreWebView2NewWindowRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2NewWindowRequestedEventHandlerIUnknownRelease(this *ICoreWebView2NewWindowRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2NewWindowRequestedEventHandlerInvoke(this *ICoreWebView2NewWindowRequestedEventHandler, sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr {
	return this.impl.NewWindowRequested(sender, args)
}

type _ICoreWebView2NewWindowRequestedEventHandlerImpl interface {
	_IUnknownImpl
	NewWindowRequested(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr
}

var _ICoreWebView2NewWindowRequestedEventHandlerFn = _ICoreWebView2NewWindowRequestedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2NewWindowRequestedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2NewWindowRequestedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2NewWindowRequestedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2NewWindowRequestedEventHandlerInvoke),
}

func NewICoreWebView2NewWindowRequestedEventHandler(impl _ICoreWebView2NewWindowRequestedEventHandlerImpl) *ICoreWebView2NewWindowRequestedEventHandler {
	return &ICoreWebView2NewWindowRequestedEventHandler{
		vtbl: &_ICoreWebView2NewWindowRequestedEventHandlerFn,
		impl: impl,
	}
}
