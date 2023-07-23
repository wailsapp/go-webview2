package types

import "strings"

var idlTypeToGoType = map[string]string{
	"IUnknown":               "_IUnknown",
	"EventRegistrationToken": "EventRegistrationToken",
	"LPWSTR":                 "string",
	"LPCWSTR":                "string",
	"HRESULT":                "uintptr",
	"UINT64":                 "uint64",
	"UINT32":                 "uint32",
	"UINT":                   "uint",
	"INT":                    "int",
	"BOOL":                   "bool",
	"BYTE":                   "uint8",
	"double":                 "float64",
}

func IdlTypeToGoType(input string) string {
	result := idlTypeToGoType[input]
	if result == "" {
		return input
	}
	return result
}

func defaultErrorValue(outputType string) string {
	if strings.HasPrefix(outputType, "uint") ||
		strings.HasPrefix(outputType, "int") {
		return "0"
	}
	if strings.HasPrefix(outputType, "float") {
		return "0.0"
	}
	switch outputType {
	case "string":
		return `""`
	case "bool":
		return "false"
	}
	return "nil"
}
