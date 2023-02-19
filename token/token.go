package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"
	NULL   = "NULL"
	MACRO  = "MACRO"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="
	BANG   = "!"

	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	LET      = "LET"
	FUNCTION = "FUNCTION"
	RETURN   = "RETURN"

	IF    = "IF"
	ELSE  = "ELSE"
	TRUE  = "TRUE"
	FALSE = "FALSE"
)

var keywords = map[string]TokenType{
	"let":    LET,
	"fn":     FUNCTION,
	"return": RETURN,

	"if":    IF,
	"else":  ELSE,
	"true":  TRUE,
	"false": FALSE,
	"null":  NULL,
	"macro": MACRO,
}

func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}

	return IDENT
}
