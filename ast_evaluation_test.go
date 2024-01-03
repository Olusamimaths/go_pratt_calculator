package main

import (
	"testing"
)

func TestASTEvaluation(t *testing.T) {
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

	if expression.Evaluate() != 7 {
		t.Errorf("wrong expression. got=%f", expression.Evaluate())
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

	if expression1.Evaluate() != 3 {
		t.Errorf("Test case 1 failed. got=%f", expression1.Evaluate())
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

	if expression2.Evaluate() != 24 {
		t.Errorf("Test case 2 failed. got=%f", expression2.Evaluate())
	}

	// Test case 3: Edge case - Single number
	expression3 := &NumberLiteral{
		Token: Token{Type: NUMBER, Literal: "42"},
		Value: 42,
	}

	if expression3.Evaluate() != 42 {
		t.Errorf("Test case 3 failed. got=%f", expression3.Evaluate())
	}
}

func TestASTEvaluationWithExponent(t *testing.T) {
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

	if expression1.Evaluate() != 9 {
		t.Errorf("Test case 1 failed. got=%f", expression1.Evaluate())
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

	if expression2.Evaluate() != 162 {
		t.Errorf("Test case 2 failed. got=%f", expression2.Evaluate())
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

	if expression3.Evaluate() != 15625 {
		t.Errorf("Test case 3 failed. got=%f", expression3.Evaluate())
	}
}
