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
	Match string
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
