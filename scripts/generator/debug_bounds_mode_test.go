package generator

import (
	"bytes"
	"testing"
	"generator/idl"
)

func TestBoundsModeEnum(t *testing.T) {
	// Test just the COREWEBVIEW2_BOUNDS_MODE enum
	idlContent := `
/// Mode for how the Bounds property is interpreted in relation to the RasterizationScale property.
[v1_enum]
typedef enum COREWEBVIEW2_BOUNDS_MODE {
  /// Bounds property represents raw pixels. Physical size of Webview is not impacted by RasterizationScale.
  COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS,
  /// Bounds property represents logical pixels and the RasterizationScale property is used to get the physical size of the WebView.
  COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE,
} COREWEBVIEW2_BOUNDS_MODE;
`

	parser := idl.NewParser(bytes.NewReader([]byte(idlContent)))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed to parse COREWEBVIEW2_BOUNDS_MODE enum: %v", err)
	}

	if len(ast.Enums) != 1 {
		t.Fatalf("Expected 1 enum, got %d", len(ast.Enums))
	}

	enum := ast.Enums[0]
	if enum.Name != "COREWEBVIEW2_BOUNDS_MODE" {
		t.Errorf("Expected enum name 'COREWEBVIEW2_BOUNDS_MODE', got '%s'", enum.Name)
	}

	if len(enum.Values) != 2 {
		t.Errorf("Expected 2 enum values, got %d", len(enum.Values))
	}
}

func TestWithPrecedingForwardDeclarations(t *testing.T) {
	// Test with some forward declarations before the enum
	idlContent := `
interface ICoreWebView2;
interface ICoreWebView2Controller;

/// Mode for how the Bounds property is interpreted in relation to the RasterizationScale property.
[v1_enum]
typedef enum COREWEBVIEW2_BOUNDS_MODE {
  /// Bounds property represents raw pixels. Physical size of Webview is not impacted by RasterizationScale.
  COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS,
  /// Bounds property represents logical pixels and the RasterizationScale property is used to get the physical size of the WebView.
  COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE,
} COREWEBVIEW2_BOUNDS_MODE;
`

	parser := idl.NewParser(bytes.NewReader([]byte(idlContent)))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed to parse with forward declarations: %v", err)
	}

	if len(ast.Enums) != 1 {
		t.Fatalf("Expected 1 enum, got %d", len(ast.Enums))
	}
}

func TestWithRealContext(t *testing.T) {
	// Test with the real context that's causing issues
	idlContent := `
interface ICoreWebView2WebResourceResponseView;
interface ICoreWebView2WebResourceResponseViewGetContentCompletedHandler;
interface ICoreWebView2WindowFeatures;

// Enums and structs


/// Mode for how the Bounds property is interpreted in relation to the RasterizationScale property.
[v1_enum]
typedef enum COREWEBVIEW2_BOUNDS_MODE {
  /// Bounds property represents raw pixels. Physical size of Webview is not impacted by RasterizationScale.
  COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS,
  /// Bounds property represents logical pixels and the RasterizationScale property is used to get the physical size of the WebView.
  COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE,
} COREWEBVIEW2_BOUNDS_MODE;
`

	parser := idl.NewParser(bytes.NewReader([]byte(idlContent)))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed to parse with real context: %v", err)
	}

	if len(ast.Enums) != 1 {
		t.Fatalf("Expected 1 enum, got %d", len(ast.Enums))
	}
}

func TestDebugTokensAroundFailure(t *testing.T) {
	// Let's see what tokens are being generated around the failure point
	idlContent := `// Enums and structs


/// Mode for how the Bounds property is interpreted in relation to the RasterizationScale property.
[v1_enum]
typedef enum COREWEBVIEW2_BOUNDS_MODE {`

	scanner := idl.NewScanner(bytes.NewReader([]byte(idlContent)))
	
	t.Log("Tokens generated:")
	for i := 0; i < 20; i++ {
		token := scanner.NextToken()
		t.Logf("  %d: Type=%d, Value='%s'", i, token.Type, token.Value)
		if token.Type == 0 { // EOF
			break
		}
	}
}