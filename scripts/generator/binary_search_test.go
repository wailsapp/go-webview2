package generator

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"generator/idl"
)

func TestBinarySearchIDLError(t *testing.T) {
	// Use binary search to find where the parser fails
	idlPath := "../WebView2.1.0.3296.44.idl"
	
	file, err := os.Open(idlPath)
	if err != nil {
		t.Fatalf("Failed to open IDL file: %v", err)
	}
	defer file.Close()

	// Read all lines
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Binary search for the failing line
	low := 0
	high := len(lines)
	
	for low < high {
		mid := (low + high) / 2
		
		// Create test content with lines 0 to mid
		var content strings.Builder
		for i := 0; i <= mid && i < len(lines); i++ {
			content.WriteString(lines[i] + "\n")
		}
		
		// Make sure we close the library if it was opened
		contentStr := content.String()
		if strings.Contains(contentStr, "library WebView2 {") && !strings.Contains(contentStr, "\n}") {
			contentStr += "\n}\n"
		}
		
		// Try parsing
		parser := idl.NewParser(strings.NewReader(contentStr))
		_, err := parser.Parse()
		
		if err != nil {
			t.Logf("Error at line %d: %v", mid, err)
			t.Logf("Line %d content: %s", mid, lines[mid])
			if mid > 0 {
				t.Logf("Line %d content: %s", mid-1, lines[mid-1])
			}
			if mid > 1 {
				t.Logf("Line %d content: %s", mid-2, lines[mid-2])
			}
			high = mid
		} else {
			t.Logf("Successfully parsed up to line %d", mid)
			low = mid + 1
		}
	}
	
	if low < len(lines) {
		t.Logf("First failing line is around: %d", low)
		if low < len(lines) {
			t.Logf("Line %d: %s", low, lines[low])
		}
		if low-1 >= 0 {
			t.Logf("Line %d: %s", low-1, lines[low-1])
		}
		if low-2 >= 0 {
			t.Logf("Line %d: %s", low-2, lines[low-2])
		}
	} else {
		t.Logf("No failing line found - this shouldn't happen!")
	}
}