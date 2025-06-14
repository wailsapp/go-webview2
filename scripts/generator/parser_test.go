package generator

import (
	"os"
	"strings"
	"testing"

	"generator/idl"
)

func TestParseSimpleEnum(t *testing.T) {
	idlContent := `typedef enum COREWEBVIEW2_BOUNDS_MODE {
		COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS,
		COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE,
	} COREWEBVIEW2_BOUNDS_MODE;`

	parser := idl.NewParser(strings.NewReader(idlContent))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Error parsing simple enum: %v", err)
	}

	if len(ast.Enums) != 1 {
		t.Fatalf("Expected 1 enum, got %d", len(ast.Enums))
	}

	enum := ast.Enums[0]
	if enum.Name != "COREWEBVIEW2_BOUNDS_MODE" {
		t.Errorf("Expected enum name 'COREWEBVIEW2_BOUNDS_MODE', got '%s'", enum.Name)
	}

	if len(enum.Values) != 2 {
		t.Errorf("Expected 2 enum values, got %d", len(enum.Values))
	}
}

func TestParseEnumWithAttributes(t *testing.T) {
	idlContent := `/// Mode for how the Bounds property is interpreted
	[v1_enum]
	typedef enum COREWEBVIEW2_BOUNDS_MODE {
		/// Bounds property represents raw pixels
		COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS,
		/// Bounds property represents logical pixels  
		COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE,
	} COREWEBVIEW2_BOUNDS_MODE;`

	parser := idl.NewParser(strings.NewReader(idlContent))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Error parsing enum with attributes: %v", err)
	}

	if len(ast.Enums) != 1 {
		t.Fatalf("Expected 1 enum, got %d", len(ast.Enums))
	}

	enum := ast.Enums[0]
	if len(enum.Attributes) == 0 {
		t.Error("Expected enum to have attributes")
	}
}

func TestParseForwardDeclaration(t *testing.T) {
	idlContent := `interface ICoreWebView2;
	interface ICoreWebView2Environment;`

	parser := idl.NewParser(strings.NewReader(idlContent))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Error parsing forward declarations: %v", err)
	}

	// Forward declarations should be skipped, not added to interfaces
	if len(ast.Interfaces) != 0 {
		t.Errorf("Expected 0 interfaces (forward declarations should be skipped), got %d", len(ast.Interfaces))
	}
}

func TestParseInterface(t *testing.T) {
	idlContent := `[uuid(b99369f3-9b11-47b5-bc6f-8e7895fcea17), object, pointer_default(unique)]
	interface ICoreWebView2CompletedHandler : IUnknown {
		HRESULT Invoke([in] HRESULT errorCode, [in] LPCWSTR result);
	}`

	parser := idl.NewParser(strings.NewReader(idlContent))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Error parsing interface: %v", err)
	}

	if len(ast.Interfaces) != 1 {
		t.Fatalf("Expected 1 interface, got %d", len(ast.Interfaces))
	}

	iface := ast.Interfaces[0]
	if iface.Name != "ICoreWebView2CompletedHandler" {
		t.Errorf("Expected interface name 'ICoreWebView2CompletedHandler', got '%s'", iface.Name)
	}

	if iface.Parent != "IUnknown" {
		t.Errorf("Expected parent 'IUnknown', got '%s'", iface.Parent)
	}

	if iface.UUID != "b99369f3-9b11-47b5-bc6f-8e7895fcea17" {
		t.Errorf("Expected UUID 'b99369f3-9b11-47b5-bc6f-8e7895fcea17', got '%s'", iface.UUID)
	}

	if len(iface.Methods) != 1 {
		t.Errorf("Expected 1 method, got %d", len(iface.Methods))
	}
}

func TestParseImports(t *testing.T) {
	idlContent := `import "objidl.idl";
	import "oaidl.idl";
	import "EventToken.idl";`

	parser := idl.NewParser(strings.NewReader(idlContent))
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Error parsing imports: %v", err)
	}

	if len(ast.Imports) != 3 {
		t.Fatalf("Expected 3 imports, got %d", len(ast.Imports))
	}

	expectedImports := []string{"objidl.idl", "oaidl.idl", "EventToken.idl"}
	for i, imp := range ast.Imports {
		if imp.Path != expectedImports[i] {
			t.Errorf("Expected import path '%s', got '%s'", expectedImports[i], imp.Path)
		}
	}
}

func TestTypeMapping(t *testing.T) {
	tests := []struct {
		idlType    string
		expectedGo string
	}{
		{"HRESULT", "syscall.Errno"},
		{"HWND", "uintptr"},
		{"LPWSTR", "*uint16"},
		{"BOOL", "int32"},
		{"DWORD", "uint32"},
	}

	for _, test := range tests {
		idlType := &idl.Type{Name: test.idlType, Kind: idl.TypeBasic}
		goType := idl.GetGoType(idlType)

		if goType != test.expectedGo {
			t.Errorf("Type mapping for %s: expected %s, got %s", test.idlType, test.expectedGo, goType)
		}
	}
}

func TestVersionDetection(t *testing.T) {
	tests := []struct {
		name        string
		version     int
		isVersioned bool
	}{
		{"ICoreWebView2", 1, false},
		{"ICoreWebView2_2", 2, true},
		{"ICoreWebView2_10", 10, true},
		{"ICoreWebView2Environment", 1, false},
	}

	for _, test := range tests {
		iface := &idl.Interface{Name: test.name}

		if iface.IsVersioned() != test.isVersioned {
			t.Errorf("IsVersioned for %s: expected %t, got %t", test.name, test.isVersioned, iface.IsVersioned())
		}

		if iface.GetVersion() != test.version {
			t.Errorf("GetVersion for %s: expected %d, got %d", test.name, test.version, iface.GetVersion())
		}
	}
}

// Benchmark parsing performance
func BenchmarkParseSimpleEnum(b *testing.B) {
	idlContent := `typedef enum COREWEBVIEW2_BOUNDS_MODE {
		COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS,
		COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE,
	} COREWEBVIEW2_BOUNDS_MODE;`

	for i := 0; i < b.N; i++ {
		parser := idl.NewParser(strings.NewReader(idlContent))
		_, err := parser.Parse()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Test parsing a real WebView2 IDL subset
func TestParseWebView2Subset(t *testing.T) {
	subsetPath := "test_subset.idl"
	if _, err := os.Stat(subsetPath); os.IsNotExist(err) {
		t.Skip("test_subset.idl not found, skipping WebView2 subset test")
	}

	file, err := os.Open(subsetPath)
	if err != nil {
		t.Fatalf("Error opening test subset: %v", err)
	}
	defer file.Close()

	parser := idl.NewParser(file)
	ast, err := parser.Parse()

	if err != nil {
		t.Fatalf("Error parsing WebView2 subset: %v", err)
	}

	// Basic validation
	if len(ast.Imports) == 0 {
		t.Error("Expected imports in WebView2 subset")
	}

	if len(ast.Enums) == 0 {
		t.Error("Expected enums in WebView2 subset")
	}

	t.Logf("Successfully parsed WebView2 subset: %d imports, %d enums, %d interfaces, %d typedefs",
		len(ast.Imports), len(ast.Enums), len(ast.Interfaces), len(ast.TypeDefs))
}
