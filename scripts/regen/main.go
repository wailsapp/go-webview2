// Command regen runs the generator against a local IDL and writes the result to a chosen
// directory. update_version_mapping.go can only regenerate as a side effect of checking
// Microsoft's release-notes page for a NEW version, which needs the network and rewrites the
// tree in place -- neither of which suits comparing generator output against the committed
// files.
//
// Usage: go run ./regen -idl WebView2.1.0.2903.40.idl -out /tmp/baseline
package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"updater/generator"
)

func main() {
	idl := flag.String("idl", "WebView2.1.0.2903.40.idl", "IDL file to generate from")
	out := flag.String("out", "", "directory to write the generated files to (required)")
	flag.Parse()
	if *out == "" {
		log.Fatal("-out is required")
	}

	data, err := os.ReadFile(*idl)
	if err != nil {
		log.Fatal(err)
	}
	files, err := generator.ParseIDL(data)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(*out, 0o755); err != nil {
		log.Fatal(err)
	}
	// Clear the previously generated files, because reusing a directory across two IDL versions
	// otherwise leaves behind the ones only the earlier version produced, and `diff -r` then reads
	// as though they were still being generated -- which defeats the one thing this command is for.
	//
	// Only non-test .go files, NOT the whole directory: -out is usually pkg/webview2 itself, which
	// also holds hand-written tests. Everything the generator emits is a non-test .go file, so that
	// line is exactly the boundary between "derived, safe to delete" and "written by a person".
	// RemoveAll here deleted marshal_windows_test.go, silently, and the package went back to
	// reporting "no test files".
	stale, err := filepath.Glob(filepath.Join(*out, "*.go"))
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range stale {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}
		if err := os.Remove(f); err != nil {
			log.Fatal(err)
		}
	}
	for _, f := range files {
		if err := os.WriteFile(filepath.Join(*out, f.FileName), f.Content.Bytes(), 0o644); err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("wrote %d files to %s", len(files), *out)
}
