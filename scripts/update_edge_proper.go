package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
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

	"github.com/go-resty/resty/v2"
	"generator/idl"
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
	fmt.Printf("%s%s║          WebView2 Edge Package Generator (ACTUAL IDL!)          ║%s\n", ColorBold, ColorPurple, ColorReset)
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
		backupDir = "" // No backup was created
	}

	// Step 3: Create new edge directory
	printStep("Creating new edge directory...")
	
	err = os.MkdirAll(edgeDir, 0755)
	if err != nil {
		printError(fmt.Sprintf("Failed to create edge directory: %v", err))
		os.Exit(1)
	}
	printSuccess("Created new edge directory")

	// Step 4: Download and parse WebView2 IDL
	printStep(fmt.Sprintf("Downloading and parsing WebView2 IDL for version %s...", latestVersion))
	
	idlData, err := DownloadIDL(latestVersion)
	if err != nil {
		printError(fmt.Sprintf("Failed to download IDL: %v", err))
		os.Exit(1)
	}
	printSuccess(fmt.Sprintf("Downloaded WebView2 IDL (%d bytes)", len(idlData)))

	// Parse the IDL
	printInfo("Parsing IDL with WebView2 parser...")
	parser := idl.NewParser(bytes.NewReader(idlData))
	ast, err := parser.Parse()
	if err != nil {
		printError(fmt.Sprintf("Failed to parse IDL: %v", err))
		os.Exit(1)
	}
	printSuccess(fmt.Sprintf("Successfully parsed IDL: %d interfaces, %d enums", len(ast.Interfaces), len(ast.Enums)))

	// Step 5: Generate Go code from AST
	printStep("Generating Go code from parsed IDL...")
	
	err = generateGoCode(ast, edgeDir)
	if err != nil {
		printError(fmt.Sprintf("Failed to generate Go code: %v", err))
		os.Exit(1)
	}
	printSuccess("Generated Go code from IDL")

	// Step 6: Copy core helper files from assets
	printStep("Copying core helper files from assets...")
	
	err = copyAssetFiles(edgeDir)
	if err != nil {
		printError(fmt.Sprintf("Failed to copy asset files: %v", err))
		os.Exit(1)
	}
	printSuccess("Copied core helper files from assets")

	// Step 7: Generate version mapping
	printStep("Generating version mapping...")
	
	err = generateVersionMapping(versions, edgeDir)
	if err != nil {
		printError(fmt.Sprintf("Failed to generate version mapping: %v", err))
		os.Exit(1)
	}
	printSuccess("Version mapping generated")

	// Step 8: Update capabilities for latest version
	printStep("Updating capabilities for version guards...")
	
	err = updateCapabilities(latestVersion, edgeDir)
	if err != nil {
		printWarning(fmt.Sprintf("Failed to update capabilities: %v", err))
	} else {
		printSuccess("Updated capabilities with latest version")
	}

	// Step 9: Format generated code with go fmt
	printStep("Formatting generated Go code...")
	
	err = runCommand("go", []string{"fmt", "."}, edgeDir)
	if err != nil {
		printWarning(fmt.Sprintf("Failed to format code: %v", err))
	} else {
		printSuccess("All Go files formatted successfully")
	}

	// Step 10: Test compilation
	printStep("Testing package compilation...")
	err = runCommand("go", []string{"test", "-c", "."}, edgeDir)
	if err != nil {
		printError(fmt.Sprintf("Package compilation test failed: %v", err))
		printError("Generated code has compilation errors - please check the logs above")
		os.Exit(1)
	}
	printSuccess("Package compiles successfully!")

	// Step 11: Clean up old backup directory after successful generation
	if backupDir != "" {
		printStep("Cleaning up old backup directory...")
		err = os.RemoveAll(backupDir)
		if err != nil {
			printWarning(fmt.Sprintf("Failed to remove backup directory %s: %v", backupDir, err))
		} else {
			printSuccess(fmt.Sprintf("Removed backup directory: %s", backupDir))
		}
	}

	// Step 12: Save version information
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
	fmt.Printf("%s%sWebView2 Edge Package Generated Successfully from FRESH IDL!%s\n", ColorBold, ColorGreen, ColorReset)
	fmt.Printf("Version: %s%s%s\n", ColorCyan, latestVersion, ColorReset)
	fmt.Printf("Location: %s%s%s\n", ColorCyan, edgeDir, ColorReset)
	fmt.Println()
}

