//go:build windows

package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type iCoreWebView2WebMessageReceivedEventArgsVtbl struct {
	_IUnknownVtbl
	GetSource                ComProc
	GetWebMessageAsJSON      ComProc
	TryGetWebMessageAsString ComProc
	GetAdditionalObjects     ComProc
}

type ICoreWebView2WebMessageReceivedEventArgs struct {
	vtbl *iCoreWebView2WebMessageReceivedEventArgsVtbl
}

func (i *ICoreWebView2WebMessageReceivedEventArgs) GetAdditionalObjects() (*ICoreWebView2ObjectCollectionView, error) {
	var err error
	var value *ICoreWebView2ObjectCollectionView
	_, _, err = i.vtbl.GetAdditionalObjects.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}
