//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Vtbl struct {
	_IUnknownVtbl
	GetSettings                            ComProc
	GetSource                              ComProc
	Navigate                               ComProc
	NavigateToString                       ComProc
	AddNavigationStarting                  ComProc
	RemoveNavigationStarting               ComProc
	AddContentLoading                      ComProc
	RemoveContentLoading                   ComProc
	AddSourceChanged                       ComProc
	RemoveSourceChanged                    ComProc
	AddHistoryChanged                      ComProc
	RemoveHistoryChanged                   ComProc
	AddNavigationCompleted                 ComProc
	RemoveNavigationCompleted              ComProc
	AddFrameNavigationStarting             ComProc
	RemoveFrameNavigationStarting          ComProc
	AddFrameNavigationCompleted            ComProc
	RemoveFrameNavigationCompleted         ComProc
	AddScriptDialogOpening                 ComProc
	RemoveScriptDialogOpening              ComProc
	AddPermissionRequested                 ComProc
	RemovePermissionRequested              ComProc
	AddProcessFailed                       ComProc
	RemoveProcessFailed                    ComProc
	AddScriptToExecuteOnDocumentCreated    ComProc
	RemoveScriptToExecuteOnDocumentCreated ComProc
	ExecuteScript                          ComProc
	CapturePreview                         ComProc
	Reload                                 ComProc
	PostWebMessageAsJson                   ComProc
	PostWebMessageAsString                 ComProc
	AddWebMessageReceived                  ComProc
	RemoveWebMessageReceived               ComProc
	CallDevToolsProtocolMethod             ComProc
	GetBrowserProcessId                    ComProc
	GetCanGoBack                           ComProc
	GetCanGoForward                        ComProc
	GoBack                                 ComProc
	GoForward                              ComProc
	GetDevToolsProtocolEventReceiver       ComProc
	Stop                                   ComProc
	AddNewWindowRequested                  ComProc
	RemoveNewWindowRequested               ComProc
	AddDocumentTitleChanged                ComProc
	RemoveDocumentTitleChanged             ComProc
	GetDocumentTitle                       ComProc
	AddHostObjectToScript                  ComProc
	RemoveHostObjectFromScript             ComProc
	OpenDevToolsWindow                     ComProc
	AddContainsFullScreenElementChanged    ComProc
	RemoveContainsFullScreenElementChanged ComProc
	GetContainsFullScreenElement           ComProc
	AddWebResourceRequested                ComProc
	RemoveWebResourceRequested             ComProc
	AddWebResourceRequestedFilter          ComProc
	RemoveWebResourceRequestedFilter       ComProc
	AddWindowCloseRequested                ComProc
	RemoveWindowCloseRequested             ComProc
}

type ICoreWebView2 struct {
	vtbl *_ICoreWebView2Vtbl
}

func (i *ICoreWebView2) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2) GetSettings() (*ICoreWebView2Settings, error) {
	var err error

	var settings *ICoreWebView2Settings

	_, _, err = i.vtbl.GetSettings.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&settings)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return settings, nil
}

