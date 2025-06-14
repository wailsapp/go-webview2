package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/scanner"
	"time"
)

const URL = "https://raw.githubusercontent.com/MicrosoftDocs/edge-developer/master/microsoft-edge/webview2/release-notes/index.md"

//go:embed latest_version.txt
var latestVersionProcessed string

type Version struct {
	Number         string
	ReleaseNotes   string
	RuntimeVersion string
	Notes          []string
}

const debug = false

// ANSI color codes for pretty output
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorBold   = "\033[1m"
)

func printStep(step string) {
	fmt.Printf("%s%s[STEP]%s %s%s%s\n", ColorBold, ColorBlue, ColorReset, ColorCyan, step, ColorReset)
}

func printSuccess(msg string) {
	fmt.Printf("%s%s[SUCCESS]%s %s%s%s\n", ColorBold, ColorGreen, ColorReset, ColorGreen, msg, ColorReset)
}

func printWarning(msg string) {
	fmt.Printf("%s%s[WARNING]%s %s%s%s\n", ColorBold, ColorYellow, ColorReset, ColorYellow, msg, ColorReset)
}

func printError(msg string) {
	fmt.Printf("%s%s[ERROR]%s %s%s%s\n", ColorBold, ColorRed, ColorReset, ColorRed, msg, ColorReset)
}

func printInfo(msg string) {
	fmt.Printf("%s%s[INFO]%s %s\n", ColorBold, ColorWhite, ColorReset, msg)
}

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

// runCommand executes a command in the specified directory
func runCommand(command string, args []string, dir string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		printError(fmt.Sprintf("Command failed: %s %s", command, strings.Join(args, " ")))
		printError(fmt.Sprintf("Output: %s", string(output)))
		return err
	}
	
	return nil
}

