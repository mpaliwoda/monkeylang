package parser

import (
	"fmt"
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

func TestIndentifierExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf(
			"program.Statements expected to have 1 statement, got %d",
			len(program.Statements),
		)
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf(
			"statement: not *ast.ExpressionStatement, got %T",
			program.Statements[0],
		)
	}

	identifier, ok := statement.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression not *ast.Identifier. got=%T", statement.Expression)
	}

	if identifier.Value != "foobar" {
		t.Fatalf("identifier.Value not foobar. got=%T", statement.Expression)
	}

	if identifier.TokenLiteral() != "foobar" {
		t.Errorf("identifier.TokenLiteral not %s. got=%s", "foobar", identifier.TokenLiteral())
	}
}

func TestIntegerExpression(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf(
			"program.Statements expected to have 1 statement, got %d",
			len(program.Statements),
		)
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf(
			"statement: not *ast.ExpressionStatement, got %T",
			program.Statements[0],
		)
	}

	literal, ok := statement.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("literal not *ast.IntegerLiteral. got=%T", statement.Expression)
	}

	if literal.Value != 5 {
		t.Fatalf("literal.Value not 5. got=%T", statement.Expression)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not 5. got=%s", literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
	}

	for _, testCase := range prefixTests {
		l := lexer.New(testCase.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf(
				"program.Statements expected to have 1 statement, got %d",
				len(program.Statements),
			)
		}

		statement, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Errorf(
				"statement: not *ast.ExpressionStatement, got %T",
				program.Statements[0],
			)
		}

		expression, ok := statement.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("expression not *ast.PrefixExpression. got=%T", statement.Expression)
		}

		if expression.Operator != testCase.operator {
			t.Fatalf("expression.Operator not %s. got=%T", testCase.operator, statement.Expression)
		}

		if !testLiteralExpression(t, expression.Right, testCase.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, literal ast.Expression, expectedValue int64) bool {
	integer, ok := literal.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("literal not and ast.IntegerLiteral, got=%T", literal)
		return false
	}

	if integer.Value != expectedValue {
		t.Errorf(
			"integer.Value expected to be %d, got %d",
			expectedValue,
			integer.Value,
		)
		return false
	}

	if integer.TokenLiteral() != fmt.Sprintf("%d", expectedValue) {
		t.Errorf(
			"integer.TokenLiteral() expected to be %d, got %s",
			expectedValue,
			integer.TokenLiteral(),
		)
		return false
	}

	return true
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, testCase := range infixTests {
		l := lexer.New(testCase.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf(
				"program.Statements expected to have 1 statement, got %d",
				len(program.Statements),
			)
		}

		statement, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Errorf(
				"statement: not *ast.ExpressionStatement, got %T",
				program.Statements[0],
			)
		}

		if !testInfixExpression(
			t,
			statement.Expression,
			testCase.leftValue,
			testCase.operator,
			testCase.rightValue,
		) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

func TestBooleanExpression(t *testing.T) {
	tests := []struct {
		input           string
		expectedBoolean bool
	}{
		{"true;", true},
		{"false;", false},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements. got=%d",
				len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		boolean, ok := stmt.Expression.(*ast.Boolean)
		if !ok {
			t.Fatalf("exp not *ast.Boolean. got=%T", stmt.Expression)
		}
		if boolean.Value != tt.expectedBoolean {
			t.Errorf("boolean.Value not %t. got=%t", tt.expectedBoolean,
				boolean.Value)
		}
	}
}

func testIdentifier(t *testing.T, expression ast.Expression, expectedValue string) bool {
	identifier, ok := expression.(*ast.Identifier)
	if !ok {
		t.Errorf("expression not *ast.Identifier, got %T", expression)
		return false
	}

	if identifier.Value != expectedValue {
		t.Errorf(
			"identifier.Value not %s, got %s",
			expectedValue,
			identifier.Value,
		)
		return false
	}

	if identifier.TokenLiteral() != expectedValue {
		t.Errorf(
			"identifier.TokenLiteral() not %s, got %s",
			expectedValue,
			identifier.TokenLiteral(),
		)
	}

	return true
}

func testLiteralExpression(
	t *testing.T,
	expression ast.Expression,
	expected interface{},
) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, expression, int64(v))
	case int64:
		return testIntegerLiteral(t, expression, v)
	case string:
		return testIdentifier(t, expression, v)
	case bool:
		return testBooleanLiteral(t, expression, v)
	}

	t.Errorf("type of expression not handled. got=%T", expression)
	return false
}

func testInfixExpression(
	t *testing.T,
	exp ast.Expression,
	left interface{},
	operator string,
	right interface{},
) bool {

	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.InfixExpression. got=%T(%s)", exp, exp)
		return false
	}

	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
		return false
	}

	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}

	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	bo, ok := exp.(*ast.Boolean)
	if !ok {
		t.Errorf("exp not *ast.Boolean. got=%T", exp)
		return false
	}

	if bo.Value != value {
		t.Errorf("bo.Value not %t. got=%t", value, bo.Value)
		return false
	}

	if bo.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("bo.TokenLiteral not %t. got=%s",
			value, bo.TokenLiteral())
		return false
	}

	return true
}
