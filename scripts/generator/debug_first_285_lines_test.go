package generator

import (
	"bytes"
	"os"
	"testing"
	"generator/idl"
)

func TestFirst285Lines(t *testing.T) {
	// Test exactly the first 285 lines (which worked) vs 286 lines (which fails)
	idlData, err := os.ReadFile("../WebView2.1.0.3296.44.idl")
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	lines := bytes.Split(idlData, []byte("\n"))
	
	// Test 285 lines (should work)
	var chunk285 []byte
	for i := 0; i < 285 && i < len(lines); i++ {
		chunk285 = append(chunk285, lines[i]...)
		chunk285 = append(chunk285, '\n')
	}
	
	parser := idl.NewParser(bytes.NewReader(chunk285))
	_, err = parser.Parse()
	if err != nil {
		t.Errorf("285 lines failed unexpectedly: %v", err)
	} else {
		t.Log("285 lines parsed successfully")
	}
	
	// Test 286 lines (should fail)
	var chunk286 []byte
	for i := 0; i < 286 && i < len(lines); i++ {
		chunk286 = append(chunk286, lines[i]...)
		chunk286 = append(chunk286, '\n')
	}
	
	t.Logf("Line 286 content: '%s'", string(lines[285]))
	
	parser = idl.NewParser(bytes.NewReader(chunk286))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("286 lines failed as expected: %v", err)
		
		// Let's try to understand what's happening by parsing just line 286 in isolation
		onlyLine286 := `typedef enum COREWEBVIEW2_BOUNDS_MODE {`
		parser = idl.NewParser(bytes.NewReader([]byte(onlyLine286)))
		_, err = parser.Parse()
		if err != nil {
			t.Logf("Line 286 alone also fails: %v", err)
		} else {
			t.Log("Line 286 alone works - this suggests a context issue")
		}
	} else {
		t.Error("286 lines unexpectedly succeeded")
	}
}

func TestReconstructionFromLines(t *testing.T) {
	// Let's try to reconstruct the problem by adding lines one by one from line 280
	idlData, err := os.ReadFile("../WebView2.1.0.3296.44.idl")
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	lines := bytes.Split(idlData, []byte("\n"))
	
	// Start with a base that we know works (first 280 lines)
	var base []byte
	for i := 0; i < 280 && i < len(lines); i++ {
		base = append(base, lines[i]...)
		base = append(base, '\n')
	}
	
	// Test the base
	parser := idl.NewParser(bytes.NewReader(base))
	_, err = parser.Parse()
	if err != nil {
		t.Fatalf("Base 280 lines failed: %v", err)
	}
	t.Log("Base 280 lines work")
	
	// Now add lines one by one and test
	for lineNum := 280; lineNum < 290 && lineNum < len(lines); lineNum++ {
		testContent := make([]byte, len(base))
		copy(testContent, base)
		
		// Add lines up to lineNum
		for i := 280; i <= lineNum && i < len(lines); i++ {
			testContent = append(testContent, lines[i]...)
			testContent = append(testContent, '\n')
		}
		
		parser = idl.NewParser(bytes.NewReader(testContent))
		_, err = parser.Parse()
		if err != nil {
			t.Logf("Failed at line %d ('%s'): %v", lineNum+1, string(lines[lineNum]), err)
			return
		} else {
			t.Logf("Line %d ('%s') works", lineNum+1, string(lines[lineNum]))
		}
	}
}