func main() {
	fmt.Printf("%s%s╔══════════════════════════════════════════════════════════════════╗%s\n", ColorBold, ColorPurple, ColorReset)
	fmt.Printf("%s%s║                    WebView2 Edge Package Updater                 ║%s\n", ColorBold, ColorPurple, ColorReset)
	fmt.Printf("%s%s╚══════════════════════════════════════════════════════════════════╝%s\n", ColorBold, ColorPurple, ColorReset)
	fmt.Println()

	var forced bool
	if len(os.Args) > 1 {
		forced = os.Args[1] == "-forced"
		if forced {
			printWarning("Running in FORCED mode - will update regardless of version")
		}
	}

	// Step 1: Check for latest version
	printStep("Checking for latest WebView2 version...")
	
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

	if latestVersion == "" {
		printError("Could not determine latest WebView2 version")
		os.Exit(1)
	}

	printInfo(fmt.Sprintf("Latest WebView2 version: %s%s%s", ColorGreen, latestVersion, ColorReset))

	// Check if we need to update
	if !forced {
		latest, err := CompareBrowserVersions(latestVersion, strings.TrimSpace(latestVersionProcessed))
		if err != nil {
			printError(fmt.Sprintf("Version comparison failed: %v", err))
			os.Exit(1)
		}
		if latest != 1 {
			printSuccess("Already up to date - no new version found")
			os.Exit(0)
		}
	}

	// Step 2: Backup current edge directory
	printStep("Backing up current edge directory...")
	
	edgeDir := "../pkg/edge"
	timestamp := time.Now().Format("20060102_150405")
	backupDir := fmt.Sprintf("../pkg/edge.%s", timestamp)
	
	if _, err := os.Stat(edgeDir); err == nil {
		err = os.Rename(edgeDir, backupDir)
		if err != nil {
			printError(fmt.Sprintf("Failed to backup edge directory: %v", err))
			os.Exit(1)
		}
		printSuccess(fmt.Sprintf("Backed up edge directory to: %s", backupDir))
	} else {
		printInfo("No existing edge directory found - creating new one")
	}

	// Step 3: Create new edge directory
	printStep("Creating new edge directory...")
	
	err = os.MkdirAll(edgeDir, 0755)
	if err != nil {
		printError(fmt.Sprintf("Failed to create edge directory: %v", err))
		os.Exit(1)
	}
	printSuccess("Created new edge directory")

	// Step 4: Download WebView2 IDL
	printStep(fmt.Sprintf("Downloading WebView2 IDL for version %s...", latestVersion))
	
	idlData, err := DownloadIDL(latestVersion)
	if err != nil {
		printError(fmt.Sprintf("Failed to download IDL: %v", err))
		os.Exit(1)
	}
	printSuccess(fmt.Sprintf("Downloaded WebView2 IDL (%d bytes)", len(idlData)))

	// Step 5: Copy existing WebView2 package as base
	printStep("Copying existing WebView2 package as base...")
	
	webview2SourceDir := "../pkg/webview2"
	if _, err := os.Stat(webview2SourceDir); err != nil {
		printError(fmt.Sprintf("WebView2 source directory not found: %v", err))
		os.Exit(1)
	}
	
	// Copy all WebView2 files to edge directory
	err = filepath.Walk(webview2SourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip directories
		if info.IsDir() {
			return nil
		}
		
		// Only copy .go files
		if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}
		
		// Read source file
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		
		// Fix package name from webview2 to edge
		content := strings.Replace(string(data), "package webview2", "package edge", 1)
		
		// Write to edge directory
		relPath, err := filepath.Rel(webview2SourceDir, path)
		if err != nil {
			return err
		}
		
		dstPath := filepath.Join(edgeDir, relPath)
		err = os.WriteFile(dstPath, []byte(content), 0644)
		if err != nil {
			return err
		}
		
		return nil
	})
	
	if err != nil {
		printError(fmt.Sprintf("Failed to copy WebView2 package: %v", err))
		os.Exit(1)
	}
	
	// Count copied files
	copiedFiles, _ := filepath.Glob(filepath.Join(edgeDir, "*.go"))
	printSuccess(fmt.Sprintf("Copied %d Go files from WebView2 package", len(copiedFiles)))

	// Step 6: Copy core helper files from edge-old
	printStep("Copying core helper files from edge-old...")
	
	// Core files that should be preserved from old edge directory
	// Note: Exclude files that conflict with generated ones (com.go, guid.go, IStream.go, corewebview2.go)
	coreFiles := []string{
		"capabilities.go",
		"capabilities_test.go", 
		"chromium.go",
		"chromium_386.go",
		"chromium_amd64.go", 
		"chromium_arm64.go",
		"cookies_test.go",
		"create_env_go.go",
		"create_env_native.go",
		"ICoreWebViewSettings.go",
	}
	
	// First try edge-old directory, then backup directory
	edgeOldDir := "../pkg/edge-old"
	var sourceDir string
	if _, err := os.Stat(edgeOldDir); err == nil {
		sourceDir = edgeOldDir
		printInfo("Using edge-old directory as source")
	} else if _, err := os.Stat(backupDir); err == nil {
		sourceDir = backupDir
		printInfo("Using backup directory as source")
	} else {
		printWarning("No source directory found for core files")
	}
	
	if sourceDir != "" {
		copiedCount := 0
		for _, coreFile := range coreFiles {
			srcPath := filepath.Join(sourceDir, coreFile)
			dstPath := filepath.Join(edgeDir, coreFile)
			
			if data, err := os.ReadFile(srcPath); err == nil {
				// Fix package name if copying from edge-old (which has package edge_old)
				content := string(data)
				if sourceDir == edgeOldDir {
					content = strings.Replace(content, "package edge_old", "package edge", 1)
				}
				
				err = os.WriteFile(dstPath, []byte(content), 0644)
				if err != nil {
					printWarning(fmt.Sprintf("Failed to copy core file %s: %v", coreFile, err))
				} else {
					printInfo(fmt.Sprintf("Copied core file: %s", coreFile))
					copiedCount++
				}
			} else {
				printInfo(fmt.Sprintf("Core file not found (skipping): %s", coreFile))
			}
		}
		printSuccess(fmt.Sprintf("Copied %d core helper files", copiedCount))
	}
	
	// Generate com.go from template in generator
	printInfo("Generating com.go from template...")
	comTemplatePath := "generator/types/templates/com.tmpl"
	if comTemplateData, err := os.ReadFile(comTemplatePath); err == nil {
		// Replace template placeholder with actual package name
		comContent := strings.Replace(string(comTemplateData), "{{.PackageName}}", "edge", 1)
		err = os.WriteFile(filepath.Join(edgeDir, "com.go"), []byte(comContent), 0644)
		if err != nil {
			printWarning(fmt.Sprintf("Failed to write com.go: %v", err))
		} else {
			printSuccess("Generated com.go from template successfully")
		}
	} else {
		printWarning(fmt.Sprintf("Failed to read com template: %v", err))
		printInfo("com.go already exists from webview2 copy")
	}

	// Step 7: Generate version mapping
	printStep("Generating version mapping...")
	
	var buffer strings.Builder
	buffer.WriteString("//go:build windows\n\n")
	buffer.WriteString("package edge\n\n")
	buffer.WriteString("type Version struct {\n")
	buffer.WriteString("\tSDKVersion         string\n")
	buffer.WriteString("\tReleaseNotes           string\n")
	buffer.WriteString("\tRuntimeVersion string\n")
	buffer.WriteString("\tNotes          string\n")
	buffer.WriteString("}\n\n")
	buffer.WriteString("var versionMapping = map[string]Version{\n")
	for _, version := range versions {
		buffer.WriteString(fmt.Sprintf("\t\"%s\": {\n", version.Number))
		buffer.WriteString(fmt.Sprintf("\t\tSDKVersion:     \"%s\",\n", version.Number))
		buffer.WriteString(fmt.Sprintf("\t\tReleaseNotes:   \"%s\",\n", version.ReleaseNotes))
		buffer.WriteString(fmt.Sprintf("\t\tRuntimeVersion: \"%s\",\n", version.RuntimeVersion))
		buffer.WriteString("\t\tNotes: ")
		buffer.WriteString(fmt.Sprintf("\t\t\t`%s`,\n", strings.Replace(strings.Join(version.Notes, "\n"), "`", "'", -1)))
		buffer.WriteString("\t},\n")
	}
	buffer.WriteString("}\n")

	// Write the buffer to edge/version_map.go
	err = os.WriteFile(filepath.Join(edgeDir, "version_map.go"), []byte(buffer.String()), 0644)
	if err != nil {
		printError(fmt.Sprintf("Failed to write version mapping: %v", err))
		os.Exit(1)
	}
	printSuccess("Version mapping generated")

	// Step 8: Update capabilities for latest version
	printStep("Updating capabilities for version guards...")
	
	capabilitiesPath := filepath.Join(edgeDir, "capabilities.go")
	if _, err := os.Stat(capabilitiesPath); err == nil {
		// Read the current capabilities file
		capData, err := os.ReadFile(capabilitiesPath)
		if err == nil {
			capContent := string(capData)
			
			// Add a new capability for the latest version if it doesn't exist
			latestCapabilityComment := fmt.Sprintf("// WebView2 Runtime Version %s", latestVersion)
			if !strings.Contains(capContent, latestCapabilityComment) {
				// Find the position to insert the new capability (after imports and before constants)
				insertPos := strings.Index(capContent, "// WebView2 Runtime Version")
				if insertPos > 0 {
					newCapability := fmt.Sprintf("%s (Released: %s)\nconst (\n\tLatestFeatures = Capability(\"%s\") // Latest WebView2 features\n)\n\n", 
						latestCapabilityComment, 
						time.Now().Format("January 2006"),
						latestVersion)
					
					updatedContent := capContent[:insertPos] + newCapability + capContent[insertPos:]
					
					err = os.WriteFile(capabilitiesPath, []byte(updatedContent), 0644)
					if err != nil {
						printWarning(fmt.Sprintf("Failed to update capabilities: %v", err))
					} else {
						printSuccess("Updated capabilities with latest version")
					}
				}
			} else {
				printInfo("Capabilities already include latest version")
			}
		}
	} else {
		printWarning("capabilities.go not found - skipping capability update")
	}

	// Step 9: Format generated code with go fmt
	printStep("Formatting generated Go code...")
	
	// Change to the edge directory and run go fmt
	err = runCommand("go", []string{"fmt", "."}, edgeDir)
	if err != nil {
		printWarning(fmt.Sprintf("Failed to format code: %v", err))
	} else {
		printSuccess("All Go files formatted successfully")
	}

	// Step 10: Save version information
	printStep("Saving version information...")
	
	// Save the latest release notes to a file
	if len(versions) > 0 {
		latestReleaseNotes := fmt.Sprintf("Version: %s\nRuntime Version: %s\nRelease Notes URL: %s\n\nNotes:\n%s",
			versions[0].Number,
			versions[0].RuntimeVersion,
			versions[0].ReleaseNotes,
			strings.Join(versions[0].Notes, "\n"))
		err = os.WriteFile("latest_release_notes.txt", []byte(latestReleaseNotes), 0644)
		if err != nil {
			printWarning(fmt.Sprintf("Failed to write release notes: %v", err))
		} else {
			printInfo("Release notes saved to latest_release_notes.txt")
		}
	}

	// Save the version to latest_version.txt
	err = os.WriteFile("latest_version.txt", []byte(latestVersion), 0644)
	if err != nil {
		printError(fmt.Sprintf("Failed to save version: %v", err))
		os.Exit(1)
	}

	// Final success message
	fmt.Println()
	fmt.Printf("%s%s╔══════════════════════════════════════════════════════════════════╗%s\n", ColorBold, ColorGreen, ColorReset)
	fmt.Printf("%s%s║                         UPDATE COMPLETE!                        ║%s\n", ColorBold, ColorGreen, ColorReset)
	fmt.Printf("%s%s╚══════════════════════════════════════════════════════════════════╝%s\n", ColorBold, ColorGreen, ColorReset)
	fmt.Printf("%s%sWebView2 Edge Package Updated Successfully!%s\n", ColorBold, ColorGreen, ColorReset)
	fmt.Printf("Version: %s%s%s\n", ColorCyan, latestVersion, ColorReset)
	fmt.Printf("Location: %s%s%s\n", ColorCyan, edgeDir, ColorReset)
	if backupDir != "" {
		fmt.Printf("Backup: %s%s%s\n", ColorYellow, backupDir, ColorReset)
	}
	fmt.Println()
}

