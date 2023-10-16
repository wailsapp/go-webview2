package edge

import "github.com/wailsapp/go-webview2/webviewloader"

type Capability string

var UnsupportedCapabilityError = &unsupportedCapabilityError{}

type unsupportedCapabilityError struct{}

func (u *unsupportedCapabilityError) Error() string {
	return "unsupported capability"
}

// Capabilities is a list of capabilities with their corresponding minimum runtime version
// Internal Capabilities are not exposed to the user
// Larger capabilities such as DragAndDrop should be exported with a capital letter
const (
	GetAdditionalObjects = Capability("113.0.1774.30")
	SwipeNavigation      = Capability("94.0.992.31")
	AllowExternalDrop    = Capability("100.0.1185.39")
)

func hasCapability(webview2RuntimeVersion string, capability Capability) bool {
	result, err := webviewloader.CompareBrowserVersions(webview2RuntimeVersion, string(capability))
	if err != nil {
		return false
	}
	return result >= 0
}
