package ast

import "github.com/mpaliwoda/monkeylang/token"

type Null struct {
	Token token.Token
}

func (n *Null) expressionNode()      {}
func (n *Null) TokenLiteral() string { return n.Token.Literal }
func (n *Null) String() string       { return n.Token.Literal }
