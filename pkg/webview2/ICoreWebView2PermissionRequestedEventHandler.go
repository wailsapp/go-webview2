//go:build windows

package webview2

type _ICoreWebView2PermissionRequestedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2PermissionRequestedEventHandler struct {
	vtbl *_ICoreWebView2PermissionRequestedEventHandlerVtbl
	impl _ICoreWebView2PermissionRequestedEventHandlerImpl
}

func (i *ICoreWebView2PermissionRequestedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2PermissionRequestedEventHandlerIUnknownQueryInterface(this *ICoreWebView2PermissionRequestedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2PermissionRequestedEventHandlerIUnknownAddRef(this *ICoreWebView2PermissionRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2PermissionRequestedEventHandlerIUnknownRelease(this *ICoreWebView2PermissionRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2PermissionRequestedEventHandlerInvoke(this *ICoreWebView2PermissionRequestedEventHandler, sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) uintptr {
	return this.impl.PermissionRequested(sender, args)
}

type _ICoreWebView2PermissionRequestedEventHandlerImpl interface {
	_IUnknownImpl
	PermissionRequested(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) uintptr
}

var _ICoreWebView2PermissionRequestedEventHandlerFn = _ICoreWebView2PermissionRequestedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2PermissionRequestedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2PermissionRequestedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2PermissionRequestedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2PermissionRequestedEventHandlerInvoke),
}

func NewICoreWebView2PermissionRequestedEventHandler(impl _ICoreWebView2PermissionRequestedEventHandlerImpl) *ICoreWebView2PermissionRequestedEventHandler {
	return &ICoreWebView2PermissionRequestedEventHandler{
		vtbl: &_ICoreWebView2PermissionRequestedEventHandlerFn,
		impl: impl,
	}
}
