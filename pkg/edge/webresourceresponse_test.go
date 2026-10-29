//go:build windows

package edge

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/wailsapp/go-webview2/internal/w32"
	"golang.org/x/sys/windows"
)

func TestCreateWebResourceResponseContent(t *testing.T) {
	err := windows.CoInitializeEx(0, windows.COINIT_APARTMENTTHREADED)
	if err != nil {
		t.Fatalf("Failed to initialize COM: %v", err)
	}
	defer windows.CoUninitialize()

	var hinstance windows.Handle
	_ = windows.GetModuleHandleEx(0, nil, &hinstance)

	className, _ := windows.UTF16PtrFromString("WebView2ResponseTest")
	wc := w32.WndClassExW{
		CbSize:        uint32(unsafe.Sizeof(w32.WndClassExW{})),
		HInstance:     hinstance,
		LpszClassName: className,
		LpfnWndProc:   windows.NewCallback(w32.DefWindowProc),
	}
	_, _, _ = w32.User32RegisterClassExW.Call(uintptr(unsafe.Pointer(&wc)))

	windowName, _ := windows.UTF16PtrFromString("WebView2 Response Test Window")
	hwnd, _, _ := w32.User32CreateWindowExW.Call(
		0,
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)),
		0xCF0000, // WS_OVERLAPPEDWINDOW
		uintptr(w32.CW_USEDEFAULT),
		uintptr(w32.CW_USEDEFAULT),
		640,
		480,
		0,
		0,
		uintptr(hinstance),
		0,
	)
	if hwnd == 0 {
		t.Fatal("Failed to create window")
	}
	defer w32.DestroyWindow(hwnd)

	chromium := NewChromium()
	require.NotNil(t, chromium, "Chromium instance should not be nil")
	require.True(t, chromium.Embed(hwnd), "WebView2 initialization should succeed")

	env := chromium.Environment()
	require.NotNil(t, env, "Environment should not be nil")

	content := []byte("<!doctype html><title>response test</title><p>body</p>")
	response, err := env.CreateWebResourceResponse(content, 200, "OK", "Content-Type: text/html")
	require.NoError(t, err, "Should create response without error")
	require.NotNil(t, response, "Response should not be nil")
	defer response.Release()

	var stream *IStream
	hr, _, _ := response.vtbl.GetContent.Call(
		uintptr(unsafe.Pointer(response)),
		uintptr(unsafe.Pointer(&stream)),
	)
	require.Equal(t, windows.S_OK, windows.Handle(hr), "GetContent should succeed")
	require.NotNil(t, stream, "Response should carry a content stream")
	defer stream.Release()

	buf := make([]byte, len(content)+16)
	n, _ := stream.Read(buf)
	require.Equal(t, content, buf[:n], "Response content should match the input")
}
