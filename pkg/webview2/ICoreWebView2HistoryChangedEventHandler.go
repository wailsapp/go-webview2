//go:build windows

package webview2

type _ICoreWebView2HistoryChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2HistoryChangedEventHandler struct {
	vtbl *_ICoreWebView2HistoryChangedEventHandlerVtbl
	impl _ICoreWebView2HistoryChangedEventHandlerImpl
}

func (i *ICoreWebView2HistoryChangedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2HistoryChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2HistoryChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2HistoryChangedEventHandlerIUnknownAddRef(this *ICoreWebView2HistoryChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2HistoryChangedEventHandlerIUnknownRelease(this *ICoreWebView2HistoryChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2HistoryChangedEventHandlerInvoke(this *ICoreWebView2HistoryChangedEventHandler, sender *ICoreWebView2, args *_IUnknown) uintptr {
	return this.impl.HistoryChanged(sender, args)
}

type _ICoreWebView2HistoryChangedEventHandlerImpl interface {
	_IUnknownImpl
	HistoryChanged(sender *ICoreWebView2, args *_IUnknown) uintptr
}

var _ICoreWebView2HistoryChangedEventHandlerFn = _ICoreWebView2HistoryChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2HistoryChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2HistoryChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2HistoryChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2HistoryChangedEventHandlerInvoke),
}

func NewICoreWebView2HistoryChangedEventHandler(impl _ICoreWebView2HistoryChangedEventHandlerImpl) *ICoreWebView2HistoryChangedEventHandler {
	return &ICoreWebView2HistoryChangedEventHandler{
		vtbl: &_ICoreWebView2HistoryChangedEventHandlerFn,
		impl: impl,
	}
}
