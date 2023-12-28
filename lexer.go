package main

import (
	"fmt"
	"regexp"
	"strings"
)

var EOF = Token{
	Type:  "",
	Match: "",
}

type Lexer struct {
	Tokens   []Token
	Position int
}

func (l *Lexer) Peek() Token {
	if l.Position >= len(l.Tokens) {
		return EOF
	}

	return l.Tokens[l.Position]
}

func (l *Lexer) Next() Token {
	nextIndex := l.Position + 1
	if nextIndex >= len(l.Tokens) {
		return EOF
	}

	return l.Tokens[nextIndex]
}

func (l *Lexer) Expect(t string) bool {
	token := l.Next()
	if token.Type != t {
		panic(fmt.Sprintf("Invalid token encountered: %+v\n", token))
	}
	return true
}

func Lex(s string) Lexer {
	tokens := []struct {
		Type string
		Re   *regexp.Regexp
	}{
		{Type: "(", Re: regexp.MustCompile(`\(`)},
		{Type: ")", Re: regexp.MustCompile(`\)`)},
		{Type: "^", Re: regexp.MustCompile(`\^`)},
		{Type: "/", Re: regexp.MustCompile(`/`)},
		{Type: "*", Re: regexp.MustCompile(`\*`)},
		{Type: "-", Re: regexp.MustCompile(`-`)},
		{Type: "+", Re: regexp.MustCompile(`\+`)},

		{Type: "NUMBER", Re: regexp.MustCompile(`(?:\d+(?:\.\d*)?|\.\d+)`)},
		{Type: "IDENTIFIER", Re: regexp.MustCompile(`[A-Za-z]+`)},
	}

	s = strings.ReplaceAll(s, " ", "")
	var tkns []Token

	for len(s) > 0 {
		var tokenType string
		var match string

		for _, t := range tokens {
			if t.Re.MatchString(s) {
				tokenType = t.Type
				match = t.Re.FindString(s)
				break
			}
		}

		tkns = append(tkns, Token{Type: tokenType, Match: match})
		s = s[len(match):]
	}

	return Lexer{Tokens: tkns, Position: 0}
}
