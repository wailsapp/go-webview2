//go:build windows

package webview2

type _ICoreWebView2EstimatedEndTimeChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2EstimatedEndTimeChangedEventHandler struct {
	vtbl *_ICoreWebView2EstimatedEndTimeChangedEventHandlerVtbl
	impl _ICoreWebView2EstimatedEndTimeChangedEventHandlerImpl
}

func (i *ICoreWebView2EstimatedEndTimeChangedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2EstimatedEndTimeChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2EstimatedEndTimeChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2EstimatedEndTimeChangedEventHandlerIUnknownAddRef(this *ICoreWebView2EstimatedEndTimeChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2EstimatedEndTimeChangedEventHandlerIUnknownRelease(this *ICoreWebView2EstimatedEndTimeChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2EstimatedEndTimeChangedEventHandlerInvoke(this *ICoreWebView2EstimatedEndTimeChangedEventHandler, sender *ICoreWebView2DownloadOperation, args *_IUnknown) uintptr {
	return this.impl.EstimatedEndTimeChanged(sender, args)
}

type _ICoreWebView2EstimatedEndTimeChangedEventHandlerImpl interface {
	_IUnknownImpl
	EstimatedEndTimeChanged(sender *ICoreWebView2DownloadOperation, args *_IUnknown) uintptr
}

var _ICoreWebView2EstimatedEndTimeChangedEventHandlerFn = _ICoreWebView2EstimatedEndTimeChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2EstimatedEndTimeChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2EstimatedEndTimeChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2EstimatedEndTimeChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2EstimatedEndTimeChangedEventHandlerInvoke),
}

func NewICoreWebView2EstimatedEndTimeChangedEventHandler(impl _ICoreWebView2EstimatedEndTimeChangedEventHandlerImpl) *ICoreWebView2EstimatedEndTimeChangedEventHandler {
	return &ICoreWebView2EstimatedEndTimeChangedEventHandler{
		vtbl: &_ICoreWebView2EstimatedEndTimeChangedEventHandlerFn,
		impl: impl,
	}
}
