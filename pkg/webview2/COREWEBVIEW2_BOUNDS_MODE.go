//go:build windows

package webview2

type COREWEBVIEW2_BOUNDS_MODE uint32

const (
	COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS          = 0
	COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE = 1
)
