package main

import (
	"math"
	"testing"
)

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		number   float64
	}{
		{"-10", "-", 10},
		{"-6.57", "-", 6.57},
		{"-0.000000001", "-", 0.000000001},
	}

	for _, tt := range prefixTests {
		lexer := NewLexer(tt.input)
		parser := NewParser(lexer)
		calculator := parser.Parse()

		exp, ok := calculator.Expression.(*PrefixExpression)
		if !ok {
			t.Fatalf("expression is not PrefixExpression. got=%T", calculator)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("expression.Operator is not '%s', got '%s'", tt.operator, exp.Operator)
		}

		if !testNumberLiteral(t, exp.Right, tt.number) {
			return
		}
	}
}

func testNumberLiteral(t *testing.T, right Expression, expectedValue float64) bool {
	number, ok := right.(*NumberLiteral)
	if !ok {
		t.Errorf("right is not NumberLiteral. got=%T", right)
		return false
	}

	if !nearlyEqual(number.Value, expectedValue) {
		t.Errorf("number is not correct. expected=%f, got=%f", expectedValue, number.Value)
		return false
	}

	return true
}

func nearlyEqual(a, b float64, epsilon ...float64) bool {
	defaultEpsilon := 1e-9 // Set your default epsilon value here

	// Use the provided epsilon or the default value if not provided
	var eps float64
	if len(epsilon) > 0 {
		eps = epsilon[0]
	} else {
		eps = defaultEpsilon
	}

	// already equal?
	if a == b {
		return true
	}

	diff := math.Abs(a - b)
	if a == 0.0 || b == 0.0 || diff < math.SmallestNonzeroFloat64 {
		return diff < eps*math.SmallestNonzeroFloat64
	}

	return diff/(math.Abs(a)+math.Abs(b)) < eps
}
