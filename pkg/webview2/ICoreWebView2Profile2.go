//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Profile2Vtbl struct {
	_IUnknownVtbl
	ClearBrowsingData            ComProc
	ClearBrowsingDataInTimeRange ComProc
	ClearBrowsingDataAll         ComProc
}

type ICoreWebView2Profile2 struct {
	vtbl *_ICoreWebView2Profile2Vtbl
}

func (i *ICoreWebView2Profile2) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Profile2) ClearBrowsingData(dataKinds COREWEBVIEW2_BROWSING_DATA_KINDS, handler *ICoreWebView2ClearBrowsingDataCompletedHandler) error {
	var err error

	_, _, err = i.vtbl.ClearBrowsingData.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(dataKinds),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Profile2) ClearBrowsingDataInTimeRange(dataKinds COREWEBVIEW2_BROWSING_DATA_KINDS, startTime float64, endTime float64, handler *ICoreWebView2ClearBrowsingDataCompletedHandler) error {
	var err error

	_, _, err = i.vtbl.ClearBrowsingDataInTimeRange.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(dataKinds),
		uintptr(unsafe.Pointer(&startTime)),
		uintptr(unsafe.Pointer(&endTime)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Profile2) ClearBrowsingDataAll(handler *ICoreWebView2ClearBrowsingDataCompletedHandler) error {
	var err error

	_, _, err = i.vtbl.ClearBrowsingDataAll.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
