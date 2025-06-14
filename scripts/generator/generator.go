package generator

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"text/template"

	"generator/idl"
)

// Generator handles code generation from IDL AST to Go code
type Generator struct {
	templates map[string]*template.Template
}

// NewGenerator creates a new code generator
func NewGenerator() (*Generator, error) {
	g := &Generator{
		templates: make(map[string]*template.Template),
	}

	if err := g.loadTemplates(); err != nil {
		return nil, fmt.Errorf("failed to load templates: %w", err)
	}

	return g, nil
}

// loadTemplates loads all Go code generation templates
func (g *Generator) loadTemplates() error {
	templateDir := filepath.Join("codegen", "templates")

	templateFiles := map[string]string{
		"package":   filepath.Join(templateDir, "package_template.go.tmpl"),
		"enum":      filepath.Join(templateDir, "enum_template.go.tmpl"),
		"interface": filepath.Join(templateDir, "interface_template.go.tmpl"),
		"typedef":   filepath.Join(templateDir, "typedef_template.go.tmpl"),
	}

	for name, path := range templateFiles {
		tmpl, err := template.ParseFiles(path)
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", name, err)
		}
		g.templates[name] = tmpl
	}

	return nil
}

// GenerateCode generates Go code from IDL AST
func (g *Generator) GenerateCode(ast *idl.AST, version string) (string, error) {
	var buf bytes.Buffer

	// Generate package header
	if err := g.generatePackageHeader(&buf, version); err != nil {
		return "", fmt.Errorf("failed to generate package header: %w", err)
	}

	// Generate enums
	for _, enum := range ast.Enums {
		if err := g.generateEnum(&buf, enum); err != nil {
			return "", fmt.Errorf("failed to generate enum %s: %w", enum.Name, err)
		}
	}

	// Generate typedefs
	for _, typedef := range ast.TypeDefs {
		if err := g.generateTypedef(&buf, typedef); err != nil {
			return "", fmt.Errorf("failed to generate typedef %s: %w", typedef.Name, err)
		}
	}

	// Generate interfaces
	for _, iface := range ast.Interfaces {
		if err := g.generateInterface(&buf, iface); err != nil {
			return "", fmt.Errorf("failed to generate interface %s: %w", iface.Name, err)
		}
	}

	return buf.String(), nil
}

// generatePackageHeader generates the package declaration and imports
func (g *Generator) generatePackageHeader(w io.Writer, version string) error {
	data := struct {
		Version string
	}{
		Version: version,
	}

	return g.templates["package"].Execute(w, data)
}

// generateEnum generates Go code for an IDL enum
func (g *Generator) generateEnum(w io.Writer, enum *idl.Enum) error {
	// Convert enum values to template data
	values := make([]struct {
		Name  string
		Value string
	}, len(enum.Values))

	for i, value := range enum.Values {
		values[i] = struct {
			Name  string
			Value string
		}{
			Name:  value.Name,
			Value: fmt.Sprintf("%d", value.Value),
		}
	}

	data := struct {
		Name   string
		Values []struct {
			Name  string
			Value string
		}
	}{
		Name:   enum.Name,
		Values: values,
	}

	return g.templates["enum"].Execute(w, data)
}

// generateTypedef generates Go code for an IDL typedef
func (g *Generator) generateTypedef(w io.Writer, typedef *idl.TypeDef) error {
	goType := idl.GetGoType(typedef.Type)

	data := struct {
		Name   string
		GoType string
	}{
		Name:   typedef.Name,
		GoType: goType,
	}

	return g.templates["typedef"].Execute(w, data)
}

// generateInterface generates Go code for an IDL interface
func (g *Generator) generateInterface(w io.Writer, iface *idl.Interface) error {
	// Convert methods to template data
	methods := make([]struct {
		Name       string
		Comment    string
		Parameters []struct {
			Name              string
			GoType            string
			SyscallConversion string
		}
		ReturnType       string
		ReturnConversion string
	}, len(iface.Methods))

	for i, method := range iface.Methods {
		// Convert parameters
		params := make([]struct {
			Name              string
			GoType            string
			SyscallConversion string
		}, len(method.Parameters))

		for j, param := range method.Parameters {
			goType := idl.GetGoType(param.Type)
			syscallConv := idl.GetConversionCode(param.Type, param.Name, true)

			params[j] = struct {
				Name              string
				GoType            string
				SyscallConversion string
			}{
				Name:              param.Name,
				GoType:            goType,
				SyscallConversion: syscallConv,
			}
		}

		// Determine return type and conversion
		returnType := "error"
		returnConversion := ""

		if method.ReturnType.Name != "HRESULT" {
			returnType = idl.GetGoType(method.ReturnType)
			returnConversion = idl.GetConversionCode(method.ReturnType, "ret", false)
		}

		methods[i] = struct {
			Name       string
			Comment    string
			Parameters []struct {
				Name              string
				GoType            string
				SyscallConversion string
			}
			ReturnType       string
			ReturnConversion string
		}{
			Name:             method.Name,
			Comment:          strings.Join(method.Comments, " "),
			Parameters:       params,
			ReturnType:       returnType,
			ReturnConversion: returnConversion,
		}
	}

	data := struct {
		Name    string
		Methods []struct {
			Name       string
			Comment    string
			Parameters []struct {
				Name              string
				GoType            string
				SyscallConversion string
			}
			ReturnType       string
			ReturnConversion string
		}
	}{
		Name:    iface.Name,
		Methods: methods,
	}

	return g.templates["interface"].Execute(w, data)
}
