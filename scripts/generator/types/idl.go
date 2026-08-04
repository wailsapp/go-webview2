package types

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/leaanthony/slicer"
	"go/format"
	"log"
	"strings"
	"text/template"
)

type GeneratedFile struct {
	FileName string
	Package  string
	Content  *bytes.Buffer
}

type IDL struct {
	Imports   []*Import  `parser:"@@*"`
	Libraries []*Library `parser:"@@*"`
}

func (i *IDL) Process() error {
	for _, library := range i.Libraries {
		err := library.Process()
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *IDL) Generate() ([]*GeneratedFile, error) {
	// Accumulate across libraries rather than returning inside the loop. Every WebView2 IDL
	// declares exactly one, so the old early return was never wrong -- it just read as a loop while
	// behaving like an index, which is the kind of thing that stops being harmless quietly.
	var all []*GeneratedFile
	for _, library := range i.Libraries {
		files, err := library.Generate()
		if err != nil {
			return nil, err
		}
		all = append(all, files...)
	}
	return gofmtAll(all)
}

// gofmtAll formats every generated file, which the generator did not previously do although the
// committed tree is gofmt-clean -- so somebody was formatting the output by hand afterwards. That
// left ~180 files differing from a fresh generation by nothing but import order and blank lines,
// which is enough noise to hide a real change in a regeneration diff, and hiding real changes in
// regeneration diffs is how this package accumulated hand-patched output in the first place.
//
// The error path is a second, unlooked-for benefit: a template that emits invalid Go now fails the
// generator by name instead of writing a file that only fails later at `go build`.
func gofmtAll(files []*GeneratedFile) ([]*GeneratedFile, error) {
	for _, f := range files {
		formatted, err := format.Source(f.Content.Bytes())
		if err != nil {
			return nil, fmt.Errorf("generated %s is not valid Go: %w", f.FileName, err)
		}
		f.Content = bytes.NewBuffer(formatted)
	}
	return files, nil
}

type Import struct {
	Name string `parser:"'import' @(!';')* ';'"`
}

type LibraryHeader struct {
	UUID string `parser:"'uuid' '(' @UUID ')' ',' 'version' '(' Int ('.' Int)? ')'"`
}

type Library struct {
	Header       *LibraryHeader `parser:"'[' @@ ']'"`
	Name         string         `parser:"'library' @Ident"`
	Declarations []*Declaration `parser:"'{' @@* '}'"`

	// private
	forewardInterfaceDeclarations slicer.StringSlicer
	enums                         slicer.StringSlicer
	packageName                   string
	interfaces                    map[string]*InterfaceDeclaration
}

func (l *Library) Process() error {
	l.packageName = strings.ToLower(l.Name)
	// Index the interfaces AND the enums before processing any of them: resolving an inheritance
	// chain needs every declaration to be findable by name, and nothing guarantees a base -- or an
	// enum -- is declared first.
	//
	// The enums were previously registered as each one was processed, while Param.IsEnum() is
	// consulted while processing an INTERFACE. An enum declared after the interface that uses it
	// therefore looked like an unknown type, which used to mean a silently wrong &address and now
	// means the generator stops with advice that does not apply. Microsoft's IDLs happen to put
	// every enum first, which is the only reason this never fired.
	l.interfaces = map[string]*InterfaceDeclaration{}
	for _, declaration := range l.Declarations {
		if declaration.Interface != nil {
			l.interfaces[declaration.Interface.Name] = declaration.Interface
		}
		if declaration.Enum != nil {
			l.enums.Add(declaration.Enum.Name)
		}
	}
	for _, declaration := range l.Declarations {
		err := declaration.Process(l)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *Library) Generate() ([]*GeneratedFile, error) {
	result := l.GenerateDefaultFiles()

	for _, declaration := range l.Declarations {
		generatedFile, err := declaration.Generate()
		if err != nil {
			return nil, err
		}
		if generatedFile != nil {
			result = append(result, generatedFile)
		}
	}

	return result, nil
}

func (l *Library) addInterfaceName(interfaceName string) {
	l.forewardInterfaceDeclarations.Add(interfaceName)
}

func (l *Library) GenerateDefaultFiles() []*GeneratedFile {
	data := struct {
		PackageName string
	}{
		PackageName: l.packageName,
	}

	var result []*GeneratedFile
	var buf bytes.Buffer

	templateData, err := templates.ReadFile("templates/com.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err := template.New("COM").Parse(string(templateData))
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(&buf, &data)
	if err != nil {
		log.Fatal(err)
	}

	result = append(result, &GeneratedFile{
		FileName: "com.go",
		Package:  l.packageName,
		Content:  &buf,
	})

	return result

}

type Declaration struct {
	InterfaceForewardDecl string                `parser:"'interface' @Ident ';'"`
	Enum                  *EnumDeclaration      `parser:"| '[' 'v1_enum' ']' @@"`
	Struct                *StructDeclaration    `parser:"| @@"`
	Interface             *InterfaceDeclaration `parser:"| @@"`
	CppQuote              string                `parser:"| 'cpp_quote' '(' @String ')'"`

	// Private
	library *Library
}

func (d *Declaration) Process(l *Library) error {
	d.library = l
	if d.Enum != nil {
		return d.Enum.Process(d)
	}
	if d.Struct != nil {
		return d.Struct.Process(d)
	}
	if d.Interface != nil {
		return d.Interface.Process(d)
	}
	if d.CppQuote != "" {
		return nil
	}
	if d.InterfaceForewardDecl != "" {
		l.addInterfaceName(d.InterfaceForewardDecl)
		return nil
	}
	return errors.New("unknown declaration to process")
}

func (d *Declaration) Generate() (*GeneratedFile, error) {

	var buffer bytes.Buffer
	var packageName = strings.ToLower(d.library.Name)
	var filename string

	if d.Enum != nil {
		err := d.Enum.Generate(packageName, &buffer)
		if err != nil {
			return nil, err
		}
		filename = d.Enum.Name + ".go"
	}
	if d.Struct != nil {
		err := d.Struct.Generate(packageName, &buffer)
		if err != nil {
			return nil, err
		}
		filename = d.Struct.Name + ".go"
	}
	if d.Interface != nil {
		err := d.Interface.Generate(packageName, &buffer)
		if err != nil {
			return nil, err
		}
		filename = d.Interface.Name + ".go"
	}
	if d.CppQuote != "" {
		return nil, nil
	}
	if d.InterfaceForewardDecl != "" {
		return nil, nil
	}
	//f := filepath.Join(packageDir, filename)
	//err := os.WriteFile(f, buffer.Bytes(), 0755)
	return &GeneratedFile{
		FileName: filename,
		Package:  packageName,
		Content:  &buffer,
	}, nil

}
