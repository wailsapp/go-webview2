package generator

import (
	"strings"
	"testing"

	"generator/idl"
)

func TestParameterAttributesParsing(t *testing.T) {
	testIDL := `[uuid("12345678-1234-1234-1234-123456789012")]
interface ITestInterface : IUnknown {
    HRESULT TestMethod([out, retval] BOOL* value);
}`

	parser := idl.NewParser(strings.NewReader(testIDL))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Failed to parse IDL with parameter attributes: %v", err)
	}

	t.Logf("Parsed successfully: %d interfaces", len(ast.Interfaces))
}
