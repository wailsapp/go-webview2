package generator

import (
	"bufio"
	"os"
	"testing"

	"generator/idl"
)

func TestDebugFullWebView2IDLParsing(t *testing.T) {
	// Test parsing of the full WebView2 IDL file with debugging
	idlPath := "../WebView2.1.0.3296.44.idl"
	
	file, err := os.Open(idlPath)
	if err != nil {
		t.Fatalf("Failed to open IDL file %s: %v", idlPath, err)
	}
	defer file.Close()

	// First, let's read the file line by line to find typedef declarations
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	lineNum := 0
	
	t.Log("Looking for typedef declarations:")
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if line == "" {
			continue
		}
		
		// Look for typedef lines and a few lines around them
		if lineNum >= 280 && lineNum <= 320 {
			t.Logf("Line %d: %s", lineNum, line)
		}
	}

	// Now try parsing with better error context
	file.Seek(0, 0)
	parser := idl.NewParser(file)
	ast, err := parser.Parse()
	if err != nil {
		t.Logf("Parser failed with error: %v", err)
		
		// Let's try to get more context about where it failed
		file.Seek(0, 0)
		content, _ := os.ReadFile(idlPath)
		t.Logf("First 2000 characters of IDL file:")
		if len(content) > 2000 {
			t.Logf("%s", string(content[:2000]))
		} else {
			t.Logf("%s", string(content))
		}
		
		t.Fatalf("Failed to parse: %v", err)
	}

	t.Logf("Successfully parsed full WebView2 IDL")
	t.Logf("Interfaces: %d", len(ast.Interfaces))
	t.Logf("Enums: %d", len(ast.Enums))
	t.Logf("TypeDefs: %d", len(ast.TypeDefs))
}