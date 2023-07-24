//go:build windows

package webview2

type _ICoreWebView2DocumentTitleChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2DocumentTitleChangedEventHandler struct {
	vtbl *_ICoreWebView2DocumentTitleChangedEventHandlerVtbl
	impl _ICoreWebView2DocumentTitleChangedEventHandlerImpl
}

func (i *ICoreWebView2DocumentTitleChangedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2DocumentTitleChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2DocumentTitleChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2DocumentTitleChangedEventHandlerIUnknownAddRef(this *ICoreWebView2DocumentTitleChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2DocumentTitleChangedEventHandlerIUnknownRelease(this *ICoreWebView2DocumentTitleChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2DocumentTitleChangedEventHandlerInvoke(this *ICoreWebView2DocumentTitleChangedEventHandler, sender *ICoreWebView2, args *_IUnknown) uintptr {
	return this.impl.DocumentTitleChanged(sender, args)
}

type _ICoreWebView2DocumentTitleChangedEventHandlerImpl interface {
	_IUnknownImpl
	DocumentTitleChanged(sender *ICoreWebView2, args *_IUnknown) uintptr
}

var _ICoreWebView2DocumentTitleChangedEventHandlerFn = _ICoreWebView2DocumentTitleChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2DocumentTitleChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2DocumentTitleChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2DocumentTitleChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2DocumentTitleChangedEventHandlerInvoke),
}

func NewICoreWebView2DocumentTitleChangedEventHandler(impl _ICoreWebView2DocumentTitleChangedEventHandlerImpl) *ICoreWebView2DocumentTitleChangedEventHandler {
	return &ICoreWebView2DocumentTitleChangedEventHandler{
		vtbl: &_ICoreWebView2DocumentTitleChangedEventHandlerFn,
		impl: impl,
	}
}
