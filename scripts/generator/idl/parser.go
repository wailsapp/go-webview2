package idl

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

// Parser represents an IDL parser
type Parser struct {
	scanner *Scanner
	ast     *AST
	current Token
	peek    Token
}

// NewParser creates a new IDL parser
func NewParser(reader io.Reader) *Parser {
	scanner := NewScanner(reader)
	p := &Parser{
		scanner: scanner,
		ast:     &AST{},
	}

	// Read first two tokens
	p.nextToken()
	p.nextToken()

	return p
}

// Parse parses the IDL file and returns an AST
func (p *Parser) Parse() (*AST, error) {
	for p.current.Type != TokenEOF {
		if err := p.parseTopLevel(); err != nil {
			return nil, err
		}
	}

	return p.ast, nil
}

// parseTopLevel parses top-level declarations
func (p *Parser) parseTopLevel() error {
	switch p.current.Type {
	case TokenImport:
		return p.parseImport()
	case TokenLibrary:
		return p.parseLibrary()
	case TokenInterface:
		return p.parseInterface()
	case TokenEnum:
		return p.parseEnum()
	case TokenTypedef:
		return p.parseTypedef()
	case TokenLBracket:
		// Parse attributes and then the declaration
		attrs, err := p.parseAttributes()
		if err != nil {
			return err
		}
		return p.parseAttributedDeclaration(attrs)
	case TokenComment:
		// Skip comments at top level
		p.nextToken()
		return nil
	case TokenIdentifier:
		// Handle special directives like cpp_quote
		if p.current.Value == "cpp_quote" {
			return p.skipCppQuote()
		}
		// Skip other unknown identifiers
		p.nextToken()
		return nil
	default:
		// Skip unknown tokens
		p.nextToken()
		return nil
	}
}

// skipCppQuote skips cpp_quote directives
func (p *Parser) skipCppQuote() error {
	// Skip "cpp_quote"
	p.nextToken()

	// Skip the opening parenthesis and string content
	if p.current.Type == TokenLParen {
		p.nextToken()

		// Find the closing parenthesis
		parenCount := 1
		for parenCount > 0 && p.current.Type != TokenEOF {
			if p.current.Type == TokenLParen {
				parenCount++
			} else if p.current.Type == TokenRParen {
				parenCount--
			}
			p.nextToken()
		}
	}

	return nil
}

// parseImport parses import statements
func (p *Parser) parseImport() error {
	if !p.expectToken(TokenImport) {
		return fmt.Errorf("expected 'import'")
	}

	if p.current.Type != TokenString {
		return fmt.Errorf("expected import path string")
	}

	path := strings.Trim(p.current.Value, "\"")
	p.nextToken()

	if !p.expectToken(TokenSemicolon) {
		return fmt.Errorf("expected ';' after import")
	}

	// Extract filename from path
	parts := strings.Split(path, "/")
	filename := parts[len(parts)-1]

	imp := &Import{
		Path: path,
		File: filename,
	}

	p.ast.Imports = append(p.ast.Imports, imp)
	return nil
}

// parseLibrary parses library declarations
func (p *Parser) parseLibrary() error {
	if !p.expectToken(TokenLibrary) {
		return fmt.Errorf("expected 'library'")
	}

	if p.current.Type != TokenIdentifier {
		return fmt.Errorf("expected library name")
	}

	name := p.current.Value
	p.nextToken()

	if !p.expectToken(TokenLBrace) {
		return fmt.Errorf("expected '{' after library name")
	}

	library := &Library{Name: name}

	// Parse library body contents
	for p.current.Type != TokenRBrace && p.current.Type != TokenEOF {
		if err := p.parseTopLevel(); err != nil {
			return err
		}
	}

	if !p.expectToken(TokenRBrace) {
		return fmt.Errorf("expected '}' to close library")
	}

	p.ast.Library = library
	return nil
}

