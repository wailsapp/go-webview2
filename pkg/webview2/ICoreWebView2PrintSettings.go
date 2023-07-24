//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2PrintSettingsVtbl struct {
	_IUnknownVtbl
	GetOrientation                ComProc
	PutOrientation                ComProc
	GetScaleFactor                ComProc
	PutScaleFactor                ComProc
	GetPageWidth                  ComProc
	PutPageWidth                  ComProc
	GetPageHeight                 ComProc
	PutPageHeight                 ComProc
	GetMarginTop                  ComProc
	PutMarginTop                  ComProc
	GetMarginBottom               ComProc
	PutMarginBottom               ComProc
	GetMarginLeft                 ComProc
	PutMarginLeft                 ComProc
	GetMarginRight                ComProc
	PutMarginRight                ComProc
	GetShouldPrintBackgrounds     ComProc
	PutShouldPrintBackgrounds     ComProc
	GetShouldPrintSelectionOnly   ComProc
	PutShouldPrintSelectionOnly   ComProc
	GetShouldPrintHeaderAndFooter ComProc
	PutShouldPrintHeaderAndFooter ComProc
	GetHeaderTitle                ComProc
	PutHeaderTitle                ComProc
	GetFooterUri                  ComProc
	PutFooterUri                  ComProc
}

type ICoreWebView2PrintSettings struct {
	vtbl *_ICoreWebView2PrintSettingsVtbl
}

func (i *ICoreWebView2PrintSettings) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2PrintSettings) GetOrientation() (*COREWEBVIEW2_PRINT_ORIENTATION, error) {
	var err error

	var orientation *COREWEBVIEW2_PRINT_ORIENTATION

	_, _, err = i.vtbl.GetOrientation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&orientation)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return orientation, nil
}

func (i *ICoreWebView2PrintSettings) PutOrientation(orientation COREWEBVIEW2_PRINT_ORIENTATION) error {
	var err error

	_, _, err = i.vtbl.PutOrientation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(orientation),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings) GetScaleFactor() (float64, error) {
	var err error

	var scaleFactor float64

	_, _, err = i.vtbl.GetScaleFactor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&scaleFactor)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0.0, err
	}
	return scaleFactor, nil
}

func (i *ICoreWebView2PrintSettings) PutScaleFactor(scaleFactor float64) error {
	var err error

	_, _, err = i.vtbl.PutScaleFactor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&scaleFactor)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings) GetPageWidth() (float64, error) {
	var err error

	var pageWidth float64

	_, _, err = i.vtbl.GetPageWidth.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pageWidth)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0.0, err
	}
	return pageWidth, nil
}

func (i *ICoreWebView2PrintSettings) PutPageWidth(pageWidth float64) error {
	var err error

	_, _, err = i.vtbl.PutPageWidth.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pageWidth)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings) GetPageHeight() (float64, error) {
	var err error

	var pageHeight float64

	_, _, err = i.vtbl.GetPageHeight.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pageHeight)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0.0, err
	}
	return pageHeight, nil
}

func (i *ICoreWebView2PrintSettings) PutPageHeight(pageHeight float64) error {
	var err error

	_, _, err = i.vtbl.PutPageHeight.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pageHeight)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings) GetMarginTop() (float64, error) {
	var err error

	var marginTop float64

	_, _, err = i.vtbl.GetMarginTop.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&marginTop)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0.0, err
	}
	return marginTop, nil
}

func (i *ICoreWebView2PrintSettings) PutMarginTop(marginTop float64) error {
	var err error

	_, _, err = i.vtbl.PutMarginTop.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&marginTop)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings) GetMarginBottom() (float64, error) {
	var err error

	var marginBottom float64

	_, _, err = i.vtbl.GetMarginBottom.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&marginBottom)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0.0, err
	}
	return marginBottom, nil
}

func (i *ICoreWebView2PrintSettings) PutMarginBottom(marginBottom float64) error {
	var err error

	_, _, err = i.vtbl.PutMarginBottom.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&marginBottom)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings) GetMarginLeft() (float64, error) {
	var err error

	var marginLeft float64

	_, _, err = i.vtbl.GetMarginLeft.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&marginLeft)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0.0, err
	}
	return marginLeft, nil
}

func (i *ICoreWebView2PrintSettings) PutMarginLeft(marginLeft float64) error {
	var err error

	_, _, err = i.vtbl.PutMarginLeft.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&marginLeft)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings) GetMarginRight() (float64, error) {
	var err error

	var marginRight float64

	_, _, err = i.vtbl.GetMarginRight.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&marginRight)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0.0, err
	}
	return marginRight, nil
}

func (i *ICoreWebView2PrintSettings) PutMarginRight(marginRight float64) error {
	var err error

	_, _, err = i.vtbl.PutMarginRight.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&marginRight)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings) GetShouldPrintBackgrounds() (bool, error) {
	var err error

	var shouldPrintBackgrounds bool

	_, _, err = i.vtbl.GetShouldPrintBackgrounds.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&shouldPrintBackgrounds)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return shouldPrintBackgrounds, nil
}

func (i *ICoreWebView2PrintSettings) PutShouldPrintBackgrounds(shouldPrintBackgrounds bool) error {
	var err error

	_, _, err = i.vtbl.PutShouldPrintBackgrounds.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&shouldPrintBackgrounds)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings) GetShouldPrintSelectionOnly() (bool, error) {
	var err error

	var shouldPrintSelectionOnly bool

	_, _, err = i.vtbl.GetShouldPrintSelectionOnly.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&shouldPrintSelectionOnly)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return shouldPrintSelectionOnly, nil
}

func (i *ICoreWebView2PrintSettings) PutShouldPrintSelectionOnly(shouldPrintSelectionOnly bool) error {
	var err error

	_, _, err = i.vtbl.PutShouldPrintSelectionOnly.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&shouldPrintSelectionOnly)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings) GetShouldPrintHeaderAndFooter() (bool, error) {
	var err error

	var shouldPrintHeaderAndFooter bool

	_, _, err = i.vtbl.GetShouldPrintHeaderAndFooter.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&shouldPrintHeaderAndFooter)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return shouldPrintHeaderAndFooter, nil
}

func (i *ICoreWebView2PrintSettings) PutShouldPrintHeaderAndFooter(shouldPrintHeaderAndFooter bool) error {
	var err error

	_, _, err = i.vtbl.PutShouldPrintHeaderAndFooter.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&shouldPrintHeaderAndFooter)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings) GetHeaderTitle() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _headerTitle *uint16

	_, _, err = i.vtbl.GetHeaderTitle.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_headerTitle)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	headerTitle := windows.UTF16PtrToString(_headerTitle)
	windows.CoTaskMemFree(unsafe.Pointer(_headerTitle))
	return headerTitle, nil
}

func (i *ICoreWebView2PrintSettings) PutHeaderTitle(headerTitle string) error {
	var err error

	// Convert string 'headerTitle' to *uint16
	_headerTitle, err := windows.UTF16PtrFromString(headerTitle)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PutHeaderTitle.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_headerTitle)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings) GetFooterUri() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _footerUri *uint16

	_, _, err = i.vtbl.GetFooterUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_footerUri)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	footerUri := windows.UTF16PtrToString(_footerUri)
	windows.CoTaskMemFree(unsafe.Pointer(_footerUri))
	return footerUri, nil
}

func (i *ICoreWebView2PrintSettings) PutFooterUri(footerUri string) error {
	var err error

	// Convert string 'footerUri' to *uint16
	_footerUri, err := windows.UTF16PtrFromString(footerUri)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PutFooterUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_footerUri)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
