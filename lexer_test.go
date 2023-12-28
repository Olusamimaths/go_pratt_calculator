package main

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	inputs := []string{
		"1 + 2 + 3",
		"1 * 2 + 3",
		"1 - 2 * 3",
		"1 - 2 / 3",
		"1 * 2 / 3",
		"1 * 2 * 3",
		"1 / 2 / 3",
		"1 - 2 - 3",
		"1.2 - 2.34 - 3.567",
		"1 * 2.34 / 3.567",
	}

	results := map[int][]Token{
		0: {
			{Type: NUMBER, Literal: "1"},
			{Type: PLUS, Literal: "+"},
			{Type: NUMBER, Literal: "2"},
			{Type: PLUS, Literal: "+"},
			{Type: NUMBER, Literal: "3"},
		},
		1: {
			{Type: NUMBER, Literal: "1"},
			{Type: MULT, Literal: "*"},
			{Type: NUMBER, Literal: "2"},
			{Type: PLUS, Literal: "+"},
			{Type: NUMBER, Literal: "3"},
		},
		2: {
			{Type: NUMBER, Literal: "1"},
			{Type: MINUS, Literal: "-"},
			{Type: NUMBER, Literal: "2"},
			{Type: MULT, Literal: "*"},
			{Type: NUMBER, Literal: "3"},
		},
		3: {
			{Type: NUMBER, Literal: "1"},
			{Type: MINUS, Literal: "-"},
			{Type: NUMBER, Literal: "2"},
			{Type: DIVIDE, Literal: "/"},
			{Type: NUMBER, Literal: "3"},
		},
		4: {
			{Type: NUMBER, Literal: "1"},
			{Type: MULT, Literal: "*"},
			{Type: NUMBER, Literal: "2"},
			{Type: DIVIDE, Literal: "/"},
			{Type: NUMBER, Literal: "3"},
		},
		5: {
			{Type: NUMBER, Literal: "1"},
			{Type: MULT, Literal: "*"},
			{Type: NUMBER, Literal: "2"},
			{Type: MULT, Literal: "*"},
			{Type: NUMBER, Literal: "3"},
		},
		6: {
			{Type: NUMBER, Literal: "1"},
			{Type: DIVIDE, Literal: "/"},
			{Type: NUMBER, Literal: "2"},
			{Type: DIVIDE, Literal: "/"},
			{Type: NUMBER, Literal: "3"},
		},
		7: {
			{Type: NUMBER, Literal: "1"},
			{Type: MINUS, Literal: "-"},
			{Type: NUMBER, Literal: "2"},
			{Type: MINUS, Literal: "-"},
			{Type: NUMBER, Literal: "3"},
		},
		8: {
			{Type: NUMBER, Literal: "1.2"},
			{Type: MINUS, Literal: "-"},
			{Type: NUMBER, Literal: "2.34"},
			{Type: MINUS, Literal: "-"},
			{Type: NUMBER, Literal: "3.567"},
		},
		9: {
			{Type: NUMBER, Literal: "1"},
			{Type: MULT, Literal: "*"},
			{Type: NUMBER, Literal: "2.34"},
			{Type: DIVIDE, Literal: "/"},
			{Type: NUMBER, Literal: "3.567"},
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

			if tok.Literal != tt.Literal {
				t.Fatalf(
					"test [%d] token Literal wrong. expected=%q, got=%q",
					i,
					tt.Literal,
					tok.Literal,
				)
			}
		}
	}
}
