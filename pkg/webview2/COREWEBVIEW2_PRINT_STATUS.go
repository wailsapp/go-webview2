//go:build windows

package webview2

type COREWEBVIEW2_PRINT_STATUS uint32

const (
	COREWEBVIEW2_PRINT_STATUS_SUCCEEDED           = 0
	COREWEBVIEW2_PRINT_STATUS_PRINTER_UNAVAILABLE = 1
	COREWEBVIEW2_PRINT_STATUS_OTHER_ERROR         = 2
)
