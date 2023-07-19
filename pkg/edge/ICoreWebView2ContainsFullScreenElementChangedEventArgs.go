//go:build windows

package edge

type _ICoreWebView2ContainsFullScreenElementChangedEventArgsVtbl struct {
	_IUnknownVtbl
}

type ICoreWebView2ContainsFullScreenElementChangedEventArgs struct {
	vtbl *_ICoreWebView2ContainsFullScreenElementChangedEventArgsVtbl
}

func (i *ICoreWebView2ContainsFullScreenElementChangedEventArgs) AddRef() uintptr {
	return i.AddRef()
}
