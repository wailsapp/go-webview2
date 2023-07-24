//go:build windows

package webview2

type _ICoreWebView2GetFaviconCompletedHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2GetFaviconCompletedHandler struct {
	vtbl *_ICoreWebView2GetFaviconCompletedHandlerVtbl
	impl _ICoreWebView2GetFaviconCompletedHandlerImpl
}

func (i *ICoreWebView2GetFaviconCompletedHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2GetFaviconCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2GetFaviconCompletedHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2GetFaviconCompletedHandlerIUnknownAddRef(this *ICoreWebView2GetFaviconCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2GetFaviconCompletedHandlerIUnknownRelease(this *ICoreWebView2GetFaviconCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2GetFaviconCompletedHandlerInvoke(this *ICoreWebView2GetFaviconCompletedHandler, errorCode uintptr, faviconStream *IStream) uintptr {
	return this.impl.GetFaviconCompleted(errorCode, faviconStream)
}

type _ICoreWebView2GetFaviconCompletedHandlerImpl interface {
	_IUnknownImpl
	GetFaviconCompleted(errorCode uintptr, faviconStream *IStream) uintptr
}

var _ICoreWebView2GetFaviconCompletedHandlerFn = _ICoreWebView2GetFaviconCompletedHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2GetFaviconCompletedHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2GetFaviconCompletedHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2GetFaviconCompletedHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2GetFaviconCompletedHandlerInvoke),
}

func NewICoreWebView2GetFaviconCompletedHandler(impl _ICoreWebView2GetFaviconCompletedHandlerImpl) *ICoreWebView2GetFaviconCompletedHandler {
	return &ICoreWebView2GetFaviconCompletedHandler{
		vtbl: &_ICoreWebView2GetFaviconCompletedHandlerFn,
		impl: impl,
	}
}