func (i *ICoreWebView2) GetSource() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _uri *uint16

	_, _, err = i.vtbl.GetSource.Call(
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

func (i *ICoreWebView2) Navigate(uri string) error {
	var err error

	// Convert string 'uri' to *uint16
	_uri, err := windows.UTF16PtrFromString(uri)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.Navigate.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) NavigateToString(htmlContent string) error {
	var err error

	// Convert string 'htmlContent' to *uint16
	_htmlContent, err := windows.UTF16PtrFromString(htmlContent)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.NavigateToString.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_htmlContent)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddNavigationStarting(eventHandler *ICoreWebView2NavigationStartingEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddNavigationStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2) RemoveNavigationStarting(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveNavigationStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddContentLoading(eventHandler *ICoreWebView2ContentLoadingEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddContentLoading.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2) RemoveContentLoading(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveContentLoading.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddSourceChanged(eventHandler *ICoreWebView2SourceChangedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddSourceChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2) RemoveSourceChanged(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveSourceChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddHistoryChanged(eventHandler *ICoreWebView2HistoryChangedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddHistoryChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2) RemoveHistoryChanged(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveHistoryChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddNavigationCompleted(eventHandler *ICoreWebView2NavigationCompletedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddNavigationCompleted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2) RemoveNavigationCompleted(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveNavigationCompleted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddFrameNavigationStarting(eventHandler *ICoreWebView2NavigationStartingEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddFrameNavigationStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2) RemoveFrameNavigationStarting(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveFrameNavigationStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddFrameNavigationCompleted(eventHandler *ICoreWebView2NavigationCompletedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddFrameNavigationCompleted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2) RemoveFrameNavigationCompleted(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveFrameNavigationCompleted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddScriptDialogOpening(eventHandler *ICoreWebView2ScriptDialogOpeningEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddScriptDialogOpening.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2) RemoveScriptDialogOpening(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveScriptDialogOpening.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddPermissionRequested(eventHandler *ICoreWebView2PermissionRequestedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddPermissionRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2) RemovePermissionRequested(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemovePermissionRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddProcessFailed(eventHandler *ICoreWebView2ProcessFailedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddProcessFailed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2) RemoveProcessFailed(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveProcessFailed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddScriptToExecuteOnDocumentCreated(javaScript string, handler *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler) error {
	var err error

	// Convert string 'javaScript' to *uint16
	_javaScript, err := windows.UTF16PtrFromString(javaScript)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.AddScriptToExecuteOnDocumentCreated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_javaScript)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveScriptToExecuteOnDocumentCreated(id string) error {
	var err error

	// Convert string 'id' to *uint16
	_id, err := windows.UTF16PtrFromString(id)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.RemoveScriptToExecuteOnDocumentCreated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_id)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) ExecuteScript(javaScript string, handler *ICoreWebView2ExecuteScriptCompletedHandler) error {
	var err error

	// Convert string 'javaScript' to *uint16
	_javaScript, err := windows.UTF16PtrFromString(javaScript)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.ExecuteScript.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_javaScript)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) CapturePreview(imageFormat COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT, imageStream *IStream, handler *ICoreWebView2CapturePreviewCompletedHandler) error {
	var err error

	_, _, err = i.vtbl.CapturePreview.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(imageFormat),
		uintptr(unsafe.Pointer(imageStream)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) Reload() error {
	var err error

	_, _, err = i.vtbl.Reload.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) PostWebMessageAsJson(webMessageAsJson string) error {
	var err error

	// Convert string 'webMessageAsJson' to *uint16
	_webMessageAsJson, err := windows.UTF16PtrFromString(webMessageAsJson)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PostWebMessageAsJson.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_webMessageAsJson)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) PostWebMessageAsString(webMessageAsString string) error {
	var err error

	// Convert string 'webMessageAsString' to *uint16
	_webMessageAsString, err := windows.UTF16PtrFromString(webMessageAsString)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.PostWebMessageAsString.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_webMessageAsString)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddWebMessageReceived(handler *ICoreWebView2WebMessageReceivedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddWebMessageReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2) RemoveWebMessageReceived(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveWebMessageReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) CallDevToolsProtocolMethod(methodName string, parametersAsJson string, handler *ICoreWebView2CallDevToolsProtocolMethodCompletedHandler) error {
	var err error

	// Convert string 'methodName' to *uint16
	_methodName, err := windows.UTF16PtrFromString(methodName)
	if err != nil {
		return err
	}

	// Convert string 'parametersAsJson' to *uint16
	_parametersAsJson, err := windows.UTF16PtrFromString(parametersAsJson)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.CallDevToolsProtocolMethod.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_methodName)),
		uintptr(unsafe.Pointer(_parametersAsJson)),
		uintptr(unsafe.Pointer(handler)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) GetBrowserProcessId() (*uint32, error) {
	var err error

	var value *uint32

	_, _, err = i.vtbl.GetBrowserProcessId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}

func (i *ICoreWebView2) GetCanGoBack() (bool, error) {
	var err error

	var canGoBack bool

	_, _, err = i.vtbl.GetCanGoBack.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&canGoBack)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return canGoBack, nil
}

func (i *ICoreWebView2) GetCanGoForward() (bool, error) {
	var err error

	var canGoForward bool

	_, _, err = i.vtbl.GetCanGoForward.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&canGoForward)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return canGoForward, nil
}

func (i *ICoreWebView2) GoBack() error {
	var err error

	_, _, err = i.vtbl.GoBack.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) GoForward() error {
	var err error

	_, _, err = i.vtbl.GoForward.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) GetDevToolsProtocolEventReceiver(eventName string) (*ICoreWebView2DevToolsProtocolEventReceiver, error) {
	var err error

	// Convert string 'eventName' to *uint16
	_eventName, err := windows.UTF16PtrFromString(eventName)
	if err != nil {
		return nil, err
	}

	var receiver *ICoreWebView2DevToolsProtocolEventReceiver

	_, _, err = i.vtbl.GetDevToolsProtocolEventReceiver.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_eventName)),
		uintptr(unsafe.Pointer(&receiver)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return receiver, nil
}

func (i *ICoreWebView2) Stop() error {
	var err error

	_, _, err = i.vtbl.Stop.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddNewWindowRequested(eventHandler *ICoreWebView2NewWindowRequestedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddNewWindowRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2) RemoveNewWindowRequested(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveNewWindowRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddDocumentTitleChanged(eventHandler *ICoreWebView2DocumentTitleChangedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddDocumentTitleChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2) RemoveDocumentTitleChanged(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveDocumentTitleChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) GetDocumentTitle() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _title *uint16

	_, _, err = i.vtbl.GetDocumentTitle.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_title)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	title := windows.UTF16PtrToString(_title)
	windows.CoTaskMemFree(unsafe.Pointer(_title))
	return title, nil
}

func (i *ICoreWebView2) AddHostObjectToScript(name string, object *VARIANT) error {
	var err error

	// Convert string 'name' to *uint16
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.AddHostObjectToScript.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(object)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveHostObjectFromScript(name string) error {
	var err error

	// Convert string 'name' to *uint16
	_name, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.RemoveHostObjectFromScript.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) OpenDevToolsWindow() error {
	var err error

	_, _, err = i.vtbl.OpenDevToolsWindow.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddContainsFullScreenElementChanged(eventHandler *ICoreWebView2ContainsFullScreenElementChangedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddContainsFullScreenElementChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2) RemoveContainsFullScreenElementChanged(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveContainsFullScreenElementChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) GetContainsFullScreenElement() (bool, error) {
	var err error

	var containsFullScreenElement bool

	_, _, err = i.vtbl.GetContainsFullScreenElement.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&containsFullScreenElement)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return containsFullScreenElement, nil
}

func (i *ICoreWebView2) AddWebResourceRequested(eventHandler *ICoreWebView2WebResourceRequestedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddWebResourceRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2) RemoveWebResourceRequested(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveWebResourceRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddWebResourceRequestedFilter(uri string, resourceContext COREWEBVIEW2_WEB_RESOURCE_CONTEXT) error {
	var err error

	// Convert string 'uri' to *uint16
	_uri, err := windows.UTF16PtrFromString(uri)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.AddWebResourceRequestedFilter.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
		uintptr(resourceContext),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) RemoveWebResourceRequestedFilter(uri string, resourceContext COREWEBVIEW2_WEB_RESOURCE_CONTEXT) error {
	var err error

	// Convert string 'uri' to *uint16
	_uri, err := windows.UTF16PtrFromString(uri)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.RemoveWebResourceRequestedFilter.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
		uintptr(resourceContext),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2) AddWindowCloseRequested(eventHandler *ICoreWebView2WindowCloseRequestedEventHandler) (*EventRegistrationToken, error) {
	var err error

	var token *EventRegistrationToken

	_, _, err = i.vtbl.AddWindowCloseRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return token, nil
}

func (i *ICoreWebView2) RemoveWindowCloseRequested(token EventRegistrationToken) error {
	var err error

	_, _, err = i.vtbl.RemoveWindowCloseRequested.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
