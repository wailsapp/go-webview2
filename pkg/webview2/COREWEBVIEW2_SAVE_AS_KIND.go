//go:build windows

package webview2

type COREWEBVIEW2_SAVE_AS_KIND uint32

const (
	COREWEBVIEW2_SAVE_AS_KIND_DEFAULT     = 0
	COREWEBVIEW2_SAVE_AS_KIND_HTML_ONLY   = 1
	COREWEBVIEW2_SAVE_AS_KIND_SINGLE_FILE = 2
	COREWEBVIEW2_SAVE_AS_KIND_COMPLETE    = 3
)
