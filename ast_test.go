package main

import (
	"testing"
)

func TestAST(t *testing.T) {
	expression := &InfixExpression{
		Token: {Type: PLUS, TokenLiteral: "+"},
		Left: &NumberLiteral{
			Token: {Type: NUMBER, TokenLiteral: "1"},
			Value: 1,
		},
		Right: &InfixExpression{
			Token: {Type: MULT, TokenLiteral: "*"},
			Left: &NumberLiteral{
				Token: {Type: NUMBER, TokenLiteral: "2"},
				Value: 2,
			},
			Right: &NumberLiteral{
				Token: {Type: NUMBER, TokenLiteral: "3"},
				Value: 3,
			},
		},
	}

	if expression.String() != "1 + 2 * 3" {
		t.Errorf("wrong expression. got=%q", expression.String())
	}
}
