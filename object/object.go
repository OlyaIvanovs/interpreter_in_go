package object
 
import (
	"fmt" 
)
 
type ObjectType string

const (
	INTEGER_OBG = "INTEGER"
	BOOLEAN_OBG = "BOOLEAN"
	NULL_OBG    = "NULL"
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
	return INTEGER_OBG
}

// Boolean
type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value) 
}

func (b *Boolean)  Type() ObjectType {
	return BOOLEAN_OBG
}

// Null
type Null struct {}

func (n *Null) Inspect() string {
	return "null"
}

func (n *Null) Type() ObjectType {
	return NULL_OBG
}