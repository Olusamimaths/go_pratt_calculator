package main

import (
	"fmt"
	"strconv"
)

type (
	nudFn func() Expression
	ledFn func(left Expression) Expression
)

type (
	nudFns map[TokenType]nudFn
	ledFns map[TokenType]ledFn
)

type Parser struct {
	l *Lexer

	curToken  Token
	peekToken Token

	NUDS nudFns
	LEDS ledFns

	errors []string
}

func NewParser(l *Lexer) *Parser {
	parser := &Parser{l: l}

	parser.NUDS = make(nudFns)
	parser.registerNud(NUMBER, parser.parseNumberLiteral)
	parser.registerNud(MINUS, parser.parsePrefixExpression)

	parser.LEDS = make(ledFns)
	parser.registerLed(PLUS, parser.parseInfixExpression)
	parser.registerLed(MINUS, parser.parseInfixExpression)
	parser.registerLed(MULT, parser.parseInfixExpression)
	parser.registerLed(DIVIDE, parser.parseInfixExpression)
	parser.registerLed(EXPONENTIAL, parser.parseInfixExpression)

	parser.initializeTokens()
	return parser
}

func (p *Parser) registerNud(tokenType TokenType, fn nudFn) {
	p.NUDS[tokenType] = fn
}

func (p *Parser) registerLed(tokenType TokenType, fn ledFn) {
	p.LEDS[tokenType] = fn
}

func (p *Parser) Parse() *Calculator {
	calculator := &Calculator{}

	calculator.Expression = p.parseExpression(0)

	return calculator
}

func (p *Parser) parseNumberLiteral() Expression {
	expression := &NumberLiteral{Token: p.curToken}

	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as a number", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	expression.Value = value

	return expression
}

func (p *Parser) parsePrefixExpression() Expression {
	expression := &PrefixExpression{
		Operator: p.curToken.Literal,
		Token:    p.curToken,
	}

	p.nextToken()

	prefixBP := BPS[PREFIX]
	expression.Right = p.parseExpression(prefixBP)

	return expression
}

func (p *Parser) parseInfixExpression(left Expression) Expression {
	expression := &InfixExpression{
		Operator: p.curToken.Literal,
		Token:    p.curToken,
		Left:     left,
	}

	curTokenType := p.curToken.Type

	bindingPower := p.currentBindingPower()
	p.nextToken()

	if curTokenType == EXPONENTIAL {
		expression.Right = p.parseExpression(bindingPower - 1)
	} else {
		expression.Right = p.parseExpression(bindingPower)
	}

	return expression
}

// The core of the Pratt Parsing Algorithm
func (p *Parser) parseExpression(rbp int) Expression {
	nud := p.NUDS[p.curToken.Type]
	if nud == nil {
		msg := fmt.Sprintf("could not find prefix parser for token type=%q", p.curToken.Type)
		p.errors = append(p.errors, msg)
		return nil
	}
	left := nud()

	for p.peekBindingPower() > rbp {
		led := p.LEDS[p.peekToken.Type]
		if led == nil {
			return left
		}
		p.nextToken()

		left = led(left)
	}
	return left
}

func (p *Parser) peekBindingPower() int {
	if bps, ok := BPS[p.peekToken.Type]; ok {
		return bps
	}

	return 0
}

func (p *Parser) currentBindingPower() int {
	if bps, ok := BPS[p.curToken.Type]; ok {
		return bps
	}

	return 0
}

func (p *Parser) initializeTokens() {
	p.nextToken() // first call sets p.peekToken vial p.l.NextToken
	p.nextToken() // second call sets p.curToken to value of p.peekToken from first call,
	// then p.peekToken is set to next token
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}
