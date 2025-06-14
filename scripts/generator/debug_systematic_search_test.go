package generator

import (
	"bytes"
	"os"
	"testing"
	"generator/idl"
)

func TestSystematicLineByLineSearch(t *testing.T) {
	// Read the IDL file
	idlData, err := os.ReadFile("../WebView2.1.0.3296.44.idl")
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	lines := bytes.Split(idlData, []byte("\n"))
	t.Logf("Total lines in file: %d", len(lines))
	
	// Start with small chunks and work our way up
	chunkSizes := []int{10, 20, 50, 100, 150, 200, 250, 280, 285, 286, 287, 300}
	
	for _, chunkSize := range chunkSizes {
		if chunkSize > len(lines) {
			chunkSize = len(lines)
		}
		
		var chunk []byte
		for i := 0; i < chunkSize; i++ {
			chunk = append(chunk, lines[i]...)
			chunk = append(chunk, '\n')
		}
		
		parser := idl.NewParser(bytes.NewReader(chunk))
		_, err = parser.Parse()
		if err != nil {
			t.Logf("FAILURE at chunk size %d lines: %v", chunkSize, err)
			
			// If this is the first failure, let's examine what's different
			if chunkSize == 10 {
				t.Log("Even 10 lines fail - there's an issue very early in the file")
				t.Log("First 10 lines:")
				for i := 0; i < 10 && i < len(lines); i++ {
					t.Logf("  %d: %s", i+1, string(lines[i]))
				}
			} else {
				// Find the previous working size
				prevSize := 0
				for _, size := range chunkSizes {
					if size < chunkSize {
						prevSize = size
					} else {
						break
					}
				}
				t.Logf("Previous size %d worked, %d failed", prevSize, chunkSize)
				
				// Show the lines that caused the failure
				t.Logf("Lines that caused failure (%d to %d):", prevSize+1, chunkSize)
				for i := prevSize; i < chunkSize && i < len(lines); i++ {
					t.Logf("  %d: %s", i+1, string(lines[i]))
				}
			}
			return
		} else {
			t.Logf("SUCCESS: chunk size %d lines parsed successfully", chunkSize)
		}
	}
	
	t.Log("All chunks parsed successfully - no failure found in these sizes")
}

func TestFirstFewLines(t *testing.T) {
	// Let's specifically test the very beginning of the file
	idlData, err := os.ReadFile("../WebView2.1.0.3296.44.idl")
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	lines := bytes.Split(idlData, []byte("\n"))
	
	// Test just the first line
	parser := idl.NewParser(bytes.NewReader([]byte(string(lines[0]) + "\n")))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("First line failed: %v", err)
		t.Logf("First line content: '%s'", string(lines[0]))
	} else {
		t.Log("First line succeeded")
	}
	
	// Test first 3 lines
	var first3 []byte
	for i := 0; i < 3 && i < len(lines); i++ {
		first3 = append(first3, lines[i]...)
		first3 = append(first3, '\n')
	}
	
	parser = idl.NewParser(bytes.NewReader(first3))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("First 3 lines failed: %v", err)
		t.Logf("First 3 lines:")
		for i := 0; i < 3 && i < len(lines); i++ {
			t.Logf("  %d: %s", i+1, string(lines[i]))
		}
	} else {
		t.Log("First 3 lines succeeded")
	}
}