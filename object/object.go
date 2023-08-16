package object
 
import (
	"fmt" 
	"strings"
	
	"github.com/OlyaIvanovs/interpreter_in_go/ast"
)
 
type ObjectType string

const (
	INTEGER_OBJ		= "INTEGER"
	BOOLEAN_OBJ 	= "BOOLEAN"
	NULL_OBJ    	= "NULL"
	RETURN_OBJ    	= "RETURN_VALUE"
	ERROR_OBJ 		= "ERROR"
	FUNCTION_OBJ	= "FUNCTION"
	STRING_OBJ	    = "STRING"
	BUILTIN_OBJ	    = "BUILTIN"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer
type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value) 
}
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// String
type String struct {
	Value string
}

func (i *String) Inspect() string {
	return i.Value 
}
func (i *String) Type() ObjectType {
	return STRING_OBJ
}

// Boolean
type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value) 
}
func (b *Boolean)  Type() ObjectType {
	return BOOLEAN_OBJ
}

// Null
type Null struct {}

func (n *Null) Inspect() string {
	return "null"
}
func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

// Return
type ReturnValue struct {
	Value Object 
}

func (rv *ReturnValue) Type() ObjectType {
	return RETURN_OBJ
}
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// Error
type Error struct {
	Message string
}

func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

// Function
type Function struct {
	Parameters 	[]*ast.Identifier
	Body		*ast.BlockStatement
	Env 		*Environment
}

func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}
func (f *Function) Inspect() string {
	var out strings.Builder
	
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	
	out.WriteString("fn(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n")
	
	return out.String()
}

// Built-in  function
type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string { return "builtin function" }