//go:build windows

package webview2

type _ICoreWebView2FrameNameChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2FrameNameChangedEventHandler struct {
	vtbl *_ICoreWebView2FrameNameChangedEventHandlerVtbl
	impl _ICoreWebView2FrameNameChangedEventHandlerImpl
}

func (i *ICoreWebView2FrameNameChangedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2FrameNameChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FrameNameChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FrameNameChangedEventHandlerIUnknownAddRef(this *ICoreWebView2FrameNameChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FrameNameChangedEventHandlerIUnknownRelease(this *ICoreWebView2FrameNameChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FrameNameChangedEventHandlerInvoke(this *ICoreWebView2FrameNameChangedEventHandler, sender *ICoreWebView2Frame, args *_IUnknown) uintptr {
	return this.impl.FrameNameChanged(sender, args)
}

type _ICoreWebView2FrameNameChangedEventHandlerImpl interface {
	_IUnknownImpl
	FrameNameChanged(sender *ICoreWebView2Frame, args *_IUnknown) uintptr
}

var _ICoreWebView2FrameNameChangedEventHandlerFn = _ICoreWebView2FrameNameChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2FrameNameChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2FrameNameChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2FrameNameChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2FrameNameChangedEventHandlerInvoke),
}

func NewICoreWebView2FrameNameChangedEventHandler(impl _ICoreWebView2FrameNameChangedEventHandlerImpl) *ICoreWebView2FrameNameChangedEventHandler {
	return &ICoreWebView2FrameNameChangedEventHandler{
		vtbl: &_ICoreWebView2FrameNameChangedEventHandlerFn,
		impl: impl,
	}
}
