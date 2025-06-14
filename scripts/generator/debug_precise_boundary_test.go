package generator

import (
	"bytes"
	"os"
	"testing"
	"generator/idl"
)

func TestFindExactFailurePoint(t *testing.T) {
	// Read the entire IDL file
	idlData, err := os.ReadFile("../WebView2.1.0.3296.44.idl")
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	// Split into lines
	lines := bytes.Split(idlData, []byte("\n"))
	
	// Try parsing progressively larger chunks, starting from line 1
	// until we hit the failure
	chunkSizes := []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}
	
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
			t.Logf("Failure occurs at chunk size %d lines: %v", chunkSize, err)
			
			// If this is the first failure, try to narrow it down more
			if chunkSize > 100 {
				prevChunkSize := chunkSize - 100
				if chunkSize == 400 {
					prevChunkSize = 300
				}
				
				// Test the previous chunk to confirm it worked
				var prevChunk []byte
				for i := 0; i < prevChunkSize; i++ {
					prevChunk = append(prevChunk, lines[i]...)
					prevChunk = append(prevChunk, '\n')
				}
				
				parser = idl.NewParser(bytes.NewReader(prevChunk))
				_, prevErr := parser.Parse()
				if prevErr == nil {
					t.Logf("Previous chunk size %d lines parsed successfully", prevChunkSize)
					t.Logf("Failure occurs between lines %d and %d", prevChunkSize+1, chunkSize)
					
					// Show the problematic lines
					t.Logf("Lines around the failure point:")
					start := prevChunkSize - 5
					if start < 0 {
						start = 0
					}
					end := chunkSize
					if end > len(lines) {
						end = len(lines)
					}
					
					for i := start; i < end; i++ {
						marker := "   "
						if i >= prevChunkSize {
							marker = ">>>"  // Mark the new lines that caused failure
						}
						t.Logf("%s %d: %s", marker, i+1, string(lines[i]))
					}
				} else {
					t.Logf("Previous chunk also failed: %v", prevErr)
				}
			}
			return
		} else {
			t.Logf("Chunk size %d lines parsed successfully", chunkSize)
		}
	}
	
	t.Log("All chunks parsed successfully - no failure found")
}

func TestSpecificFailureLines(t *testing.T) {
	// Based on the previous test, let's test lines 300-400 more carefully
	idlData, err := os.ReadFile("../WebView2.1.0.3296.44.idl")
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	lines := bytes.Split(idlData, []byte("\n"))
	
	// Try a range that should include the problematic area but start cleanly
	// Let's start from line 280 to capture the complete context
	startLine := 280
	endLine := 420
	
	var chunk []byte
	for i := startLine-1; i < endLine && i < len(lines); i++ {
		chunk = append(chunk, lines[i]...)
		chunk = append(chunk, '\n')
	}
	
	t.Logf("Testing lines %d to %d:", startLine, endLine)
	
	parser := idl.NewParser(bytes.NewReader(chunk))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Parse error: %v", err)
		
		// Show a few lines around what we think is the problem
		t.Logf("Content around the problem area:")
		for i := startLine-1; i < startLine+40 && i < len(lines); i++ {
			t.Logf("  %d: %s", i+1, string(lines[i]))
		}
	} else {
		t.Log("Parse succeeded")
	}
}