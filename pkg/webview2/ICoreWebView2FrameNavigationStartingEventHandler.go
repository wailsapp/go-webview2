//go:build windows

package webview2

type _ICoreWebView2FrameNavigationStartingEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2FrameNavigationStartingEventHandler struct {
	vtbl *_ICoreWebView2FrameNavigationStartingEventHandlerVtbl
	impl _ICoreWebView2FrameNavigationStartingEventHandlerImpl
}

func (i *ICoreWebView2FrameNavigationStartingEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2FrameNavigationStartingEventHandlerIUnknownQueryInterface(this *ICoreWebView2FrameNavigationStartingEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FrameNavigationStartingEventHandlerIUnknownAddRef(this *ICoreWebView2FrameNavigationStartingEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FrameNavigationStartingEventHandlerIUnknownRelease(this *ICoreWebView2FrameNavigationStartingEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FrameNavigationStartingEventHandlerInvoke(this *ICoreWebView2FrameNavigationStartingEventHandler, sender *ICoreWebView2Frame, args *ICoreWebView2NavigationStartingEventArgs) uintptr {
	return this.impl.FrameNavigationStarting(sender, args)
}

type _ICoreWebView2FrameNavigationStartingEventHandlerImpl interface {
	_IUnknownImpl
	FrameNavigationStarting(sender *ICoreWebView2Frame, args *ICoreWebView2NavigationStartingEventArgs) uintptr
}

var _ICoreWebView2FrameNavigationStartingEventHandlerFn = _ICoreWebView2FrameNavigationStartingEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2FrameNavigationStartingEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2FrameNavigationStartingEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2FrameNavigationStartingEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2FrameNavigationStartingEventHandlerInvoke),
}

func NewICoreWebView2FrameNavigationStartingEventHandler(impl _ICoreWebView2FrameNavigationStartingEventHandlerImpl) *ICoreWebView2FrameNavigationStartingEventHandler {
	return &ICoreWebView2FrameNavigationStartingEventHandler{
		vtbl: &_ICoreWebView2FrameNavigationStartingEventHandlerFn,
		impl: impl,
	}
}
