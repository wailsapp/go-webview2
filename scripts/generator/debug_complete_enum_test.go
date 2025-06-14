package generator

import (
	"bytes"
	"testing"
	"generator/idl"
)

func TestCompleteFirstEnum(t *testing.T) {
	// Test the complete first enum with its context
	idlContent := `/// Comment
[v1_enum]
typedef enum COREWEBVIEW2_BOUNDS_MODE {
  /// Bounds property represents raw pixels. Physical size of Webview is not impacted by RasterizationScale.
  COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS,
  /// Bounds property represents logical pixels and the RasterizationScale property is used to get the physical size of the WebView.
  COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE,
} COREWEBVIEW2_BOUNDS_MODE;`

	parser := idl.NewParser(bytes.NewReader([]byte(idlContent)))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Complete first enum failed: %v", err)
	}
	
	if len(ast.Enums) != 1 {
		t.Errorf("Expected 1 enum, got %d", len(ast.Enums))
	}
	
	if ast.Enums[0].Name != "COREWEBVIEW2_BOUNDS_MODE" {
		t.Errorf("Expected enum name 'COREWEBVIEW2_BOUNDS_MODE', got '%s'", ast.Enums[0].Name)
	}
}

func TestCompleteFirstEnumPlusNext(t *testing.T) {
	// Test the complete first enum plus the start of the next one
	idlContent := `/// Comment
[v1_enum]
typedef enum COREWEBVIEW2_BOUNDS_MODE {
  /// Bounds property represents raw pixels. Physical size of Webview is not impacted by RasterizationScale.
  COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS,
  /// Bounds property represents logical pixels and the RasterizationScale property is used to get the physical size of the WebView.
  COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE,
} COREWEBVIEW2_BOUNDS_MODE;

/// Specifies the browser process exit type used in the
/// ICoreWebView2BrowserProcessExitedEventArgs interface.
[v1_enum]
typedef enum COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND {`

	parser := idl.NewParser(bytes.NewReader([]byte(idlContent)))
	_, err := parser.Parse()
	if err != nil {
		t.Logf("Complete first enum + incomplete second failed: %v", err)
	} else {
		t.Error("Should have failed due to incomplete second enum")
	}
}

func TestRealLines286To291(t *testing.T) {
	// Test the exact lines from the real file that should work
	idlContent := `typedef enum COREWEBVIEW2_BOUNDS_MODE {
  /// Bounds property represents raw pixels. Physical size of Webview is not impacted by RasterizationScale.
  COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS,
  /// Bounds property represents logical pixels and the RasterizationScale property is used to get the physical size of the WebView.
  COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE,
} COREWEBVIEW2_BOUNDS_MODE;`

	parser := idl.NewParser(bytes.NewReader([]byte(idlContent)))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Real lines 286-291 failed: %v", err)
	}
	
	if len(ast.Enums) != 1 {
		t.Errorf("Expected 1 enum, got %d", len(ast.Enums))
	}
}

func TestRealLines285To291(t *testing.T) {
	// Test including the attribute line
	idlContent := `[v1_enum]
typedef enum COREWEBVIEW2_BOUNDS_MODE {
  /// Bounds property represents raw pixels. Physical size of Webview is not impacted by RasterizationScale.
  COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS,
  /// Bounds property represents logical pixels and the RasterizationScale property is used to get the physical size of the WebView.
  COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE,
} COREWEBVIEW2_BOUNDS_MODE;`

	parser := idl.NewParser(bytes.NewReader([]byte(idlContent)))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Real lines 285-291 failed: %v", err)
	}
	
	if len(ast.Enums) != 1 {
		t.Errorf("Expected 1 enum, got %d", len(ast.Enums))
	}
	
	enum := ast.Enums[0]
	if len(enum.Attributes) != 1 {
		t.Errorf("Expected 1 attribute, got %d", len(enum.Attributes))
	}
	
	if enum.Attributes[0].Name != "v1_enum" {
		t.Errorf("Expected attribute 'v1_enum', got '%s'", enum.Attributes[0].Name)
	}
}