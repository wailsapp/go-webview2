package generator

import (
	"bytes"
	"os"
	"testing"
	"generator/idl"
)

func TestParserTransitionIssue(t *testing.T) {
	// Test if there's an issue with parser state after many forward declarations
	
	// Read the first 284 lines (which work)
	idlData, err := os.ReadFile("../WebView2.1.0.3296.44.idl")
	if err != nil {
		t.Fatalf("Failed to read IDL file: %v", err)
	}

	lines := bytes.Split(idlData, []byte("\n"))
	
	// Build content with first 284 lines
	var content284 []byte
	for i := 0; i < 284 && i < len(lines); i++ {
		content284 = append(content284, lines[i]...)
		content284 = append(content284, '\n')
	}
	
	// Test that 284 lines work
	parser := idl.NewParser(bytes.NewReader(content284))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("284 lines failed: %v", err)
	}
	t.Logf("284 lines parsed successfully: %d interfaces, %d enums", len(ast.Interfaces), len(ast.Enums))
	
	// Now add line 285 ([v1_enum])
	content285 := make([]byte, len(content284))
	copy(content285, content284)
	content285 = append(content285, lines[284]...)
	content285 = append(content285, '\n')
	
	parser = idl.NewParser(bytes.NewReader(content285))
	ast, err = parser.Parse()
	if err != nil {
		t.Fatalf("285 lines failed: %v", err)
	}
	t.Logf("285 lines parsed successfully: %d interfaces, %d enums", len(ast.Interfaces), len(ast.Enums))
	
	// The key test: what happens when we have 285 lines + incomplete line 286
	content286Incomplete := make([]byte, len(content285))
	copy(content286Incomplete, content285)
	// Add just the beginning of line 286
	content286Incomplete = append(content286Incomplete, []byte("typedef")...)
	
	parser = idl.NewParser(bytes.NewReader(content286Incomplete))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("285 lines + 'typedef' failed: %v", err)
	} else {
		t.Log("285 lines + 'typedef' succeeded")
	}
	
	// Test with "typedef enum"
	content286IncompleteEnum := make([]byte, len(content285))
	copy(content286IncompleteEnum, content285)
	content286IncompleteEnum = append(content286IncompleteEnum, []byte("typedef enum")...)
	
	parser = idl.NewParser(bytes.NewReader(content286IncompleteEnum))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("285 lines + 'typedef enum' failed: %v", err)
	} else {
		t.Log("285 lines + 'typedef enum' succeeded")
	}
	
	// Test with the full line 286
	content286Full := make([]byte, len(content285))
	copy(content286Full, content285)
	content286Full = append(content286Full, lines[285]...)
	content286Full = append(content286Full, '\n')
	
	parser = idl.NewParser(bytes.NewReader(content286Full))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("286 lines failed: %v", err)
		
		// This is the key failure case - let's see what line 286 actually contains
		t.Logf("Line 286 content: '%s'", string(lines[285]))
		
		// Let's try parsing just the problematic transition
		transitionContent := `[v1_enum]
typedef enum COREWEBVIEW2_BOUNDS_MODE {`
		
		parser = idl.NewParser(bytes.NewReader([]byte(transitionContent)))
		_, err = parser.Parse()
		if err != nil {
			t.Logf("Transition content alone failed: %v", err)
		} else {
			t.Log("Transition content alone succeeded")
		}
		
	} else {
		t.Log("286 lines succeeded")
	}
}

func TestParserStateAfterInterfaces(t *testing.T) {
	// Investigate what happens to parser state after processing interface declarations
	
	// Create content similar to what's before line 285
	interfaceDecls := `interface ICoreWebView2WebResourceResponse;
interface ICoreWebView2WebResourceResponseReceivedEventArgs;
interface ICoreWebView2WebResourceResponseView;
interface ICoreWebView2WebResourceResponseViewGetContentCompletedHandler;
interface ICoreWebView2WindowFeatures;

// Enums and structs


/// Mode for how the Bounds property is interpreted in relation to the RasterizationScale property.`

	// Test that this part works
	parser := idl.NewParser(bytes.NewReader([]byte(interfaceDecls)))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Interface declarations failed: %v", err)
	}
	t.Logf("Interface declarations parsed: %d interfaces", len(ast.Interfaces))
	
	// Now add the attribute and see what happens
	withAttribute := interfaceDecls + "\n[v1_enum]"
	parser = idl.NewParser(bytes.NewReader([]byte(withAttribute)))
	ast, err = parser.Parse()
	if err != nil {
		t.Fatalf("Interface declarations + attribute failed: %v", err)
	}
	t.Log("Interface declarations + attribute succeeded")
	
	// Now add typedef enum and see what happens
	withTypedefEnum := withAttribute + "\ntypedef enum TEST {"
	parser = idl.NewParser(bytes.NewReader([]byte(withTypedefEnum)))
	_, err = parser.Parse()
	if err != nil {
		t.Logf("Interface declarations + attribute + typedef enum failed: %v", err)
	} else {
		t.Log("Interface declarations + attribute + typedef enum succeeded")
	}
}