package parser

import (
	"github.com/DimRev/monkey-lang/ast"
	"github.com/DimRev/monkey-lang/lexer"
	"github.com/DimRev/monkey-lang/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	// read two tokens so both curToken and peekToken are set

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
