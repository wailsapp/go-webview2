//go:build windows

package webview2

type _ICoreWebView2CustomItemSelectedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2CustomItemSelectedEventHandler struct {
	vtbl *_ICoreWebView2CustomItemSelectedEventHandlerVtbl
	impl _ICoreWebView2CustomItemSelectedEventHandlerImpl
}

func (i *ICoreWebView2CustomItemSelectedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2CustomItemSelectedEventHandlerIUnknownQueryInterface(this *ICoreWebView2CustomItemSelectedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2CustomItemSelectedEventHandlerIUnknownAddRef(this *ICoreWebView2CustomItemSelectedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2CustomItemSelectedEventHandlerIUnknownRelease(this *ICoreWebView2CustomItemSelectedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2CustomItemSelectedEventHandlerInvoke(this *ICoreWebView2CustomItemSelectedEventHandler, sender *ICoreWebView2ContextMenuItem, args *_IUnknown) uintptr {
	return this.impl.CustomItemSelected(sender, args)
}

type _ICoreWebView2CustomItemSelectedEventHandlerImpl interface {
	_IUnknownImpl
	CustomItemSelected(sender *ICoreWebView2ContextMenuItem, args *_IUnknown) uintptr
}

var _ICoreWebView2CustomItemSelectedEventHandlerFn = _ICoreWebView2CustomItemSelectedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2CustomItemSelectedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2CustomItemSelectedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2CustomItemSelectedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2CustomItemSelectedEventHandlerInvoke),
}

func NewICoreWebView2CustomItemSelectedEventHandler(impl _ICoreWebView2CustomItemSelectedEventHandlerImpl) *ICoreWebView2CustomItemSelectedEventHandler {
	return &ICoreWebView2CustomItemSelectedEventHandler{
		vtbl: &_ICoreWebView2CustomItemSelectedEventHandlerFn,
		impl: impl,
	}
}
