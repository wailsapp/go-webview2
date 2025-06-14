package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"generator/idl"
)

func TestCodeGeneration(t *testing.T) {
	// Create test IDL content
	testIDL := `
    typedef enum {
        TEST_VALUE_1 = 0,
        TEST_VALUE_2 = 1
    } TestEnum;
    
    typedef DWORD TestTypedef;
    
    [uuid("12345678-1234-1234-1234-123456789012")]
    interface ITestInterface : IUnknown {
        HRESULT TestMethod([in] BOOL value);
        DWORD GetTestValue();
    };
    `

	// Parse the IDL
	parser := idl.NewParser(strings.NewReader(testIDL))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed to parse test IDL: %v", err)
	}

	// Generate code
	generator, err := NewGenerator()
	if err != nil {
		t.Fatalf("Failed to create generator: %v", err)
	}

	code, err := generator.GenerateCode(ast, "test-version")
	if err != nil {
		t.Fatalf("Failed to generate code: %v", err)
	}

	// Verify generated code contains expected elements
	testCases := []struct {
		name     string
		expected string
	}{
		{"Package header", "package webview2"},
		{"Enum type", "type TestEnum uint32"},
		{"Enum constants", "TEST_VALUE_1 TestEnum = 0"},
		{"Typedef", "type TestTypedef uint32"},
		{"Interface type", "type ITestInterface struct"},
		{"Method", "func (obj *ITestInterface) TestMethod"},
		{"VTable", "type ITestInterfaceVTable struct"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if !strings.Contains(code, tc.expected) {
				t.Errorf("Generated code missing %s: expected to contain %q", tc.name, tc.expected)
				t.Logf("Generated code:\n%s", code)
			}
		})
	}
}

func TestGeneratedCodeCompilation(t *testing.T) {
	// Skip if not on Windows (since we're generating Windows-specific code)
	if os.Getenv("GOOS") != "windows" && os.Getenv("GOOS") != "" {
		t.Skip("Skipping Windows-specific compilation test")
	}

	// Create test IDL content
	testIDL := `
    typedef enum {
        SIMPLE_ENUM_VALUE = 42
    } SimpleEnum;
    `

	// Parse and generate
	parser := idl.NewParser(strings.NewReader(testIDL))
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed to parse test IDL: %v", err)
	}

	generator, err := NewGenerator()
	if err != nil {
		t.Fatalf("Failed to create generator: %v", err)
	}

	code, err := generator.GenerateCode(ast, "test-version")
	if err != nil {
		t.Fatalf("Failed to generate code: %v", err)
	}

	// Write to temporary file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "generated_test.go")

	if err := os.WriteFile(testFile, []byte(code), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Create a basic go.mod for the test
	goMod := `module test
go 1.20

require github.com/wailsapp/go-webview2 v0.0.0
replace github.com/wailsapp/go-webview2 => ../../../..
`

	goModFile := filepath.Join(tmpDir, "go.mod")
	if err := os.WriteFile(goModFile, []byte(goMod), 0644); err != nil {
		t.Fatalf("Failed to write go.mod: %v", err)
	}

	// TODO: Add actual compilation test using go/build or exec.Command
	// For now, we'll just verify the file was created
	if _, err := os.Stat(testFile); err != nil {
		t.Errorf("Generated test file not found: %v", err)
	}
}

func TestSyscallGeneration(t *testing.T) {
	// Test syscall wrapper generation
	generator := NewSyscallGenerator()

	// Create a test method
	method := &idl.Method{
		Name:       "TestMethod",
		Comments:   []string{"Test method comment"},
		ReturnType: &idl.Type{Name: "HRESULT", Kind: idl.TypeBasic},
		Parameters: []*idl.Parameter{
			{
				Name: "value",
				Type: &idl.Type{Name: "BOOL", Kind: idl.TypeBasic},
			},
			{
				Name: "stringPtr",
				Type: &idl.Type{Name: "LPWSTR", Kind: idl.TypeBasic},
			},
		},
	}

	// Test method signature generation
	signature := generator.GenerateMethodSignature(method)
	expected := "TestMethod(value int32, stringPtr *uint16) error"
	if signature != expected {
		t.Errorf("Method signature mismatch. Expected: %q, Got: %q", expected, signature)
	}

	// Test syscall parameter generation
	params := generator.GenerateSyscallParams(method)
	expectedParams := []string{
		"uintptr(unsafe.Pointer(obj))",
		"uintptr(boolToUint32(value))",
		"uintptr(unsafe.Pointer(stringPtr))",
	}

	if len(params) != len(expectedParams) {
		t.Errorf("Parameter count mismatch. Expected: %d, Got: %d", len(expectedParams), len(params))
	}

	for i, param := range params {
		if i < len(expectedParams) && param != expectedParams[i] {
			t.Errorf("Parameter %d mismatch. Expected: %q, Got: %q", i, expectedParams[i], param)
		}
	}
}

func BenchmarkCodeGeneration(b *testing.B) {
	// Simple benchmark for code generation
	testIDL := `
    typedef enum {
        BENCH_VALUE_1 = 0,
        BENCH_VALUE_2 = 1,
        BENCH_VALUE_3 = 2
    } BenchEnum;
    `

	parser := idl.NewParser(strings.NewReader(testIDL))
	ast, err := parser.Parse()
	if err != nil {
		b.Fatalf("Failed to parse test IDL: %v", err)
	}

	generator, err := NewGenerator()
	if err != nil {
		b.Fatalf("Failed to create generator: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := generator.GenerateCode(ast, "bench-version")
		if err != nil {
			b.Fatalf("Code generation failed: %v", err)
		}
	}
}
