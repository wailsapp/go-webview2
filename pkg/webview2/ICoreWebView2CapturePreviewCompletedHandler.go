//go:build windows

package webview2

type _ICoreWebView2CapturePreviewCompletedHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2CapturePreviewCompletedHandler struct {
	vtbl *_ICoreWebView2CapturePreviewCompletedHandlerVtbl
	impl _ICoreWebView2CapturePreviewCompletedHandlerImpl
}

func (i *ICoreWebView2CapturePreviewCompletedHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2CapturePreviewCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2CapturePreviewCompletedHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2CapturePreviewCompletedHandlerIUnknownAddRef(this *ICoreWebView2CapturePreviewCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2CapturePreviewCompletedHandlerIUnknownRelease(this *ICoreWebView2CapturePreviewCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2CapturePreviewCompletedHandlerInvoke(this *ICoreWebView2CapturePreviewCompletedHandler, errorCode uintptr) uintptr {
	return this.impl.CapturePreviewCompleted(errorCode)
}

type _ICoreWebView2CapturePreviewCompletedHandlerImpl interface {
	_IUnknownImpl
	CapturePreviewCompleted(errorCode uintptr) uintptr
}

var _ICoreWebView2CapturePreviewCompletedHandlerFn = _ICoreWebView2CapturePreviewCompletedHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2CapturePreviewCompletedHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2CapturePreviewCompletedHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2CapturePreviewCompletedHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2CapturePreviewCompletedHandlerInvoke),
}

func NewICoreWebView2CapturePreviewCompletedHandler(impl _ICoreWebView2CapturePreviewCompletedHandlerImpl) *ICoreWebView2CapturePreviewCompletedHandler {
	return &ICoreWebView2CapturePreviewCompletedHandler{
		vtbl: &_ICoreWebView2CapturePreviewCompletedHandlerFn,
		impl: impl,
	}
}
