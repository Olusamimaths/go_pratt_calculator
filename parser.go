package main

type (
	nudFn func()
	ledFn func(left Token)
)

type (
	nudFns map[string]nudFn
	ledFns map[string]ledFn
)

type Parser struct {
	l *Lexer

	curToken  Token
	peekToken Token

	NUDS nudFns
	LEDS ledFns
}
