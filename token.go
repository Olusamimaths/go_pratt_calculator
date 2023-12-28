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
)

type Token struct {
	Type  string
	Literal string
}

// The order of the slice matters; when there is potential ambiguity in the input
// The Lexer will try to match input with the regex in order the regex are defined
// The first matching token will be choosen
var TokenMatchers = []struct {
	Type string
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

type BindingPowers map[string]int

var BPS = BindingPowers{
	NUMBER:      0,
	IDENTIFIER:  0,
	RPAREN:      0,
	PLUS:        10,
	MINUS:       10,
	MULT:        20,
	DIVIDE:      20,
	EXPONENTIAL: 30,
	LPAREN:      40,
}
