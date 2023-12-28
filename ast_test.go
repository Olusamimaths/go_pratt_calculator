package main

import (
	"testing"
)

func TestAST(t *testing.T) {
	expression := &InfixExpression{
		Token: Token{Type: PLUS, Literal: "+"},
		Left: &NumberLiteral{
			Token: Token{Type: NUMBER, Literal: "1"},
			Value: 1,
		},
		Right: &InfixExpression{
			Token: Token{Type: MULT, Literal: "*"},
			Left: &NumberLiteral{
				Token: Token{Type: NUMBER, Literal: "2"},
				Value: 2,
			},
			Right: &NumberLiteral{
				Token: Token{Type: NUMBER, Literal: "3"},
				Value: 3,
			},
		},
	}

	if expression.String() != "1 + (2 * 3)" {
		t.Errorf("wrong expression. got=%q", expression.String())
	}
}
