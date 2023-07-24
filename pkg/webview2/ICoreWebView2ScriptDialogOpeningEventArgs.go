//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2ScriptDialogOpeningEventArgsVtbl struct {
	_IUnknownVtbl
	GetUri         ComProc
	GetKind        ComProc
	GetMessage     ComProc
	Accept         ComProc
	GetDefaultText ComProc
	GetResultText  ComProc
	PutResultText  ComProc
	GetDeferral    ComProc
}

type ICoreWebView2ScriptDialogOpeningEventArgs struct {
	vtbl *_ICoreWebView2ScriptDialogOpeningEventArgsVtbl
}

func (i *ICoreWebView2ScriptDialogOpeningEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2ScriptDialogOpeningEventArgs) GetUri() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _uri *uint16

	_, _, err = i.vtbl.GetUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	uri := windows.UTF16PtrToString(_uri)
	windows.CoTaskMemFree(unsafe.Pointer(_uri))
	return uri, nil
}

func (i *ICoreWebView2ScriptDialogOpeningEventArgs) GetKind() (*COREWEBVIEW2_SCRIPT_DIALOG_KIND, error) {
	var err error

	var kind *COREWEBVIEW2_SCRIPT_DIALOG_KIND

	_, _, err = i.vtbl.GetKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&kind)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return kind, nil
}

func (i *ICoreWebView2ScriptDialogOpeningEventArgs) GetMessage() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _message *uint16

	_, _, err = i.vtbl.GetMessage.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_message)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	message := windows.UTF16PtrToString(_message)
	windows.CoTaskMemFree(unsafe.Pointer(_message))
	return message, nil
}

func (i *ICoreWebView2ScriptDialogOpeningEventArgs) Accept() error {
	var err error

	_, _, err = i.vtbl.Accept.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2ScriptDialogOpeningEventArgs) GetDefaultText() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _defaultText *uint16

	_, _, err = i.vtbl.GetDefaultText.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_defaultText)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	defaultText := windows.UTF16PtrToString(_defaultText)
	windows.CoTaskMemFree(unsafe.Pointer(_defaultText))
	return defaultText, nil
}

func (i *ICoreWebView2ScriptDialogOpeningEventArgs) GetResultText() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _resultText *uint16

	_, _, err = i.vtbl.GetResultText.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_resultText)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	resultText := windows.UTF16PtrToString(_resultText)
	windows.CoTaskMemFree(unsafe.Pointer(_resultText))
	return resultText, nil
}

func (i *ICoreWebView2ScriptDialogOpeningEventArgs) PutResultText(resultText string) error {
	var err error

	// Convert string 'resultText' to *uint16
	_resultText, err := windows.UTF16PtrFromString(resultText)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PutResultText.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_resultText)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2ScriptDialogOpeningEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var err error

	var deferral *ICoreWebView2Deferral

	_, _, err = i.vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&deferral)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return deferral, nil
}
