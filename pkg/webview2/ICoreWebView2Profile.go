//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2ProfileVtbl struct {
	_IUnknownVtbl
	GetProfileName               ComProc
	GetIsInPrivateModeEnabled    ComProc
	GetProfilePath               ComProc
	GetDefaultDownloadFolderPath ComProc
	PutDefaultDownloadFolderPath ComProc
	GetPreferredColorScheme      ComProc
	PutPreferredColorScheme      ComProc
}

type ICoreWebView2Profile struct {
	vtbl *_ICoreWebView2ProfileVtbl
}

func (i *ICoreWebView2Profile) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Profile) GetProfileName() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetProfileName.Call(
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

func (i *ICoreWebView2Profile) GetIsInPrivateModeEnabled() (bool, error) {
	var err error

	var value bool

	_, _, err = i.vtbl.GetIsInPrivateModeEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return value, nil
}

func (i *ICoreWebView2Profile) GetProfilePath() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetProfilePath.Call(
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

func (i *ICoreWebView2Profile) GetDefaultDownloadFolderPath() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetDefaultDownloadFolderPath.Call(
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

func (i *ICoreWebView2Profile) PutDefaultDownloadFolderPath(value string) error {
	var err error

	// Convert string 'value' to *uint16
	_value, err := windows.UTF16PtrFromString(value)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PutDefaultDownloadFolderPath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2Profile) GetPreferredColorScheme() (*COREWEBVIEW2_PREFERRED_COLOR_SCHEME, error) {
	var err error

	var value *COREWEBVIEW2_PREFERRED_COLOR_SCHEME

	_, _, err = i.vtbl.GetPreferredColorScheme.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2Profile) PutPreferredColorScheme(value COREWEBVIEW2_PREFERRED_COLOR_SCHEME) error {
	var err error

	_, _, err = i.vtbl.PutPreferredColorScheme.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
