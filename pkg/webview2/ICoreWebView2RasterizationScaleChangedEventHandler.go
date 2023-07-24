//go:build windows

package webview2

type _ICoreWebView2RasterizationScaleChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2RasterizationScaleChangedEventHandler struct {
	vtbl *_ICoreWebView2RasterizationScaleChangedEventHandlerVtbl
	impl _ICoreWebView2RasterizationScaleChangedEventHandlerImpl
}

func (i *ICoreWebView2RasterizationScaleChangedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2RasterizationScaleChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2RasterizationScaleChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2RasterizationScaleChangedEventHandlerIUnknownAddRef(this *ICoreWebView2RasterizationScaleChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2RasterizationScaleChangedEventHandlerIUnknownRelease(this *ICoreWebView2RasterizationScaleChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2RasterizationScaleChangedEventHandlerInvoke(this *ICoreWebView2RasterizationScaleChangedEventHandler, sender *ICoreWebView2Controller, args *_IUnknown) uintptr {
	return this.impl.RasterizationScaleChanged(sender, args)
}

type _ICoreWebView2RasterizationScaleChangedEventHandlerImpl interface {
	_IUnknownImpl
	RasterizationScaleChanged(sender *ICoreWebView2Controller, args *_IUnknown) uintptr
}

var _ICoreWebView2RasterizationScaleChangedEventHandlerFn = _ICoreWebView2RasterizationScaleChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2RasterizationScaleChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2RasterizationScaleChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2RasterizationScaleChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2RasterizationScaleChangedEventHandlerInvoke),
}

func NewICoreWebView2RasterizationScaleChangedEventHandler(impl _ICoreWebView2RasterizationScaleChangedEventHandlerImpl) *ICoreWebView2RasterizationScaleChangedEventHandler {
	return &ICoreWebView2RasterizationScaleChangedEventHandler{
		vtbl: &_ICoreWebView2RasterizationScaleChangedEventHandlerFn,
		impl: impl,
	}
}
