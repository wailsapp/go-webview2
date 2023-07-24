//go:build windows

package webview2

type _ICoreWebView2NewBrowserVersionAvailableEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2NewBrowserVersionAvailableEventHandler struct {
	vtbl *_ICoreWebView2NewBrowserVersionAvailableEventHandlerVtbl
	impl _ICoreWebView2NewBrowserVersionAvailableEventHandlerImpl
}

func (i *ICoreWebView2NewBrowserVersionAvailableEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2NewBrowserVersionAvailableEventHandlerIUnknownQueryInterface(this *ICoreWebView2NewBrowserVersionAvailableEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2NewBrowserVersionAvailableEventHandlerIUnknownAddRef(this *ICoreWebView2NewBrowserVersionAvailableEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2NewBrowserVersionAvailableEventHandlerIUnknownRelease(this *ICoreWebView2NewBrowserVersionAvailableEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2NewBrowserVersionAvailableEventHandlerInvoke(this *ICoreWebView2NewBrowserVersionAvailableEventHandler, sender *ICoreWebView2Environment, args *_IUnknown) uintptr {
	return this.impl.NewBrowserVersionAvailable(sender, args)
}

type _ICoreWebView2NewBrowserVersionAvailableEventHandlerImpl interface {
	_IUnknownImpl
	NewBrowserVersionAvailable(sender *ICoreWebView2Environment, args *_IUnknown) uintptr
}

var _ICoreWebView2NewBrowserVersionAvailableEventHandlerFn = _ICoreWebView2NewBrowserVersionAvailableEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2NewBrowserVersionAvailableEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2NewBrowserVersionAvailableEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2NewBrowserVersionAvailableEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2NewBrowserVersionAvailableEventHandlerInvoke),
}

func NewICoreWebView2NewBrowserVersionAvailableEventHandler(impl _ICoreWebView2NewBrowserVersionAvailableEventHandlerImpl) *ICoreWebView2NewBrowserVersionAvailableEventHandler {
	return &ICoreWebView2NewBrowserVersionAvailableEventHandler{
		vtbl: &_ICoreWebView2NewBrowserVersionAvailableEventHandlerFn,
		impl: impl,
	}
}
