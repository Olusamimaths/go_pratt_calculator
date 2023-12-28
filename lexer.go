package main

import (
	"fmt"
	"strings"
)

// Represents end of file (EOF)
var EOF = Token{
	Type:  "",
	Value: "",
}

// A Lexer contains all the tokens after lexing a given string
// the Position holds the current read position
type Lexer struct {
	Tokens   []Token
	Position int
}

// Returns the value of the current token Position in lexer points to
// if Position is out of bonds, returns EOF token
func (l *Lexer) CurToken() Token {
	if l.Position >= len(l.Tokens) {
		return EOF
	}

	return l.Tokens[l.Position]
}

// Returns the next token after the current one, allowing us to peek one position forward
func (l *Lexer) NextToken() Token {
	l.Position = l.Position + 1
	if l.Position >= len(l.Tokens) {
		return EOF
	}

	return l.Tokens[l.Position]
}

// Checks that the next token is of a given type
func (l *Lexer) Expect(t string) bool {
	token := l.NextToken()
	if token.Type != t {
		panic(fmt.Sprintf("Invalid token encountered: %+v\n", token))
	}
	return true
}

// Creates a new lexer for a given string
// The string is tokenized and the Lexer is initialized from the tokens in the string
func NewLexer(s string) Lexer {
	// remove all whitespaces
	s = strings.ReplaceAll(s, " ", "")
	var tkns []Token

	for len(s) > 0 {
		var tokenType string
		var match string

		for _, t := range Tokens {
			if t.Re.MatchString(s) {
				tokenType = t.Type
				match = t.Re.FindString(s)
				break
			}
		}

		tkns = append(tkns, Token{Type: tokenType, Value: match})
		// this removes the matched string from the input string s
		// and assigns s to the rest of the string
		s = s[len(match):]
	}

	return Lexer{Tokens: tkns, Position: 0}
}
