//go:build windows

package webview2

type _ICoreWebView2FrameDestroyedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2FrameDestroyedEventHandler struct {
	vtbl *_ICoreWebView2FrameDestroyedEventHandlerVtbl
	impl _ICoreWebView2FrameDestroyedEventHandlerImpl
}

func (i *ICoreWebView2FrameDestroyedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2FrameDestroyedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FrameDestroyedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FrameDestroyedEventHandlerIUnknownAddRef(this *ICoreWebView2FrameDestroyedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FrameDestroyedEventHandlerIUnknownRelease(this *ICoreWebView2FrameDestroyedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FrameDestroyedEventHandlerInvoke(this *ICoreWebView2FrameDestroyedEventHandler, sender *ICoreWebView2Frame, args *_IUnknown) uintptr {
	return this.impl.FrameDestroyed(sender, args)
}

type _ICoreWebView2FrameDestroyedEventHandlerImpl interface {
	_IUnknownImpl
	FrameDestroyed(sender *ICoreWebView2Frame, args *_IUnknown) uintptr
}

var _ICoreWebView2FrameDestroyedEventHandlerFn = _ICoreWebView2FrameDestroyedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2FrameDestroyedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2FrameDestroyedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2FrameDestroyedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2FrameDestroyedEventHandlerInvoke),
}

func NewICoreWebView2FrameDestroyedEventHandler(impl _ICoreWebView2FrameDestroyedEventHandlerImpl) *ICoreWebView2FrameDestroyedEventHandler {
	return &ICoreWebView2FrameDestroyedEventHandler{
		vtbl: &_ICoreWebView2FrameDestroyedEventHandlerFn,
		impl: impl,
	}
}
