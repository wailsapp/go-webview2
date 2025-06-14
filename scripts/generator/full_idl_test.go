package generator

import (
	"os"
	"testing"

	"generator/idl"
)

func TestFullWebView2IDLParsing(t *testing.T) {
	// Test parsing of the full WebView2 IDL file
	idlPath := "../WebView2.1.0.3296.44.idl"
	
	file, err := os.Open(idlPath)
	if err != nil {
		t.Fatalf("Failed to open IDL file %s: %v", idlPath, err)
	}
	defer file.Close()

	parser := idl.NewParser(file)
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed to parse full WebView2 IDL: %v", err)
	}

	t.Logf("Successfully parsed full WebView2 IDL")
	t.Logf("Interfaces: %d", len(ast.Interfaces))
	t.Logf("Enums: %d", len(ast.Enums))
	t.Logf("TypeDefs: %d", len(ast.TypeDefs))
}