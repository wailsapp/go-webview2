package generator

import (
	"os"
	"strings"
	"testing"

	"generator/idl"
)

func TestTokenizeDebugAroundLine285(t *testing.T) {
	// Read the IDL file and examine the tokens around line 285
	idlPath := "../WebView2.1.0.3296.44.idl"
	
	content, err := os.ReadFile(idlPath)
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	
	// Extract lines 280-295 to see the exact content and tokens
	var testContent strings.Builder
	for i := 279; i < 295 && i < len(lines); i++ { // lines are 0-indexed
		testContent.WriteString(lines[i] + "\n")
	}
	
	testStr := testContent.String()
	t.Logf("Content around line 285:\n%s", testStr)
	
	// Now scan the tokens
	scanner := idl.NewScanner(strings.NewReader(testStr))
	tokenCount := 0
	for {
		token := scanner.NextToken()
		if token.Type == idl.TokenEOF {
			break
		}
		tokenCount++
		if tokenCount < 50 { // Show first 50 tokens
			t.Logf("Token %d: Type=%d, Value='%s', Line=%d, Col=%d", 
				tokenCount, int(token.Type), token.Value, token.Line, token.Col)
		}
	}
	
	t.Logf("Total tokens scanned: %d", tokenCount)
}