func DownloadIDL(version string) ([]byte, error) {
	// Look for the file locally: WebView2.version.idl
	data, err := os.ReadFile("WebView2." + version + ".idl")
	if err == nil {
		printInfo("Using cached IDL file")
		return data, nil
	}

	// URL for the nuget package: https://www.nuget.org/api/v2/package/Microsoft.Web.WebView2/<version>
	// Download the package to the current directory
	client := resty.New()
	printInfo(fmt.Sprintf("Downloading NuGet package: Microsoft.Web.WebView2/%s", version))
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
	
	if len(idlData) == 0 {
		return nil, fmt.Errorf("WebView2.idl not found in NuGet package")
	}
	
	// Write IDL to disk for caching
	err = os.WriteFile("WebView2."+version+".idl", idlData, 0755)
	if err != nil {
		printWarning(fmt.Sprintf("Failed to cache IDL file: %v", err))
	} else {
		printInfo("IDL file cached for future use")
	}
	
	return idlData, nil
}

// CompareBrowserVersions will compare the 2 given versions and return:
//
//	-1 = v1 < v2
//	 0 = v1 == v2
//	 1 = v1 > v2
func CompareBrowserVersions(v1 string, v2 string) (int, error) {
	v, err := parseVersion(v1)
	if err != nil {
		return 0, fmt.Errorf("v1 invalid: %w", err)
	}

	w, err := parseVersion(v2)
	if err != nil {
		return 0, fmt.Errorf("v2 invalid: %w", err)
	}

	return v.compare(w), nil
}