// parseInterface parses interface declarations
func (p *Parser) parseInterface() error {
	if !p.expectToken(TokenInterface) {
		return fmt.Errorf("expected 'interface'")
	}

	if p.current.Type != TokenIdentifier {
		return fmt.Errorf("expected interface name")
	}

	name := p.current.Value
	p.nextToken()

	// Check for forward declaration (interface Name;)
	if p.current.Type == TokenSemicolon {
		p.nextToken()
		// This is just a forward declaration, skip it
		return nil
	}

	iface := &Interface{Name: name}

	// Parse inheritance
	if p.current.Type == TokenColon {
		p.nextToken()
		if p.current.Type != TokenIdentifier {
			return fmt.Errorf("expected parent interface name")
		}
		iface.Parent = p.current.Value
		p.nextToken()
	}

	if !p.expectToken(TokenLBrace) {
		return fmt.Errorf("expected '{' after interface declaration")
	}

	// Parse interface body
	for p.current.Type != TokenRBrace && p.current.Type != TokenEOF {
		if err := p.parseInterfaceMember(iface); err != nil {
			return err
		}
	}

	if !p.expectToken(TokenRBrace) {
		return fmt.Errorf("expected '}' to close interface")
	}

	// Optional semicolon
	if p.current.Type == TokenSemicolon {
		p.nextToken()
	}

	p.ast.Interfaces = append(p.ast.Interfaces, iface)
	return nil
}

// parseInterfaceMember parses interface members (methods and properties)
func (p *Parser) parseInterfaceMember(iface *Interface) error {
	var comments []string
	var attributes []*Attribute

	// Collect comments
	for p.current.Type == TokenComment {
		comments = append(comments, p.current.Value)
		p.nextToken()
	}

	// Parse attributes
	if p.current.Type == TokenLBracket {
		attrs, err := p.parseAttributes()
		if err != nil {
			return err
		}
		attributes = attrs
	}

	// Parse return type
	returnType, err := p.parseType()
	if err != nil {
		return err
	}

	// Parse method/property name
	if p.current.Type != TokenIdentifier {
		return fmt.Errorf("expected method or property name")
	}

	name := p.current.Value
	p.nextToken()

	if p.current.Type == TokenLParen {
		// It's a method
		method := &Method{
			Name:       name,
			ReturnType: returnType,
			Attributes: attributes,
			Comments:   comments,
		}

		if err := p.parseMethodParameters(method); err != nil {
			return err
		}

		if !p.expectToken(TokenSemicolon) {
			return fmt.Errorf("expected ';' after method declaration")
		}

		iface.Methods = append(iface.Methods, method)
	} else {
		// It's a property - handle [propget]/[propput] attributes
		prop := &Property{
			Name:       name,
			Type:       returnType,
			Attributes: attributes,
			Comments:   comments,
		}

		// Check for getter/setter attributes
		for _, attr := range attributes {
			if attr.Name == "propget" {
				prop.Getter = true
			} else if attr.Name == "propput" {
				prop.Setter = true
			}
		}

		if !p.expectToken(TokenSemicolon) {
			return fmt.Errorf("expected ';' after property declaration")
		}

		iface.Properties = append(iface.Properties, prop)
	}

	return nil
}

// parseMethodParameters parses method parameters
func (p *Parser) parseMethodParameters(method *Method) error {
	if !p.expectToken(TokenLParen) {
		return fmt.Errorf("expected '(' for method parameters")
	}

	// Handle empty parameter list
	if p.current.Type == TokenRParen {
		p.nextToken()
		return nil
	}

	for {
		param, err := p.parseParameter()
		if err != nil {
			return err
		}

		method.Parameters = append(method.Parameters, param)

		if p.current.Type == TokenComma {
			p.nextToken()
			continue
		} else if p.current.Type == TokenRParen {
			p.nextToken()
			break
		} else {
			return fmt.Errorf("expected ',' or ')' in parameter list")
		}
	}

	return nil
}

// parseParameter parses a method parameter
func (p *Parser) parseParameter() (*Parameter, error) {
	var attributes []*Attribute

	// Parse parameter attributes
	if p.current.Type == TokenLBracket {
		attrs, err := p.parseAttributes()
		if err != nil {
			return nil, err
		}
		attributes = attrs
	}

	// Parse parameter type
	paramType, err := p.parseType()
	if err != nil {
		return nil, err
	}

	// Parse parameter name
	var name string
	if p.current.Type == TokenIdentifier {
		name = p.current.Value
		p.nextToken()
	}

	param := &Parameter{
		Name:       name,
		Type:       paramType,
		Attributes: attributes,
		Direction:  ParamIn, // Default
	}

	// Determine parameter direction from attributes
	for _, attr := range attributes {
		switch attr.Name {
		case "in":
			param.Direction = ParamIn
		case "out":
			param.Direction = ParamOut
		case "retval":
			param.Direction = ParamRetVal
			param.RetVal = true
		}

		if len(attr.Values) > 0 {
			for _, val := range attr.Values {
				if val == "retval" {
					param.RetVal = true
				}
			}
		}
	}

	return param, nil
}

