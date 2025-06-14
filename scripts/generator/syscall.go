package generator

import (
	"fmt"
	"strings"

	"generator/idl"
)

// SyscallGenerator handles Windows syscall wrapper generation
type SyscallGenerator struct{}

// NewSyscallGenerator creates a new syscall generator
func NewSyscallGenerator() *SyscallGenerator {
	return &SyscallGenerator{}
}

// GenerateMethodSignature generates the Go method signature for a COM method
func (sg *SyscallGenerator) GenerateMethodSignature(method *idl.Method) string {
	var params []string

	for _, param := range method.Parameters {
		goType := idl.GetGoType(param.Type)
		params = append(params, fmt.Sprintf("%s %s", param.Name, goType))
	}

	returnType := "error"
	if method.ReturnType.Name != "HRESULT" {
		returnType = idl.GetGoType(method.ReturnType)
	}

	return fmt.Sprintf("%s(%s) %s", method.Name, strings.Join(params, ", "), returnType)
}

// GenerateSyscallParams generates syscall parameter conversions
func (sg *SyscallGenerator) GenerateSyscallParams(method *idl.Method) []string {
	var params []string

	// Always start with 'this' pointer
	params = append(params, "uintptr(unsafe.Pointer(obj))")

	for _, param := range method.Parameters {
		conversion := idl.GetConversionCode(param.Type, param.Name, true)
		params = append(params, conversion)
	}

	return params
}

// GenerateReturnConversion generates return value conversion
func (sg *SyscallGenerator) GenerateReturnConversion(method *idl.Method) string {
	if method.ReturnType.Name == "HRESULT" {
		return `if ret != 0 {
        return syscall.Errno(ret)
    }
    return nil`
	}

	return fmt.Sprintf("return %s", idl.GetConversionCode(method.ReturnType, "ret", false))
}

// GenerateVTableEntry generates a VTable entry for a method
func (sg *SyscallGenerator) GenerateVTableEntry(method *idl.Method, offset int) string {
	return fmt.Sprintf("    %s uintptr // Offset %d", method.Name, offset)
}

// GenerateSyscallWrapper generates a complete syscall wrapper for a method
func (sg *SyscallGenerator) GenerateSyscallWrapper(interfaceName string, method *idl.Method) string {
	signature := sg.GenerateMethodSignature(method)
	params := sg.GenerateSyscallParams(method)
	returnConv := sg.GenerateReturnConversion(method)

	template := `// %s %s
func (obj *%s) %s {
    ret, _, _ := syscall.SyscallN(
        obj.VTable().%s,
        %s,
    )
    %s
}`

	return fmt.Sprintf(template,
		method.Name,
		strings.Join(method.Comments, " "),
		interfaceName,
		signature,
		method.Name,
		strings.Join(params, ",\n        "),
		returnConv,
	)
}

// CalculateVTableSize calculates the size of a VTable for an interface
func (sg *SyscallGenerator) CalculateVTableSize(iface *idl.Interface) int {
	// Start with IUnknown methods (QueryInterface, AddRef, Release)
	baseSize := 3
	return baseSize + len(iface.Methods)
}
