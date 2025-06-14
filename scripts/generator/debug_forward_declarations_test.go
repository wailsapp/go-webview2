package generator

import (
	"bytes"
	"strings"
	"testing"
	"generator/idl"
)

func TestManyForwardDeclarations(t *testing.T) {
	// Create a test with many forward declarations like in the real file
	var content strings.Builder
	
	// Add many forward declarations
	for i := 0; i < 100; i++ {
		content.WriteString("interface ICoreWebView2Test")
		content.WriteString(strings.Repeat("A", i%10)) // Make them unique
		content.WriteString(";\n")
	}
	
	// Add the enum that's failing
	content.WriteString(`
/// Mode for how the Bounds property is interpreted in relation to the RasterizationScale property.
[v1_enum]
typedef enum COREWEBVIEW2_BOUNDS_MODE {
  /// Bounds property represents raw pixels. Physical size of Webview is not impacted by RasterizationScale.
  COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS,
  /// Bounds property represents logical pixels and the RasterizationScale property is used to get the physical size of the WebView.
  COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE,
} COREWEBVIEW2_BOUNDS_MODE;
`)

	parser := idl.NewParser(bytes.NewReader([]byte(content.String())))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed to parse with many forward declarations: %v", err)
	}

	if len(ast.Enums) != 1 {
		t.Errorf("Expected 1 enum, got %d", len(ast.Enums))
	}
}

func TestWithExactRealContext(t *testing.T) {
	// Test with the exact sequence from lines 275-291
	idlContent := `interface ICoreWebView2WebResourceResponse;
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
`

	parser := idl.NewParser(bytes.NewReader([]byte(idlContent)))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed to parse with exact real context: %v", err)
	}

	if len(ast.Enums) != 1 {
		t.Errorf("Expected 1 enum, got %d", len(ast.Enums))
	}
}

func TestDebugParserStateAfterForwardDecls(t *testing.T) {
	// Let's test the exact failure case step by step
	
	// First, test the forward declarations work
	idlContent1 := `interface ICoreWebView2WebResourceResponse;
interface ICoreWebView2WebResourceResponseReceivedEventArgs;
interface ICoreWebView2WebResourceResponseView;
interface ICoreWebView2WebResourceResponseViewGetContentCompletedHandler;
interface ICoreWebView2WindowFeatures;`

	parser := idl.NewParser(bytes.NewReader([]byte(idlContent1)))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Forward declarations failed: %v", err)
	}
	t.Logf("Forward declarations parsed successfully: %d interfaces", len(ast.Interfaces))

	// Now test with the comment line
	idlContent2 := `interface ICoreWebView2WebResourceResponse;
interface ICoreWebView2WebResourceResponseReceivedEventArgs;
interface ICoreWebView2WebResourceResponseView;
interface ICoreWebView2WebResourceResponseViewGetContentCompletedHandler;
interface ICoreWebView2WindowFeatures;

// Enums and structs`

	parser = idl.NewParser(bytes.NewReader([]byte(idlContent2)))
	ast, err = parser.Parse()
	if err != nil {
		t.Fatalf("Forward declarations + comment failed: %v", err)
	}
	t.Logf("Forward declarations + comment parsed successfully")

	// Now test with empty lines
	idlContent3 := `interface ICoreWebView2WebResourceResponse;
interface ICoreWebView2WebResourceResponseReceivedEventArgs;
interface ICoreWebView2WebResourceResponseView;
interface ICoreWebView2WebResourceResponseViewGetContentCompletedHandler;
interface ICoreWebView2WindowFeatures;

// Enums and structs



`

	parser = idl.NewParser(bytes.NewReader([]byte(idlContent3)))
	ast, err = parser.Parse()
	if err != nil {
		t.Fatalf("Forward declarations + comment + empty lines failed: %v", err)
	}
	t.Logf("Forward declarations + comment + empty lines parsed successfully")

	// Now test with the triple slash comment
	idlContent4 := `interface ICoreWebView2WebResourceResponse;
interface ICoreWebView2WebResourceResponseReceivedEventArgs;
interface ICoreWebView2WebResourceResponseView;
interface ICoreWebView2WebResourceResponseViewGetContentCompletedHandler;
interface ICoreWebView2WindowFeatures;

// Enums and structs


/// Mode for how the Bounds property is interpreted in relation to the RasterizationScale property.`

	parser = idl.NewParser(bytes.NewReader([]byte(idlContent4)))
	ast, err = parser.Parse()
	if err != nil {
		t.Fatalf("Forward declarations + comments failed: %v", err)
	}
	t.Logf("Forward declarations + comments parsed successfully")

	// Now test with the attributes
	idlContent5 := `interface ICoreWebView2WebResourceResponse;
interface ICoreWebView2WebResourceResponseReceivedEventArgs;
interface ICoreWebView2WebResourceResponseView;
interface ICoreWebView2WebResourceResponseViewGetContentCompletedHandler;
interface ICoreWebView2WindowFeatures;

// Enums and structs


/// Mode for how the Bounds property is interpreted in relation to the RasterizationScale property.
[v1_enum]`

	parser = idl.NewParser(bytes.NewReader([]byte(idlContent5)))
	ast, err = parser.Parse()
	if err != nil {
		t.Fatalf("Forward declarations + comments + attributes failed: %v", err)
	}
	t.Logf("Forward declarations + comments + attributes parsed successfully")

	// Finally test with the typedef start
	idlContent6 := `interface ICoreWebView2WebResourceResponse;
interface ICoreWebView2WebResourceResponseReceivedEventArgs;
interface ICoreWebView2WebResourceResponseView;
interface ICoreWebView2WebResourceResponseViewGetContentCompletedHandler;
interface ICoreWebView2WindowFeatures;

// Enums and structs


/// Mode for how the Bounds property is interpreted in relation to the RasterizationScale property.
[v1_enum]
typedef`

	parser = idl.NewParser(bytes.NewReader([]byte(idlContent6)))
	ast, err = parser.Parse()
	if err != nil {
		t.Fatalf("Forward declarations + comments + attributes + typedef failed: %v", err)
	}
	t.Logf("Forward declarations + comments + attributes + typedef parsed successfully")
}