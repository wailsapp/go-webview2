package generator

import (
	"os"
	"strings"
	"testing"

	"generator/idl"
)

func TestLibraryToFirstEnum(t *testing.T) {
	// Test parsing from the library declaration to just past the first enum
	idlPath := "../WebView2.1.0.3296.44.idl"
	
	content, err := os.ReadFile(idlPath)
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	
	// Find the library line (around line 11)
	libraryStart := -1
	for i, line := range lines {
		if strings.Contains(line, "library WebView2") {
			libraryStart = i
			break
		}
	}
	
	if libraryStart == -1 {
		t.Fatal("Could not find library declaration")
	}
	
	// Find the first enum (around line 286)
	firstEnumEnd := -1
	for i := libraryStart; i < len(lines); i++ {
		if strings.Contains(lines[i], "} COREWEBVIEW2_BOUNDS_MODE;") {
			firstEnumEnd = i
			break
		}
	}
	
	if firstEnumEnd == -1 {
		t.Fatal("Could not find first enum end")
	}
	
	// Include the complete imports and library structure up to the first enum
	var testContent strings.Builder
	
	// Add imports
	for i := 0; i < libraryStart; i++ {
		if strings.Contains(lines[i], "import") || strings.Contains(lines[i], "[uuid") {
			testContent.WriteString(lines[i] + "\n")
		}
	}
	
	// Add from library start to first enum end
	for i := libraryStart; i <= firstEnumEnd; i++ {
		testContent.WriteString(lines[i] + "\n")
	}
	
	// Close the library
	testContent.WriteString("}\n")
	
	testStr := testContent.String()
	t.Logf("Content from library to first enum (first 1000 chars):\n%s", testStr[:min(len(testStr), 1000)])
	
	// Now try parsing
	parser := idl.NewParser(strings.NewReader(testStr))
	ast, err := parser.Parse()
	
	if err != nil {
		t.Fatalf("Failed to parse library to first enum: %v", err)
	}
	
	t.Logf("Successfully parsed library to first enum!")
	t.Logf("Enums found: %d", len(ast.Enums))
	if len(ast.Enums) > 0 {
		t.Logf("First enum: %s", ast.Enums[0].Name)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}