//go:build windows

package edge

import (
	"unsafe"

	"github.com/wailsapp/go-webview2/pkg/webview2"
)

// PutIsReputationCheckingRequired controls SmartScreen through the generic
// Settings8 binding, releasing both acquired settings interfaces on every path.
func (e *Chromium) PutIsReputationCheckingRequired(required bool) error {
	settings, err := e.GetSettings()
	if err != nil {
		return err
	}
	defer settings.Release()
	reputation, err := (*webview2.ICoreWebView2Settings)(unsafe.Pointer(settings)).GetICoreWebView2Settings8()
	if err != nil {
		return err
	}
	defer reputation.Vtbl.CallRelease(unsafe.Pointer(reputation))
	return reputation.PutIsReputationCheckingRequired(required)
}
