//go:build windows

package webview2

type _ICoreWebView2SetPermissionStateCompletedHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2SetPermissionStateCompletedHandler struct {
	vtbl *_ICoreWebView2SetPermissionStateCompletedHandlerVtbl
	impl _ICoreWebView2SetPermissionStateCompletedHandlerImpl
}

func (i *ICoreWebView2SetPermissionStateCompletedHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2SetPermissionStateCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2SetPermissionStateCompletedHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2SetPermissionStateCompletedHandlerIUnknownAddRef(this *ICoreWebView2SetPermissionStateCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2SetPermissionStateCompletedHandlerIUnknownRelease(this *ICoreWebView2SetPermissionStateCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2SetPermissionStateCompletedHandlerInvoke(this *ICoreWebView2SetPermissionStateCompletedHandler, errorCode uintptr) uintptr {
	return this.impl.SetPermissionStateCompleted(errorCode)
}

type _ICoreWebView2SetPermissionStateCompletedHandlerImpl interface {
	_IUnknownImpl
	SetPermissionStateCompleted(errorCode uintptr) uintptr
}

var _ICoreWebView2SetPermissionStateCompletedHandlerFn = _ICoreWebView2SetPermissionStateCompletedHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2SetPermissionStateCompletedHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2SetPermissionStateCompletedHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2SetPermissionStateCompletedHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2SetPermissionStateCompletedHandlerInvoke),
}

func NewICoreWebView2SetPermissionStateCompletedHandler(impl _ICoreWebView2SetPermissionStateCompletedHandlerImpl) *ICoreWebView2SetPermissionStateCompletedHandler {
	return &ICoreWebView2SetPermissionStateCompletedHandler{
		vtbl: &_ICoreWebView2SetPermissionStateCompletedHandlerFn,
		impl: impl,
	}
}
