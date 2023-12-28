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
	}

	results := map[int][]Token{
		0: {
			{Type: NUMBER, TokenLiteral: "1"},
			{Type: PLUS, TokenLiteral: "+"},
			{Type: NUMBER, TokenLiteral: "2"},
			{Type: PLUS, TokenLiteral: "+"},
			{Type: NUMBER, TokenLiteral: "3"},
		},
		1: {
			{Type: NUMBER, TokenLiteral: "1"},
			{Type: MULT, TokenLiteral: "*"},
			{Type: NUMBER, TokenLiteral: "2"},
			{Type: PLUS, TokenLiteral: "+"},
			{Type: NUMBER, TokenLiteral: "3"},
		},
		2: {
			{Type: NUMBER, TokenLiteral: "1"},
			{Type: MINUS, TokenLiteral: "-"},
			{Type: NUMBER, TokenLiteral: "2"},
			{Type: MULT, TokenLiteral: "*"},
			{Type: NUMBER, TokenLiteral: "3"},
		},
		3: {
			{Type: NUMBER, TokenLiteral: "1"},
			{Type: MINUS, TokenLiteral: "-"},
			{Type: NUMBER, TokenLiteral: "2"},
			{Type: DIVIDE, TokenLiteral: "/"},
			{Type: NUMBER, TokenLiteral: "3"},
		},
		4: {
			{Type: NUMBER, TokenLiteral: "1"},
			{Type: MULT, TokenLiteral: "*"},
			{Type: NUMBER, TokenLiteral: "2"},
			{Type: DIVIDE, TokenLiteral: "/"},
			{Type: NUMBER, TokenLiteral: "3"},
		},
		5: {
			{Type: NUMBER, TokenLiteral: "1"},
			{Type: MULT, TokenLiteral: "*"},
			{Type: NUMBER, TokenLiteral: "2"},
			{Type: MULT, TokenLiteral: "*"},
			{Type: NUMBER, TokenLiteral: "3"},
		},
		6: {
			{Type: NUMBER, TokenLiteral: "1"},
			{Type: DIVIDE, TokenLiteral: "/"},
			{Type: NUMBER, TokenLiteral: "2"},
			{Type: DIVIDE, TokenLiteral: "/"},
			{Type: NUMBER, TokenLiteral: "3"},
		},
		7: {
			{Type: NUMBER, TokenLiteral: "1"},
			{Type: MINUS, TokenLiteral: "-"},
			{Type: NUMBER, TokenLiteral: "2"},
			{Type: MINUS, TokenLiteral: "-"},
			{Type: NUMBER, TokenLiteral: "3"},
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

			if tok.TokenLiteral != tt.TokenLiteral {
				t.Fatalf("test [%d] token TokenLiteral wrong. expected=%q, got=%q", i, tt.TokenLiteral, tok.TokenLiteral)
			}
		}
	}
}
