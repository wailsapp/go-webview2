//go:build windows

package main

import (
	"fmt"

	"github.com/wailsapp/go-webview2/pkg/edge"
)

// Example showing how to disable pinch zoom in WebView2
func main() {
	fmt.Println("WebView2 Pinch Zoom Control Example")
	fmt.Println("===================================")

	// Create a WebView2 instance
	chromium := edge.NewChromium()

	// Method 1: Disable pinch zoom during initialization
	// This is the recommended approach as it ensures the setting
	// is applied as soon as the WebView2 is ready
	pinchZoomEnabled := false
	chromium.PinchZoomEnabled = &pinchZoomEnabled
	
	fmt.Println("✓ Pinch zoom disabled during initialization")

	// Set up other properties if needed
	chromium.Debug = true

	// In a real application, you would:
	// 1. Create a window (hwnd)
	// 2. Call chromium.Embed(hwnd) to initialize WebView2
	// 3. The pinch zoom setting will be automatically applied

	// Method 2: Disable pinch zoom after initialization (runtime control)
	// This can be used to enable/disable pinch zoom dynamically
	
	// Note: This would only work after chromium.Embed() is called
	// err := chromium.PutIsPinchZoomEnabled(false)
	// if err != nil {
	//     log.Printf("Error disabling pinch zoom: %v", err)
	// }

	// Method 3: Check current pinch zoom setting
	// Note: This would only work after chromium.Embed() is called
	// enabled, err := chromium.GetIsPinchZoomEnabled()
	// if err != nil {
	//     log.Printf("Error getting pinch zoom status: %v", err)
	// } else {
	//     fmt.Printf("Pinch zoom is currently: %t\n", enabled)
	// }

	fmt.Println("\nUsage Examples:")
	fmt.Println("1. Set before initialization: chromium.PinchZoomEnabled = &false")
	fmt.Println("2. Set after initialization: chromium.PutIsPinchZoomEnabled(false)")
	fmt.Println("3. Check current setting: enabled, err := chromium.GetIsPinchZoomEnabled()")
	
	fmt.Println("\nThis setting prevents users from zooming the webpage using:")
	fmt.Println("- Ctrl + mouse wheel")
	fmt.Println("- Pinch gestures on touchpads/touchscreens")
	fmt.Println("- Keyboard shortcuts like Ctrl+Plus/Ctrl+Minus")
}