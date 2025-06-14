package generator

import (
	"bytes"
	"os"
	"testing"
	"generator/idl"
)

func TestDebugParserTrace(t *testing.T) {
	// Let's test with a smaller subset that includes the problematic enum
	// First, let's just test around lines 300-400 of the IDL file
	idlData, err := os.ReadFile("../WebView2.1.0.3296.44.idl")
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	// Extract just the lines around the problematic enum (300-400)
	lines := bytes.Split(idlData, []byte("\n"))
	if len(lines) < 400 {
		t.Fatalf("IDL file too short")
	}

	// Take lines 300-400 (0-indexed)
	var subset []byte
	for i := 299; i < 400 && i < len(lines); i++ {
		subset = append(subset, lines[i]...)
		subset = append(subset, '\n')
	}

	t.Logf("Testing subset of IDL file (lines 300-400):")
	t.Logf("%s", string(subset))

	// Try to parse this subset
	parser := idl.NewParser(bytes.NewReader(subset))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Parse error on subset: %v", err)
	} else {
		t.Log("Subset parsed successfully")
	}
}

func TestDebugSpecificProblemArea(t *testing.T) {
	// Test the exact problematic sequence we identified
	idlContent := `
/// A ProcessFailed event will also be sent to listening WebViews from the
/// ICoreWebView2Environment associated to the failed process.
COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND_FAILED,
} COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND;

/// Specifies the datatype for the
/// ICoreWebView2Profile2::ClearBrowsingData method.
[v1_enum]
typedef enum COREWEBVIEW2_BROWSING_DATA_KINDS {
  /// Specifies file systems data.
  COREWEBVIEW2_BROWSING_DATA_KINDS_FILE_SYSTEMS = 0x1,
  /// Specifies data stored by the IndexedDB DOM feature.
  COREWEBVIEW2_BROWSING_DATA_KINDS_INDEXED_DB = 0x2,
  /// Specifies service workers registered for an origin, and clear will result in
  /// termination and deregistration of them.
  COREWEBVIEW2_BROWSING_DATA_KINDS_SERVICE_WORKERS = 0x8000,
} COREWEBVIEW2_BROWSING_DATA_KINDS;
cpp_quote("DEFINE_ENUM_FLAG_OPERATORS(COREWEBVIEW2_BROWSING_DATA_KINDS)")

/// Specifies the image format for the ICoreWebView2::CapturePreview method.
[v1_enum]
typedef enum COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT {
  /// Indicates that the PNG image format is used.
  COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT_PNG,
  /// Indicates the JPEG image format is used.
  COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT_JPEG,
} COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT;
`

	parser := idl.NewParser(bytes.NewReader([]byte(idlContent)))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed to parse problem area: %v", err)
	}

	t.Logf("Successfully parsed problem area:")
	t.Logf("  Enums: %d", len(ast.Enums))
	for i, enum := range ast.Enums {
		t.Logf("    %d: %s (%d values)", i, enum.Name, len(enum.Values))
	}
}

func TestProgressiveExpansion(t *testing.T) {
	// Test by progressively adding more content until we find the breaking point
	
	// Start with just the basic enum
	test1 := `
[v1_enum]
typedef enum COREWEBVIEW2_BROWSING_DATA_KINDS {
  COREWEBVIEW2_BROWSING_DATA_KINDS_FILE_SYSTEMS = 0x1,
} COREWEBVIEW2_BROWSING_DATA_KINDS;
`
	parser := idl.NewParser(bytes.NewReader([]byte(test1)))
	_, err := parser.Parse()
	if err != nil {
		t.Errorf("Test1 failed: %v", err)
	} else {
		t.Log("Test1 passed: basic enum")
	}

	// Add cpp_quote
	test2 := `
[v1_enum]
typedef enum COREWEBVIEW2_BROWSING_DATA_KINDS {
  COREWEBVIEW2_BROWSING_DATA_KINDS_FILE_SYSTEMS = 0x1,
} COREWEBVIEW2_BROWSING_DATA_KINDS;
cpp_quote("DEFINE_ENUM_FLAG_OPERATORS(COREWEBVIEW2_BROWSING_DATA_KINDS)")
`
	parser = idl.NewParser(bytes.NewReader([]byte(test2)))
	_, err = parser.Parse()
	if err != nil {
		t.Errorf("Test2 failed: %v", err)
	} else {
		t.Log("Test2 passed: enum with cpp_quote")
	}

	// Add next enum
	test3 := `
[v1_enum]
typedef enum COREWEBVIEW2_BROWSING_DATA_KINDS {
  COREWEBVIEW2_BROWSING_DATA_KINDS_FILE_SYSTEMS = 0x1,
} COREWEBVIEW2_BROWSING_DATA_KINDS;
cpp_quote("DEFINE_ENUM_FLAG_OPERATORS(COREWEBVIEW2_BROWSING_DATA_KINDS)")

[v1_enum]
typedef enum COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT {
  COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT_PNG,
} COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT;
`
	parser = idl.NewParser(bytes.NewReader([]byte(test3)))
	_, err = parser.Parse()
	if err != nil {
		t.Errorf("Test3 failed: %v", err)
	} else {
		t.Log("Test3 passed: two enums with cpp_quote")
	}
}