package types

import (
	"io"
	"log"
	"text/template"
)

type StructDeclaration struct {
	Name   string         `parser:"'typedef' 'struct' @Ident '{' "`
	Fields []*StructField `parser:" (@@)+ '}' Ident ';'"`

	// private
	decl *Declaration
}

func (d *StructDeclaration) Process(decl *Declaration) error {
	d.decl = decl
	for _, f := range d.Fields {
		f.Process()
	}
	return nil
}

func (d *StructDeclaration) Generate(packageName string, w io.Writer) error {
	data := struct {
		PackageName string
		Name        string
		Fields      []*StructField
	}{
		PackageName: packageName,
		Name:        d.Name,
		Fields:      d.Fields,
	}
	templateData, err := templates.ReadFile("templates/struct.tmpl")
	if err != nil {
		return err
	}
	tmpl, err := template.New("Struct").Parse(string(templateData))
	if err != nil {
		log.Fatalln(err)
	}
	return tmpl.Execute(w, &data)
}

type StructField struct {
	Type string `parser:"@('UINT32' | 'BOOL' | 'BYTE')"`
	Name string `parser:"@Ident ';'"`

	// process
	GoType string
}

// idlStructFieldToGoType is deliberately separate from IdlTypeToGoType, which is written for
// PARAMETERS. A parameter's BOOL is converted at the boundary, so mapping it to Go's bool is a
// kindness to callers and costs nothing. A struct FIELD has no boundary: the struct is a memory
// layout that the callee writes into directly, so every field must be the width the C declaration
// says it is.
//
// BOOL is a 4-byte int. Mapping it to Go's 1-byte bool made COREWEBVIEW2_PHYSICAL_KEY_STATUS 12
// bytes against a native 24, and its only use is
// ICoreWebView2AcceleratorKeyPressedEventArgs.GetPhysicalKeyStatus, which hands WebView2 the
// address of that 12-byte local. So every call wrote 12 bytes past the end of a heap object -- and
// even the bytes that landed inside were misread, because the last three flags sat at native
// offsets 12/16/20 while Go looked for them at 9/10/11, making them permanently false.
var idlStructFieldToGoType = map[string]string{
	"BOOL": "int32",
}

func (s *StructField) Process() {
	if goType, ok := idlStructFieldToGoType[s.Type]; ok {
		s.GoType = goType
		return
	}
	s.GoType = IdlTypeToGoType(s.Type)
}
