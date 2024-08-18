package compiler

import (
	"fmt"

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
		c.emit(code.OpPop)

	case *ast.PrefixExpression:
		err := c.Complie(node.Right)
		if err != nil {
			return err
		}

		switch node.Operator {
		case "!":
			c.emit(code.OpBang)
		case "-":
			c.emit(code.OpMinus)
		default:
			return fmt.Errorf("unknown operator %s", node.Operator)
		}

	case *ast.InfixExpression:
		if node.Operator == "<" {
			err := c.Complie(node.Right)
			if err != nil {
				return err
			}

			err = c.Complie(node.Left)
			if err != nil {
				return err
			}

			c.emit(code.OpGreaterThan)
			return nil
		}

		err := c.Complie(node.Left)
		if err != nil {
			return err
		}

		err = c.Complie(node.Right)
		if err != nil {
			return err
		}

		switch node.Operator {
		case "+":
			c.emit(code.OpAdd)
		case "-":
			c.emit(code.OpSub)
		case "*":
			c.emit(code.OpMul)
		case "/":
			c.emit(code.OpDiv)
		case ">":
			c.emit(code.OpGreaterThan)
		case "==":
			c.emit(code.OpEquals)
		case "!=":
			c.emit(code.OpNotEquals)
		default:
			return fmt.Errorf("unknown operator %s", node.Operator)
		}

	case *ast.IntegerLiteral:
		integer := &object.Integer{Value: node.Value}
		c.emit(code.OpConstant, c.addConstant(integer))

	case *ast.Boolean:
		if node.Value {
			c.emit(code.OpTrue)
		} else {
			c.emit(code.OpFalse)
		}
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
