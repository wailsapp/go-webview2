package idl

// BasicTypes maps IDL basic types to their Go equivalents
var BasicTypes = map[string]string{
	// Basic types
	"void":    "void",
	"VOID":    "void",
	"HRESULT": "syscall.Errno",

	// Integer types
	"BOOL":      "int32",
	"BYTE":      "byte",
	"WORD":      "uint16",
	"DWORD":     "uint32",
	"QWORD":     "uint64",
	"INT":       "int32",
	"UINT":      "uint32",
	"LONG":      "int32",
	"ULONG":     "uint32",
	"LONGLONG":  "int64",
	"ULONGLONG": "uint64",
	"SHORT":     "int16",
	"USHORT":    "uint16",
	"CHAR":      "int8",
	"UCHAR":     "uint8",
	"WCHAR":     "uint16",

	// Floating point
	"FLOAT":  "float32",
	"DOUBLE": "float64",

	// String types
	"LPSTR":    "*byte",
	"LPCSTR":   "*byte",
	"LPWSTR":   "*uint16",
	"LPCWSTR":  "*uint16",
	"BSTR":     "*uint16",
	"OLECHAR":  "uint16",
	"LOLESTR":  "*uint16",
	"LPOLESTR": "*uint16",

	// Handle types - use uintptr for syscalls
	"HANDLE":    "uintptr",
	"HWND":      "uintptr",
	"HDC":       "uintptr",
	"HBITMAP":   "uintptr",
	"HICON":     "uintptr",
	"HCURSOR":   "uintptr",
	"HBRUSH":    "uintptr",
	"HPEN":      "uintptr",
	"HFONT":     "uintptr",
	"HMODULE":   "uintptr",
	"HINSTANCE": "uintptr",
	"HKEY":      "uintptr",

	// Pointer types
	"PVOID":   "unsafe.Pointer",
	"LPVOID":  "unsafe.Pointer",
	"LPCVOID": "unsafe.Pointer",

	// COM types
	"IUnknown":  "*combridge.IUnknown",
	"IDispatch": "*combridge.IDispatch",
	"VARIANT":   "combridge.Variant",
	"SAFEARRAY": "*combridge.SafeArray",
	"GUID":      "combridge.GUID",
	"CLSID":     "combridge.GUID",
	"IID":       "combridge.GUID",
	"REFIID":    "*combridge.GUID",
	"REFCLSID":  "*combridge.GUID",
	"REFGUID":   "*combridge.GUID",

	// Event types
	"EventRegistrationToken": "combridge.EventToken",

	// Size and position types
	"SIZE":     "combridge.Size",
	"POINT":    "combridge.Point",
	"RECT":     "combridge.Rect",
	"COLORREF": "uint32",
}

// WindowsTypes maps Windows-specific types to their syscall equivalents
var WindowsTypes = map[string]string{
	"HWND":      "uintptr",
	"HDC":       "uintptr",
	"HBITMAP":   "uintptr",
	"HICON":     "uintptr",
	"HCURSOR":   "uintptr",
	"HBRUSH":    "uintptr",
	"HPEN":      "uintptr",
	"HFONT":     "uintptr",
	"HMODULE":   "uintptr",
	"HINSTANCE": "uintptr",
	"HANDLE":    "uintptr",
	"HKEY":      "uintptr",
}

// COMTypes represents COM interface types
var COMTypes = map[string]string{
	"IUnknown":  "*combridge.IUnknown",
	"IDispatch": "*combridge.IDispatch",
}

// IsBasicType checks if a type name is a basic IDL type
func IsBasicType(typeName string) bool {
	_, exists := BasicTypes[typeName]
	return exists
}

// IsWindowsType checks if a type name is a Windows handle type
func IsWindowsType(typeName string) bool {
	_, exists := WindowsTypes[typeName]
	return exists
}

// IsCOMType checks if a type name is a COM interface type
func IsCOMType(typeName string) bool {
	_, exists := COMTypes[typeName]
	return exists
}

// IsInterfaceType checks if a type name looks like an interface
func IsInterfaceType(typeName string) bool {
	return len(typeName) > 1 && typeName[0] == 'I' &&
		(typeName[1] >= 'A' && typeName[1] <= 'Z')
}

// IsEnumType checks if a type name looks like an enum
func IsEnumType(typeName string) bool {
	return len(typeName) > 12 && typeName[:12] == "COREWEBVIEW2"
}

// IsEventHandlerType checks if a type name is an event handler
func IsEventHandlerType(typeName string) bool {
	return len(typeName) > 12 &&
		(typeName[len(typeName)-12:] == "EventHandler" ||
			typeName[len(typeName)-16:] == "CompletedHandler")
}

