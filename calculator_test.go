package main

import (
	"testing"
)

func TestParenthesisEvaluation(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{
			"(2 + 3)",
			5,
		},
		{
			"(3.0 * 2.5)",
			7.5,
		},
		{
			"(2 + (1 + 2))",
			5,
		},
		{
			"(2.5 ^ 3.0) * 4.2",
			65.625,
		},
		{
			"((1 + 2) * 3) ^ 2",
			81,
		},
		{
			"(2 ^ 3) * (4 + 1)",
			40,
		},
		{
			"(2 ^ (2 ^ 3))",
			256,
		},
		{
			"((2.0 + 3.5) * 4.1)",
			22.55,
		},
		{
			"(2 ^ 2.5) / 4.0",
			1.4142135623731,
		},
		{
			"(2 ^ (3 ^ 2)) * (4 + 1)",
			2560,
		},
	}

	for _, tt := range tests {
		lexer := NewLexer(tt.input)
		parser := NewParser(lexer)
		calculator := parser.Parse()

		actual := calculator.Evaluate()
		if !NearlyEqual(actual, tt.expected) {
			t.Errorf("input=%q, expected=%f, got=%f", tt.input, tt.expected, actual)
		}
	}
}

func TestExponentialEvaluation(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{
			"2^3",
			8,
		},
		{
			"3.0 ^ 2.5",
			15.58845726812,
		},
		{
			"2.5 ^ 3.0 * 4.2",
			65.625,
		},
		{
			"2 ^ 3 * 4 + 1",
			33,
		},
		{
			"2 ^ 2.5 / 4.0",
			1.4142135623731,
		},
		{
			"2 ^ 3 ^ 2",
			512,
		},
	}

	for _, tt := range tests {
		lexer := NewLexer(tt.input)
		parser := NewParser(lexer)
		calculator := parser.Parse()
		checkParserErrors(t, parser)

		actual := calculator.Evaluate()
		if !NearlyEqual(actual, tt.expected) {
			t.Errorf("input=%q, expected=%f, got=%f", tt.input, tt.expected, actual)
		}
	}
}

func TestOperatorPrecedenceEvaluation(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{
			"-2 * 3",
			-6,
		},
		{
			"1 + 2 + 3",
			6,
		},
		{
			"1 - 2 + 3",
			2,
		},
		{
			"1 * 2 * 3",
			6,
		},
		{
			"1 + 2 / 3",
			1.6666666666667,
		},
		{
			"1 + 2 * 3 + 4 / 5 - 6",
			1.8,
		},
		{
			"-1 * -4",
			4,
		},
		{
			"2",
			2,
		},
		{
			"2 + 3 * 4",
			14,
		},
		{
			"2 - 3 / 4",
			1.25,
		},
		{
			"2 * 3 / 4",
			1.5,
		},
		{
			"0 * 5 + 3",
			3,
		},
		{
			"-2 * 3.5",
			-7,
		},
		{
			"1.0 + 2.5 + 3.3",
			6.8,
		},
		{
			"1.5 - 2.2 + 3.1",
			2.4,
		},
		{
			"1.2 * 2.0 * 3.5",
			8.4,
		},
		{
			"1.0 + 2.5 / 3.0",
			1.8333333333333,
		},
		{
			"1.0 + 2.5 * 3.0 + 4.7 / 5.2 - 6.0",
			3.4038461538461,
		},
		{
			"-1.5 * -4.2",
			6.3,
		},
		{
			"2.5",
			2.5,
		},
		{
			"2.5 + 3.0 * 4.2",
			15.1,
		},
		{
			"2.8 - 3.1 / 4.0",
			2.025,
		},
		{
			"2.5 * 3.2 / 4.5",
			1.7777777777778,
		},
		{
			"0.0 * 5.3 + 3.9",
			3.9,
		},
	}

	for _, tt := range tests {
		lexer := NewLexer(tt.input)
		parser := NewParser(lexer)
		calculator := parser.Parse()
		checkParserErrors(t, parser)

		actual := calculator.Evaluate()
		if !NearlyEqual(actual, tt.expected) {
			t.Errorf("input=%q, expected=%f, got=%f", tt.input, tt.expected, actual)
		}
	}
}
