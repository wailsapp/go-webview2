package idl

import (
	"fmt"
	"strings"
)

// AST represents the Abstract Syntax Tree for an IDL file
type AST struct {
	Library    *Library
	Imports    []*Import
	Interfaces []*Interface
	Enums      []*Enum
	TypeDefs   []*TypeDef
}

// Library represents a COM library declaration
type Library struct {
	Name       string
	UUID       string
	Version    string
	Attributes []*Attribute
}

// Import represents an import statement
type Import struct {
	Path string
	File string
}

// Interface represents a COM interface
type Interface struct {
	Name       string
	UUID       string
	Parent     string
	Methods    []*Method
	Properties []*Property
	Attributes []*Attribute
	Comments   []string
}

// Method represents an interface method
type Method struct {
	Name       string
	ReturnType *Type
	Parameters []*Parameter
	Attributes []*Attribute
	Comments   []string
}

// Property represents an interface property
type Property struct {
	Name       string
	Type       *Type
	Getter     bool
	Setter     bool
	Attributes []*Attribute
	Comments   []string
}

// Parameter represents a method parameter
type Parameter struct {
	Name       string
	Type       *Type
	Direction  ParameterDirection
	Attributes []*Attribute
	Optional   bool
	RetVal     bool
}

// ParameterDirection indicates parameter direction
type ParameterDirection int

const (
	ParamIn ParameterDirection = iota
	ParamOut
	ParamInOut
	ParamRetVal
)

func (d ParameterDirection) String() string {
	switch d {
	case ParamIn:
		return "in"
	case ParamOut:
		return "out"
	case ParamInOut:
		return "in,out"
	case ParamRetVal:
		return "retval"
	default:
		return "unknown"
	}
}

// Enum represents an enumeration
type Enum struct {
	Name       string
	Values     []*EnumValue
	Attributes []*Attribute
	Comments   []string
}

// EnumValue represents an enum value
type EnumValue struct {
	Name     string
	Value    int64
	Comments []string
}

// TypeDef represents a type definition
type TypeDef struct {
	Name       string
	Type       *Type
	Attributes []*Attribute
	Comments   []string
}

// Type represents a data type
type Type struct {
	Name      string
	Kind      TypeKind
	Pointer   bool
	Array     bool
	Const     bool
	Reference bool
	Generic   []*Type // For generic types like IVector<T>
}

// TypeKind represents the kind of type
type TypeKind int

const (
	TypeBasic TypeKind = iota
	TypeInterface
	TypeEnum
	TypeStruct
	TypeUnion
	TypePointer
	TypeArray
	TypeGeneric
)

func (k TypeKind) String() string {
	switch k {
	case TypeBasic:
		return "basic"
	case TypeInterface:
		return "interface"
	case TypeEnum:
		return "enum"
	case TypeStruct:
		return "struct"
	case TypeUnion:
		return "union"
	case TypePointer:
		return "pointer"
	case TypeArray:
		return "array"
	case TypeGeneric:
		return "generic"
	default:
		return "unknown"
	}
}

// Attribute represents an IDL attribute
type Attribute struct {
	Name   string
	Values []string
}

// Helper methods for the AST

// String returns a string representation of the type
func (t *Type) String() string {
	var parts []string

	if t.Const {
		parts = append(parts, "const")
	}

	parts = append(parts, t.Name)

	if t.Pointer {
		parts = append(parts, "*")
	}

	if t.Reference {
		parts = append(parts, "&")
	}

	if t.Array {
		parts = append(parts, "[]")
	}

	if len(t.Generic) > 0 {
		var genericParts []string
		for _, g := range t.Generic {
			genericParts = append(genericParts, g.String())
		}
		parts = append(parts, fmt.Sprintf("<%s>", strings.Join(genericParts, ", ")))
	}

	return strings.Join(parts, " ")
}

// IsBasicType returns true if this is a basic type (int, string, etc.)
func (t *Type) IsBasicType() bool {
	return t.Kind == TypeBasic
}

// IsInterface returns true if this is an interface type
func (t *Type) IsInterface() bool {
	return t.Kind == TypeInterface
}

// IsEnum returns true if this is an enum type
func (t *Type) IsEnum() bool {
	return t.Kind == TypeEnum
}

// IsPointer returns true if this is a pointer type
func (t *Type) IsPointer() bool {
	return t.Pointer || t.Kind == TypePointer
}

// HasAttribute checks if an interface has a specific attribute
func (i *Interface) HasAttribute(name string) bool {
	return hasAttribute(i.Attributes, name)
}

// GetAttribute gets an attribute value by name
func (i *Interface) GetAttribute(name string) *Attribute {
	return getAttribute(i.Attributes, name)
}

// HasAttribute checks if a method has a specific attribute
func (m *Method) HasAttribute(name string) bool {
	return hasAttribute(m.Attributes, name)
}

// GetAttribute gets an attribute value by name
func (m *Method) GetAttribute(name string) *Attribute {
	return getAttribute(m.Attributes, name)
}

// HasAttribute checks if a parameter has a specific attribute
func (p *Parameter) HasAttribute(name string) bool {
	return hasAttribute(p.Attributes, name)
}

// GetAttribute gets an attribute value by name
func (p *Parameter) GetAttribute(name string) *Attribute {
	return getAttribute(p.Attributes, name)
}

// Helper functions
func hasAttribute(attrs []*Attribute, name string) bool {
	for _, attr := range attrs {
		if attr.Name == name {
			return true
		}
	}
	return false
}

func getAttribute(attrs []*Attribute, name string) *Attribute {
	for _, attr := range attrs {
		if attr.Name == name {
			return attr
		}
	}
	return nil
}

// IsEventHandler returns true if this interface is an event handler
func (i *Interface) IsEventHandler() bool {
	return strings.HasSuffix(i.Name, "EventHandler") ||
		strings.HasSuffix(i.Name, "CompletedHandler")
}

// IsVersioned returns true if this interface has version information
func (i *Interface) IsVersioned() bool {
	// Check for version suffixes like _2, _3, _10, etc.
	if !strings.Contains(i.Name, "_") {
		return false
	}

	parts := strings.Split(i.Name, "_")
	if len(parts) < 2 {
		return false
	}

	// Check if the last part is a number
	lastPart := parts[len(parts)-1]
	var version int
	n, err := fmt.Sscanf(lastPart, "%d", &version)
	return n == 1 && err == nil && version > 0
}

// GetVersion extracts version number from interface name (e.g., ICoreWebView2_3 -> 3)
func (i *Interface) GetVersion() int {
	if !i.IsVersioned() {
		return 1
	}

	parts := strings.Split(i.Name, "_")
	if len(parts) < 2 {
		return 1
	}

	// Try to parse the version number
	versionStr := parts[len(parts)-1]
	var version int
	fmt.Sscanf(versionStr, "%d", &version)
	if version == 0 {
		return 1
	}

	return version
}

// GetBaseInterfaceName returns the base interface name without version suffix
func (i *Interface) GetBaseInterfaceName() string {
	if !i.IsVersioned() {
		return i.Name
	}

	parts := strings.Split(i.Name, "_")
	if len(parts) < 2 {
		return i.Name
	}

	return strings.Join(parts[:len(parts)-1], "_")
}
