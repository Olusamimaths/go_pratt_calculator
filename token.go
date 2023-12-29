package main

import (
	"regexp"
)

const (
	LPAREN = "("
	RPAREN = ")"

	EXPONENTIAL = "^"
	DIVIDE      = "/"
	MULT        = "*"
	MINUS       = "-"
	PLUS        = "+"

	NUMBER     = "NUMBER"
	IDENTIFIER = "IDENTIFIER"

	PREFIX = "PREFIX"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// The order of the slice matters; when there is potential ambiguity in the input
// The Lexer will try to match input with the regex in order the regex are defined
// The first matching token will be choosen
var TokenMatchers = []struct {
	Type TokenType
	Re   *regexp.Regexp
}{
	{Type: NUMBER, Re: regexp.MustCompile(`(?:\d+(?:\.\d*)?|\.\d+)`)},
	{Type: IDENTIFIER, Re: regexp.MustCompile(`[A-Za-z]+`)},
	{Type: PLUS, Re: regexp.MustCompile(`\+`)},
	{Type: MINUS, Re: regexp.MustCompile(`-`)},
	{Type: MULT, Re: regexp.MustCompile(`\*`)},
	{Type: DIVIDE, Re: regexp.MustCompile(`\/`)},
	{Type: EXPONENTIAL, Re: regexp.MustCompile(`\^`)},
	{Type: LPAREN, Re: regexp.MustCompile(`\(`)},
	{Type: RPAREN, Re: regexp.MustCompile(`\)`)},
}

type BindingPowers map[TokenType]int

var BPS = BindingPowers{
	NUMBER:      10,
	IDENTIFIER:  10,
	RPAREN:      10,
	PLUS:        20,
	MINUS:       20,
	MULT:        30,
	DIVIDE:      30,
	EXPONENTIAL: 40,
	LPAREN:      50,
	PREFIX:      80,
}
