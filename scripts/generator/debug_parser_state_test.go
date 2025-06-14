package generator

import (
	"bytes"
	"testing"
	"generator/idl"
)

func TestParserStateIssue(t *testing.T) {
	// The key insight is that the error "expected '}' to close enum" suggests
	// the parser thinks it's inside an enum when it hits [v1_enum]
	
	// Let's test the exact sequence that leads to this state
	idlContent := `[v1_enum]
typedef enum COREWEBVIEW2_BOUNDS_MODE {`

	parser := idl.NewParser(bytes.NewReader([]byte(idlContent)))
	_, err := parser.Parse()
	if err != nil {
		t.Logf("Parse error: %v", err)
	}
	
	// The error suggests that when parseAttributedDeclaration is called with [v1_enum],
	// and then it sees 'typedef', it calls parseTypedefEnumWithAttributes
	// But somehow the parser thinks it's already inside an enum
	
	// Let's test step by step
	
	// Test 1: Just the attribute
	parser = idl.NewParser(bytes.NewReader([]byte("[v1_enum]")))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Just attribute failed: %v", err)
	} else {
		t.Log("Just attribute succeeded")
	}
	
	// Test 2: Attribute + typedef
	parser = idl.NewParser(bytes.NewReader([]byte("[v1_enum]\ntypedef")))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Attribute + typedef failed: %v", err)
	} else {
		t.Log("Attribute + typedef succeeded")
	}
	
	// Test 3: Attribute + typedef + enum
	parser = idl.NewParser(bytes.NewReader([]byte("[v1_enum]\ntypedef enum")))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Attribute + typedef + enum failed: %v", err)
	} else {
		t.Log("Attribute + typedef + enum succeeded")
	}
	
	// Test 4: The actual failing sequence
	parser = idl.NewParser(bytes.NewReader([]byte("[v1_enum]\ntypedef enum COREWEBVIEW2_BOUNDS_MODE")))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Full sequence failed: %v", err)
	} else {
		t.Log("Full sequence succeeded")
	}
}

func TestTokenSequenceAroundFailure(t *testing.T) {
	// Let's see exactly what tokens are being generated
	idlContent := `[v1_enum]
typedef enum COREWEBVIEW2_BOUNDS_MODE {`

	scanner := idl.NewScanner(bytes.NewReader([]byte(idlContent)))
	
	t.Log("Token sequence:")
	for i := 0; i < 10; i++ {
		token := scanner.NextToken()
		t.Logf("  %d: Type=%d, Value='%s'", i, token.Type, token.Value)
		if token.Type == 0 { // EOF
			break
		}
	}
}

func TestMinimalEnum(t *testing.T) {
	// Test the absolute minimum enum that should work
	idlContent := `typedef enum TEST { VALUE } TEST;`

	parser := idl.NewParser(bytes.NewReader([]byte(idlContent)))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Minimal enum failed: %v", err)
	}
	
	if len(ast.Enums) != 1 {
		t.Errorf("Expected 1 enum, got %d", len(ast.Enums))
	}
}

func TestWithAttributeMinimal(t *testing.T) {
	// Test the minimal enum with attributes
	idlContent := `[v1_enum]
typedef enum TEST { VALUE } TEST;`

	parser := idl.NewParser(bytes.NewReader([]byte(idlContent)))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Attributed minimal enum failed: %v", err)
	}
	
	if len(ast.Enums) != 1 {
		t.Errorf("Expected 1 enum, got %d", len(ast.Enums))
	}
}