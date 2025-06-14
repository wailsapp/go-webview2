package generator

import (
	"strings"
	"testing"
	"generator/idl"
)

func TestDebugFullIDLParsing(t *testing.T) {
	// Test with the exact sequence from lines 307-370 that's causing issues
	idlContent := `
[v1_enum]
typedef enum COREWEBVIEW2_BROWSING_DATA_KINDS {
  COREWEBVIEW2_BROWSING_DATA_KINDS_FILE_SYSTEMS = 0x1,
  COREWEBVIEW2_BROWSING_DATA_KINDS_SERVICE_WORKERS = 0x8000,
} COREWEBVIEW2_BROWSING_DATA_KINDS;
cpp_quote("DEFINE_ENUM_FLAG_OPERATORS(COREWEBVIEW2_BROWSING_DATA_KINDS)")

/// Specifies the image format for the ICoreWebView2::CapturePreview method.
[v1_enum]
typedef enum COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT {
  COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT_PNG,
  COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT_JPEG,
} COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT;
`

	reader := strings.NewReader(idlContent)
	parser := idl.NewParser(reader)

	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed to parse full sequence: %v", err)
	}

	// Should have 2 enums
	if len(ast.Enums) != 2 {
		t.Fatalf("Expected 2 enums, got %d", len(ast.Enums))
	}

	// First enum should be COREWEBVIEW2_BROWSING_DATA_KINDS
	enum1 := ast.Enums[0]
	if enum1.Name != "COREWEBVIEW2_BROWSING_DATA_KINDS" {
		t.Errorf("Expected first enum name 'COREWEBVIEW2_BROWSING_DATA_KINDS', got '%s'", enum1.Name)
	}

	// Second enum should be COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT
	enum2 := ast.Enums[1]
	if enum2.Name != "COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT" {
		t.Errorf("Expected second enum name 'COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT', got '%s'", enum2.Name)
	}
}

func TestDebugTokenSequence(t *testing.T) {
	// Test just the critical closing sequence
	idlContent := `
} COREWEBVIEW2_BROWSING_DATA_KINDS;
cpp_quote("DEFINE_ENUM_FLAG_OPERATORS(COREWEBVIEW2_BROWSING_DATA_KINDS)")
`

	reader := strings.NewReader(idlContent)
	scanner := idl.NewScanner(reader)

	// Manually tokenize to see what tokens are generated
	var tokens []idl.Token
	for {
		token := scanner.NextToken()
		tokens = append(tokens, token)
		if token.Type == 0 { // TokenEOF
			break
		}
	}

	t.Logf("Tokens found:")
	for i, token := range tokens {
		t.Logf("  %d: Type=%d, Value='%s'", i, token.Type, token.Value)
	}

	// Should see: RBrace, Identifier(COREWEBVIEW2_BROWSING_DATA_KINDS), Semicolon, Identifier(cpp_quote), etc.
	if len(tokens) < 4 {
		t.Errorf("Expected at least 4 tokens, got %d", len(tokens))
		return
	}

	// Verify we get the expected sequence
	expectedSequence := []string{"}", "COREWEBVIEW2_BROWSING_DATA_KINDS", ";", "cpp_quote"}
	for i, expected := range expectedSequence {
		if i >= len(tokens) {
			t.Errorf("Missing token %d: expected '%s'", i, expected)
			continue
		}
		if tokens[i].Value != expected {
			t.Errorf("Token %d: expected '%s', got '%s'", i, expected, tokens[i].Value)
		}
	}
}

func TestDebugWithCommentsAndCommas(t *testing.T) {
	// Test the exact problematic enum with comments and trailing comma
	idlContent := `
[v1_enum]
typedef enum COREWEBVIEW2_BROWSING_DATA_KINDS {
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
		t.Fatalf("Failed to parse enum with trailing comma and comments: %v", err)
	}

	if len(ast.Enums) != 1 {
		t.Fatalf("Expected 1 enum, got %d", len(ast.Enums))
	}

	enum := ast.Enums[0]
	if len(enum.Values) != 1 {
		t.Errorf("Expected 1 enum value, got %d", len(enum.Values))
	}

	if enum.Values[0].Name != "COREWEBVIEW2_BROWSING_DATA_KINDS_SERVICE_WORKERS" {
		t.Errorf("Expected enum value 'COREWEBVIEW2_BROWSING_DATA_KINDS_SERVICE_WORKERS', got '%s'", enum.Values[0].Name)
	}
}