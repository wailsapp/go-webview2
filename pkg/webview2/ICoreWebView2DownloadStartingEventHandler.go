//go:build windows

package webview2

type _ICoreWebView2DownloadStartingEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2DownloadStartingEventHandler struct {
	vtbl *_ICoreWebView2DownloadStartingEventHandlerVtbl
	impl _ICoreWebView2DownloadStartingEventHandlerImpl
}

func (i *ICoreWebView2DownloadStartingEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2DownloadStartingEventHandlerIUnknownQueryInterface(this *ICoreWebView2DownloadStartingEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2DownloadStartingEventHandlerIUnknownAddRef(this *ICoreWebView2DownloadStartingEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2DownloadStartingEventHandlerIUnknownRelease(this *ICoreWebView2DownloadStartingEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2DownloadStartingEventHandlerInvoke(this *ICoreWebView2DownloadStartingEventHandler, sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) uintptr {
	return this.impl.DownloadStarting(sender, args)
}

type _ICoreWebView2DownloadStartingEventHandlerImpl interface {
	_IUnknownImpl
	DownloadStarting(sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) uintptr
}

var _ICoreWebView2DownloadStartingEventHandlerFn = _ICoreWebView2DownloadStartingEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2DownloadStartingEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2DownloadStartingEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2DownloadStartingEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2DownloadStartingEventHandlerInvoke),
}

func NewICoreWebView2DownloadStartingEventHandler(impl _ICoreWebView2DownloadStartingEventHandlerImpl) *ICoreWebView2DownloadStartingEventHandler {
	return &ICoreWebView2DownloadStartingEventHandler{
		vtbl: &_ICoreWebView2DownloadStartingEventHandlerFn,
		impl: impl,
	}
}