// parseEnum parses enum declarations
func (p *Parser) parseEnum() error {
	if !p.expectToken(TokenEnum) {
		return fmt.Errorf("expected 'enum'")
	}

	if p.current.Type != TokenIdentifier {
		return fmt.Errorf("expected enum name")
	}

	name := p.current.Value
	p.nextToken()

	if !p.expectToken(TokenLBrace) {
		return fmt.Errorf("expected '{' after enum name")
	}

	enum := &Enum{Name: name}

	// Parse enum values
	currentValue := int64(0)
	for p.current.Type != TokenRBrace && p.current.Type != TokenEOF {
		// Skip all comments inside enum
		for p.current.Type == TokenComment {
			p.nextToken()
		}

		// Check if we've reached the end after skipping comments
		if p.current.Type == TokenRBrace {
			break
		}

		if p.current.Type != TokenIdentifier {
			return fmt.Errorf("expected enum value name")
		}

		valueName := p.current.Value
		p.nextToken()

		value := currentValue
		if p.current.Type == TokenAssign {
			p.nextToken()
			if p.current.Type == TokenNumber {
				var err error
				value, err = strconv.ParseInt(p.current.Value, 0, 64)
				if err != nil {
					return fmt.Errorf("invalid enum value: %s", p.current.Value)
				}
				p.nextToken()
			}
		}

		enumValue := &EnumValue{
			Name:  valueName,
			Value: value,
		}

		enum.Values = append(enum.Values, enumValue)
		currentValue = value + 1

		if p.current.Type == TokenComma {
			p.nextToken()
		} else if p.current.Type == TokenRBrace {
			break
		}
	}

	if !p.expectToken(TokenRBrace) {
		return fmt.Errorf("expected '}' to close enum")
	}

	// Handle optional enum name after closing brace
	if p.current.Type == TokenIdentifier {
		p.nextToken()
	}

	if p.current.Type == TokenSemicolon {
		p.nextToken()
	}

	p.ast.Enums = append(p.ast.Enums, enum)
	return nil
}

// parseTypedef parses typedef declarations
func (p *Parser) parseTypedef() error {
	if !p.expectToken(TokenTypedef) {
		return fmt.Errorf("expected 'typedef'")
	}

	// Handle "typedef enum" special case
	if p.current.Type == TokenEnum {
		return p.parseTypedefEnum()
	}

	// Parse the type being aliased
	sourceType, err := p.parseType()
	if err != nil {
		return err
	}

	// Parse the new type name
	if p.current.Type != TokenIdentifier {
		return fmt.Errorf("expected typedef name")
	}

	name := p.current.Value
	p.nextToken()

	if !p.expectToken(TokenSemicolon) {
		return fmt.Errorf("expected ';' after typedef")
	}

	typedef := &TypeDef{
		Name: name,
		Type: sourceType,
	}

	p.ast.TypeDefs = append(p.ast.TypeDefs, typedef)
	return nil
}

// parseTypedefEnum parses "typedef enum { ... } NAME;" declarations
func (p *Parser) parseTypedefEnum() error {
	if !p.expectToken(TokenEnum) {
		return fmt.Errorf("expected 'enum'")
	}

	// Optional enum name before brace
	var enumName string
	if p.current.Type == TokenIdentifier {
		enumName = p.current.Value
		p.nextToken()
	}

	if !p.expectToken(TokenLBrace) {
		return fmt.Errorf("expected '{' after enum")
	}

	enum := &Enum{Name: enumName}

	// Parse enum values
	currentValue := int64(0)
	for p.current.Type != TokenRBrace && p.current.Type != TokenEOF {
		// Skip all comments inside enum
		for p.current.Type == TokenComment {
			p.nextToken()
		}

		// Check if we've reached the end after skipping comments
		if p.current.Type == TokenRBrace {
			break
		}

		if p.current.Type != TokenIdentifier {
			return fmt.Errorf("expected enum value name")
		}

		valueName := p.current.Value
		p.nextToken()

		value := currentValue
		if p.current.Type == TokenAssign {
			p.nextToken()
			if p.current.Type == TokenNumber {
				var err error
				value, err = strconv.ParseInt(p.current.Value, 0, 64)
				if err != nil {
					return fmt.Errorf("invalid enum value: %s", p.current.Value)
				}
				p.nextToken()
			}
		}

		enumValue := &EnumValue{
			Name:  valueName,
			Value: value,
		}

		enum.Values = append(enum.Values, enumValue)
		currentValue = value + 1

		if p.current.Type == TokenComma {
			p.nextToken()
		} else if p.current.Type == TokenRBrace {
			break
		}
	}

	if !p.expectToken(TokenRBrace) {
		return fmt.Errorf("expected '}' to close enum")
	}

	// Parse the typedef name after the closing brace
	if p.current.Type == TokenIdentifier {
		finalName := p.current.Value
		if enumName == "" {
			enum.Name = finalName
		}
		p.nextToken()
	}

	// Skip any cpp_quote directives that might appear before the semicolon
	for p.current.Type == TokenIdentifier && p.current.Value == "cpp_quote" {
		if err := p.skipCppQuote(); err != nil {
			return fmt.Errorf("failed to skip cpp_quote: %w", err)
		}
	}

	if !p.expectToken(TokenSemicolon) {
		return fmt.Errorf("expected ';' after typedef enum")
	}

	p.ast.Enums = append(p.ast.Enums, enum)
	return nil
}