type version struct {
	major int
	minor int
	patch int
	build int

	channel string
}

func (v version) String() string {
	vv := fmt.Sprintf("%d.%d.%d.%d", v.major, v.minor, v.patch, v.build)
	if v.channel != "" {
		vv += " " + v.channel
	}

	return vv
}

func (v version) compare(o version) int {
	if c := compareInt(v.major, o.major); c != 0 {
		return c
	}
	if c := compareInt(v.minor, o.minor); c != 0 {
		return c
	}
	if c := compareInt(v.patch, o.patch); c != 0 {
		return c
	}
	return compareInt(v.build, o.build)
}

func parseVersion(v string) (version, error) {
	var p version

	// Split away channel information...
	if i := strings.Index(v, " "); i > 0 {
		p.channel = v[i+1:]
		v = v[:i]
	}

	vv := strings.Split(v, ".")
	if len(vv) > 4 {
		return p, fmt.Errorf("too many version parts")
	}

	var err error
	vv, p.major, err = parseInt(vv)
	if err != nil {
		return p, fmt.Errorf("bad major version: %w", err)
	}

	vv, p.minor, err = parseInt(vv)
	if err != nil {
		return p, fmt.Errorf("bad minor version: %w", err)
	}

	vv, p.patch, err = parseInt(vv)
	if err != nil {
		return p, fmt.Errorf("bad patch version: %w", err)
	}

	_, p.build, err = parseInt(vv)
	if err != nil {
		return p, fmt.Errorf("bad build version: %w", err)
	}

	return p, nil
}

func parseInt(v []string) ([]string, int, error) {
	if len(v) == 0 {
		return nil, 0, nil
	}

	p, err := strconv.ParseInt(v[0], 10, 32)
	if err != nil {
		return nil, 0, err
	}
	return v[1:], int(p), nil
}

func compareInt(v1, v2 int) int {
	if v1 == v2 {
		return 0
	}
	if v1 < v2 {
		return -1
	} else {
		return +1
	}
}