package types

import (
	"io"
	"log"
	"strconv"
	"strings"
	"text/template"
)

type EnumDeclaration struct {
	Name   string       `parser:"'typedef' 'enum' @Ident '{'"`
	Values []*EnumValue `parser:" (@@)+ '}' Ident ';'"`

	// private
	decl *Declaration
}

type EnumValue struct {
	Key   string         `parser:"@Ident"`
	Value *EnumValueDecl `parser:"('=' @@)? ','?"`
}

type EnumValueDecl struct {
	Value     string `parser:"@Hex | @Int"`
	LeftShift *int   `parser:"('<' '<' @Int)?"`
}

func (e *EnumValueDecl) Process() {
	if e.LeftShift != nil {
		e.Value += " << " + strconv.Itoa(*e.LeftShift)
	}
}

// asInt evaluates the initialiser forms the grammar accepts -- a decimal or hex literal, with an
// optional "<< N" -- so that the NEXT enumerator can continue counting from it. Reports false for
// anything it cannot evaluate, which is not an error: the caller has a correct fallback.
func (e *EnumValueDecl) asInt() (int64, bool) {
	text := strings.TrimSpace(e.Value)
	shift := uint64(0)
	if base, sh, found := strings.Cut(text, "<<"); found {
		n, err := strconv.ParseUint(strings.TrimSpace(sh), 0, 8)
		if err != nil {
			return 0, false
		}
		text, shift = strings.TrimSpace(base), n
	}
	// ParseInt with base 0 takes both "5" and "0x4".
	n, err := strconv.ParseInt(text, 0, 64)
	if err != nil {
		return 0, false
	}
	return n << shift, true
}

func (d *EnumDeclaration) Process(decl *Declaration) error {
	d.decl = decl
	// An enumerator with no initialiser is PREVIOUS + 1 in C, not its ordinal position. Using the
	// index gave the right answer only for an enum that sets no values at all, or sets them to
	// their own positions -- which every shipped WebView2 enum happens to do, which is why this
	// never showed. "A = 5, B" produced B = 1 rather than 6: a wrong constant that compiles and
	// goes straight to WebView2.
	prev, prevKnown := int64(-1), true
	for index, value := range d.Values {
		if value.Value != nil {
			value.Value.Process()
			prev, prevKnown = value.Value.asInt()
			continue
		}
		switch {
		case prevKnown:
			prev++
			value.Value = &EnumValueDecl{Value: strconv.FormatInt(prev, 10)}
		default:
			// The previous initialiser was an expression this cannot evaluate. Naming the previous
			// enumerator is still exact, because Go constant declarations allow it.
			value.Value = &EnumValueDecl{Value: d.Values[index-1].Key + " + 1"}
		}
	}
	// The name is registered in Library.Process's pre-pass, before any interface is processed --
	// see the comment there. Registering it again here would be too late to be useful.
	return nil
}

func (d *EnumDeclaration) Generate(packageName string, w io.Writer) error {
	data := struct {
		PackageName string
		Name        string
		Values      []*EnumValue
	}{
		PackageName: packageName,
		Name:        d.Name,
		Values:      d.Values,
	}
	templateData, err := templates.ReadFile("templates/enum.tmpl")
	if err != nil {
		return err
	}
	tmpl, err := template.New("Enum").Parse(string(templateData))
	if err != nil {
		log.Fatalln(err)
	}
	return tmpl.Execute(w, &data)
}
