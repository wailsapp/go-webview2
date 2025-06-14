package generator

import (
	"strings"
	"testing"
	"generator/idl"
)

func TestCoreWebView2BrowsingDataKinds(t *testing.T) {
	// Exact enum definition from the IDL file
	idlContent := `
/// Specifies the datatype for the
/// ICoreWebView2Profile2::ClearBrowsingData method.
[v1_enum]
typedef enum COREWEBVIEW2_BROWSING_DATA_KINDS {
  /// Specifies file systems data.
  COREWEBVIEW2_BROWSING_DATA_KINDS_FILE_SYSTEMS = 0x1,
  /// Specifies data stored by the IndexedDB DOM feature.
  COREWEBVIEW2_BROWSING_DATA_KINDS_INDEXED_DB = 0x2,
  /// Specifies data stored by the localStorage DOM API.
  COREWEBVIEW2_BROWSING_DATA_KINDS_LOCAL_STORAGE = 0x4,
  /// Specifies data stored by the Web SQL database DOM API.
  COREWEBVIEW2_BROWSING_DATA_KINDS_WEB_SQL = 0x8,
  /// Specifies data stored by the CacheStorage DOM API.
  COREWEBVIEW2_BROWSING_DATA_KINDS_CACHE_STORAGE = 0x10,
  /// Specifies DOM storage data, now and future. This browsing data kind is
  /// inclusive of COREWEBVIEW2_BROWSING_DATA_KINDS_FILE_SYSTEMS,
  /// COREWEBVIEW2_BROWSING_DATA_KINDS_INDEXED_DB,
  /// COREWEBVIEW2_BROWSING_DATA_KINDS_LOCAL_STORAGE,
  /// COREWEBVIEW2_BROWSING_DATA_KINDS_WEB_SQL,
  /// COREWEBVIEW2_BROWSING_DATA_KINDS_SERVICE_WORKERS,
  /// COREWEBVIEW2_BROWSING_DATA_KINDS_CACHE_STORAGE,
  /// and some other data kinds not listed yet to keep consistent with
  /// [DOM-accessible storage](https://www.w3.org/TR/clear-site-data/#storage).
  COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_DOM_STORAGE = 0x20,
  /// Specifies HTTP cookies data.
  COREWEBVIEW2_BROWSING_DATA_KINDS_COOKIES = 0x40,
  /// Specifies all site data, now and future. This browsing data kind
  /// is inclusive of COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_DOM_STORAGE and
  /// COREWEBVIEW2_BROWSING_DATA_KINDS_COOKIES. New site data types
  /// may be added to this data kind in the future.
  COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_SITE = 0x80,
  /// Specifies disk cache.
  COREWEBVIEW2_BROWSING_DATA_KINDS_DISK_CACHE = 0x100,
  /// Specifies download history data.
  COREWEBVIEW2_BROWSING_DATA_KINDS_DOWNLOAD_HISTORY = 0x200,
  /// Specifies general autofill form data.
  /// This excludes password information and includes information like:
  /// names, street and email addresses, phone numbers, and arbitrary input.
  /// This also includes payment data.
  COREWEBVIEW2_BROWSING_DATA_KINDS_GENERAL_AUTOFILL = 0x400,
  /// Specifies password autosave data.
  COREWEBVIEW2_BROWSING_DATA_KINDS_PASSWORD_AUTOSAVE = 0x800,
  /// Specifies browsing history data.
  COREWEBVIEW2_BROWSING_DATA_KINDS_BROWSING_HISTORY = 0x1000,
  /// Specifies settings data.
  COREWEBVIEW2_BROWSING_DATA_KINDS_SETTINGS = 0x2000,
  /// Specifies profile data that should be wiped to make it look like a new profile.
  /// This does not delete account-scoped data like passwords but will remove access
  /// to account-scoped data by signing the user out.
  /// Specifies all profile data, now and future. New profile data types may be added
  /// to this data kind in the future.
  /// This browsing data kind is inclusive of COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_SITE,
  /// COREWEBVIEW2_BROWSING_DATA_KINDS_DISK_CACHE,
  /// COREWEBVIEW2_BROWSING_DATA_KINDS_DOWNLOAD_HISTORY,
  /// COREWEBVIEW2_BROWSING_DATA_KINDS_GENERAL_AUTOFILL,
  /// COREWEBVIEW2_BROWSING_DATA_KINDS_PASSWORD_AUTOSAVE,
  /// COREWEBVIEW2_BROWSING_DATA_KINDS_BROWSING_HISTORY, and
  /// COREWEBVIEW2_BROWSING_DATA_KINDS_SETTINGS.
  COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_PROFILE = 0x4000,
  /// Specifies service workers registered for an origin, and clear will result in
  /// termination and deregistration of them.
  COREWEBVIEW2_BROWSING_DATA_KINDS_SERVICE_WORKERS = 0x8000,
} COREWEBVIEW2_BROWSING_DATA_KINDS;
cpp_quote("DEFINE_ENUM_FLAG_OPERATORS(COREWEBVIEW2_BROWSING_DATA_KINDS)")
`

	reader := strings.NewReader(idlContent)
	parser := idl.NewParser(reader)

	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed to parse COREWEBVIEW2_BROWSING_DATA_KINDS: %v", err)
	}

	// Verify the enum was parsed correctly
	if len(ast.Enums) != 1 {
		t.Fatalf("Expected 1 enum, got %d", len(ast.Enums))
	}

	enum := ast.Enums[0]
	if enum.Name != "COREWEBVIEW2_BROWSING_DATA_KINDS" {
		t.Errorf("Expected enum name 'COREWEBVIEW2_BROWSING_DATA_KINDS', got '%s'", enum.Name)
	}

	// Check that we have all the expected values
	expectedValues := []string{
		"COREWEBVIEW2_BROWSING_DATA_KINDS_FILE_SYSTEMS",
		"COREWEBVIEW2_BROWSING_DATA_KINDS_INDEXED_DB",
		"COREWEBVIEW2_BROWSING_DATA_KINDS_LOCAL_STORAGE",
		"COREWEBVIEW2_BROWSING_DATA_KINDS_WEB_SQL",
		"COREWEBVIEW2_BROWSING_DATA_KINDS_CACHE_STORAGE",
		"COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_DOM_STORAGE",
		"COREWEBVIEW2_BROWSING_DATA_KINDS_COOKIES",
		"COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_SITE",
		"COREWEBVIEW2_BROWSING_DATA_KINDS_DISK_CACHE",
		"COREWEBVIEW2_BROWSING_DATA_KINDS_DOWNLOAD_HISTORY",
		"COREWEBVIEW2_BROWSING_DATA_KINDS_GENERAL_AUTOFILL",
		"COREWEBVIEW2_BROWSING_DATA_KINDS_PASSWORD_AUTOSAVE",
		"COREWEBVIEW2_BROWSING_DATA_KINDS_BROWSING_HISTORY",
		"COREWEBVIEW2_BROWSING_DATA_KINDS_SETTINGS",
		"COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_PROFILE",
		"COREWEBVIEW2_BROWSING_DATA_KINDS_SERVICE_WORKERS",
	}

	if len(enum.Values) != len(expectedValues) {
		t.Errorf("Expected %d enum values, got %d", len(expectedValues), len(enum.Values))
	}

	for i, expected := range expectedValues {
		if i >= len(enum.Values) {
			t.Errorf("Missing enum value: %s", expected)
			continue
		}
		if enum.Values[i].Name != expected {
			t.Errorf("Expected enum value '%s', got '%s'", expected, enum.Values[i].Name)
		}
	}

	// Verify the hex values are correct
	expectedHexValues := []int64{0x1, 0x2, 0x4, 0x8, 0x10, 0x20, 0x40, 0x80, 0x100, 0x200, 0x400, 0x800, 0x1000, 0x2000, 0x4000, 0x8000}
	for i, expected := range expectedHexValues {
		if i >= len(enum.Values) {
			continue
		}
		if enum.Values[i].Value != expected {
			t.Errorf("Expected enum value %d to have value 0x%x, got 0x%x", i, expected, enum.Values[i].Value)
		}
	}
}

