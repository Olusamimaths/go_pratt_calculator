package main

type Node interface {
	TokenLiteral() string
}

type PrefixExpression struct {
	Token    Token
	Operator string
	Right    Node
}

func (p *PrefixExpression) TokenLiteral() string {
	return p.Token.TokenLiteral
}

type InfixExpression struct {
	Token    Token
	Left     Node
	Operator string
	Right    Node
}

func (i *InfixExpression) TokenLiteral() string {
	return i.Token.TokenLiteral
}
