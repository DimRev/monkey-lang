package compiler

import (
	"github.com/DimRev/monkey-lang/ast"
	"github.com/DimRev/monkey-lang/code"
	"github.com/DimRev/monkey-lang/object"
)

type Compiler struct {
	instructions code.Instructions
	constants    []object.Object
}

func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants:    []object.Object{},
	}
}

func (c *Compiler) Compile(node ast.Node) error {
	return nil
}

func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		instructions: c.instructions,
		constants:    c.constants,
	}
}

type Bytecode struct {
	instructions code.Instructions
	constants    []object.Object
}
