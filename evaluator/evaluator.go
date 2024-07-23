package evaluator

import (
	"github.com/DimRev/monkey-lang/ast"
	"github.com/DimRev/monkey-lang/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// STATEMENTS
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	// EXPRESSIONS
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
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
