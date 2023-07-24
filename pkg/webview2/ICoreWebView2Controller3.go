//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Controller3Vtbl struct {
	_IUnknownVtbl
	GetRasterizationScale              ComProc
	PutRasterizationScale              ComProc
	GetShouldDetectMonitorScaleChanges ComProc
	PutShouldDetectMonitorScaleChanges ComProc
	AddRasterizationScaleChanged       ComProc
	RemoveRasterizationScaleChanged    ComProc
	GetBoundsMode                      ComProc
	PutBoundsMode                      ComProc
}

type ICoreWebView2Controller3 struct {
	vtbl *_ICoreWebView2Controller3Vtbl
}

func (i *ICoreWebView2Controller3) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Controller3) GetRasterizationScale() (float64, error) {
	var err error

	var scale float64

	_, _, err = i.vtbl.GetRasterizationScale.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&scale)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0.0, err
	}
	return scale, nil
}

func (i *ICoreWebView2Controller3) PutRasterizationScale(scale float64) error {
	var err error

	_, _, err = i.vtbl.PutRasterizationScale.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&scale)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller3) GetShouldDetectMonitorScaleChanges() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetShouldDetectMonitorScaleChanges.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2Controller3) PutShouldDetectMonitorScaleChanges(value bool) error {
	var err error

	_, _, err = i.vtbl.PutShouldDetectMonitorScaleChanges.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller3) AddRasterizationScaleChanged(eventHandler *ICoreWebView2RasterizationScaleChangedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddRasterizationScaleChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2Controller3) RemoveRasterizationScaleChanged(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveRasterizationScaleChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Controller3) GetBoundsMode() (*COREWEBVIEW2_BOUNDS_MODE, error) {
	var err error

	var boundsMode *COREWEBVIEW2_BOUNDS_MODE

	_, _, err = i.vtbl.GetBoundsMode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&boundsMode)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return boundsMode, nil
}

func (i *ICoreWebView2Controller3) PutBoundsMode(boundsMode COREWEBVIEW2_BOUNDS_MODE) error {
	var err error

	_, _, err = i.vtbl.PutBoundsMode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(boundsMode),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
