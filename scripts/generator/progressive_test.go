package generator

import (
	"os"
	"strings"
	"testing"

	"generator/idl"
)

func TestProgressiveParsing(t *testing.T) {
	// Test progressively larger sections of the IDL file
	idlPath := "../WebView2.1.0.3296.44.idl"
	
	content, err := os.ReadFile(idlPath)
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	
	// Test at different line counts
	testLines := []int{350, 400, 500, 600, 700, 800, 1000, 1500, 2000}
	
	for _, lineCount := range testLines {
		if lineCount >= len(lines) {
			lineCount = len(lines) - 1
		}
		
		// Create test content
		var testContent strings.Builder
		for i := 0; i < lineCount; i++ {
			testContent.WriteString(lines[i] + "\n")
		}
		
		// Make sure we close the library properly
		contentStr := testContent.String()
		if strings.Contains(contentStr, "library WebView2 {") && !strings.Contains(contentStr, "\n}") {
			contentStr += "\n}\n"
		}
		
		// Try parsing
		parser := idl.NewParser(strings.NewReader(contentStr))
		ast, err := parser.Parse()
		
		if err != nil {
			t.Logf("FAILED at line %d: %v", lineCount, err)
			t.Logf("Line %d content: %s", lineCount, lines[lineCount-1])
			
			// Show a few lines around the failure
			start := max(0, lineCount-5)
			end := minInt(len(lines), lineCount+5)
			t.Logf("Context around line %d:", lineCount)
			for i := start; i < end; i++ {
				prefix := "  "
				if i == lineCount-1 {
					prefix = ">>>"
				}
				t.Logf("%s %d: %s", prefix, i+1, lines[i])
			}
			break
		} else {
			t.Logf("SUCCESS up to line %d (enums: %d, interfaces: %d)", lineCount, len(ast.Enums), len(ast.Interfaces))
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}