// parseAttributes parses attribute lists [attr1, attr2(value)]
func (p *Parser) parseAttributes() ([]*Attribute, error) {
	var attributes []*Attribute

	if !p.expectToken(TokenLBracket) {
		return nil, fmt.Errorf("expected '[' for attributes")
	}

	for p.current.Type != TokenRBracket && p.current.Type != TokenEOF {
		if p.current.Type != TokenIdentifier {
			return nil, fmt.Errorf("expected attribute name")
		}

		attr := &Attribute{Name: p.current.Value}
		p.nextToken()

		// Parse attribute values
		if p.current.Type == TokenLParen {
			p.nextToken()

			for p.current.Type != TokenRParen && p.current.Type != TokenEOF {
				var value string
				if p.current.Type == TokenIdentifier || p.current.Type == TokenString || p.current.Type == TokenNumber {
					value = p.current.Value
					p.nextToken()

					// Handle UUIDs by concatenating dash-separated parts
					if attr.Name == "uuid" {
						for p.current.Type != TokenRParen && p.current.Type != TokenComma && p.current.Type != TokenEOF {
							if p.current.Type == TokenIdentifier || p.current.Type == TokenNumber {
								value += p.current.Value
							} else {
								// Handle dashes and other punctuation in UUIDs
								value += p.current.Value
							}
							p.nextToken()
						}
					}
				} else {
					// Skip unknown tokens
					p.nextToken()
					continue
				}

				attr.Values = append(attr.Values, value)

				if p.current.Type == TokenComma {
					p.nextToken()
				} else if p.current.Type == TokenRParen {
					break
				}
			}

			if !p.expectToken(TokenRParen) {
				return nil, fmt.Errorf("expected ')' after attribute values")
			}
		}

		attributes = append(attributes, attr)

		if p.current.Type == TokenComma {
			p.nextToken()
		} else if p.current.Type == TokenRBracket {
			break
		}
	}

	if !p.expectToken(TokenRBracket) {
		return nil, fmt.Errorf("expected ']' to close attributes")
	}

	return attributes, nil
}

// parseAttributedDeclaration parses declarations that start with attributes
func (p *Parser) parseAttributedDeclaration(attrs []*Attribute) error {
	switch p.current.Type {
	case TokenInterface:
		iface, err := p.parseInterfaceWithAttributes(attrs)
		if err != nil {
			return err
		}
		p.ast.Interfaces = append(p.ast.Interfaces, iface)
	case TokenEnum:
		enum, err := p.parseEnumWithAttributes(attrs)
		if err != nil {
			return err
		}
		p.ast.Enums = append(p.ast.Enums, enum)
	case TokenTypedef:
		// Need to check if this is a typedef enum to handle both parts
		if p.peek.Type == TokenEnum {
			// This is a typedef enum, parse it specially
			p.nextToken() // consume 'typedef'
			enum, err := p.parseTypedefEnumWithAttributes(attrs)
			if err != nil {
				return err
			}

			// Add the enum to the AST
			p.ast.Enums = append(p.ast.Enums, enum)

			// Also create a typedef entry
			typedef := &TypeDef{
				Name:       enum.Name,
				Type:       &Type{Name: enum.Name, Kind: TypeEnum},
				Attributes: attrs,
			}
			p.ast.TypeDefs = append(p.ast.TypeDefs, typedef)
		} else {
			// Regular typedef
			typedef, err := p.parseTypedefWithAttributes(attrs)
			if err != nil {
				return err
			}

			p.ast.TypeDefs = append(p.ast.TypeDefs, typedef)
		}
	default:
		// Skip unknown attributed declarations
		p.nextToken()
	}

	return nil
}

