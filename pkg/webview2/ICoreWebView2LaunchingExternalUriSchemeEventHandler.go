//go:build windows

package webview2

type _ICoreWebView2LaunchingExternalUriSchemeEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2LaunchingExternalUriSchemeEventHandler struct {
	vtbl *_ICoreWebView2LaunchingExternalUriSchemeEventHandlerVtbl
	impl _ICoreWebView2LaunchingExternalUriSchemeEventHandlerImpl
}

func (i *ICoreWebView2LaunchingExternalUriSchemeEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2LaunchingExternalUriSchemeEventHandlerIUnknownQueryInterface(this *ICoreWebView2LaunchingExternalUriSchemeEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2LaunchingExternalUriSchemeEventHandlerIUnknownAddRef(this *ICoreWebView2LaunchingExternalUriSchemeEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2LaunchingExternalUriSchemeEventHandlerIUnknownRelease(this *ICoreWebView2LaunchingExternalUriSchemeEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2LaunchingExternalUriSchemeEventHandlerInvoke(this *ICoreWebView2LaunchingExternalUriSchemeEventHandler, sender *ICoreWebView2, args *ICoreWebView2LaunchingExternalUriSchemeEventArgs) uintptr {
	return this.impl.LaunchingExternalUriScheme(sender, args)
}

type _ICoreWebView2LaunchingExternalUriSchemeEventHandlerImpl interface {
	_IUnknownImpl
	LaunchingExternalUriScheme(sender *ICoreWebView2, args *ICoreWebView2LaunchingExternalUriSchemeEventArgs) uintptr
}

var _ICoreWebView2LaunchingExternalUriSchemeEventHandlerFn = _ICoreWebView2LaunchingExternalUriSchemeEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2LaunchingExternalUriSchemeEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2LaunchingExternalUriSchemeEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2LaunchingExternalUriSchemeEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2LaunchingExternalUriSchemeEventHandlerInvoke),
}

func NewICoreWebView2LaunchingExternalUriSchemeEventHandler(impl _ICoreWebView2LaunchingExternalUriSchemeEventHandlerImpl) *ICoreWebView2LaunchingExternalUriSchemeEventHandler {
	return &ICoreWebView2LaunchingExternalUriSchemeEventHandler{
		vtbl: &_ICoreWebView2LaunchingExternalUriSchemeEventHandlerFn,
		impl: impl,
	}
}
