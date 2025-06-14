package generator

import (
	"strings"
	"testing"

	"generator/idl"
)

func TestTypedefEnumWithCppQuote(t *testing.T) {
	// This is the exact pattern from WebView2 IDL that was failing
	testIDL := `[v1_enum]
typedef enum COREWEBVIEW2_BROWSING_DATA_KINDS {
  COREWEBVIEW2_BROWSING_DATA_KINDS_FILE_SYSTEMS = 0x1,
  COREWEBVIEW2_BROWSING_DATA_KINDS_INDEXED_DB = 0x2,
  COREWEBVIEW2_BROWSING_DATA_KINDS_LOCAL_STORAGE = 0x4,
  COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_PROFILE = 0x4000,
} COREWEBVIEW2_BROWSING_DATA_KINDS;
cpp_quote("DEFINE_ENUM_FLAG_OPERATORS(COREWEBVIEW2_BROWSING_DATA_KINDS)")`

	parser := idl.NewParser(strings.NewReader(testIDL))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Failed to parse typedef enum with cpp_quote: %v", err)
	}

	if len(ast.Enums) != 1 {
		t.Fatalf("Expected 1 enum, got %d", len(ast.Enums))
	}

	enum := ast.Enums[0]
	if enum.Name != "COREWEBVIEW2_BROWSING_DATA_KINDS" {
		t.Errorf("Expected enum name 'COREWEBVIEW2_BROWSING_DATA_KINDS', got '%s'", enum.Name)
	}

	if len(enum.Values) != 4 {
		t.Errorf("Expected 4 enum values, got %d", len(enum.Values))
	}

	t.Logf("Successfully parsed typedef enum with cpp_quote: %s with %d values", enum.Name, len(enum.Values))
}

func TestTypedefEnumWithoutAttributesAndCppQuote(t *testing.T) {
	// Test without attributes but with cpp_quote
	testIDL := `typedef enum COREWEBVIEW2_TEST {
  VALUE1,
  VALUE2,
} COREWEBVIEW2_TEST;
cpp_quote("DEFINE_ENUM_FLAG_OPERATORS(COREWEBVIEW2_TEST)")`

	parser := idl.NewParser(strings.NewReader(testIDL))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Failed to parse typedef enum without attributes but with cpp_quote: %v", err)
	}

	if len(ast.Enums) != 1 {
		t.Fatalf("Expected 1 enum, got %d", len(ast.Enums))
	}

	enum := ast.Enums[0]
	if enum.Name != "COREWEBVIEW2_TEST" {
		t.Errorf("Expected enum name 'COREWEBVIEW2_TEST', got '%s'", enum.Name)
	}

	t.Logf("Successfully parsed typedef enum without attributes but with cpp_quote: %s", enum.Name)
}