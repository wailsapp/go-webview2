package generator

import (
	"bytes"
	"os"
	"testing"
	"generator/idl"
)

func TestFindActualParsingFailure(t *testing.T) {
	// Let's create a version of the parser that gives us more information
	// about where exactly it fails
	
	idlData, err := os.ReadFile("../WebView2.1.0.3296.44.idl")
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	// Use the original idl parser directly instead of ParseIDL
	parser := idl.NewParser(bytes.NewReader(idlData))
	ast, err := parser.Parse()
	if err != nil {
		t.Logf("Direct parser failed: %v", err)
		
		// Let's try to understand the context better
		// The error should be the raw parser error, not wrapped by ParseIDL
		if err.Error() == "expected ';' after typedef" {
			t.Log("This is the raw 'expected ';' after typedef' error")
		} else {
			t.Logf("Different error from direct parser: %s", err.Error())
		}
		
		// Show what was parsed before the failure (if ast is not nil)
		if ast != nil {
			t.Logf("Parsed before failure:")
			t.Logf("  Interfaces: %d", len(ast.Interfaces))
			t.Logf("  Enums: %d", len(ast.Enums))
			t.Logf("  TypeDefs: %d", len(ast.TypeDefs))
			
			if len(ast.Enums) > 0 {
				t.Logf("  Last enum parsed: %s", ast.Enums[len(ast.Enums)-1].Name)
			}
		} else {
			t.Log("AST is nil after parse failure")
		}
		
	} else {
		t.Error("Direct parser unexpectedly succeeded")
	}
}

func TestDetermineWhichTypedefPath(t *testing.T) {
	// Let's figure out which typedef parsing path is being taken
	// by creating test cases for each path
	
	// Path 1: parseTypedef (regular typedef, not attributed)
	content1 := `typedef int MyInt;`
	parser := idl.NewParser(bytes.NewReader([]byte(content1)))
	_, err := parser.Parse()
	if err != nil {
		t.Logf("Path 1 (parseTypedef) error: %v", err)
	} else {
		t.Log("Path 1 (parseTypedef) succeeded")
	}
	
	// Path 2: parseTypedefWithAttributes (attributed typedef)
	content2 := `[attr]
typedef int MyInt;`
	parser = idl.NewParser(bytes.NewReader([]byte(content2)))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Path 2 (parseTypedefWithAttributes) error: %v", err)
	} else {
		t.Log("Path 2 (parseTypedefWithAttributes) succeeded")
	}
	
	// Path 3: parseTypedefDeclaration (called from parseTypedefWithAttributes)
	// This one is harder to trigger directly
	
	// Path 4: parseTypedefEnum (attributed typedef enum)
	content4 := `[v1_enum]
typedef enum TEST { VALUE } TEST;`
	parser = idl.NewParser(bytes.NewReader([]byte(content4)))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Path 4 (parseTypedefEnum) error: %v", err)
	} else {
		t.Log("Path 4 (parseTypedefEnum) succeeded")
	}
	
	// Path 5: What if there's no semicolon?
	content5 := `[v1_enum]
typedef enum TEST { VALUE } TEST`
	parser = idl.NewParser(bytes.NewReader([]byte(content5)))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Path 5 (missing semicolon) error: %v", err)
	}
}