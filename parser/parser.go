package parser

import (
	"fmt"

	"github.com/mpaliwoda/interpreter-book/ast"
	"github.com/mpaliwoda/interpreter-book/lexer"
	"github.com/mpaliwoda/interpreter-book/token"
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	p.NextToken()
	p.NextToken()

	return p
}

func (p *Parser) NextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for !p.currentTokenIs(token.EOF) {
		statement := p.ParseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		p.NextToken()
	}

	return program
}

func (p *Parser) ParseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(expectedTokenType token.TokenType) {
	msg := fmt.Sprintf(
		"expected next token to be %s, got %s instead",
		expectedTokenType,
		p.peekToken.Type,
	)

	p.errors = append(p.errors, msg)
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.currentToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// skipping expressions for nau
	for !p.currentTokenIs(token.SEMICOLON) {
		p.NextToken()
	}

	return statement
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{Token: p.currentToken}
	p.NextToken()

	for !p.currentTokenIs(token.SEMICOLON) {
		p.NextToken()
	}

	return statement
}

func (p *Parser) currentTokenIs(expectedTokenType token.TokenType) bool {
	return p.currentToken.Type == expectedTokenType
}

func (p *Parser) peekTokenIs(expectedTokenType token.TokenType) bool {
	return p.peekToken.Type == expectedTokenType
}

func (p *Parser) expectPeek(expectedTokenType token.TokenType) bool {
	if p.peekTokenIs(expectedTokenType) {
		p.NextToken()
		return true
	}

	p.peekError(expectedTokenType)
	return false
}
