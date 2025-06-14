package generator

import (
	"bytes"
	"os"
	"testing"
	"generator/idl"
)

func TestParseActualIDLFile(t *testing.T) {
	// Read the actual IDL file that's failing
	idlData, err := os.ReadFile("../WebView2.1.0.3296.44.idl")
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	// Try to parse it with our parser
	files, err := ParseIDL(idlData)
	if err != nil {
		t.Fatalf("Failed to parse actual IDL file: %v", err)
	}

	if len(files) == 0 {
		t.Error("No files generated")
	}
}

func TestParseActualIDLFileDirectly(t *testing.T) {
	// Read the actual IDL file that's failing
	idlData, err := os.ReadFile("../WebView2.1.0.3296.44.idl")
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	// Try to parse it directly with the IDL parser
	parser := idl.NewParser(bytes.NewReader(idlData))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed to parse actual IDL file directly: %v", err)
	}

	// Log some stats about what was parsed
	t.Logf("Parsed AST stats:")
	t.Logf("  Interfaces: %d", len(ast.Interfaces))
	t.Logf("  Enums: %d", len(ast.Enums))
	t.Logf("  TypeDefs: %d", len(ast.TypeDefs))

	// Look for the problematic enum
	var browsingDataKindsEnum *idl.Enum
	for _, enum := range ast.Enums {
		if enum.Name == "COREWEBVIEW2_BROWSING_DATA_KINDS" {
			browsingDataKindsEnum = enum
			break
		}
	}

	if browsingDataKindsEnum == nil {
		t.Error("COREWEBVIEW2_BROWSING_DATA_KINDS enum not found")
	} else {
		t.Logf("Found COREWEBVIEW2_BROWSING_DATA_KINDS with %d values", len(browsingDataKindsEnum.Values))
	}
}