// parseInterfaceWithAttributes parses an interface with pre-parsed attributes
func (p *Parser) parseInterfaceWithAttributes(attrs []*Attribute) (*Interface, error) {
	if !p.expectToken(TokenInterface) {
		return nil, fmt.Errorf("expected 'interface'")
	}

	if p.current.Type != TokenIdentifier {
		return nil, fmt.Errorf("expected interface name")
	}

	name := p.current.Value
	p.nextToken()

	iface := &Interface{
		Name:       name,
		Attributes: attrs,
	}

	// Extract UUID from attributes
	for _, attr := range attrs {
		if attr.Name == "uuid" && len(attr.Values) > 0 {
			iface.UUID = strings.Trim(attr.Values[0], "\"")
		}
	}

	// Parse inheritance
	if p.current.Type == TokenColon {
		p.nextToken()
		if p.current.Type != TokenIdentifier {
			return nil, fmt.Errorf("expected parent interface name")
		}
		iface.Parent = p.current.Value
		p.nextToken()
	}

	if !p.expectToken(TokenLBrace) {
		return nil, fmt.Errorf("expected '{' after interface declaration")
	}

	// Parse interface body
	for p.current.Type != TokenRBrace && p.current.Type != TokenEOF {
		if err := p.parseInterfaceMember(iface); err != nil {
			return nil, err
		}
	}

	if !p.expectToken(TokenRBrace) {
		return nil, fmt.Errorf("expected '}' to close interface")
	}

	// Optional semicolon
	if p.current.Type == TokenSemicolon {
		p.nextToken()
	}

	return iface, nil
}

// parseEnumWithAttributes parses an enum with pre-parsed attributes
func (p *Parser) parseEnumWithAttributes(attrs []*Attribute) (*Enum, error) {
	enum, err := p.parseEnumDeclaration()
	if err != nil {
		return nil, err
	}

	enum.Attributes = attrs
	return enum, nil
}

// parseEnumDeclaration parses just the enum declaration part
func (p *Parser) parseEnumDeclaration() (*Enum, error) {
	if !p.expectToken(TokenEnum) {
		return nil, fmt.Errorf("expected 'enum'")
	}

	if p.current.Type != TokenIdentifier {
		return nil, fmt.Errorf("expected enum name")
	}

	name := p.current.Value
	p.nextToken()

	if !p.expectToken(TokenLBrace) {
		return nil, fmt.Errorf("expected '{' after enum name")
	}

	enum := &Enum{Name: name}

	// Parse enum values
	currentValue := int64(0)
	for p.current.Type != TokenRBrace && p.current.Type != TokenEOF {
		if p.current.Type != TokenIdentifier {
			return nil, fmt.Errorf("expected enum value name")
		}

		valueName := p.current.Value
		p.nextToken()

		value := currentValue
		if p.current.Type == TokenAssign {
			p.nextToken()
			if p.current.Type == TokenNumber {
				var err error
				value, err = strconv.ParseInt(p.current.Value, 0, 64)
				if err != nil {
					return nil, fmt.Errorf("invalid enum value: %s", p.current.Value)
				}
				p.nextToken()
			}
		}

		enumValue := &EnumValue{
			Name:  valueName,
			Value: value,
		}

		enum.Values = append(enum.Values, enumValue)
		currentValue = value + 1

		if p.current.Type == TokenComma {
			p.nextToken()
		} else if p.current.Type == TokenRBrace {
			break
		}
	}

	if !p.expectToken(TokenRBrace) {
		return nil, fmt.Errorf("expected '}' to close enum")
	}

	// Handle optional enum name after closing brace
	if p.current.Type == TokenIdentifier {
		p.nextToken()
	}

	if p.current.Type == TokenSemicolon {
		p.nextToken()
	}

	return enum, nil
}

