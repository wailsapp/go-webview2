//go:build windows

package webview2

type COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL uint32

const (
	COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL_NORMAL = 0
	COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL_LOW    = 1
)