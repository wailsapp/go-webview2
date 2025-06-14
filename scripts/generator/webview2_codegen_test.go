package generator

import (
	"os"
	"strings"
	"testing"

	"generator/idl"
)

func TestWebView2CodeGeneration(t *testing.T) {
	// Test code generation with real WebView2 IDL subset
	subsetPath := "test_subset.idl"
	if _, err := os.Stat(subsetPath); os.IsNotExist(err) {
		t.Skip("test_subset.idl not found, skipping WebView2 code generation test")
	}

	file, err := os.Open(subsetPath)
	if err != nil {
		t.Fatalf("Error opening test subset: %v", err)
	}
	defer file.Close()

	// Parse the IDL
	parser := idl.NewParser(file)
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Error parsing WebView2 subset: %v", err)
	}

	// Generate code
	generator, err := NewGenerator()
	if err != nil {
		t.Fatalf("Failed to create generator: %v", err)
	}

	code, err := generator.GenerateCode(ast, "webview2-subset")
	if err != nil {
		t.Fatalf("Failed to generate code: %v", err)
	}

	// Verify generated code contains expected WebView2 elements
	testCases := []struct {
		name     string
		expected string
	}{
		{"Package header", "package webview2"},
		{"WebView2 enum", "type COREWEBVIEW2_BOUNDS_MODE uint32"},
		{"Enum constants", "COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS"},
		{"Interface type", "type ICoreWebView2Environment struct"},
		{"Method", "func (obj *ICoreWebView2Environment) CreateCoreWebView2Controller"},
		{"VTable", "type ICoreWebView2EnvironmentVTable struct"},
		{"Core interface", "type ICoreWebView2 struct"},
		{"Navigation method", "func (obj *ICoreWebView2) Navigate"},
		{"Settings interface", "type ICoreWebView2Settings struct"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if !strings.Contains(code, tc.expected) {
				t.Errorf("Generated code missing %s: expected to contain %q", tc.name, tc.expected)
			}
		})
	}

	// Log some statistics
	t.Logf("Generated code length: %d characters", len(code))
	t.Logf("Number of interfaces generated: %d", len(ast.Interfaces))
	t.Logf("Number of enums generated: %d", len(ast.Enums))

	// Optionally save generated code for inspection
	if testing.Verbose() {
		outputPath := "generated_webview2_subset.go"
		if err := os.WriteFile(outputPath, []byte(code), 0644); err == nil {
			t.Logf("Generated code saved to %s", outputPath)
		}
	}
}

func BenchmarkWebView2CodeGeneration(b *testing.B) {
	// Benchmark WebView2 code generation performance
	subsetPath := "test_subset.idl"
	if _, err := os.Stat(subsetPath); os.IsNotExist(err) {
		b.Skip("test_subset.idl not found, skipping WebView2 code generation benchmark")
	}

	file, err := os.Open(subsetPath)
	if err != nil {
		b.Fatalf("Error opening test subset: %v", err)
	}
	defer file.Close()

	// Parse once outside the benchmark
	parser := idl.NewParser(file)
	ast, err := parser.Parse()
	if err != nil {
		b.Fatalf("Error parsing WebView2 subset: %v", err)
	}

	generator, err := NewGenerator()
	if err != nil {
		b.Fatalf("Failed to create generator: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := generator.GenerateCode(ast, "webview2-benchmark")
		if err != nil {
			b.Fatalf("Code generation failed: %v", err)
		}
	}
}