func TestCoreWebView2BrowsingDataKindsWithCppQuote(t *testing.T) {
	// Test with the cpp_quote directive that appears after the enum
	idlContent := `
[v1_enum]
typedef enum COREWEBVIEW2_BROWSING_DATA_KINDS {
  COREWEBVIEW2_BROWSING_DATA_KINDS_FILE_SYSTEMS = 0x1,
  COREWEBVIEW2_BROWSING_DATA_KINDS_INDEXED_DB = 0x2,
  COREWEBVIEW2_BROWSING_DATA_KINDS_SERVICE_WORKERS = 0x8000,
} COREWEBVIEW2_BROWSING_DATA_KINDS;
cpp_quote("DEFINE_ENUM_FLAG_OPERATORS(COREWEBVIEW2_BROWSING_DATA_KINDS)")
`

	reader := strings.NewReader(idlContent)
	parser := idl.NewParser(reader)

	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed to parse COREWEBVIEW2_BROWSING_DATA_KINDS with cpp_quote: %v", err)
	}

	// Verify the enum was parsed correctly
	if len(ast.Enums) != 1 {
		t.Fatalf("Expected 1 enum, got %d", len(ast.Enums))
	}

	enum := ast.Enums[0]
	if enum.Name != "COREWEBVIEW2_BROWSING_DATA_KINDS" {
		t.Errorf("Expected enum name 'COREWEBVIEW2_BROWSING_DATA_KINDS', got '%s'", enum.Name)
	}

	if len(enum.Values) != 3 {
		t.Errorf("Expected 3 enum values, got %d", len(enum.Values))
	}
}

func TestMinimalBrowsingDataKinds(t *testing.T) {
	// Test with just the problematic part
	idlContent := `
typedef enum COREWEBVIEW2_BROWSING_DATA_KINDS {
  COREWEBVIEW2_BROWSING_DATA_KINDS_FILE_SYSTEMS = 0x1,
} COREWEBVIEW2_BROWSING_DATA_KINDS;
`

	reader := strings.NewReader(idlContent)
	parser := idl.NewParser(reader)

	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed to parse minimal COREWEBVIEW2_BROWSING_DATA_KINDS: %v", err)
	}

	// Verify the enum was parsed correctly
	if len(ast.Enums) != 1 {
		t.Fatalf("Expected 1 enum, got %d", len(ast.Enums))
	}

	enum := ast.Enums[0]
	if enum.Name != "COREWEBVIEW2_BROWSING_DATA_KINDS" {
		t.Errorf("Expected enum name 'COREWEBVIEW2_BROWSING_DATA_KINDS', got '%s'", enum.Name)
	}
}