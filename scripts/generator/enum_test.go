package generator

import (
	"bytes"
	"embed"
	_ "embed"
	"flag"
	"github.com/matryer/is"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"updater/generator/types"
)

//go:embed testfiles/*
var testfiles embed.FS

func testfile(path string) *bytes.Buffer {
	f, err := testfiles.ReadFile("testfiles/" + path)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(f)
}

// update rewrites the goldens in testfiles/ from the generator's current output:
//
//	go test ./generator -update    # then review the diff, then run without -update
//
// The goldens are the only executable record of what the templates are supposed to produce, so a
// template change is not finished until they are regenerated and the diff read. Every test used to
// carry this as a commented-out os.WriteFile loop, which made regenerating them an edit-run-revert
// cycle across seven files -- enough friction that fixing generated output by hand looks like the
// cheaper option. It is not: the next regeneration silently reverts it.
var update = flag.Bool("update", false, "rewrite testfiles/ goldens from generator output")

// updateGoldens writes each generated file as its golden when -update is set, and reports whether
// it did. Callers return immediately if so: the goldens are embedded at compile time, so the
// assertions in the same run would still be comparing against the previous build's copies.
// Call it after the com.go strip, or com.go acquires a golden no test asserts on.
func updateGoldens(t *testing.T, files []*types.GeneratedFile) bool {
	t.Helper()
	if !*update {
		return false
	}
	for _, f := range files {
		name := filepath.Join("testfiles", f.FileName+".txt")
		if err := os.WriteFile(name, f.Content.Bytes(), 0o644); err != nil {
			t.Fatal(err)
		}
		t.Logf("updated %s", name)
	}
	return true
}

func makeOutput(input string) *bytes.Buffer {
	var buf bytes.Buffer
	// Normalise newlines
	input = strings.ReplaceAll(input, "\r\n", "\n")
	buf.Write([]byte(input))
	return &buf
}

var testData = []byte(`

[uuid(26d34152-879f-4065-bea2-3daa2cfadfb8), version(1.0)]
library WebView2 {

[v1_enum]
typedef enum COREWEBVIEW2_PREFERRED_COLOR_SCHEME {
    /// Auto color scheme.
    COREWEBVIEW2_PREFERRED_COLOR_SCHEME_AUTO,

    /// Light color scheme.
    COREWEBVIEW2_PREFERRED_COLOR_SCHEME_LIGHT,

    /// Dark color scheme.
    COREWEBVIEW2_PREFERRED_COLOR_SCHEME_DARK
} COREWEBVIEW2_PREFERRED_COLOR_SCHEME;


[v1_enum]
typedef enum COREWEBVIEW2_PREFERRED_COLOR_SCHEME1 {
    /// Auto color scheme.
    COREWEBVIEW2_PREFERRED_COLOR_SCHEME_AUTO1 = 1,

    /// Light color scheme.
    COREWEBVIEW2_PREFERRED_COLOR_SCHEME_LIGHT1 = 2,

    /// Dark color scheme.
    COREWEBVIEW2_PREFERRED_COLOR_SCHEME_DARK1 = 3,
} COREWEBVIEW2_PREFERRED_COLOR_SCHEME;


[v1_enum]
typedef enum COREWEBVIEW2_PREFERRED_COLOR_SCHEME2 {
   /// Auto color scheme.
   COREWEBVIEW2_PREFERRED_COLOR_SCHEME_AUTO2 = 1 << 1,

   /// Light color scheme.
   COREWEBVIEW2_PREFERRED_COLOR_SCHEME_LIGHT2 = 1 << 2,

   /// Dark color scheme.
   COREWEBVIEW2_PREFERRED_COLOR_SCHEME_DARK2 = 1 << 3
} COREWEBVIEW2_PREFERRED_COLOR_SCHEME;

}`)

func TestEnum(t *testing.T) {

	i := is.New(t)

	var buf bytes.Buffer
	buf.Write(testData)

	idl, err := Parser.Parse("", &buf)
	i.NoErr(err)

	err = idl.Process()
	i.NoErr(err)

	files, err := idl.Generate()
	i.NoErr(err)

	// Remove the `com.go` filename
	files = files[1:]

	if updateGoldens(t, files) {
		return
	}

	expected := []*types.GeneratedFile{
		{
			FileName: "COREWEBVIEW2_PREFERRED_COLOR_SCHEME.go",
			Package:  "webview2",
			Content:  testfile("COREWEBVIEW2_PREFERRED_COLOR_SCHEME.go.txt"),
		},
		{
			FileName: "COREWEBVIEW2_PREFERRED_COLOR_SCHEME1.go",
			Package:  "webview2",
			Content:  testfile("COREWEBVIEW2_PREFERRED_COLOR_SCHEME1.go.txt"),
		},
		{
			FileName: "COREWEBVIEW2_PREFERRED_COLOR_SCHEME2.go",
			Package:  "webview2",
			Content:  testfile("COREWEBVIEW2_PREFERRED_COLOR_SCHEME2.go.txt"),
		},
	}

	require.ElementsMatch(t, files, expected)

}