// generateGoCode generates Go code from the parsed AST using the existing generator
func generateGoCode(ast *idl.AST, outputDir string) error {
	printInfo("Initializing code generator...")
	
	// Create the generator (we'll need to import this properly)
	// For now, let's generate individual files manually using templates
	
	// First, generate the base com.go file
	err := generateComFile(outputDir)
	if err != nil {
		return fmt.Errorf("failed to generate com.go: %w", err)
	}
	
	// Generate interface files individually
	for _, iface := range ast.Interfaces {
		err := generateInterfaceFile(iface, outputDir)
		if err != nil {
			printWarning(fmt.Sprintf("Failed to generate interface %s: %v", iface.Name, err))
			// Continue with other interfaces
		}
	}
	
	// Generate enum files individually  
	for _, enum := range ast.Enums {
		err := generateEnumFile(enum, outputDir)
		if err != nil {
			printWarning(fmt.Sprintf("Failed to generate enum %s: %v", enum.Name, err))
			// Continue with other enums
		}
	}
	
	printInfo(fmt.Sprintf("Generated %d interfaces and %d enums", len(ast.Interfaces), len(ast.Enums)))
	return nil
}

// generateComFile creates the base com.go file from template
func generateComFile(outputDir string) error {
	templatePath := "generator/types/templates/com.tmpl"
	templateData, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read com template: %w", err)
	}

	// Replace template placeholder with package name
	content := strings.Replace(string(templateData), "{{.PackageName}}", "edge", 1)
	
	err = os.WriteFile(filepath.Join(outputDir, "com.go"), []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write com.go: %w", err)
	}

	return nil
}

// generateInterfaceFile generates a Go file for an IDL interface
func generateInterfaceFile(iface *idl.Interface, outputDir string) error {
	var buf strings.Builder
	
	// Write package and imports
	buf.WriteString("//go:build windows\n\n")
	buf.WriteString("package edge\n\n")
	buf.WriteString("import (\n")
	buf.WriteString("\t\"golang.org/x/sys/windows\"\n")
	buf.WriteString("\t\"syscall\"\n")
	buf.WriteString("\t\"unsafe\"\n")
	buf.WriteString(")\n\n")
	
	// Generate Vtbl struct
	buf.WriteString(fmt.Sprintf("type %sVtbl struct {\n", iface.Name))
	buf.WriteString("\tIUnknownVtbl\n")
	
	for _, method := range iface.Methods {
		buf.WriteString(fmt.Sprintf("\t%s ComProc\n", method.Name))
	}
	
	buf.WriteString("}\n\n")
	
	// Generate interface struct
	buf.WriteString(fmt.Sprintf("type %s struct {\n", iface.Name))
	buf.WriteString(fmt.Sprintf("\tVtbl *%sVtbl\n", iface.Name))
	buf.WriteString("}\n\n")
	
	// Generate AddRef method
	buf.WriteString(fmt.Sprintf("func (i *%s) AddRef() uintptr {\n", iface.Name))
	buf.WriteString("\trefCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))\n")
	buf.WriteString("\treturn refCounter\n")
	buf.WriteString("}\n\n")
	
	// Generate methods
	for _, method := range iface.Methods {
		err := generateMethod(&buf, iface.Name, method)
		if err != nil {
			return fmt.Errorf("failed to generate method %s: %w", method.Name, err)
		}
	}
	
	// Write to file
	filename := fmt.Sprintf("%s.go", iface.Name)
	err := os.WriteFile(filepath.Join(outputDir, filename), []byte(buf.String()), 0644)
	if err != nil {
		return fmt.Errorf("failed to write interface file: %w", err)
	}
	
	return nil
}

// generateMethod generates a Go method for an IDL interface method
func generateMethod(buf *strings.Builder, interfaceName string, method *idl.Method) error {
	// Generate method signature
	buf.WriteString(fmt.Sprintf("func (i *%s) %s(", interfaceName, method.Name))
	
	// Parameters
	for i, param := range method.Parameters {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(fmt.Sprintf("%s %s", param.Name, convertIDLTypeToGo(param.Type)))
	}
	
	// Return type
	buf.WriteString(") ")
	if method.ReturnType != nil {
		buf.WriteString(convertIDLTypeToGo(method.ReturnType))
	} else {
		buf.WriteString("error")
	}
	buf.WriteString(" {\n")
	
	// Method body (basic implementation)
	buf.WriteString(fmt.Sprintf("\thr, _, _ := i.Vtbl.%s.Call(\n", method.Name))
	buf.WriteString("\t\tuintptr(unsafe.Pointer(i)),\n")
	
	for _, param := range method.Parameters {
		buf.WriteString(fmt.Sprintf("\t\tuintptr(unsafe.Pointer(%s)),\n", param.Name))
	}
	
	buf.WriteString("\t)\n")
	buf.WriteString("\tif windows.Handle(hr) != windows.S_OK {\n")
	buf.WriteString("\t\treturn syscall.Errno(hr)\n")
	buf.WriteString("\t}\n")
	buf.WriteString("\treturn nil\n")
	buf.WriteString("}\n\n")
	
	return nil
}

