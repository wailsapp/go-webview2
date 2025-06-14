package generator

import (
	"bytes"
	"os"
	"testing"
	"generator/idl"
)

func TestExactFailureLocation(t *testing.T) {
	// Read the IDL file
	idlData, err := os.ReadFile("../WebView2.1.0.3296.44.idl")
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	lines := bytes.Split(idlData, []byte("\n"))
	
	// Test the exact area where we saw the failure
	startLine := 280
	endLine := 320
	
	var chunk []byte
	for i := startLine-1; i < endLine && i < len(lines); i++ {
		chunk = append(chunk, lines[i]...)
		chunk = append(chunk, '\n')
	}
	
	t.Logf("Testing lines %d to %d:", startLine, endLine)
	
	// Show the content we're parsing
	for i := startLine-1; i < endLine && i < len(lines); i++ {
		t.Logf("  %d: %s", i+1, string(lines[i]))
	}
	
	parser := idl.NewParser(bytes.NewReader(chunk))
	ast, err := parser.Parse()
	if err != nil {
		t.Errorf("Parse error: %v", err)
	} else {
		t.Logf("Parse succeeded! Found %d enums", len(ast.Enums))
		for _, enum := range ast.Enums {
			t.Logf("  Enum: %s with %d values", enum.Name, len(enum.Values))
		}
	}
}

func TestBinarySearchForFailure(t *testing.T) {
	// Read the IDL file
	idlData, err := os.ReadFile("../WebView2.1.0.3296.44.idl")
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	lines := bytes.Split(idlData, []byte("\n"))
	
	// We know it works at 200 lines and fails at 300, so binary search between 200-300
	low := 200
	high := 300
	
	for low < high {
		mid := (low + high) / 2
		
		var chunk []byte
		for i := 0; i < mid && i < len(lines); i++ {
			chunk = append(chunk, lines[i]...)
			chunk = append(chunk, '\n')
		}
		
		parser := idl.NewParser(bytes.NewReader(chunk))
		_, err = parser.Parse()
		if err != nil {
			t.Logf("Lines 1-%d failed: %v", mid, err)
			high = mid
		} else {
			t.Logf("Lines 1-%d succeeded", mid)
			low = mid + 1
		}
	}
	
	t.Logf("Failure occurs at line %d", low)
	
	// Show lines around the failure point
	if low > 0 && low <= len(lines) {
		start := low - 10
		end := low + 10
		if start < 0 {
			start = 0
		}
		if end > len(lines) {
			end = len(lines)
		}
		
		t.Logf("Lines around failure point:")
		for i := start; i < end; i++ {
			marker := "   "
			if i+1 == low {
				marker = ">>>"
			}
			t.Logf("%s %d: %s", marker, i+1, string(lines[i]))
		}
	}
}