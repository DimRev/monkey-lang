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
	switch node := node.(type) {
	case *ast.Program:
		for _, s := range node.Statements {
			err := c.Complie(s)
			if err != nil {
				return err
			}
		}

	case *ast.ExpressionStatement:
		err := c.Complie(node.Expression)
		if err != nil {
			return err
		}

	case *ast.InfixExpression:
		err := c.Complie(node.Left)
		if err != nil {
			return err
		}

		err = c.Complie(node.Right)
		if err != nil {
			return err
		}

	case *ast.IntegerLiteral:
		integer := &object.Integer{Value: node.Value}
		c.emit(code.OpConstant, c.addConstant(integer))
	}

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

func (c *Complier) addConstant(obj object.Object) int {
	c.constants = append(c.constants, obj)
	return len(c.constants) - 1
}

func (c *Complier) addInstruction(ins []byte) int {
	posNewInstruction := len(c.instructions)
	c.instructions = append(c.instructions, ins...)
	return posNewInstruction
}

func (c *Complier) emit(op code.Opcode, operands ...int) int {
	ins := code.Make(op, operands...)
	pos := c.addInstruction(ins)
	return pos
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}
