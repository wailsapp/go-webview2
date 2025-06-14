package generator

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"generator/idl"
)

func TestPreciseDebugParsing(t *testing.T) {
	// Test parsing step by step to find the exact location of failure
	idlPath := "../WebView2.1.0.3296.44.idl"
	
	// First, read and parse just the first few enum declarations to see which one fails
	file, err := os.Open(idlPath)
	if err != nil {
		t.Fatalf("Failed to open IDL file: %v", err)
	}
	defer file.Close()

	// Extract just the first few typedef enum declarations with their cpp_quote directives
	var content strings.Builder
	scanner := bufio.NewScanner(file)
	inLibrary := false
	enumCount := 0
	targetEnums := 5 // Test first 5 enums
	
	for scanner.Scan() {
		line := scanner.Text()
		
		// Add headers and library declaration
		if strings.Contains(line, "import") || strings.Contains(line, "library") || strings.Contains(line, "[uuid") {
			content.WriteString(line + "\n")
			if strings.Contains(line, "library") {
				inLibrary = true
			}
			continue
		}
		
		if !inLibrary {
			continue
		}
		
		// Add enum declarations and related content
		if strings.Contains(line, "typedef enum") || strings.Contains(line, "[v1_enum]") ||
		   strings.Contains(line, "cpp_quote") ||
		   (strings.TrimSpace(line) != "" && enumCount < targetEnums && 
		    (strings.Contains(line, "COREWEBVIEW2_") || strings.Contains(line, "///") || 
		     strings.Contains(line, "}") || strings.Contains(line, "="))) {
			content.WriteString(line + "\n")
			
			if strings.Contains(line, "typedef enum") {
				enumCount++
			}
			
			if enumCount >= targetEnums && strings.Contains(line, ";") {
				break
			}
		}
	}
	
	// Close the library
	content.WriteString("}\n")
	
	testIDL := content.String()
	t.Logf("Test IDL content:\n%s", testIDL)
	
	// Now try parsing this subset
	parser := idl.NewParser(strings.NewReader(testIDL))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed to parse IDL subset: %v", err)
	}

	t.Logf("Successfully parsed IDL subset!")
	t.Logf("Enums found: %d", len(ast.Enums))
	for i, enum := range ast.Enums {
		t.Logf("  Enum %d: %s (%d values)", i+1, enum.Name, len(enum.Values))
	}
}