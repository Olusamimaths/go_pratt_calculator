package main

import (
	"math"
	"testing"
)

func TestExponentialParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"2^3",
			"(2 ^ 3)",
		},
		{
			"3.0 ^ 2.5",
			"(3.0 ^ 2.5)",
		},
		{
			"2.5 ^ 3.0 * 4.2",
			"((2.5 ^ 3.0) * 4.2)",
		},
		{
			"2 ^ 3 * 4 + 1",
			"(((2 ^ 3) * 4) + 1)",
		},
		{
			"2 ^ 2.5 / 4.0",
			"((2 ^ 2.5) / 4.0)",
		},
		{
			"2 ^ 3 ^ 2",
			"(2 ^ (3 ^ 2))",
		},
	}

	for _, tt := range tests {
		lexer := NewLexer(tt.input)
		parser := NewParser(lexer)
		calculator := parser.Parse()
		checkParserErrors(t, parser)

		actual := calculator.String()
		if actual != tt.expected {
			t.Errorf("input=%q, expected=%q, got=%q", tt.input, tt.expected, actual)
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-2 * 3",
			"((-2) * 3)",
		},
		{
			"1 + 2 + 3",
			"((1 + 2) + 3)",
		},
		{
			"1 - 2 + 3",
			"((1 - 2) + 3)",
		},
		{
			"1 * 2 * 3",
			"((1 * 2) * 3)",
		},
		{
			"1 + 2 / 3",
			"(1 + (2 / 3))",
		},
		{
			"1 + 2 * 3 + 4 / 5 - 6",
			"(((1 + (2 * 3)) + (4 / 5)) - 6)",
		},
		{
			"-1 * -4",
			"((-1) * (-4))",
		},
		{
			"2",
			"2",
		},
		{
			"2 + 3 * 4",
			"(2 + (3 * 4))",
		},
		{
			"2 - 3 / 4",
			"(2 - (3 / 4))",
		},
		{
			"2 * 3 / 4",
			"((2 * 3) / 4)",
		},
		{
			"0 * 5 + 3",
			"((0 * 5) + 3)",
		},
		{
			"-2 * 3.5",
			"((-2) * 3.5)",
		},
		{
			"1.0 + 2.5 + 3.3",
			"((1.0 + 2.5) + 3.3)",
		},
		{
			"1.5 - 2.2 + 3.1",
			"((1.5 - 2.2) + 3.1)",
		},
		{
			"1.2 * 2.0 * 3.5",
			"((1.2 * 2.0) * 3.5)",
		},
		{
			"1.0 + 2.5 / 3.0",
			"(1.0 + (2.5 / 3.0))",
		},
		{
			"1.0 + 2.5 * 3.0 + 4.7 / 5.2 - 6.0",
			"(((1.0 + (2.5 * 3.0)) + (4.7 / 5.2)) - 6.0)",
		},
		{
			"-1.5 * -4.2",
			"((-1.5) * (-4.2))",
		},
		{
			"2.5",
			"2.5",
		},
		{
			"2.5 + 3.0 * 4.2",
			"(2.5 + (3.0 * 4.2))",
		},
		{
			"2.8 - 3.1 / 4.0",
			"(2.8 - (3.1 / 4.0))",
		},
		{
			"2.5 * 3.2 / 4.5",
			"((2.5 * 3.2) / 4.5)",
		},
		{
			"0.0 * 5.3 + 3.9",
			"((0.0 * 5.3) + 3.9)",
		},
	}

	for _, tt := range tests {
		lexer := NewLexer(tt.input)
		parser := NewParser(lexer)
		calculator := parser.Parse()
		checkParserErrors(t, parser)

		actual := calculator.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

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
		checkParserErrors(t, parser)

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
	defaultEpsilon := 1e-9

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

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}