// GetGoType converts an IDL type to its Go equivalent
func GetGoType(idlType *Type) string {
	if idlType == nil {
		return "interface{}"
	}

	// Handle basic types first
	if goType, exists := BasicTypes[idlType.Name]; exists {
		if idlType.Pointer {
			return "*" + goType
		}
		return goType
	}

	// Handle Windows types for syscalls
	if goType, exists := WindowsTypes[idlType.Name]; exists {
		return goType
	}

	// Handle interfaces
	if IsInterfaceType(idlType.Name) {
		if idlType.Pointer {
			return "*" + idlType.Name
		}
		return "*" + idlType.Name
	}

	// Handle enums
	if IsEnumType(idlType.Name) {
		return idlType.Name
	}

	// Handle generic types
	if len(idlType.Generic) > 0 {
		// For now, treat generics as interface{}
		return "interface{}"
	}

	// Handle arrays
	if idlType.Array {
		elementType := &Type{
			Name:    idlType.Name,
			Kind:    idlType.Kind,
			Pointer: false,
			Array:   false,
		}
		return "[]" + GetGoType(elementType)
	}

	// Handle pointers
	if idlType.Pointer {
		baseType := &Type{
			Name:    idlType.Name,
			Kind:    idlType.Kind,
			Pointer: false,
		}
		return "*" + GetGoType(baseType)
	}

	// Default to the type name as-is
	return idlType.Name
}

// GetSyscallType converts an IDL type to its Windows syscall equivalent
func GetSyscallType(idlType *Type) string {
	if idlType == nil {
		return "uintptr"
	}

	// Handle Windows handle types
	if IsWindowsType(idlType.Name) {
		return "uintptr"
	}

	// Handle basic types that need syscall conversion
	switch idlType.Name {
	case "HRESULT":
		return "uintptr"
	case "BOOL":
		return "uintptr"
	case "DWORD":
		return "uintptr"
	case "LPWSTR", "LPCWSTR":
		return "uintptr"
	case "LPSTR", "LPCSTR":
		return "uintptr"
	case "PVOID", "LPVOID", "LPCVOID":
		return "uintptr"
	}

	// Handle interfaces (they become COM pointers)
	if IsInterfaceType(idlType.Name) {
		return "uintptr"
	}

	// Handle enums (they're usually DWORDs)
	if IsEnumType(idlType.Name) {
		return "uintptr"
	}

	// Handle pointers
	if idlType.Pointer {
		return "uintptr"
	}

	// Default for syscalls
	return "uintptr"
}

// NeedsConversion checks if a type needs conversion between Go and syscall
func NeedsConversion(idlType *Type) bool {
	if idlType == nil {
		return false
	}

	// String types need conversion
	switch idlType.Name {
	case "LPWSTR", "LPCWSTR", "LPSTR", "LPCSTR", "BSTR":
		return true
	}

	// Interface types need conversion
	if IsInterfaceType(idlType.Name) {
		return true
	}

	// Handle types need conversion
	if IsWindowsType(idlType.Name) {
		return true
	}

	// BOOL needs conversion
	if idlType.Name == "BOOL" {
		return true
	}

	return false
}

// GetConversionCode generates Go code to convert between types
func GetConversionCode(idlType *Type, varName string, toSyscall bool) string {
	if !NeedsConversion(idlType) {
		return varName
	}

	if toSyscall {
		// Convert Go type to syscall type
		switch idlType.Name {
		case "LPWSTR", "LPCWSTR":
			return "uintptr(unsafe.Pointer(" + varName + "))"
		case "LPSTR", "LPCSTR":
			return "uintptr(unsafe.Pointer(" + varName + "))"
		case "BOOL":
			return "uintptr(boolToUint32(" + varName + "))"
		}

		if IsInterfaceType(idlType.Name) {
			return "uintptr(unsafe.Pointer(" + varName + "))"
		}

		if IsWindowsType(idlType.Name) {
			return "uintptr(" + varName + ")"
		}
	} else {
		// Convert syscall type to Go type
		switch idlType.Name {
		case "LPWSTR":
			return "(*uint16)(unsafe.Pointer(" + varName + "))"
		case "LPCWSTR":
			return "(*uint16)(unsafe.Pointer(" + varName + "))"
		case "LPSTR":
			return "(*byte)(unsafe.Pointer(" + varName + "))"
		case "LPCSTR":
			return "(*byte)(unsafe.Pointer(" + varName + "))"
		case "BOOL":
			return "uint32ToBool(uint32(" + varName + "))"
		}

		if IsInterfaceType(idlType.Name) {
			return "(*" + idlType.Name + ")(unsafe.Pointer(" + varName + "))"
		}

		if IsWindowsType(idlType.Name) {
			return idlType.Name + "(" + varName + ")"
		}
	}

	return varName
}
