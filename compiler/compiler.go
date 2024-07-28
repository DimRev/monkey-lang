package compiler

import (
	"github.com/DimRev/monkey-lang/ast"
	"github.com/DimRev/monkey-lang/code"
	"github.com/DimRev/monkey-lang/lexer"
	"github.com/DimRev/monkey-lang/object"
	"github.com/DimRev/monkey-lang/parser"
)

type Complier struct {
	instructions code.Instructions
	constants    []object.Object
}

func New() *Complier {
	return &Complier{
		instructions: code.Instructions{},
		constants:    []object.Object{},
	}
}

func (c *Complier) Complie(node ast.Node) error {
	return nil
}

type Bytecode struct {
	Instructions code.Instructions
	Constants    []object.Object
}

func (c *Complier) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}
