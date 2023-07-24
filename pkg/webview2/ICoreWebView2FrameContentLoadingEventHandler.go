//go:build windows

package webview2

type _ICoreWebView2FrameContentLoadingEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2FrameContentLoadingEventHandler struct {
	vtbl *_ICoreWebView2FrameContentLoadingEventHandlerVtbl
	impl _ICoreWebView2FrameContentLoadingEventHandlerImpl
}

func (i *ICoreWebView2FrameContentLoadingEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2FrameContentLoadingEventHandlerIUnknownQueryInterface(this *ICoreWebView2FrameContentLoadingEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FrameContentLoadingEventHandlerIUnknownAddRef(this *ICoreWebView2FrameContentLoadingEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FrameContentLoadingEventHandlerIUnknownRelease(this *ICoreWebView2FrameContentLoadingEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FrameContentLoadingEventHandlerInvoke(this *ICoreWebView2FrameContentLoadingEventHandler, sender *ICoreWebView2Frame, args *ICoreWebView2ContentLoadingEventArgs) uintptr {
	return this.impl.FrameContentLoading(sender, args)
}

type _ICoreWebView2FrameContentLoadingEventHandlerImpl interface {
	_IUnknownImpl
	FrameContentLoading(sender *ICoreWebView2Frame, args *ICoreWebView2ContentLoadingEventArgs) uintptr
}

var _ICoreWebView2FrameContentLoadingEventHandlerFn = _ICoreWebView2FrameContentLoadingEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2FrameContentLoadingEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2FrameContentLoadingEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2FrameContentLoadingEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2FrameContentLoadingEventHandlerInvoke),
}

func NewICoreWebView2FrameContentLoadingEventHandler(impl _ICoreWebView2FrameContentLoadingEventHandlerImpl) *ICoreWebView2FrameContentLoadingEventHandler {
	return &ICoreWebView2FrameContentLoadingEventHandler{
		vtbl: &_ICoreWebView2FrameContentLoadingEventHandlerFn,
		impl: impl,
	}
}
