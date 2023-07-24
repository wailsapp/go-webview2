//go:build windows

package webview2

type _ICoreWebView2FrameNavigationCompletedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2FrameNavigationCompletedEventHandler struct {
	vtbl *_ICoreWebView2FrameNavigationCompletedEventHandlerVtbl
	impl _ICoreWebView2FrameNavigationCompletedEventHandlerImpl
}

func (i *ICoreWebView2FrameNavigationCompletedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2FrameNavigationCompletedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FrameNavigationCompletedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FrameNavigationCompletedEventHandlerIUnknownAddRef(this *ICoreWebView2FrameNavigationCompletedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FrameNavigationCompletedEventHandlerIUnknownRelease(this *ICoreWebView2FrameNavigationCompletedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FrameNavigationCompletedEventHandlerInvoke(this *ICoreWebView2FrameNavigationCompletedEventHandler, sender *ICoreWebView2Frame, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	return this.impl.FrameNavigationCompleted(sender, args)
}

type _ICoreWebView2FrameNavigationCompletedEventHandlerImpl interface {
	_IUnknownImpl
	FrameNavigationCompleted(sender *ICoreWebView2Frame, args *ICoreWebView2NavigationCompletedEventArgs) uintptr
}

var _ICoreWebView2FrameNavigationCompletedEventHandlerFn = _ICoreWebView2FrameNavigationCompletedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2FrameNavigationCompletedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2FrameNavigationCompletedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2FrameNavigationCompletedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2FrameNavigationCompletedEventHandlerInvoke),
}

func NewICoreWebView2FrameNavigationCompletedEventHandler(impl _ICoreWebView2FrameNavigationCompletedEventHandlerImpl) *ICoreWebView2FrameNavigationCompletedEventHandler {
	return &ICoreWebView2FrameNavigationCompletedEventHandler{
		vtbl: &_ICoreWebView2FrameNavigationCompletedEventHandlerFn,
		impl: impl,
	}
}
