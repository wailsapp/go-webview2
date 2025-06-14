//go:build windows

package main

import (
	"fmt"
	"log"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/wailsapp/go-webview2/internal/w32"
	"github.com/wailsapp/go-webview2/webviewloader"
	"golang.org/x/sys/windows"
)

const (
	windowWidth  = 1024
	windowHeight = 768
	
	// Win32 constants not in w32 package
	SW_SHOW        = 5
	WM_QUIT        = 0x0012
	CS_HREDRAW     = 0x0002
	CS_VREDRAW     = 0x0001
	IDC_ARROW      = 32512
	COLOR_WINDOW   = 5
	PM_REMOVE      = 0x0001
)

var (
	user32 = windows.NewLazySystemDLL("user32")
	peekMessage = user32.NewProc("PeekMessageW")
	loadCursor = user32.NewProc("LoadCursorW")
)

func main() {
	runtime.LockOSThread()

	// Create main window
	hwnd := createWindow()
	if hwnd == 0 {
		log.Fatal("Failed to create window")
	}

	// Test WebView2 environment creation directly
	fmt.Println("Testing WebView2 environment creation...")
	
	// Test if WebView2 runtime is available
	version, err := webviewloader.GetAvailableCoreWebView2BrowserVersionString("")
	if err != nil {
		log.Fatal("WebView2 runtime not available:", err)
	}
	
	fmt.Printf("WebView2 runtime version: %s\n", version)
	fmt.Println("WebView2 test successful! WebView2 runtime is properly installed.")

	// Show the window
	w32.User32ShowWindow.Call(hwnd, SW_SHOW)
	w32.User32UpdateWindow.Call(hwnd)

	// Simple message loop
	var msg w32.Msg
	
	for i := 0; i < 30; i++ { // Run for a few seconds
		ret, _, _ := peekMessage.Call(
			uintptr(unsafe.Pointer(&msg)),
			0,
			0,
			0,
			PM_REMOVE,
		)

		if ret != 0 {
			if msg.Message == WM_QUIT {
				break
			}
			w32.User32TranslateMessage.Call(uintptr(unsafe.Pointer(&msg)))
			w32.User32DispatchMessageW.Call(uintptr(unsafe.Pointer(&msg)))
		}
		
		runtime.Gosched()
	}
	
	fmt.Println("Test completed successfully!")
}

func createWindow() uintptr {
	className, _ := syscall.UTF16PtrFromString("WebView2Test")
	windowName, _ := syscall.UTF16PtrFromString("WebView2 Runtime Test")

	// Register window class  
	var wc w32.WndClassExW
	wc.CbSize = uint32(unsafe.Sizeof(wc))
	wc.Style = CS_HREDRAW | CS_VREDRAW
	wc.LpfnWndProc = syscall.NewCallback(wndProc)
	process, _ := windows.GetCurrentProcess()
	wc.HInstance = windows.Handle(process)
	
	cursor, _, _ := loadCursor.Call(0, uintptr(IDC_ARROW))
	wc.HCursor = windows.Handle(cursor)
	wc.HbrBackground = windows.Handle(COLOR_WINDOW + 1)
	wc.LpszClassName = className

	if ret, _, _ := w32.User32RegisterClassExW.Call(uintptr(unsafe.Pointer(&wc))); ret == 0 {
		return 0
	}

	// Create window
	hwnd, _, _ := w32.User32CreateWindowExW.Call(
		0,
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)),
		w32.WSOverlappedWindow,
		w32.CW_USEDEFAULT,
		w32.CW_USEDEFAULT,
		windowWidth,
		windowHeight,
		0,
		0,
		uintptr(process),
		0,
	)

	return hwnd
}

func wndProc(hwnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case w32.WMDestroy:
		w32.User32PostQuitMessage.Call(0)
		return 0
	default:
		ret, _, _ := w32.User32DefWindowProcW.Call(hwnd, uintptr(msg), wParam, lParam)
		return ret
	}
}