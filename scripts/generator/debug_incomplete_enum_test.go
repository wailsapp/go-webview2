package generator

import (
	"bytes"
	"testing"
	"generator/idl"
)

func TestIncompleteEnumTheory(t *testing.T) {
	// Test what happens when we parse an incomplete enum followed by more content
	
	// This should fail because the enum is incomplete
	idlContent1 := `[v1_enum]
typedef enum COREWEBVIEW2_BOUNDS_MODE {
interface ICoreWebView2Test;`

	parser := idl.NewParser(bytes.NewReader([]byte(idlContent1)))
	_, err := parser.Parse()
	if err != nil {
		t.Logf("Incomplete enum + interface failed as expected: %v", err)
	} else {
		t.Error("Incomplete enum + interface unexpectedly succeeded")
	}
	
	// This should also fail with a different error
	idlContent2 := `[v1_enum]
typedef enum COREWEBVIEW2_BOUNDS_MODE {
[v1_enum]
typedef enum TEST { VALUE } TEST;`

	parser = idl.NewParser(bytes.NewReader([]byte(idlContent2)))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Incomplete enum + complete enum failed: %v", err)
	} else {
		t.Error("Incomplete enum + complete enum unexpectedly succeeded")
	}
	
	// The real issue might be that when we test line 286 in isolation:
	// "typedef enum COREWEBVIEW2_BOUNDS_MODE {"
	// The parser enters parseAttributedDeclaration due to some leftover state
	
	// Let me test if there's a parser state issue by parsing multiple times
	for i := 0; i < 3; i++ {
		parser = idl.NewParser(bytes.NewReader([]byte("typedef enum TEST {")))
		_, err = parser.Parse()
		t.Logf("Attempt %d: %v", i+1, err)
	}
}

func TestParserReset(t *testing.T) {
	// Test if the parser state is properly reset between parses
	
	// First parse - this should fail
	parser := idl.NewParser(bytes.NewReader([]byte("typedef enum TEST {")))
	_, err := parser.Parse()
	if err != nil {
		t.Logf("First parse failed as expected: %v", err)
	}
	
	// Second parse with a new parser - this should also fail the same way
	parser2 := idl.NewParser(bytes.NewReader([]byte("typedef enum TEST {")))
	_, err2 := parser2.Parse()
	if err2 != nil {
		t.Logf("Second parse failed as expected: %v", err2)
		if err.Error() != err2.Error() {
			t.Errorf("Different error messages: '%s' vs '%s'", err.Error(), err2.Error())
		}
	}
	
	// Third parse with complete enum - this should work
	parser3 := idl.NewParser(bytes.NewReader([]byte("typedef enum TEST { VALUE } TEST;")))
	_, err3 := parser3.Parse()
	if err3 != nil {
		t.Errorf("Complete enum failed: %v", err3)
	} else {
		t.Log("Complete enum succeeded")
	}
}

func TestExactReproduction(t *testing.T) {
	// Now let's try to exactly reproduce the original issue
	// The original error was "expected ';' after typedef" but our tests show "expected '}' to close enum"
	// Maybe the issue is in the update script path vs our test path
	
	// Let's create the exact context that the update script would see
	
	// Build up content line by line like the failing case
	lines := []string{
		"// Some content",
		"interface ICoreWebView2;",
		"",
		"// More content", 
		"/// Comment",
		"[v1_enum]",
		"typedef enum COREWEBVIEW2_BOUNDS_MODE {",
		"  COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS,",
		"} COREWEBVIEW2_BOUNDS_MODE;",
	}
	
	// Test progressively adding lines to see when it breaks
	for i := 0; i < len(lines); i++ {
		var content bytes.Buffer
		for j := 0; j <= i; j++ {
			content.WriteString(lines[j])
			content.WriteString("\n")
		}
		
		parser := idl.NewParser(bytes.NewReader(content.Bytes()))
		_, err := parser.Parse()
		if err != nil {
			t.Logf("Failed at line %d ('%s'): %v", i+1, lines[i], err)
		} else {
			t.Logf("Succeeded through line %d ('%s')", i+1, lines[i])
		}
	}
}