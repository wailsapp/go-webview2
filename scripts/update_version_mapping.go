package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
	"regexp"
	"strings"
	"text/scanner"
)

const URL = "https://raw.githubusercontent.com/MicrosoftDocs/edge-developer/master/microsoft-edge/webview2/release-notes.md"

type Version struct {
	Number         string
	ReleaseNotes   string
	RuntimeVersion string
	Notes          []string
}

const debug = false

func getDoc() []byte {
	if debug {
		data, err := os.ReadFile("test.md")
		if err != nil {
			log.Fatal(err)
		}
		return data
	}
	// GET the URL
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Body()
}

func extractVersion(in string) string {
	// Match the version numbers by format: 1.0.774.44, 0.9.515-prerelease,
	regex := regexp.MustCompile(`\d+\.\d+\.\d+(\.\d+|-prerelease)`)
	version := regex.Find([]byte(in))
	return string(version)
}

func main() {

	var buf bytes.Buffer
	data := getDoc()
	buf.Write(data)
	var s scanner.Scanner
	s.Init(&buf)

	r := bufio.NewReader(&buf)

	var err error
	var line []byte
	var versions []*Version
	var currentVersion *Version
	var nomnom bool
	for err == nil {
		line, _, err = r.ReadLine()

		// Check if line starts with `[NuGet package for WebView2 `
		l := string(line)

		if strings.HasPrefix(l, `## `) {
			nomnom = false
			continue
		}

		if currentVersion != nil && nomnom {
			currentVersion.Notes = append(currentVersion.Notes, l)
		}

		if strings.HasPrefix(l, `[NuGet package for WebView2 `) {
			version := extractVersion(l)
			if version == "" {
				continue
			}
			if currentVersion != nil {
				versions = append(versions, currentVersion)
			}
			currentVersion = &Version{
				Number: version,
			}
			continue
		}
		if strings.HasSuffix(strings.TrimSpace(l), "or higher.") {
			if currentVersion != nil {
				currentVersion.RuntimeVersion = extractVersion(l)
				currentVersion.ReleaseNotes = `https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#` + strings.Replace(currentVersion.Number, ".", "", -1)
				nomnom = true
			}
			continue
		}

		if strings.HasPrefix(l, `Release Date:`) {
			if currentVersion != nil {
				currentVersion.Notes = append(currentVersion.Notes, strings.Trim(l, " "))
			}
		}
	}

	var buffer strings.Builder
	buffer.WriteString("//go:build windows\n\n")
	buffer.WriteString("package webviewloader\n\n")
	buffer.WriteString("type Version struct {\n")
	buffer.WriteString("	SDKVersion         string\n")
	buffer.WriteString("	ReleaseNotes           string\n")
	buffer.WriteString("	RuntimeVersion string\n")
	buffer.WriteString("	Notes          string\n")
	buffer.WriteString("}\n\n")
	buffer.WriteString("var versionMapping = map[string]Version{\n")
	for _, version := range versions {
		buffer.WriteString(fmt.Sprintf("	\"%s\": {\n", version.Number))
		buffer.WriteString(fmt.Sprintf("		SDKVersion:     \"%s\",\n", version.Number))
		buffer.WriteString(fmt.Sprintf("		ReleaseNotes:   \"%s\",\n", version.ReleaseNotes))
		buffer.WriteString(fmt.Sprintf("		RuntimeVersion: \"%s\",\n", version.RuntimeVersion))
		buffer.WriteString("		Notes: ")
		buffer.WriteString(fmt.Sprintf("			`%s`,\n", strings.Replace(strings.Join(version.Notes, "\n"), "`", "'", -1)))
		buffer.WriteString("	},\n")
	}
	buffer.WriteString("}\n")

	// Write the buffer to ../webviewloader/version_map.go
	err = os.WriteFile("../webviewloader/version_map.go", []byte(buffer.String()), 0644)
	if err != nil {
		log.Fatal(err)
	}

}
