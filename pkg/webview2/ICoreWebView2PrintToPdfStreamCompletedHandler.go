//go:build windows

package webview2

type _ICoreWebView2PrintToPdfStreamCompletedHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2PrintToPdfStreamCompletedHandler struct {
	vtbl *_ICoreWebView2PrintToPdfStreamCompletedHandlerVtbl
	impl _ICoreWebView2PrintToPdfStreamCompletedHandlerImpl
}

func (i *ICoreWebView2PrintToPdfStreamCompletedHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2PrintToPdfStreamCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2PrintToPdfStreamCompletedHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2PrintToPdfStreamCompletedHandlerIUnknownAddRef(this *ICoreWebView2PrintToPdfStreamCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2PrintToPdfStreamCompletedHandlerIUnknownRelease(this *ICoreWebView2PrintToPdfStreamCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2PrintToPdfStreamCompletedHandlerInvoke(this *ICoreWebView2PrintToPdfStreamCompletedHandler, errorCode uintptr, pdfStream *IStream) uintptr {
	return this.impl.PrintToPdfStreamCompleted(errorCode, pdfStream)
}

type _ICoreWebView2PrintToPdfStreamCompletedHandlerImpl interface {
	_IUnknownImpl
	PrintToPdfStreamCompleted(errorCode uintptr, pdfStream *IStream) uintptr
}

var _ICoreWebView2PrintToPdfStreamCompletedHandlerFn = _ICoreWebView2PrintToPdfStreamCompletedHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2PrintToPdfStreamCompletedHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2PrintToPdfStreamCompletedHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2PrintToPdfStreamCompletedHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2PrintToPdfStreamCompletedHandlerInvoke),
}

func NewICoreWebView2PrintToPdfStreamCompletedHandler(impl _ICoreWebView2PrintToPdfStreamCompletedHandlerImpl) *ICoreWebView2PrintToPdfStreamCompletedHandler {
	return &ICoreWebView2PrintToPdfStreamCompletedHandler{
		vtbl: &_ICoreWebView2PrintToPdfStreamCompletedHandlerFn,
		impl: impl,
	}
}
