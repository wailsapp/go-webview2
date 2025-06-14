package generator

import (
	"strings"
	"testing"

	"generator/idl"
)

func TestHexEnumValues(t *testing.T) {
	// Test enum with hex values like the BROWSING_DATA_KINDS enum
	testIDL := `[v1_enum]
typedef enum COREWEBVIEW2_BROWSING_DATA_KINDS {
  /// Specifies file systems data.
  COREWEBVIEW2_BROWSING_DATA_KINDS_FILE_SYSTEMS = 0x1,
  /// Specifies data stored by the IndexedDB DOM feature.
  COREWEBVIEW2_BROWSING_DATA_KINDS_INDEXED_DB = 0x2,
  /// Specifies data stored by the localStorage DOM API.
  COREWEBVIEW2_BROWSING_DATA_KINDS_LOCAL_STORAGE = 0x4,
  COREWEBVIEW2_BROWSING_DATA_KINDS_WEB_SQL = 0x8,
} COREWEBVIEW2_BROWSING_DATA_KINDS;`

	parser := idl.NewParser(strings.NewReader(testIDL))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Failed to parse hex enum values: %v", err)
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

	// Check hex values are parsed correctly
	expectedValues := []int64{0x1, 0x2, 0x4, 0x8}
	for i, expectedValue := range expectedValues {
		if i < len(enum.Values) && enum.Values[i].Value != expectedValue {
			t.Errorf("Expected enum value %d to be %d (0x%x), got %d", 
				i, expectedValue, expectedValue, enum.Values[i].Value)
		}
	}

	t.Logf("Successfully parsed hex enum values: %s with %d values", enum.Name, len(enum.Values))
	for i, value := range enum.Values {
		t.Logf("  Value %d: %s = %d (0x%x)", i, value.Name, value.Value, value.Value)
	}
}