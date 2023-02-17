package lexer

import (
	"testing"

	"monkeylang/token"
)

func TestNextTokenBasic(t *testing.T) {
	input := "=+(){};,-/*<>"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.COMMA, ","},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.LT, "<"},
		{token.GT, ">"},
	}

	l := New(input)

	for i, testCase := range tests {
		tok := l.NextToken()

		if tok.Type != testCase.expectedType {
			t.Fatalf(
				"tests[%d] - tokentype wrong -- expected '%q', got '%q'",
				i,
				testCase.expectedType,
				tok.Type,
			)
		}

		if tok.Literal != testCase.expectedLiteral {
			t.Fatalf(
				"tests[%d] - literal wrong -- expected '%q', got '%q'",
				i,
				testCase.expectedLiteral,
				tok.Literal,
			)
		}
	}
}

func TestNextTokenExtended(t *testing.T) {
	input := `
let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y
}

let result = add(five, ten);

if (5 < 10) {
	return true;
} else {
	return false;
}

if (result == 15) {} 

if (result != 10) {}

"foobar"
"foo bar"
[1, 2];
{"foo": "bar"}
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.RBRACE, "}"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "result"},
		{token.EQ, "=="},
		{token.INT, "15"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "result"},
		{token.NOT_EQ, "!="},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, testCase := range tests {
		tok := l.NextToken()

		if tok.Type != testCase.expectedType {
			t.Fatalf(
				"tests[%d] - tokentype wrong -- expected '%q', got '%q'",
				i,
				testCase.expectedType,
				tok.Type,
			)
		}

		if tok.Literal != testCase.expectedLiteral {
			t.Fatalf(
				"tests[%d] - literal wrong -- expected '%q', got '%q'",
				i,
				testCase.expectedLiteral,
				tok.Literal,
			)
		}
	}
}
