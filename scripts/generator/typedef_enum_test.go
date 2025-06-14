package generator

import (
	"strings"
	"testing"

	"generator/idl"
)

func TestAttributedTypedefEnumParsing(t *testing.T) {
	// This is the exact pattern from WebView2 IDL that's failing
	testIDL := `[v1_enum]
typedef enum COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND {
  COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND_NORMAL,
  COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND_FAILED,
} COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND;`

	parser := idl.NewParser(strings.NewReader(testIDL))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Failed to parse attributed typedef enum: %v", err)
	}

	if len(ast.Enums) != 1 {
		t.Fatalf("Expected 1 enum, got %d", len(ast.Enums))
	}

	enum := ast.Enums[0]
	if enum.Name != "COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND" {
		t.Errorf("Expected enum name 'COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND', got '%s'", enum.Name)
	}

	if len(enum.Attributes) != 1 {
		t.Errorf("Expected 1 attribute, got %d", len(enum.Attributes))
	}

	if len(enum.Attributes) > 0 && enum.Attributes[0].Name != "v1_enum" {
		t.Errorf("Expected attribute name 'v1_enum', got '%s'", enum.Attributes[0].Name)
	}

	if len(enum.Values) != 2 {
		t.Errorf("Expected 2 enum values, got %d", len(enum.Values))
	}

	t.Logf("Successfully parsed attributed typedef enum: %s with %d values", enum.Name, len(enum.Values))
}

func TestSimpleTypedefEnumParsing(t *testing.T) {
	// Test without attributes first
	testIDL := `typedef enum COREWEBVIEW2_BOUNDS_MODE {
  COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS,
  COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE,
} COREWEBVIEW2_BOUNDS_MODE;`

	parser := idl.NewParser(strings.NewReader(testIDL))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Failed to parse simple typedef enum: %v", err)
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

	t.Logf("Successfully parsed simple typedef enum: %s with %d values", enum.Name, len(enum.Values))
}