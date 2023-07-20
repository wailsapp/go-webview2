package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"log"
	"os"
	"strings"
)

const URL = "https://raw.githubusercontent.com/MicrosoftDocs/edge-developer/master/microsoft-edge/webview2/release-notes.md"

type Version struct {
	Number         string
	Link           string
	RuntimeVersion string
	Notes          []string
}

func main() {
	// GET the URL
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the file into a tree
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(resp.Body())

	var versions []*Version

	var currentVersion *Version
	var inPromotions bool

	ast.WalkFunc(doc, func(node ast.Node, entering bool) ast.WalkStatus {
		if entering {
			switch n := node.(type) {
			case *ast.Heading:
				if n.Level == 2 {
					if n.HeadingID[0] >= '0' && n.HeadingID[0] <= '9' {
						if currentVersion != nil {
							versions = append(versions, currentVersion)
						}
						currentVersion = &Version{
							Number: strings.ReplaceAll(n.HeadingID, "-", "."),
						}
						inPromotions = false
					}
				}
			case *ast.Text:
				if len(n.Literal) == 0 {
					break
				}
				if currentVersion != nil {
					literal := string(n.Literal)
					if strings.HasPrefix(literal, "Release Date: ") {
						currentVersion.Notes = append(currentVersion.Notes, literal)
						break
					}
					if strings.Contains(literal, "or higher.") {
						runtimeVersion := literal
						runtimeVersion = strings.TrimPrefix(runtimeVersion, "For full API compatibility, this prerelease version of the WebView2 SDK requires Microsoft Edge version")
						runtimeVersion = strings.TrimPrefix(runtimeVersion, "For full API compatibility, this version of the WebView2 SDK requires WebView2 Runtime version")
						runtimeVersion = strings.TrimPrefix(runtimeVersion, "For full API compatibility, this version of the WebView2 SDK requires Microsoft Edge version")
						runtimeVersion = strings.TrimPrefix(runtimeVersion, "This version of the WebView2 SDK requires WebView2 Runtime version")
						runtimeVersion = strings.TrimPrefix(runtimeVersion, "This version of the WebView2 SDK requires Microsoft Edge version ")
						runtimeVersion = strings.TrimPrefix(runtimeVersion, "This prerelease version of the WebView2 SDK requires Microsoft Edge version ")
						runtimeVersion = strings.TrimPrefix(runtimeVersion, "This prerelease version of the WebView2 SDK requires WebView2 Runtime version ")
						runtimeVersion = strings.TrimSuffix(runtimeVersion, "or higher.")
						currentVersion.RuntimeVersion = strings.TrimSpace(runtimeVersion)
						break
					}
					if inPromotions {
						_ = literal
					}
				}
			}
		}
		return ast.GoToNext
	})

	var buffer strings.Builder
	buffer.WriteString("//go:build windows\n\n")
	buffer.WriteString("package webviewloader\n\n")
	buffer.WriteString("type Version struct {\n")
	buffer.WriteString("	Number         string\n")
	buffer.WriteString("	Link           string\n")
	buffer.WriteString("	RuntimeVersion string\n")
	buffer.WriteString("	Notes          []string\n")
	buffer.WriteString("}\n\n")
	buffer.WriteString("var versionMapping = map[string]Version{\n")
	for _, version := range versions {
		buffer.WriteString(fmt.Sprintf("	\"%s\": {\n", version.Number))
		buffer.WriteString(fmt.Sprintf("		Number:         \"%s\",\n", version.Number))
		buffer.WriteString(fmt.Sprintf("		Link:           \"%s\",\n", version.Link))
		buffer.WriteString(fmt.Sprintf("		RuntimeVersion: \"%s\",\n", version.RuntimeVersion))
		buffer.WriteString("		Notes: []string{\n")
		for _, note := range version.Notes {
			buffer.WriteString(fmt.Sprintf("			\"%s\",\n", note))
		}
		buffer.WriteString("		},\n")
		buffer.WriteString("	},\n")
	}
	buffer.WriteString("}\n")

	// Write the buffer to ../webviewloader/version_map.go
	err = os.WriteFile("../webviewloader/version_map.go", []byte(buffer.String()), 0644)
	if err != nil {
		log.Fatal(err)
	}

}
