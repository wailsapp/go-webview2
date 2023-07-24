//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2PointerInfoVtbl struct {
	_IUnknownVtbl
	GetPointerKind         ComProc
	PutPointerKind         ComProc
	GetPointerId           ComProc
	PutPointerId           ComProc
	GetFrameId             ComProc
	PutFrameId             ComProc
	GetPointerFlags        ComProc
	PutPointerFlags        ComProc
	GetPointerDeviceRect   ComProc
	PutPointerDeviceRect   ComProc
	GetDisplayRect         ComProc
	PutDisplayRect         ComProc
	GetPixelLocation       ComProc
	PutPixelLocation       ComProc
	GetHimetricLocation    ComProc
	PutHimetricLocation    ComProc
	GetPixelLocationRaw    ComProc
	PutPixelLocationRaw    ComProc
	GetHimetricLocationRaw ComProc
	PutHimetricLocationRaw ComProc
	GetTime                ComProc
	PutTime                ComProc
	GetHistoryCount        ComProc
	PutHistoryCount        ComProc
	GetInputData           ComProc
	PutInputData           ComProc
	GetKeyStates           ComProc
	PutKeyStates           ComProc
	GetPerformanceCount    ComProc
	PutPerformanceCount    ComProc
	GetButtonChangeKind    ComProc
	PutButtonChangeKind    ComProc
	GetPenFlags            ComProc
	PutPenFlags            ComProc
	GetPenMask             ComProc
	PutPenMask             ComProc
	GetPenPressure         ComProc
	PutPenPressure         ComProc
	GetPenRotation         ComProc
	PutPenRotation         ComProc
	GetPenTiltX            ComProc
	PutPenTiltX            ComProc
	GetPenTiltY            ComProc
	PutPenTiltY            ComProc
	GetTouchFlags          ComProc
	PutTouchFlags          ComProc
	GetTouchMask           ComProc
	PutTouchMask           ComProc
	GetTouchContact        ComProc
	PutTouchContact        ComProc
	GetTouchContactRaw     ComProc
	PutTouchContactRaw     ComProc
	GetTouchOrientation    ComProc
	PutTouchOrientation    ComProc
	GetTouchPressure       ComProc
	PutTouchPressure       ComProc
}

type ICoreWebView2PointerInfo struct {
	vtbl *_ICoreWebView2PointerInfoVtbl
}

func (i *ICoreWebView2PointerInfo) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2PointerInfo) GetPointerKind() (*DWORD, error) {
	var err error

	var pointerKind *DWORD

	_, _, err = i.vtbl.GetPointerKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pointerKind)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return pointerKind, nil
}

func (i *ICoreWebView2PointerInfo) PutPointerKind(pointerKind DWORD) error {
	var err error

	_, _, err = i.vtbl.PutPointerKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pointerKind)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetPointerId() (*uint32, error) {
	var err error

	var pointerId *uint32

	_, _, err = i.vtbl.GetPointerId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pointerId)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return pointerId, nil
}

func (i *ICoreWebView2PointerInfo) PutPointerId(pointerId uint32) error {
	var err error

	_, _, err = i.vtbl.PutPointerId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pointerId)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetFrameId() (*uint32, error) {
	var err error

	var frameId *uint32

	_, _, err = i.vtbl.GetFrameId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&frameId)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return frameId, nil
}

func (i *ICoreWebView2PointerInfo) PutFrameId(frameId uint32) error {
	var err error

	_, _, err = i.vtbl.PutFrameId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&frameId)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetPointerFlags() (*uint32, error) {
	var err error

	var pointerFlags *uint32

	_, _, err = i.vtbl.GetPointerFlags.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pointerFlags)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return pointerFlags, nil
}

func (i *ICoreWebView2PointerInfo) PutPointerFlags(pointerFlags uint32) error {
	var err error

	_, _, err = i.vtbl.PutPointerFlags.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pointerFlags)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetPointerDeviceRect() (*RECT, error) {
	var err error

	var pointerDeviceRect *RECT

	_, _, err = i.vtbl.GetPointerDeviceRect.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pointerDeviceRect)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return pointerDeviceRect, nil
}

func (i *ICoreWebView2PointerInfo) PutPointerDeviceRect(pointerDeviceRect RECT) error {
	var err error

	_, _, err = i.vtbl.PutPointerDeviceRect.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pointerDeviceRect)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetDisplayRect() (*RECT, error) {
	var err error

	var displayRect *RECT

	_, _, err = i.vtbl.GetDisplayRect.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&displayRect)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return displayRect, nil
}

