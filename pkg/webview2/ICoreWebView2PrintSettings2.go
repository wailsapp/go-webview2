//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2PrintSettings2Vtbl struct {
	_IUnknownVtbl
	GetPageRanges   ComProc
	PutPageRanges   ComProc
	GetPagesPerSide ComProc
	PutPagesPerSide ComProc
	GetCopies       ComProc
	PutCopies       ComProc
	GetCollation    ComProc
	PutCollation    ComProc
	GetColorMode    ComProc
	PutColorMode    ComProc
	GetDuplex       ComProc
	PutDuplex       ComProc
	GetMediaSize    ComProc
	PutMediaSize    ComProc
	GetPrinterName  ComProc
	PutPrinterName  ComProc
}

type ICoreWebView2PrintSettings2 struct {
	vtbl *_ICoreWebView2PrintSettings2Vtbl
}

func (i *ICoreWebView2PrintSettings2) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2PrintSettings2) GetPageRanges() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetPageRanges.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	value := windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2PrintSettings2) PutPageRanges(value string) error {
	var err error

	// Convert string 'value' to *uint16
	_value, err := windows.UTF16PtrFromString(value)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PutPageRanges.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings2) GetPagesPerSide() (*INT32, error) {
	var err error

	var value *INT32

	_, _, err = i.vtbl.GetPagesPerSide.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2PrintSettings2) PutPagesPerSide(value INT32) error {
	var err error

	_, _, err = i.vtbl.PutPagesPerSide.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings2) GetCopies() (*INT32, error) {
	var err error

	var value *INT32

	_, _, err = i.vtbl.GetCopies.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2PrintSettings2) PutCopies(value INT32) error {
	var err error

	_, _, err = i.vtbl.PutCopies.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings2) GetCollation() (*COREWEBVIEW2_PRINT_COLLATION, error) {
	var err error

	var value *COREWEBVIEW2_PRINT_COLLATION

	_, _, err = i.vtbl.GetCollation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2PrintSettings2) PutCollation(value COREWEBVIEW2_PRINT_COLLATION) error {
	var err error

	_, _, err = i.vtbl.PutCollation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings2) GetColorMode() (*COREWEBVIEW2_PRINT_COLOR_MODE, error) {
	var err error

	var value *COREWEBVIEW2_PRINT_COLOR_MODE

	_, _, err = i.vtbl.GetColorMode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2PrintSettings2) PutColorMode(value COREWEBVIEW2_PRINT_COLOR_MODE) error {
	var err error

	_, _, err = i.vtbl.PutColorMode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings2) GetDuplex() (*COREWEBVIEW2_PRINT_DUPLEX, error) {
	var err error

	var value *COREWEBVIEW2_PRINT_DUPLEX

	_, _, err = i.vtbl.GetDuplex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2PrintSettings2) PutDuplex(value COREWEBVIEW2_PRINT_DUPLEX) error {
	var err error

	_, _, err = i.vtbl.PutDuplex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings2) GetMediaSize() (*COREWEBVIEW2_PRINT_MEDIA_SIZE, error) {
	var err error

	var value *COREWEBVIEW2_PRINT_MEDIA_SIZE

	_, _, err = i.vtbl.GetMediaSize.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2PrintSettings2) PutMediaSize(value COREWEBVIEW2_PRINT_MEDIA_SIZE) error {
	var err error

	_, _, err = i.vtbl.PutMediaSize.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2PrintSettings2) GetPrinterName() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetPrinterName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	value := windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2PrintSettings2) PutPrinterName(value string) error {
	var err error

	// Convert string 'value' to *uint16
	_value, err := windows.UTF16PtrFromString(value)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PutPrinterName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
