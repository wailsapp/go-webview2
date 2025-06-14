package generator

import (
	"strings"
	"testing"

	"generator/idl"
)

func TestExactBrowsingDataKindsCase(t *testing.T) {
	// This is the exact enum from WebView2 IDL that has cpp_quote after it
	testIDL := `[v1_enum]
typedef enum COREWEBVIEW2_BROWSING_DATA_KINDS {
  COREWEBVIEW2_BROWSING_DATA_KINDS_FILE_SYSTEMS = 0x1,
  COREWEBVIEW2_BROWSING_DATA_KINDS_SERVICE_WORKERS = 0x8000,
} COREWEBVIEW2_BROWSING_DATA_KINDS;
cpp_quote("DEFINE_ENUM_FLAG_OPERATORS(COREWEBVIEW2_BROWSING_DATA_KINDS)")

[v1_enum]
typedef enum COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT {
  COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT_PNG,
  COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT_JPEG,
} COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT;`

	parser := idl.NewParser(strings.NewReader(testIDL))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Failed to parse exact browsing data kinds case: %v", err)
	}

	if len(ast.Enums) != 2 {
		t.Fatalf("Expected 2 enums, got %d", len(ast.Enums))
	}

	t.Logf("Successfully parsed both enums:")
	for i, enum := range ast.Enums {
		t.Logf("  Enum %d: %s (%d values)", i+1, enum.Name, len(enum.Values))
	}
}