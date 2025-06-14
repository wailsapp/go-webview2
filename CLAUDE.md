# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## About This Project

This is a Wails-maintained fork of go-webview2 that provides Windows WebView2 integration without CGo. It's specifically designed for Wails applications and is not intended as a standalone package.

## Common Commands

### Development Tasks
- `task update` - Updates WebView2 version mappings and regenerates code
- `task update:forced` - Forced update of version mappings (use when normal update fails)
- `task test` - Run all tests
- `task gofmt` - Format generated code in pkg/webview2
- `task commit` - Add all changes and commit with standard message

### Go Commands
- `go test ./...` - Run all tests
- `go run ./cmd/demo` - Run the demo application (requires WebView2 runtime)

## Architecture Overview

### Core Components

1. **`/pkg/webview2/`** - Auto-generated COM interface bindings from Microsoft WebView2 IDL files. Contains hundreds of interfaces, enums, and event handlers. Do not manually edit these files.

2. **`/pkg/edge/`** - Higher-level wrapper API that provides:
   - Simplified WebView2 interface (`chromium.go`)
   - Version capability detection (`capabilities.go`, `version_map.go`)
   - Platform-specific implementations (`chromium_*.go`)

3. **`/pkg/combridge/`** - COM object management layer:
   - Interface resolution and vtable handling
   - Reference counting and memory management
   - Type-safe Go interfaces for COM objects

4. **`/webviewloader/`** - WebView2 runtime loader:
   - Embedded WebView2Loader.dll files for x86/x64/arm64
   - Environment creation and initialization
   - Runtime version detection

5. **`/scripts/`** - Code generation system:
   - IDL parser that converts Microsoft WebView2 IDL to Go code
   - Automatic version mapping updates from WebView2 releases
   - Template-based code generation

### Code Generation Workflow

The `/pkg/webview2/` directory is entirely auto-generated. When working on the generator:
- IDL parsing logic is in `/scripts/generator/parser.go`
- Type definitions and templates are in `/scripts/generator/types/`
- Test files for generation validation are in `/scripts/generator/testfiles/`
- The main update script is `/scripts/update_version_mapping.go`

### Version Management

- `version_map.go` tracks WebView2 capabilities across runtime versions
- Capability detection allows graceful degradation for older runtimes
- Version mappings are automatically updated from Microsoft's release notes

## Key Files for Development

- **`pkg/edge/chromium.go`** - Main WebView2 wrapper implementation
- **`pkg/edge/capabilities.go`** - Runtime capability detection
- **`pkg/combridge/bridge.go`** - COM object lifecycle management
- **`webviewloader/env_create.go`** - WebView2 environment creation
- **`scripts/update_version_mapping.go`** - Version mapping automation

## Testing

- Capability tests validate version compatibility (630+ test cases)
- Generator tests ensure IDL parsing correctness
- Cookie management and WebView2 functionality tests
- Run with `go test ./...` or `task test`

## Requirements

- Go 1.20+
- Windows operating system
- WebView2 runtime installed
- Task runner for automation tasks