//go:build windows

package webview2

type _ICoreWebView2PrintCompletedHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2PrintCompletedHandler struct {
	vtbl *_ICoreWebView2PrintCompletedHandlerVtbl
	impl _ICoreWebView2PrintCompletedHandlerImpl
}

func (i *ICoreWebView2PrintCompletedHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2PrintCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2PrintCompletedHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2PrintCompletedHandlerIUnknownAddRef(this *ICoreWebView2PrintCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2PrintCompletedHandlerIUnknownRelease(this *ICoreWebView2PrintCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2PrintCompletedHandlerInvoke(this *ICoreWebView2PrintCompletedHandler, errorCode uintptr, printStatus COREWEBVIEW2_PRINT_STATUS) uintptr {
	return this.impl.PrintCompleted(errorCode, printStatus)
}

type _ICoreWebView2PrintCompletedHandlerImpl interface {
	_IUnknownImpl
	PrintCompleted(errorCode uintptr, printStatus COREWEBVIEW2_PRINT_STATUS) uintptr
}

var _ICoreWebView2PrintCompletedHandlerFn = _ICoreWebView2PrintCompletedHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2PrintCompletedHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2PrintCompletedHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2PrintCompletedHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2PrintCompletedHandlerInvoke),
}

func NewICoreWebView2PrintCompletedHandler(impl _ICoreWebView2PrintCompletedHandlerImpl) *ICoreWebView2PrintCompletedHandler {
	return &ICoreWebView2PrintCompletedHandler{
		vtbl: &_ICoreWebView2PrintCompletedHandlerFn,
		impl: impl,
	}
}
