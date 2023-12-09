package object

import "fmt"

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (integer *Integer) Inspect() string  { return fmt.Sprintf("%d", integer.Value) }
func (integer *Integer) Type() ObjectType { return INTEGER_OBJ }

type Boolean struct {
	Value bool
}

func (boolean *Boolean) Inspect() string  { return fmt.Sprintf("%t", boolean.Value) }
func (boolean *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
