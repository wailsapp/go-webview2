//go:build windows

package webview2

type _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2IsDocumentPlayingAudioChangedEventHandler struct {
	vtbl *_ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerVtbl
	impl _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerImpl
}

func (i *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerIUnknownAddRef(this *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerIUnknownRelease(this *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerInvoke(this *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler, sender *ICoreWebView2, args *_IUnknown) uintptr {
	return this.impl.IsDocumentPlayingAudioChanged(sender, args)
}

type _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerImpl interface {
	_IUnknownImpl
	IsDocumentPlayingAudioChanged(sender *ICoreWebView2, args *_IUnknown) uintptr
}

var _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerFn = _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerInvoke),
}

func NewICoreWebView2IsDocumentPlayingAudioChangedEventHandler(impl _ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerImpl) *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler {
	return &ICoreWebView2IsDocumentPlayingAudioChangedEventHandler{
		vtbl: &_ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerFn,
		impl: impl,
	}
}
