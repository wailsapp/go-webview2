//go:build windows

package webview2

type _ICoreWebView2ProcessInfosChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2ProcessInfosChangedEventHandler struct {
	vtbl *_ICoreWebView2ProcessInfosChangedEventHandlerVtbl
	impl _ICoreWebView2ProcessInfosChangedEventHandlerImpl
}

func (i *ICoreWebView2ProcessInfosChangedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2ProcessInfosChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2ProcessInfosChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ProcessInfosChangedEventHandlerIUnknownAddRef(this *ICoreWebView2ProcessInfosChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ProcessInfosChangedEventHandlerIUnknownRelease(this *ICoreWebView2ProcessInfosChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ProcessInfosChangedEventHandlerInvoke(this *ICoreWebView2ProcessInfosChangedEventHandler, sender *ICoreWebView2Environment, args *_IUnknown) uintptr {
	return this.impl.ProcessInfosChanged(sender, args)
}

type _ICoreWebView2ProcessInfosChangedEventHandlerImpl interface {
	_IUnknownImpl
	ProcessInfosChanged(sender *ICoreWebView2Environment, args *_IUnknown) uintptr
}

var _ICoreWebView2ProcessInfosChangedEventHandlerFn = _ICoreWebView2ProcessInfosChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2ProcessInfosChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2ProcessInfosChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2ProcessInfosChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2ProcessInfosChangedEventHandlerInvoke),
}

func NewICoreWebView2ProcessInfosChangedEventHandler(impl _ICoreWebView2ProcessInfosChangedEventHandlerImpl) *ICoreWebView2ProcessInfosChangedEventHandler {
	return &ICoreWebView2ProcessInfosChangedEventHandler{
		vtbl: &_ICoreWebView2ProcessInfosChangedEventHandlerFn,
		impl: impl,
	}
}
