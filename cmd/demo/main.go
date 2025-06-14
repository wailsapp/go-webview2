//go:build windows

package main

import (
	"log"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/wailsapp/go-webview2/internal/w32"
	"github.com/wailsapp/go-webview2/pkg/edge"
)

const (
	windowWidth  = 1024
	windowHeight = 768
)

func main() {
	runtime.LockOSThread()

	// Create main window
	hwnd := createWindow()
	if hwnd == 0 {
		log.Fatal("Failed to create window")
	}

	// Create WebView2 instance
	webview := edge.NewChromium()
	webview.Debug = true

	// Set message callback to handle JavaScript messages
	webview.MessageCallback = func(message string) {
		log.Printf("Received message from JavaScript: %s", message)
	}

	// Embed WebView2 in the window
	if !webview.Embed(hwnd) {
		log.Fatal("Failed to embed WebView2")
	}

	// Navigate to a test website
	webview.Navigate("https://www.google.com")

	// Show the window
	w32.User32ShowWindow.Call(hwnd, w32.SW_SHOW)
	w32.User32UpdateWindow.Call(hwnd)

	// Message loop
	var msg w32.Msg
	for {
		ret, _, _ := w32.User32GetMessageW.Call(
			uintptr(unsafe.Pointer(&msg)),
			0,
			0,
			0,
		)

		if ret == 0 { // WM_QUIT
			break
		}

		if ret == ^uintptr(0) { // -1, error
			log.Fatal("GetMessage error")
		}

		// Handle window resize
		if msg.Message == w32.WM_SIZE {
			webview.Resize()
		}

		w32.User32TranslateMessage.Call(uintptr(unsafe.Pointer(&msg)))
		w32.User32DispatchMessageW.Call(uintptr(unsafe.Pointer(&msg)))
	}
}

func createWindow() uintptr {
	className, _ := syscall.UTF16PtrFromString("WebView2Demo")
	windowName, _ := syscall.UTF16PtrFromString("WebView2 Demo - Google")

	// Register window class
	var wc w32.WndClassEx
	wc.CbSize = uint32(unsafe.Sizeof(wc))
	wc.Style = w32.CS_HREDRAW | w32.CS_VREDRAW
	wc.LpfnWndProc = syscall.NewCallback(wndProc)
	wc.HInstance = w32.GetModuleHandle()
	wc.HCursor, _, _ = w32.User32LoadCursorW.Call(0, uintptr(w32.IDC_ARROW))
	wc.HbrBackground = w32.COLOR_WINDOW + 1
	wc.LpszClassName = className

	if _, _, _ := w32.User32RegisterClassExW.Call(uintptr(unsafe.Pointer(&wc))); wc.LpszClassName == nil {
		return 0
	}

	// Create window
	hwnd, _, _ := w32.User32CreateWindowExW.Call(
		0,
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)),
		w32.WS_OVERLAPPEDWINDOW,
		w32.CW_USEDEFAULT,
		w32.CW_USEDEFAULT,
		windowWidth,
		windowHeight,
		0,
		0,
		w32.GetModuleHandle(),
		0,
	)

	return hwnd
}

func wndProc(hwnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case w32.WM_DESTROY:
		w32.User32PostQuitMessage.Call(0)
		return 0
	case w32.WM_SIZE:
		// WebView2 resize will be handled in main loop
		return 0
	default:
		ret, _, _ := w32.User32DefWindowProcW.Call(hwnd, uintptr(msg), wParam, lParam)
		return ret
	}
}
