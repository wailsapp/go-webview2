package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/wailsapp/go-webview2/webviewloader"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"text/scanner"
	"updater/generator"
)

const URL = "https://raw.githubusercontent.com/MicrosoftDocs/edge-developer/master/microsoft-edge/webview2/release-notes.md"

//go:embed latest_version.txt
var latestVersionProcessed string

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

var latestVersion string

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
			} else {
				latestVersion = version
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

	// Check if the latest version is different from the last time we ran this script
	latest, err := webviewloader.CompareBrowserVersions(latestVersion, latestVersionProcessed)
	if err != nil {
		log.Fatal(err)
	}
	if latest != 1 {
		println("No new version found")
		os.Exit(0)
	}

	println("Processing version: ", latestVersion)
	// Download Webview2 IDL for this version
	idlData, err := DownloadIDL(latestVersion)
	if err != nil {
		log.Fatal(err)
	}

	err = generator.ParseIDL(idlData, "../pkg/edge_generated")
	if err != nil {
		log.Fatal(err)
	}

	// Save the version to latest_version.txt
	err = os.WriteFile("latest_version.txt", []byte(latestVersion), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func DownloadIDL(version string) ([]byte, error) {

	// URL for the nuget package: https://www.nuget.org/api/v2/package/Microsoft.Web.WebView2/<version>
	// Download the package to the current directory
	client := resty.New()
	println("Downloading: ", fmt.Sprintf("https://www.nuget.org/api/v2/package/Microsoft.Web.WebView2/%s", version))
	resp, err := client.R().
		EnableTrace().
		Get(fmt.Sprintf("https://www.nuget.org/api/v2/package/Microsoft.Web.WebView2/%s", version))
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(resp.Body())
	zr, err := zip.NewReader(reader, int64(reader.Len()))
	if err != nil {
		return nil, err
	}

	var idlData []byte
	for _, file := range zr.File {
		if file.Name == "WebView2.idl" {
			r, err := file.Open()
			if err != nil {
				return nil, err
			}
			idlData, err = io.ReadAll(r)
			if err != nil {
				return nil, err
			}
		}
	}
	return idlData, nil
}
