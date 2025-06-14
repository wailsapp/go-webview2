package generator

import (
	"strings"
	"testing"

	"generator/idl"
)

func TestExactContextFromWebView2(t *testing.T) {
	// Test the exact context where the parser fails, including the forward declarations
	testIDL := `library WebView2 {

// Interface forward declarations
interface ICoreWebView2;
interface ICoreWebView2WebResourceResponse;
interface ICoreWebView2WebResourceResponseReceivedEventArgs;
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

}`

	parser := idl.NewParser(strings.NewReader(testIDL))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Failed to parse exact WebView2 context: %v", err)
	}

	if len(ast.Enums) != 1 {
		t.Fatalf("Expected 1 enum, got %d", len(ast.Enums))
	}

	enum := ast.Enums[0]
	if enum.Name != "COREWEBVIEW2_BOUNDS_MODE" {
		t.Errorf("Expected enum name 'COREWEBVIEW2_BOUNDS_MODE', got '%s'", enum.Name)
	}

	t.Logf("Successfully parsed exact WebView2 context with enum: %s", enum.Name)
	t.Logf("Interfaces parsed: %d", len(ast.Interfaces))
}