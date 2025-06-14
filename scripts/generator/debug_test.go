package generator

import (
	"strings"
	"testing"

	"generator/idl"
)

func TestSimpleIDLParsing(t *testing.T) {
	testIDL := `import "oaidl.idl";

typedef enum TEST_ENUM {
    TEST_VALUE_1 = 0,
    TEST_VALUE_2 = 1
} TEST_ENUM;

[uuid("12345678-1234-1234-1234-123456789012")]
interface ITestInterface : IUnknown {
    HRESULT TestMethod([in] BOOL value);
}`

	parser := idl.NewParser(strings.NewReader(testIDL))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Failed to parse simple IDL: %v", err)
	}

	t.Logf("Parsed successfully: %d imports, %d enums, %d interfaces",
		len(ast.Imports), len(ast.Enums), len(ast.Interfaces))

	if len(ast.Interfaces) > 0 && len(ast.Interfaces[0].Methods) > 0 {
		method := ast.Interfaces[0].Methods[0]
		t.Logf("First method: %s with %d parameters", method.Name, len(method.Parameters))
	}
}

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

func TestMultilineParametersParsing(t *testing.T) {
	testIDL := `[uuid("12345678-1234-1234-1234-123456789012")]
interface ITestInterface : IUnknown {
    HRESULT CreateWebResourceRequest(
        [in] LPCWSTR uri,
        [in] LPCWSTR method,
        [in] IStream* postData,
        [in] LPCWSTR headers,
        [out, retval] ICoreWebView2WebResourceRequest** request
    );
}`

	parser := idl.NewParser(strings.NewReader(testIDL))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Failed to parse IDL with multiline parameters: %v", err)
	}

	t.Logf("Parsed successfully: %d interfaces", len(ast.Interfaces))
}

func TestDoublePointerParsing(t *testing.T) {
	testIDL := `[uuid("12345678-1234-1234-1234-123456789012")]
interface ITestInterface : IUnknown {
    HRESULT TestMethod([out, retval] ITest** request);
}`

	parser := idl.NewParser(strings.NewReader(testIDL))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Failed to parse IDL with double pointer: %v", err)
	}

	t.Logf("Parsed successfully: %d interfaces", len(ast.Interfaces))
}
