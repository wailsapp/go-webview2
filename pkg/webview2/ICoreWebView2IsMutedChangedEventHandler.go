//go:build windows

package webview2

type _ICoreWebView2IsMutedChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2IsMutedChangedEventHandler struct {
	vtbl *_ICoreWebView2IsMutedChangedEventHandlerVtbl
	impl _ICoreWebView2IsMutedChangedEventHandlerImpl
}

func (i *ICoreWebView2IsMutedChangedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2IsMutedChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2IsMutedChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2IsMutedChangedEventHandlerIUnknownAddRef(this *ICoreWebView2IsMutedChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2IsMutedChangedEventHandlerIUnknownRelease(this *ICoreWebView2IsMutedChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2IsMutedChangedEventHandlerInvoke(this *ICoreWebView2IsMutedChangedEventHandler, sender *ICoreWebView2, args *_IUnknown) uintptr {
	return this.impl.IsMutedChanged(sender, args)
}

type _ICoreWebView2IsMutedChangedEventHandlerImpl interface {
	_IUnknownImpl
	IsMutedChanged(sender *ICoreWebView2, args *_IUnknown) uintptr
}

var _ICoreWebView2IsMutedChangedEventHandlerFn = _ICoreWebView2IsMutedChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2IsMutedChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2IsMutedChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2IsMutedChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2IsMutedChangedEventHandlerInvoke),
}

func NewICoreWebView2IsMutedChangedEventHandler(impl _ICoreWebView2IsMutedChangedEventHandlerImpl) *ICoreWebView2IsMutedChangedEventHandler {
	return &ICoreWebView2IsMutedChangedEventHandler{
		vtbl: &_ICoreWebView2IsMutedChangedEventHandlerFn,
		impl: impl,
	}
}
