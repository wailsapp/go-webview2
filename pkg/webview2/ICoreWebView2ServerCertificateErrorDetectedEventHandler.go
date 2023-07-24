//go:build windows

package webview2

type _ICoreWebView2ServerCertificateErrorDetectedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2ServerCertificateErrorDetectedEventHandler struct {
	vtbl *_ICoreWebView2ServerCertificateErrorDetectedEventHandlerVtbl
	impl _ICoreWebView2ServerCertificateErrorDetectedEventHandlerImpl
}

func (i *ICoreWebView2ServerCertificateErrorDetectedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2ServerCertificateErrorDetectedEventHandlerIUnknownQueryInterface(this *ICoreWebView2ServerCertificateErrorDetectedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ServerCertificateErrorDetectedEventHandlerIUnknownAddRef(this *ICoreWebView2ServerCertificateErrorDetectedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ServerCertificateErrorDetectedEventHandlerIUnknownRelease(this *ICoreWebView2ServerCertificateErrorDetectedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ServerCertificateErrorDetectedEventHandlerInvoke(this *ICoreWebView2ServerCertificateErrorDetectedEventHandler, sender *ICoreWebView2, args *ICoreWebView2ServerCertificateErrorDetectedEventArgs) uintptr {
	return this.impl.ServerCertificateErrorDetected(sender, args)
}

type _ICoreWebView2ServerCertificateErrorDetectedEventHandlerImpl interface {
	_IUnknownImpl
	ServerCertificateErrorDetected(sender *ICoreWebView2, args *ICoreWebView2ServerCertificateErrorDetectedEventArgs) uintptr
}

var _ICoreWebView2ServerCertificateErrorDetectedEventHandlerFn = _ICoreWebView2ServerCertificateErrorDetectedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2ServerCertificateErrorDetectedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2ServerCertificateErrorDetectedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2ServerCertificateErrorDetectedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2ServerCertificateErrorDetectedEventHandlerInvoke),
}

func NewICoreWebView2ServerCertificateErrorDetectedEventHandler(impl _ICoreWebView2ServerCertificateErrorDetectedEventHandlerImpl) *ICoreWebView2ServerCertificateErrorDetectedEventHandler {
	return &ICoreWebView2ServerCertificateErrorDetectedEventHandler{
		vtbl: &_ICoreWebView2ServerCertificateErrorDetectedEventHandlerFn,
		impl: impl,
	}
}