// parseTypedefWithAttributes parses a typedef with pre-parsed attributes
func (p *Parser) parseTypedefWithAttributes(attrs []*Attribute) (*TypeDef, error) {
	if !p.expectToken(TokenTypedef) {
		return nil, fmt.Errorf("expected 'typedef'")
	}

	// Handle "typedef enum" with attributes
	if p.current.Type == TokenEnum {
		enum, err := p.parseTypedefEnumWithAttributes(attrs)
		if err != nil {
			return nil, err
		}

		// The enum will be added by the caller

		return &TypeDef{
			Name:       enum.Name,
			Type:       &Type{Name: enum.Name, Kind: TypeEnum},
			Attributes: attrs,
		}, nil
	}

	// Parse the type being aliased
	sourceType, err := p.parseType()
	if err != nil {
		return nil, err
	}

	// Parse the new type name
	if p.current.Type != TokenIdentifier {
		return nil, fmt.Errorf("expected typedef name")
	}

	name := p.current.Value
	p.nextToken()

	if !p.expectToken(TokenSemicolon) {
		return nil, fmt.Errorf("expected ';' after typedef")
	}

	return &TypeDef{
		Name:       name,
		Type:       sourceType,
		Attributes: attrs,
	}, nil
}

// parseTypedefEnumWithAttributes parses "typedef enum { ... } NAME;" declarations with attributes
func (p *Parser) parseTypedefEnumWithAttributes(attrs []*Attribute) (*Enum, error) {
	if !p.expectToken(TokenEnum) {
		return nil, fmt.Errorf("expected 'enum'")
	}

	// Optional enum name before brace
	var enumName string
	if p.current.Type == TokenIdentifier {
		enumName = p.current.Value
		p.nextToken()
	}

	if !p.expectToken(TokenLBrace) {
		return nil, fmt.Errorf("expected '{' after enum")
	}

	enum := &Enum{Name: enumName, Attributes: attrs}

	// Parse enum values
	currentValue := int64(0)
	for p.current.Type != TokenRBrace && p.current.Type != TokenEOF {
		// Skip all comments inside enum
		for p.current.Type == TokenComment {
			p.nextToken()
		}

		// Check if we've reached the end after skipping comments
		if p.current.Type == TokenRBrace {
			break
		}

		if p.current.Type != TokenIdentifier {
			return nil, fmt.Errorf("expected enum value name")
		}

		valueName := p.current.Value
		p.nextToken()

		value := currentValue
		if p.current.Type == TokenAssign {
			p.nextToken()
			if p.current.Type == TokenNumber {
				var err error
				value, err = strconv.ParseInt(p.current.Value, 0, 64)
				if err != nil {
					return nil, fmt.Errorf("invalid enum value: %s", p.current.Value)
				}
				p.nextToken()
			}
		}

		enumValue := &EnumValue{
			Name:  valueName,
			Value: value,
		}

		enum.Values = append(enum.Values, enumValue)
		currentValue = value + 1

		if p.current.Type == TokenComma {
			p.nextToken()
		} else if p.current.Type == TokenRBrace {
			break
		}
	}

	if !p.expectToken(TokenRBrace) {
		return nil, fmt.Errorf("expected '}' to close enum")
	}

	// Parse the typedef name after the closing brace
	if p.current.Type == TokenIdentifier {
		finalName := p.current.Value
		if enumName == "" {
			enum.Name = finalName
		}
		p.nextToken()
	}

	// Skip any cpp_quote directives that might appear before the semicolon
	for p.current.Type == TokenIdentifier && p.current.Value == "cpp_quote" {
		if err := p.skipCppQuote(); err != nil {
			return nil, fmt.Errorf("failed to skip cpp_quote: %w", err)
		}
	}

	if !p.expectToken(TokenSemicolon) {
		return nil, fmt.Errorf("expected ';' after typedef enum")
	}

	return enum, nil
}

// parseTypedefDeclaration parses just the typedef declaration part
func (p *Parser) parseTypedefDeclaration() (*TypeDef, error) {
	if !p.expectToken(TokenTypedef) {
		return nil, fmt.Errorf("expected 'typedef'")
	}

	// Handle "typedef enum" special case
	if p.current.Type == TokenEnum {
		enum, err := p.parseEnumDeclaration()
		if err != nil {
			return nil, err
		}

		return &TypeDef{
			Name: enum.Name,
			Type: &Type{Name: enum.Name, Kind: TypeEnum},
		}, nil
	}

	// Parse the type being aliased
	sourceType, err := p.parseType()
	if err != nil {
		return nil, err
	}

	// Parse the new type name
	if p.current.Type != TokenIdentifier {
		return nil, fmt.Errorf("expected typedef name")
	}

	name := p.current.Value
	p.nextToken()

	if !p.expectToken(TokenSemicolon) {
		return nil, fmt.Errorf("expected ';' after typedef")
	}

	return &TypeDef{
		Name: name,
		Type: sourceType,
	}, nil
}

