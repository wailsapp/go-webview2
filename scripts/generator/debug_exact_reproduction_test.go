package generator

import (
	"os"
	"testing"
)

func TestExactUpdateScriptReproduction(t *testing.T) {
	// This test exactly reproduces what the update script does
	
	// Read the IDL file like the update script does
	idlData, err := os.ReadFile("../WebView2.1.0.3296.44.idl")
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	// Call ParseIDL exactly like the update script does
	files, err := ParseIDL(idlData)
	if err != nil {
		t.Logf("ParseIDL failed with error: %v", err)
		
		// This should be the exact same error as the update script
		if err.Error() != "failed to parse IDL: expected ';' after typedef" {
			t.Logf("Different error than expected. Got: '%s'", err.Error())
			t.Logf("Expected: 'failed to parse IDL: expected ';' after typedef'")
		}
	} else {
		t.Errorf("ParseIDL unexpectedly succeeded")
		if len(files) > 0 {
			t.Logf("Generated %d files", len(files))
		}
	}
}

func TestParseIDLFunction(t *testing.T) {
	// Test just the ParseIDL function with a minimal failing case
	
	// Create content that should fail
	idlContent := `[v1_enum]
typedef enum TEST {`

	files, err := ParseIDL([]byte(idlContent))
	if err != nil {
		t.Logf("ParseIDL failed as expected: %v", err)
	} else {
		t.Errorf("ParseIDL should have failed")
		if len(files) > 0 {
			t.Logf("Generated %d files", len(files))
		}
	}
}

func TestDifferentErrorSources(t *testing.T) {
	// Test different ways to get "expected ';' after typedef" error
	
	// From parseTypedef (line 510)
	content1 := `typedef int MyInt`
	_, err := ParseIDL([]byte(content1))
	if err != nil {
		t.Logf("Test 1 error: %v", err)
	}
	
	// From parseTypedefWithAttributes (line 922)  
	content2 := `[attr]
typedef int MyInt`
	_, err = ParseIDL([]byte(content2))
	if err != nil {
		t.Logf("Test 2 error: %v", err)
	}
	
	// From parseTypedefDeclaration (line 1060)
	content3 := `typedef struct { int x; } MyStruct`
	_, err = ParseIDL([]byte(content3))
	if err != nil {
		t.Logf("Test 3 error: %v", err)
	}
	
	// What about our problematic case?
	content4 := `[v1_enum]
typedef enum TEST { VALUE } TEST`
	_, err = ParseIDL([]byte(content4))
	if err != nil {
		t.Logf("Test 4 error: %v", err)
	} else {
		t.Log("Test 4 succeeded")
	}
}