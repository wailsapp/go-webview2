package types

import (
	"bytes"
	"errors"
	"github.com/leaanthony/slicer"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

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

func (i *IDL) Generate(targetDir string) error {
	for _, library := range i.Libraries {
		return library.Generate(targetDir)
	}
	return nil
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
}

func (l *Library) Process() error {
	l.packageName = strings.ToLower(l.Name)
	for _, declaration := range l.Declarations {
		err := declaration.Process(l)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *Library) Generate(targetDir string) error {
	packageDir, err := filepath.Abs(filepath.Join(targetDir, strings.ToLower(l.Name)))
	if err != nil {
		return err
	}

	_ = os.MkdirAll(packageDir, 0755)

	l.GenerateDefaultFiles(packageDir)

	for _, declaration := range l.Declarations {
		err := declaration.Generate(packageDir)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Library) addInterfaceName(interfaceName string) {
	l.forewardInterfaceDeclarations.Add(interfaceName)
}

func (l *Library) GenerateDefaultFiles(packageDir string) {
	data := struct {
		PackageName string
	}{
		PackageName: l.packageName,
	}
	templateData, err := templates.ReadFile("templates/com.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err := template.New("COM").Parse(string(templateData))
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(filepath.Join(packageDir, "com.go"))
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
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

func (d *Declaration) Generate(packageDir string) error {

	var buffer bytes.Buffer
	var packageName = strings.ToLower(d.library.Name)
	var filename string

	if d.Enum != nil {
		err := d.Enum.Generate(packageName, &buffer)
		if err != nil {
			return err
		}
		filename = d.Enum.Name + ".go"
	}
	if d.Struct != nil {
		err := d.Struct.Generate(packageName, &buffer)
		if err != nil {
			return err
		}
		filename = d.Struct.Name + ".go"
	}
	if d.Interface != nil {
		err := d.Interface.Generate(packageName, &buffer)
		if err != nil {
			return err
		}
		filename = d.Interface.Name + ".go"
	}
	if d.CppQuote != "" {
		return nil
	}
	if d.InterfaceForewardDecl != "" {
		return nil
	}
	f := filepath.Join(packageDir, filename)
	return os.WriteFile(f, buffer.Bytes(), 0755)

}