// parseType parses a type specification
func (p *Parser) parseType() (*Type, error) {
	var isConst bool
	var isPointer bool

	// Handle const modifier
	if p.current.Type == TokenConst {
		isConst = true
		p.nextToken()
	}

	// Parse base type
	if p.current.Type != TokenIdentifier {
		return nil, fmt.Errorf("expected type name")
	}

	typeName := p.current.Value
	p.nextToken()

	// Handle pointer modifier
	if p.current.Type == TokenStar {
		isPointer = true
		p.nextToken()
	}

	// Determine type kind
	var kind TypeKind
	if IsBasicType(typeName) {
		kind = TypeBasic
	} else if IsInterfaceType(typeName) {
		kind = TypeInterface
	} else if IsEnumType(typeName) {
		kind = TypeEnum
	} else {
		kind = TypeBasic // Default
	}

	return &Type{
		Name:    typeName,
		Kind:    kind,
		Const:   isConst,
		Pointer: isPointer,
	}, nil
}

// Helper methods

// nextToken advances to the next token
func (p *Parser) nextToken() {
	p.current = p.peek
	p.peek = p.scanner.NextToken()
}

// expectToken checks if current token matches expected type and advances
func (p *Parser) expectToken(expected TokenType) bool {
	if p.current.Type == expected {
		p.nextToken()
		return true
	}
	return false
}

// Token types and scanner implementation - exported for debugging

type TokenType int

const (
	TokenEOF TokenType = iota
	TokenIdentifier
	TokenNumber
	TokenString
	TokenComment

	// Keywords
	TokenImport
	TokenLibrary
	TokenInterface
	TokenEnum
	TokenTypedef
	TokenConst

	// Punctuation
	TokenLParen
	TokenRParen
	TokenLBrace
	TokenRBrace
	TokenLBracket
	TokenRBracket
	TokenSemicolon
	TokenComma
	TokenColon
	TokenStar
	TokenAssign
	TokenDash
)

type Token struct {
	Type  TokenType
	Value string
	Line  int
	Col   int
}

type Scanner struct {
	reader  *bufio.Reader
	line    int
	col     int
	current rune
	eof     bool
}

func NewScanner(reader io.Reader) *Scanner {
	s := &Scanner{
		reader: bufio.NewReader(reader),
		line:   1,
		col:    0,
	}
	s.nextRune()
	return s
}

func (s *Scanner) nextRune() {
	r, _, err := s.reader.ReadRune()
	if err != nil {
		s.eof = true
		s.current = 0
		return
	}

	if s.current == '\n' {
		s.line++
		s.col = 0
	} else {
		s.col++
	}

	s.current = r
}

func (s *Scanner) peekRune() rune {
	r, _, err := s.reader.ReadRune()
	if err != nil {
		return 0
	}
	s.reader.UnreadRune()
	return r
}

func (s *Scanner) NextToken() Token {
	for !s.eof {
		// Skip whitespace
		if unicode.IsSpace(s.current) {
			s.nextRune()
			continue
		}

		// Comments
		if s.current == '/' {
			if s.peekRune() == '/' {
				return s.scanLineComment()
			} else if s.peekRune() == '*' {
				return s.scanBlockComment()
			}
		}

		// Numbers
		if unicode.IsDigit(s.current) || s.current == '-' {
			return s.scanNumber()
		}

		// Strings
		if s.current == '"' {
			return s.scanString()
		}

		// Identifiers and keywords
		if unicode.IsLetter(s.current) || s.current == '_' {
			return s.scanIdentifier()
		}

		// Single character tokens
		switch s.current {
		case ';':
			s.nextRune()
			return Token{TokenSemicolon, ";", s.line, s.col}
		case ',':
			s.nextRune()
			return Token{TokenComma, ",", s.line, s.col}
		case ':':
			s.nextRune()
			return Token{TokenColon, ":", s.line, s.col}
		case '*':
			s.nextRune()
			return Token{TokenStar, "*", s.line, s.col}
		case '=':
			s.nextRune()
			return Token{TokenAssign, "=", s.line, s.col}
		case '-':
			s.nextRune()
			return Token{TokenDash, "-", s.line, s.col}
		case '(':
			s.nextRune()
			return Token{TokenLParen, "(", s.line, s.col}
		case ')':
			s.nextRune()
			return Token{TokenRParen, ")", s.line, s.col}
		case '{':
			s.nextRune()
			return Token{TokenLBrace, "{", s.line, s.col}
		case '}':
			s.nextRune()
			return Token{TokenRBrace, "}", s.line, s.col}
		case '[':
			s.nextRune()
			return Token{TokenLBracket, "[", s.line, s.col}
		case ']':
			s.nextRune()
			return Token{TokenRBracket, "]", s.line, s.col}
		default:
			// Skip unknown characters
			s.nextRune()
		}
	}

	return Token{TokenEOF, "", s.line, s.col}
}

