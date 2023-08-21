package object
 
import (
	"fmt" 
	"hash/fnv"
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
	ARRAY_OBJ	    = "ARRAY"
	HASH_OBJ	    = "HASH"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Hashable interface {
	HashKey() HashKey
}
type HashKey struct {
	Type ObjectType
	Value uint64
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
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: INTEGER_OBJ, Value: uint64(i.Value)}
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
func (i *String) HashKey() HashKey {
	h := fnv.New64()
	h.Write([]byte(i.Value))
	return HashKey{Type: STRING_OBJ, Value: h.Sum64()}
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
func (b *Boolean) HashKey() HashKey {
	if b.Value {
		return HashKey{Type: BOOLEAN_OBJ, Value: 1}
	}
	return HashKey{Type: BOOLEAN_OBJ, Value: 0}
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

// Array
type Array struct {
	Elements []Object
}
func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out strings.Builder
	
	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}
	
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	
	return out.String()
}

// Hash
type Hash struct {
	Pairs map[HashKey]HashPair
}
type HashPair struct {
	Key		Object	
	Value   Object
}
func (h *Hash) Type() ObjectType {return HASH_OBJ}
func (h *Hash) Inspect() string {
	var out strings.Builder
	
	pairs := []string{}
	
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}
	
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	
	return out.String()
}


