package main

import (
	"testing"
)

func TestAST(t *testing.T) {
	// Test case 0: Addition of multiplied numbers
	expression := &InfixExpression{
		Token:    Token{Type: PLUS, Literal: "+"},
		Operator: "+",
		Left: &NumberLiteral{
			Token: Token{Type: NUMBER, Literal: "1"},
			Value: 1,
		},
		Right: &InfixExpression{
			Token:    Token{Type: MULT, Literal: "*"},
			Operator: "*",
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

	if expression.String() != "(1 + (2 * 3))" {
		t.Errorf("wrong expression. got=%q", expression.String())
	}
	// Test case 1: Addition of two numbers
	expression1 := &InfixExpression{
		Token:    Token{Type: PLUS, Literal: "+"},
		Operator: "+",
		Left: &NumberLiteral{
			Token: Token{Type: NUMBER, Literal: "1"},
			Value: 1,
		},
		Right: &NumberLiteral{
			Token: Token{Type: NUMBER, Literal: "2"},
			Value: 2,
		},
	}

	if expression1.String() != "(1 + 2)" {
		t.Errorf("Test case 1 failed. got=%q", expression1.String())
	}

	// Test case 2: Multiplication of three numbers
	expression2 := &InfixExpression{
		Token:    Token{Type: MULT, Literal: "*"},
		Operator: "*",
		Left: &NumberLiteral{
			Token: Token{Type: NUMBER, Literal: "2"},
			Value: 2,
		},
		Right: &InfixExpression{
			Token:    Token{Type: MULT, Literal: "*"},
			Operator: "*",
			Left: &NumberLiteral{
				Token: Token{Type: NUMBER, Literal: "3"},
				Value: 3,
			},
			Right: &NumberLiteral{
				Token: Token{Type: NUMBER, Literal: "4"},
				Value: 4,
			},
		},
	}

	if expression2.String() != "(2 * (3 * 4))" {
		t.Errorf("Test case 2 failed. got=%q", expression2.String())
	}

	// Test case 3: Edge case - Single number
	expression3 := &NumberLiteral{
		Token: Token{Type: NUMBER, Literal: "42"},
		Value: 42,
	}

	if expression3.String() != "42" {
		t.Errorf("Test case 3 failed. got=%q", expression3.String())
	}
}

func TestASTWithExponent(t *testing.T) {
	// Test case 1: Addition and Exponential
	expression1 := &InfixExpression{
		Token:    Token{Type: PLUS, Literal: "+"},
		Operator: "+",
		Left: &NumberLiteral{
			Token: Token{Type: NUMBER, Literal: "1"},
			Value: 1,
		},
		Right: &InfixExpression{
			Token:    Token{Type: EXPONENTIAL, Literal: "^"},
			Operator: "^",
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

	if expression1.String() != "(1 + (2 ^ 3))" {
		t.Errorf("Test case 1 failed. got=%q", expression1.String())
	}

	// Test case 2: Multiplication and Exponential
	expression2 := &InfixExpression{
		Token:    Token{Type: MULT, Literal: "*"},
		Operator: "*",
		Left: &NumberLiteral{
			Token: Token{Type: NUMBER, Literal: "2"},
			Value: 2,
		},
		Right: &InfixExpression{
			Token:    Token{Type: EXPONENTIAL, Literal: "^"},
			Operator: "^",
			Left: &NumberLiteral{
				Token: Token{Type: NUMBER, Literal: "3"},
				Value: 3,
			},
			Right: &NumberLiteral{
				Token: Token{Type: NUMBER, Literal: "4"},
				Value: 4,
			},
		},
	}

	if expression2.String() != "(2 * (3 ^ 4))" {
		t.Errorf("Test case 2 failed. got=%q", expression2.String())
	}

	// Test case 3: Exponential with single number
	expression3 := &InfixExpression{
		Token:    Token{Type: EXPONENTIAL, Literal: "^"},
		Operator: "^",
		Left: &NumberLiteral{
			Token: Token{Type: NUMBER, Literal: "5"},
			Value: 5,
		},
		Right: &NumberLiteral{
			Token: Token{Type: NUMBER, Literal: "6"},
			Value: 6,
		},
	}

	if expression3.String() != "(5 ^ 6)" {
		t.Errorf("Test case 3 failed. got=%q", expression3.String())
	}

	// Test case 4: Exponential with no operator (should default to addition)
	expression4 := &InfixExpression{
		Token:    Token{Type: PLUS, Literal: "+"},
		Operator: "+",
		Left: &NumberLiteral{
			Token: Token{Type: NUMBER, Literal: "7"},
			Value: 7,
		},
		Right: &InfixExpression{
			Token:    Token{Type: EXPONENTIAL, Literal: "^"},
			Operator: "^",
			Left: &NumberLiteral{
				Token: Token{Type: NUMBER, Literal: "8"},
				Value: 8,
			},
			Right: &NumberLiteral{
				Token: Token{Type: NUMBER, Literal: "9"},
				Value: 9,
			},
		},
	}

	if expression4.String() != "(7 + (8 ^ 9))" {
		t.Errorf("Test case 4 failed. got=%q", expression4.String())
	}
}
