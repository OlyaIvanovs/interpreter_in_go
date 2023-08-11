package evaluator 


import (
	"github.com/OlyaIvanovs/interpreter_in_go/ast"
	"github.com/OlyaIvanovs/interpreter_in_go/object"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL  = &object.Null{}
)


func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return  nativeBoolToBooleanObject(node.Value)
		
	// Statements
	case *ast.Program:
		return evalStatements(node.Statements)
		
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	}
	
	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	
	for _, statement := range stmts {
		result = Eval(statement)
	}
	
	return result
}

func nativeBoolToBooleanObject(val bool) *object.Boolean{
	if val {
		return TRUE
	}
	return FALSE
}