func (i *ICoreWebView2PointerInfo) PutDisplayRect(displayRect RECT) error {
	var err error

	_, _, err = i.vtbl.PutDisplayRect.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&displayRect)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetPixelLocation() (*POINT, error) {
	var err error

	var pixelLocation *POINT

	_, _, err = i.vtbl.GetPixelLocation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pixelLocation)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return pixelLocation, nil
}

func (i *ICoreWebView2PointerInfo) PutPixelLocation(pixelLocation POINT) error {
	var err error

	_, _, err = i.vtbl.PutPixelLocation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pixelLocation)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetHimetricLocation() (*POINT, error) {
	var err error

	var himetricLocation *POINT

	_, _, err = i.vtbl.GetHimetricLocation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&himetricLocation)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return himetricLocation, nil
}

func (i *ICoreWebView2PointerInfo) PutHimetricLocation(himetricLocation POINT) error {
	var err error

	_, _, err = i.vtbl.PutHimetricLocation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&himetricLocation)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetPixelLocationRaw() (*POINT, error) {
	var err error

	var pixelLocationRaw *POINT

	_, _, err = i.vtbl.GetPixelLocationRaw.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pixelLocationRaw)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return pixelLocationRaw, nil
}

func (i *ICoreWebView2PointerInfo) PutPixelLocationRaw(pixelLocationRaw POINT) error {
	var err error

	_, _, err = i.vtbl.PutPixelLocationRaw.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pixelLocationRaw)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetHimetricLocationRaw() (*POINT, error) {
	var err error

	var himetricLocationRaw *POINT

	_, _, err = i.vtbl.GetHimetricLocationRaw.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&himetricLocationRaw)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return himetricLocationRaw, nil
}

func (i *ICoreWebView2PointerInfo) PutHimetricLocationRaw(himetricLocationRaw POINT) error {
	var err error

	_, _, err = i.vtbl.PutHimetricLocationRaw.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&himetricLocationRaw)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetTime() (*DWORD, error) {
	var err error

	var time *DWORD

	_, _, err = i.vtbl.GetTime.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&time)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return time, nil
}

func (i *ICoreWebView2PointerInfo) PutTime(time DWORD) error {
	var err error

	_, _, err = i.vtbl.PutTime.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&time)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetHistoryCount() (*uint32, error) {
	var err error

	var historyCount *uint32

	_, _, err = i.vtbl.GetHistoryCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&historyCount)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return historyCount, nil
}

func (i *ICoreWebView2PointerInfo) PutHistoryCount(historyCount uint32) error {
	var err error

	_, _, err = i.vtbl.PutHistoryCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&historyCount)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetInputData() (*INT32, error) {
	var err error

	var inputData *INT32

	_, _, err = i.vtbl.GetInputData.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&inputData)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return inputData, nil
}

func (i *ICoreWebView2PointerInfo) PutInputData(inputData INT32) error {
	var err error

	_, _, err = i.vtbl.PutInputData.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&inputData)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetKeyStates() (*DWORD, error) {
	var err error

	var keyStates *DWORD

	_, _, err = i.vtbl.GetKeyStates.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&keyStates)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return keyStates, nil
}

func (i *ICoreWebView2PointerInfo) PutKeyStates(keyStates DWORD) error {
	var err error

	_, _, err = i.vtbl.PutKeyStates.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&keyStates)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetPerformanceCount() (*uint64, error) {
	var err error

	var performanceCount *uint64

	_, _, err = i.vtbl.GetPerformanceCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&performanceCount)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return performanceCount, nil
}

func (i *ICoreWebView2PointerInfo) PutPerformanceCount(performanceCount uint64) error {
	var err error

	_, _, err = i.vtbl.PutPerformanceCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&performanceCount)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetButtonChangeKind() (*INT32, error) {
	var err error

	var buttonChangeKind *INT32

	_, _, err = i.vtbl.GetButtonChangeKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&buttonChangeKind)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return buttonChangeKind, nil
}

func (i *ICoreWebView2PointerInfo) PutButtonChangeKind(buttonChangeKind INT32) error {
	var err error

	_, _, err = i.vtbl.PutButtonChangeKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&buttonChangeKind)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetPenFlags() (*uint32, error) {
	var err error

	var penFLags *uint32

	_, _, err = i.vtbl.GetPenFlags.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&penFLags)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return penFLags, nil
}

