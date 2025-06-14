package generator

import (
	"bytes"
	"os"
	"testing"
	"generator/idl"
)

func TestLines285And286(t *testing.T) {
	// Read the IDL file and test exactly lines 285-286
	idlData, err := os.ReadFile("../WebView2.1.0.3296.44.idl")
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	lines := bytes.Split(idlData, []byte("\n"))
	
	// Test line 285 alone
	line285 := string(lines[284]) + "\n"  // lines are 0-indexed
	parser := idl.NewParser(bytes.NewReader([]byte(line285)))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Line 285 alone failed: %v", err)
	} else {
		t.Log("Line 285 alone succeeded")
	}
	t.Logf("Line 285 content: '%s'", string(lines[284]))
	
	// Test line 286 alone
	line286 := string(lines[285]) + "\n"  // lines are 0-indexed
	parser = idl.NewParser(bytes.NewReader([]byte(line286)))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Line 286 alone failed: %v", err)
	} else {
		t.Log("Line 286 alone succeeded")
	}
	t.Logf("Line 286 content: '%s'", string(lines[285]))
	
	// Test lines 285-286 together
	lines285_286 := string(lines[284]) + "\n" + string(lines[285]) + "\n"
	parser = idl.NewParser(bytes.NewReader([]byte(lines285_286)))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Lines 285-286 together failed: %v", err)
	} else {
		t.Log("Lines 285-286 together succeeded")
	}
	
	// Test with more context - lines 284-287
	var lines284_287 []byte
	for i := 283; i < 287 && i < len(lines); i++ {
		lines284_287 = append(lines284_287, lines[i]...)
		lines284_287 = append(lines284_287, '\n')
	}
	
	parser = idl.NewParser(bytes.NewReader(lines284_287))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Lines 284-287 failed: %v", err)
	} else {
		t.Log("Lines 284-287 succeeded")
	}
	
	t.Log("Lines 284-287 content:")
	for i := 283; i < 287 && i < len(lines); i++ {
		t.Logf("  %d: %s", i+1, string(lines[i]))
	}
}

func TestSpecificErrorPattern(t *testing.T) {
	// Let's test the specific pattern that causes "expected ';' after typedef" vs "expected '}' to close enum"
	
	// This should give "expected '}' to close enum"
	content1 := `typedef enum TEST {`
	parser := idl.NewParser(bytes.NewReader([]byte(content1)))
	_, err := parser.Parse()
	if err != nil {
		t.Logf("Incomplete enum gives: %v", err)
	}
	
	// This should give "expected ';' after typedef"
	content2 := `typedef int MyInt`
	parser = idl.NewParser(bytes.NewReader([]byte(content2)))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Incomplete typedef gives: %v", err)
	}
	
	// What about attributed incomplete enum?
	content3 := `[v1_enum]
typedef enum TEST {`
	parser = idl.NewParser(bytes.NewReader([]byte(content3)))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Attributed incomplete enum gives: %v", err)
	}
	
	// What if the parse failure happens after the enum definition?
	content4 := `[v1_enum]
typedef enum TEST { VALUE } TEST;
typedef int MyInt`
	parser = idl.NewParser(bytes.NewReader([]byte(content4)))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Complete enum + incomplete typedef gives: %v", err)
	}
}