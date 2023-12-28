package main

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	inputs := []string{
		"1 + 2 + 3",
		"1 * 2 + 3",
		"1 - 2 * 3",
	}

	results := map[int][]Token{
		0: {
			{Type: NUMBER, Value: "1"},
			{Type: PLUS, Value: "+"},
			{Type: NUMBER, Value: "2"},
			{Type: PLUS, Value: "+"},
			{Type: NUMBER, Value: "3"},
		},
		1: {
			{Type: NUMBER, Value: "1"},
			{Type: MULT, Value: "*"},
			{Type: NUMBER, Value: "2"},
			{Type: PLUS, Value: "+"},
			{Type: NUMBER, Value: "3"},
		},
		2: {
			{Type: NUMBER, Value: "1"},
			{Type: MINUS, Value: "-"},
			{Type: NUMBER, Value: "2"},
			{Type: MULT, Value: "*"},
			{Type: NUMBER, Value: "3"},
		},
	}

	for i, input := range inputs {
		lexer := NewLexer(input)
		result := results[i]
		for _, tt := range result {
			tok := lexer.NextToken()
			if tok.Type != tt.Type {
				t.Fatalf("test [%d] token type wrong. expected=%q, got=%q", i, tt.Type, tok.Type)
			}

			if tok.Value != tt.Value {
				t.Fatalf("test [%d] token value wrong. expected=%q, got=%q", i, tt.Value, tok.Value)
			}
		}
	}
}
