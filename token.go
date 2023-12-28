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
	Value string
}

var Tokens = []struct {
	Type string
	Re   *regexp.Regexp
}{
	{Type: LPAREN, Re: regexp.MustCompile(`\(`)},
	{Type: RPAREN, Re: regexp.MustCompile(`\)`)},
	{Type: EXPONENTIAL, Re: regexp.MustCompile(`\^`)},
	{Type: DIVIDE, Re: regexp.MustCompile(`/`)},
	{Type: MULT, Re: regexp.MustCompile(`\*`)},
	{Type: MINUS, Re: regexp.MustCompile(`-`)},
	{Type: PLUS, Re: regexp.MustCompile(`\+`)},

	{Type: NUMBER, Re: regexp.MustCompile(`(?:\d+(?:\.\d*)?|\.\d+)`)},
	{Type: IDENTIFIER, Re: regexp.MustCompile(`[A-Za-z]+`)},
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