func (s *Scanner) scanLineComment() Token {
	start := s.col
	var value strings.Builder

	// Skip // or ///
	s.nextRune()
	s.nextRune()

	// Handle /// comments by skipping the extra /
	if s.current == '/' {
		s.nextRune()
	}

	for !s.eof && s.current != '\n' {
		value.WriteRune(s.current)
		s.nextRune()
	}

	return Token{TokenComment, strings.TrimSpace(value.String()), s.line, start}
}

func (s *Scanner) scanBlockComment() Token {
	start := s.col
	var value strings.Builder

	// Skip /*
	s.nextRune()
	s.nextRune()

	for !s.eof {
		if s.current == '*' && s.peekRune() == '/' {
			s.nextRune() // Skip *
			s.nextRune() // Skip /
			break
		}
		value.WriteRune(s.current)
		s.nextRune()
	}

	return Token{TokenComment, strings.TrimSpace(value.String()), s.line, start}
}

func (s *Scanner) scanNumber() Token {
	start := s.col
	var value strings.Builder

	// Handle negative numbers
	if s.current == '-' {
		value.WriteRune(s.current)
		s.nextRune()
	}

	// Handle hex numbers
	if s.current == '0' && (s.peekRune() == 'x' || s.peekRune() == 'X') {
		value.WriteRune(s.current)
		s.nextRune()
		value.WriteRune(s.current)
		s.nextRune()

		for !s.eof && (unicode.IsDigit(s.current) ||
			(s.current >= 'a' && s.current <= 'f') ||
			(s.current >= 'A' && s.current <= 'F')) {
			value.WriteRune(s.current)
			s.nextRune()
		}
	} else {
		// Regular decimal number
		for !s.eof && (unicode.IsDigit(s.current) || s.current == '.') {
			value.WriteRune(s.current)
			s.nextRune()
		}
	}

	return Token{TokenNumber, value.String(), s.line, start}
}

func (s *Scanner) scanString() Token {
	start := s.col
	var value strings.Builder

	// Skip opening quote
	s.nextRune()

	for !s.eof && s.current != '"' {
		if s.current == '\\' {
			s.nextRune()
			if !s.eof {
				switch s.current {
				case 'n':
					value.WriteRune('\n')
				case 't':
					value.WriteRune('\t')
				case 'r':
					value.WriteRune('\r')
				case '\\':
					value.WriteRune('\\')
				case '"':
					value.WriteRune('"')
				default:
					value.WriteRune(s.current)
				}
				s.nextRune()
			}
		} else {
			value.WriteRune(s.current)
			s.nextRune()
		}
	}

	// Skip closing quote
	if s.current == '"' {
		s.nextRune()
	}

	return Token{TokenString, "\"" + value.String() + "\"", s.line, start}
}

func (s *Scanner) scanIdentifier() Token {
	start := s.col
	var value strings.Builder

	for !s.eof && (unicode.IsLetter(s.current) || unicode.IsDigit(s.current) || s.current == '_') {
		value.WriteRune(s.current)
		s.nextRune()
	}

	str := value.String()

	// Check for keywords
	var tokenType TokenType
	switch str {
	case "import":
		tokenType = TokenImport
	case "library":
		tokenType = TokenLibrary
	case "interface":
		tokenType = TokenInterface
	case "enum":
		tokenType = TokenEnum
	case "typedef":
		tokenType = TokenTypedef
	case "const":
		tokenType = TokenConst
	default:
		tokenType = TokenIdentifier
	}

	return Token{tokenType, str, s.line, start}
}
