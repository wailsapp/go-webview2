//go:build windows

package webview2

type _ICoreWebView2ClientCertificateRequestedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2ClientCertificateRequestedEventHandler struct {
	vtbl *_ICoreWebView2ClientCertificateRequestedEventHandlerVtbl
	impl _ICoreWebView2ClientCertificateRequestedEventHandlerImpl
}

func (i *ICoreWebView2ClientCertificateRequestedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2ClientCertificateRequestedEventHandlerIUnknownQueryInterface(this *ICoreWebView2ClientCertificateRequestedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ClientCertificateRequestedEventHandlerIUnknownAddRef(this *ICoreWebView2ClientCertificateRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ClientCertificateRequestedEventHandlerIUnknownRelease(this *ICoreWebView2ClientCertificateRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ClientCertificateRequestedEventHandlerInvoke(this *ICoreWebView2ClientCertificateRequestedEventHandler, sender *ICoreWebView2, args *ICoreWebView2ClientCertificateRequestedEventArgs) uintptr {
	return this.impl.ClientCertificateRequested(sender, args)
}

type _ICoreWebView2ClientCertificateRequestedEventHandlerImpl interface {
	_IUnknownImpl
	ClientCertificateRequested(sender *ICoreWebView2, args *ICoreWebView2ClientCertificateRequestedEventArgs) uintptr
}

var _ICoreWebView2ClientCertificateRequestedEventHandlerFn = _ICoreWebView2ClientCertificateRequestedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2ClientCertificateRequestedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2ClientCertificateRequestedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2ClientCertificateRequestedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2ClientCertificateRequestedEventHandlerInvoke),
}

func NewICoreWebView2ClientCertificateRequestedEventHandler(impl _ICoreWebView2ClientCertificateRequestedEventHandlerImpl) *ICoreWebView2ClientCertificateRequestedEventHandler {
	return &ICoreWebView2ClientCertificateRequestedEventHandler{
		vtbl: &_ICoreWebView2ClientCertificateRequestedEventHandlerFn,
		impl: impl,
	}
}
