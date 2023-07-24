//go:build windows

package webview2

type _ICoreWebView2FramePermissionRequestedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2FramePermissionRequestedEventHandler struct {
	vtbl *_ICoreWebView2FramePermissionRequestedEventHandlerVtbl
	impl _ICoreWebView2FramePermissionRequestedEventHandlerImpl
}

func (i *ICoreWebView2FramePermissionRequestedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2FramePermissionRequestedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FramePermissionRequestedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FramePermissionRequestedEventHandlerIUnknownAddRef(this *ICoreWebView2FramePermissionRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FramePermissionRequestedEventHandlerIUnknownRelease(this *ICoreWebView2FramePermissionRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FramePermissionRequestedEventHandlerInvoke(this *ICoreWebView2FramePermissionRequestedEventHandler, sender *ICoreWebView2Frame, args *ICoreWebView2PermissionRequestedEventArgs2) uintptr {
	return this.impl.FramePermissionRequested(sender, args)
}

type _ICoreWebView2FramePermissionRequestedEventHandlerImpl interface {
	_IUnknownImpl
	FramePermissionRequested(sender *ICoreWebView2Frame, args *ICoreWebView2PermissionRequestedEventArgs2) uintptr
}

var _ICoreWebView2FramePermissionRequestedEventHandlerFn = _ICoreWebView2FramePermissionRequestedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2FramePermissionRequestedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2FramePermissionRequestedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2FramePermissionRequestedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2FramePermissionRequestedEventHandlerInvoke),
}

func NewICoreWebView2FramePermissionRequestedEventHandler(impl _ICoreWebView2FramePermissionRequestedEventHandlerImpl) *ICoreWebView2FramePermissionRequestedEventHandler {
	return &ICoreWebView2FramePermissionRequestedEventHandler{
		vtbl: &_ICoreWebView2FramePermissionRequestedEventHandlerFn,
		impl: impl,
	}
}