func (i *ICoreWebView2PointerInfo) PutPenFlags(penFLags uint32) error {
	var err error

	_, _, err = i.vtbl.PutPenFlags.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&penFLags)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetPenMask() (*uint32, error) {
	var err error

	var penMask *uint32

	_, _, err = i.vtbl.GetPenMask.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&penMask)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return penMask, nil
}

func (i *ICoreWebView2PointerInfo) PutPenMask(penMask uint32) error {
	var err error

	_, _, err = i.vtbl.PutPenMask.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&penMask)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetPenPressure() (*uint32, error) {
	var err error

	var penPressure *uint32

	_, _, err = i.vtbl.GetPenPressure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&penPressure)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return penPressure, nil
}

func (i *ICoreWebView2PointerInfo) PutPenPressure(penPressure uint32) error {
	var err error

	_, _, err = i.vtbl.PutPenPressure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&penPressure)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetPenRotation() (*uint32, error) {
	var err error

	var penRotation *uint32

	_, _, err = i.vtbl.GetPenRotation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&penRotation)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return penRotation, nil
}

func (i *ICoreWebView2PointerInfo) PutPenRotation(penRotation uint32) error {
	var err error

	_, _, err = i.vtbl.PutPenRotation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&penRotation)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetPenTiltX() (*INT32, error) {
	var err error

	var penTiltX *INT32

	_, _, err = i.vtbl.GetPenTiltX.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&penTiltX)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return penTiltX, nil
}

func (i *ICoreWebView2PointerInfo) PutPenTiltX(penTiltX INT32) error {
	var err error

	_, _, err = i.vtbl.PutPenTiltX.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&penTiltX)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetPenTiltY() (*INT32, error) {
	var err error

	var penTiltY *INT32

	_, _, err = i.vtbl.GetPenTiltY.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&penTiltY)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return penTiltY, nil
}

func (i *ICoreWebView2PointerInfo) PutPenTiltY(penTiltY INT32) error {
	var err error

	_, _, err = i.vtbl.PutPenTiltY.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&penTiltY)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetTouchFlags() (*uint32, error) {
	var err error

	var touchFlags *uint32

	_, _, err = i.vtbl.GetTouchFlags.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&touchFlags)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return touchFlags, nil
}

func (i *ICoreWebView2PointerInfo) PutTouchFlags(touchFlags uint32) error {
	var err error

	_, _, err = i.vtbl.PutTouchFlags.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&touchFlags)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetTouchMask() (*uint32, error) {
	var err error

	var touchMask *uint32

	_, _, err = i.vtbl.GetTouchMask.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&touchMask)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return touchMask, nil
}

func (i *ICoreWebView2PointerInfo) PutTouchMask(touchMask uint32) error {
	var err error

	_, _, err = i.vtbl.PutTouchMask.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&touchMask)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetTouchContact() (*RECT, error) {
	var err error

	var touchContact *RECT

	_, _, err = i.vtbl.GetTouchContact.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&touchContact)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return touchContact, nil
}

func (i *ICoreWebView2PointerInfo) PutTouchContact(touchContact RECT) error {
	var err error

	_, _, err = i.vtbl.PutTouchContact.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&touchContact)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetTouchContactRaw() (*RECT, error) {
	var err error

	var touchContactRaw *RECT

	_, _, err = i.vtbl.GetTouchContactRaw.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&touchContactRaw)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return touchContactRaw, nil
}

func (i *ICoreWebView2PointerInfo) PutTouchContactRaw(touchContactRaw RECT) error {
	var err error

	_, _, err = i.vtbl.PutTouchContactRaw.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&touchContactRaw)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetTouchOrientation() (*uint32, error) {
	var err error

	var touchOrientation *uint32

	_, _, err = i.vtbl.GetTouchOrientation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&touchOrientation)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return touchOrientation, nil
}

func (i *ICoreWebView2PointerInfo) PutTouchOrientation(touchOrientation uint32) error {
	var err error

	_, _, err = i.vtbl.PutTouchOrientation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&touchOrientation)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PointerInfo) GetTouchPressure() (*uint32, error) {
	var err error

	var touchPressure *uint32

	_, _, err = i.vtbl.GetTouchPressure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&touchPressure)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return touchPressure, nil
}

func (i *ICoreWebView2PointerInfo) PutTouchPressure(touchPressure uint32) error {
	var err error

	_, _, err = i.vtbl.PutTouchPressure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&touchPressure)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
