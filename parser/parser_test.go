package parser

import (
	"testing"

	"github.com/mpaliwoda/interpreter-book/ast"
	"github.com/mpaliwoda/interpreter-book/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;	
let foobar = 838383;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf(
			"program.Statements expected to have 3 statements, got %d",
			len(program.Statements),
		)
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, testCase := range tests {
		statment := program.Statements[i]

		if !testLetStatement(t, statment, testCase.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser had %d errors", len(errors))
	for i, msg := range errors {
		t.Errorf("parser error %d: %s", i+1, msg)
	}

	t.FailNow()
}

func testLetStatement(t *testing.T, statement ast.Statement, expectedIdentifier string) bool {
	if statement.TokenLiteral() != "let" {
		t.Errorf("statement.TokenLiteral() not 'let', got %q", statement.TokenLiteral())
		return false
	}

	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		t.Errorf("statement not *ast.LetStatement, got %T", statement)
		return false
	}

	if letStatement.Name.Value != expectedIdentifier {
		t.Errorf(
			"letStatement.Name.Value not '%s', got %s",
			expectedIdentifier,
			letStatement.Name.Value,
		)
		return false
	}

	if letStatement.Name.TokenLiteral() != expectedIdentifier {
		t.Errorf(
			"letStatement.Name.TokenLiteral() not '%s', got %s",
			expectedIdentifier,
			letStatement.Name.TokenLiteral(),
		)
		return false
	}

	return true
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf(
			"program.Statements expected to have 3 statements, got %d",
			len(program.Statements),
		)
	}

	for i, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement)

		if !ok {
			t.Errorf(
				"returnStatement %d: not *ast.ReturnStatement, got %T",
				i,
				statement,
			)
			continue
		}

		if returnStatement.TokenLiteral() != "return" {
			t.Errorf(
				"returnStatement.TokenLiteral() %d: not return, got %q",
				i,
				returnStatement.TokenLiteral(),
			)

		}
	}
}
