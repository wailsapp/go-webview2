//go:build windows

package webview2

type _ICoreWebView2BytesReceivedChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2BytesReceivedChangedEventHandler struct {
	vtbl *_ICoreWebView2BytesReceivedChangedEventHandlerVtbl
	impl _ICoreWebView2BytesReceivedChangedEventHandlerImpl
}

func (i *ICoreWebView2BytesReceivedChangedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2BytesReceivedChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2BytesReceivedChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2BytesReceivedChangedEventHandlerIUnknownAddRef(this *ICoreWebView2BytesReceivedChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2BytesReceivedChangedEventHandlerIUnknownRelease(this *ICoreWebView2BytesReceivedChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2BytesReceivedChangedEventHandlerInvoke(this *ICoreWebView2BytesReceivedChangedEventHandler, sender *ICoreWebView2DownloadOperation, args *_IUnknown) uintptr {
	return this.impl.BytesReceivedChanged(sender, args)
}

type _ICoreWebView2BytesReceivedChangedEventHandlerImpl interface {
	_IUnknownImpl
	BytesReceivedChanged(sender *ICoreWebView2DownloadOperation, args *_IUnknown) uintptr
}

var _ICoreWebView2BytesReceivedChangedEventHandlerFn = _ICoreWebView2BytesReceivedChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2BytesReceivedChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2BytesReceivedChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2BytesReceivedChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2BytesReceivedChangedEventHandlerInvoke),
}

func NewICoreWebView2BytesReceivedChangedEventHandler(impl _ICoreWebView2BytesReceivedChangedEventHandlerImpl) *ICoreWebView2BytesReceivedChangedEventHandler {
	return &ICoreWebView2BytesReceivedChangedEventHandler{
		vtbl: &_ICoreWebView2BytesReceivedChangedEventHandlerFn,
		impl: impl,
	}
}
