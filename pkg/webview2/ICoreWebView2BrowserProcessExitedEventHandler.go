//go:build windows

package webview2

type _ICoreWebView2BrowserProcessExitedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2BrowserProcessExitedEventHandler struct {
	vtbl *_ICoreWebView2BrowserProcessExitedEventHandlerVtbl
	impl _ICoreWebView2BrowserProcessExitedEventHandlerImpl
}

func (i *ICoreWebView2BrowserProcessExitedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2BrowserProcessExitedEventHandlerIUnknownQueryInterface(this *ICoreWebView2BrowserProcessExitedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2BrowserProcessExitedEventHandlerIUnknownAddRef(this *ICoreWebView2BrowserProcessExitedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2BrowserProcessExitedEventHandlerIUnknownRelease(this *ICoreWebView2BrowserProcessExitedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2BrowserProcessExitedEventHandlerInvoke(this *ICoreWebView2BrowserProcessExitedEventHandler, sender *ICoreWebView2Environment, args *ICoreWebView2BrowserProcessExitedEventArgs) uintptr {
	return this.impl.BrowserProcessExited(sender, args)
}

type _ICoreWebView2BrowserProcessExitedEventHandlerImpl interface {
	_IUnknownImpl
	BrowserProcessExited(sender *ICoreWebView2Environment, args *ICoreWebView2BrowserProcessExitedEventArgs) uintptr
}

var _ICoreWebView2BrowserProcessExitedEventHandlerFn = _ICoreWebView2BrowserProcessExitedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2BrowserProcessExitedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2BrowserProcessExitedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2BrowserProcessExitedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2BrowserProcessExitedEventHandlerInvoke),
}

func NewICoreWebView2BrowserProcessExitedEventHandler(impl _ICoreWebView2BrowserProcessExitedEventHandlerImpl) *ICoreWebView2BrowserProcessExitedEventHandler {
	return &ICoreWebView2BrowserProcessExitedEventHandler{
		vtbl: &_ICoreWebView2BrowserProcessExitedEventHandlerFn,
		impl: impl,
	}
}
