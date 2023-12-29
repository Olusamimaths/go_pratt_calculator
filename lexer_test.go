package main

import (
	"testing"
)

type lexerTest struct {
	input          string
	expectedResult []Token
}

func TestNextToken(t *testing.T) {
	tests := []lexerTest{
		{
			input: "1 + 2 + 3",
			expectedResult: []Token{
				{Type: NUMBER, Literal: "1"},
				{Type: PLUS, Literal: "+"},
				{Type: NUMBER, Literal: "2"},
				{Type: PLUS, Literal: "+"},
				{Type: NUMBER, Literal: "3"},
			},
		},
		{
			input: "1 * 2 + 3",
			expectedResult: []Token{
				{Type: NUMBER, Literal: "1"},
				{Type: MULT, Literal: "*"},
				{Type: NUMBER, Literal: "2"},
				{Type: PLUS, Literal: "+"},
				{Type: NUMBER, Literal: "3"},
			},
		},
		{
			input: "1 - 2 * 3",
			expectedResult: []Token{
				{Type: NUMBER, Literal: "1"},
				{Type: MINUS, Literal: "-"},
				{Type: NUMBER, Literal: "2"},
				{Type: MULT, Literal: "*"},
				{Type: NUMBER, Literal: "3"},
			},
		},
		{
			input: "1 - 2 / 3",
			expectedResult: []Token{
				{Type: NUMBER, Literal: "1"},
				{Type: MINUS, Literal: "-"},
				{Type: NUMBER, Literal: "2"},
				{Type: DIVIDE, Literal: "/"},
				{Type: NUMBER, Literal: "3"},
			},
		},
		{
			input: "1 * 2 / 3",
			expectedResult: []Token{
				{Type: NUMBER, Literal: "1"},
				{Type: MULT, Literal: "*"},
				{Type: NUMBER, Literal: "2"},
				{Type: DIVIDE, Literal: "/"},
				{Type: NUMBER, Literal: "3"},
			},
		},
		{
			input: "1 * 2 * 3",
			expectedResult: []Token{
				{Type: NUMBER, Literal: "1"},
				{Type: MULT, Literal: "*"},
				{Type: NUMBER, Literal: "2"},
				{Type: MULT, Literal: "*"},
				{Type: NUMBER, Literal: "3"},
			},
		},
		{
			input: "1 / 2 / 3",
			expectedResult: []Token{
				{Type: NUMBER, Literal: "1"},
				{Type: DIVIDE, Literal: "/"},
				{Type: NUMBER, Literal: "2"},
				{Type: DIVIDE, Literal: "/"},
				{Type: NUMBER, Literal: "3"},
			},
		},
		{
			input: "1 - 2 - 3",
			expectedResult: []Token{
				{Type: NUMBER, Literal: "1"},
				{Type: MINUS, Literal: "-"},
				{Type: NUMBER, Literal: "2"},
				{Type: MINUS, Literal: "-"},
				{Type: NUMBER, Literal: "3"},
			},
		},
		{
			input: "1.2 - 2.34 - 3.567",
			expectedResult: []Token{
				{Type: NUMBER, Literal: "1.2"},
				{Type: MINUS, Literal: "-"},
				{Type: NUMBER, Literal: "2.34"},
				{Type: MINUS, Literal: "-"},
				{Type: NUMBER, Literal: "3.567"},
			},
		},
		{
			input: "1 * 2.34 / 3.567",
			expectedResult: []Token{
				{Type: NUMBER, Literal: "1"},
				{Type: MULT, Literal: "*"},
				{Type: NUMBER, Literal: "2.34"},
				{Type: DIVIDE, Literal: "/"},
				{Type: NUMBER, Literal: "3.567"},
			},
		},
	}

	for i, test := range tests {
		lexer := NewLexer(test.input)
		expectedResult := test.expectedResult

		for _, tt := range expectedResult {
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