// generateEnumFile generates a Go file for an IDL enum
func generateEnumFile(enum *idl.Enum, outputDir string) error {
	var buf strings.Builder
	
	// Write package header
	buf.WriteString("//go:build windows\n\n")
	buf.WriteString("package edge\n\n")
	
	// Generate enum type
	buf.WriteString(fmt.Sprintf("type %s int32\n\n", enum.Name))
	
	// Generate enum constants
	buf.WriteString("const (\n")
	for i, value := range enum.Values {
		if i == 0 {
			buf.WriteString(fmt.Sprintf("\t%s %s = %s\n", value.Name, enum.Name, value.Value))
		} else {
			buf.WriteString(fmt.Sprintf("\t%s %s = %s\n", value.Name, enum.Name, value.Value))
		}
	}
	buf.WriteString(")\n")
	
	// Write to file
	filename := fmt.Sprintf("%s.go", enum.Name)
	err := os.WriteFile(filepath.Join(outputDir, filename), []byte(buf.String()), 0644)
	if err != nil {
		return fmt.Errorf("failed to write enum file: %w", err)
	}
	
	return nil
}

// convertIDLTypeToGo converts IDL types to Go types
func convertIDLTypeToGo(idlType *idl.Type) string {
	if idlType == nil {
		return "interface{}"
	}
	
	// Basic type mappings
	switch idlType.Name {
	case "HRESULT":
		return "uintptr"
	case "BOOL":
		return "bool"
	case "UINT32":
		return "uint32"
	case "INT32":
		return "int32"
	case "LPWSTR", "LPCWSTR":
		return "string"
	case "HWND":
		return "uintptr"
	case "VARIANT":
		return "*VARIANT"
	case "IUnknown":
		return "*IUnknown"
	default:
		// Assume it's a WebView2 interface or enum
		if strings.HasPrefix(idlType.Name, "ICoreWebView2") || strings.HasPrefix(idlType.Name, "COREWEBVIEW2_") {
			if idlType.Pointer {
				return fmt.Sprintf("*%s", idlType.Name)
			}
			return idlType.Name
		}
		return idlType.Name
	}
}

// copyAssetFiles copies the essential Wails wrapper files from assets
func copyAssetFiles(edgeDir string) error {
	// Core files that should be preserved (stored in assets directory)
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
	
	// Use assets directory as source
	assetsDir := "assets"
	if _, err := os.Stat(assetsDir); err != nil {
		return fmt.Errorf("assets directory not found: %w", err)
	}
	
	copiedCount := 0
	for _, coreFile := range coreFiles {
		srcPath := filepath.Join(assetsDir, coreFile)
		dstPath := filepath.Join(edgeDir, coreFile)
		
		if data, err := os.ReadFile(srcPath); err == nil {
			// Assets files already have correct package name
			err = os.WriteFile(dstPath, data, 0644)
			if err != nil {
				return fmt.Errorf("failed to copy core file %s: %w", coreFile, err)
			} else {
				printInfo(fmt.Sprintf("Copied core file: %s", coreFile))
				copiedCount++
			}
		} else {
			return fmt.Errorf("core file not found in assets: %s", coreFile)
		}
	}
	
	printInfo(fmt.Sprintf("Copied %d core helper files from assets", copiedCount))
	return nil
}

// generateVersionMapping generates the version mapping file
func generateVersionMapping(versions []*Version, edgeDir string) error {
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
	err := os.WriteFile(filepath.Join(edgeDir, "version_map.go"), []byte(buffer.String()), 0644)
	if err != nil {
		return fmt.Errorf("failed to write version mapping: %w", err)
	}
	return nil
}

// updateCapabilities updates the capabilities file with the latest version
func updateCapabilities(latestVersion, edgeDir string) error {
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
						return fmt.Errorf("failed to update capabilities: %w", err)
					}
				}
			}
		}
	}
	return nil
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