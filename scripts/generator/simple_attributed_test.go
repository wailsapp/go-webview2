package generator

import (
	"strings"
	"testing"

	"generator/idl"
)

func TestSimpleAttributedTypedef(t *testing.T) {
	// Test the exact minimal case that should work
	testIDL := `[v1_enum]
typedef enum COREWEBVIEW2_BOUNDS_MODE {
  COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS,
} COREWEBVIEW2_BOUNDS_MODE;`

	parser := idl.NewParser(strings.NewReader(testIDL))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Failed to parse simple attributed typedef: %v", err)
	}

	if len(ast.Enums) != 1 {
		t.Fatalf("Expected 1 enum, got %d", len(ast.Enums))
	}

	enum := ast.Enums[0]
	if enum.Name != "COREWEBVIEW2_BOUNDS_MODE" {
		t.Errorf("Expected enum name 'COREWEBVIEW2_BOUNDS_MODE', got '%s'", enum.Name)
	}

	if len(enum.Attributes) != 1 {
		t.Errorf("Expected 1 attribute, got %d", len(enum.Attributes))
	}

	if len(enum.Attributes) > 0 && enum.Attributes[0].Name != "v1_enum" {
		t.Errorf("Expected attribute name 'v1_enum', got '%s'", enum.Attributes[0].Name)
	}

	t.Logf("Successfully parsed simple attributed typedef: %s", enum.Name)
}

func TestLibraryWithAttributedTypedef(t *testing.T) {
	// Test inside a library context
	testIDL := `library WebView2 {
[v1_enum]
typedef enum COREWEBVIEW2_BOUNDS_MODE {
  COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS,
} COREWEBVIEW2_BOUNDS_MODE;
}`

	parser := idl.NewParser(strings.NewReader(testIDL))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Failed to parse library with attributed typedef: %v", err)
	}

	if len(ast.Enums) != 1 {
		t.Fatalf("Expected 1 enum, got %d", len(ast.Enums))
	}

	enum := ast.Enums[0]
	if enum.Name != "COREWEBVIEW2_BOUNDS_MODE" {
		t.Errorf("Expected enum name 'COREWEBVIEW2_BOUNDS_MODE', got '%s'", enum.Name)
	}

	t.Logf("Successfully parsed library with attributed typedef: %s", enum.Name)
}