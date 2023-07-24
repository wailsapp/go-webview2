//go:build windows

package webview2

type _ICoreWebView2ScriptDialogOpeningEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2ScriptDialogOpeningEventHandler struct {
	vtbl *_ICoreWebView2ScriptDialogOpeningEventHandlerVtbl
	impl _ICoreWebView2ScriptDialogOpeningEventHandlerImpl
}

func (i *ICoreWebView2ScriptDialogOpeningEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2ScriptDialogOpeningEventHandlerIUnknownQueryInterface(this *ICoreWebView2ScriptDialogOpeningEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ScriptDialogOpeningEventHandlerIUnknownAddRef(this *ICoreWebView2ScriptDialogOpeningEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ScriptDialogOpeningEventHandlerIUnknownRelease(this *ICoreWebView2ScriptDialogOpeningEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ScriptDialogOpeningEventHandlerInvoke(this *ICoreWebView2ScriptDialogOpeningEventHandler, sender *ICoreWebView2, args *ICoreWebView2ScriptDialogOpeningEventArgs) uintptr {
	return this.impl.ScriptDialogOpening(sender, args)
}

type _ICoreWebView2ScriptDialogOpeningEventHandlerImpl interface {
	_IUnknownImpl
	ScriptDialogOpening(sender *ICoreWebView2, args *ICoreWebView2ScriptDialogOpeningEventArgs) uintptr
}

var _ICoreWebView2ScriptDialogOpeningEventHandlerFn = _ICoreWebView2ScriptDialogOpeningEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2ScriptDialogOpeningEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2ScriptDialogOpeningEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2ScriptDialogOpeningEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2ScriptDialogOpeningEventHandlerInvoke),
}

func NewICoreWebView2ScriptDialogOpeningEventHandler(impl _ICoreWebView2ScriptDialogOpeningEventHandlerImpl) *ICoreWebView2ScriptDialogOpeningEventHandler {
	return &ICoreWebView2ScriptDialogOpeningEventHandler{
		vtbl: &_ICoreWebView2ScriptDialogOpeningEventHandlerFn,
		impl: impl,
	}
}
