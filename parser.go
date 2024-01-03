package main

import (
	"fmt"
	"strconv"
)

type (
	nudFn func() (Expression, error)
	ledFn func(left Expression) (Expression, error)
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
	parser.registerNud(LPAREN, parser.parsePrefixExpression)
	parser.registerNud(RPAREN, parser.parsePrefixExpression)

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

	expr, err := p.parseExpression(0)
	if err != nil {
		fmt.Println(err.Error())
		panic("An error occured while parsing expression")
	}

	if expr == nil {
		panic("error occurred when parsing the whole expression")
	}
	calculator.Expression = expr
	return calculator
}

func (p *Parser) parseNumberLiteral() (Expression, error) {
	expression := &NumberLiteral{Token: p.curToken}

	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as a number", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil, err
	}

	expression.Value = value

	return expression, nil
}

func (p *Parser) parsePrefixExpression() (Expression, error) {
	if p.curToken.Type == LPAREN {
		p.nextToken()
		expr, err := p.parseExpression(0)
		if err != nil {
			msg := fmt.Sprintf("an error occured passing ( expression: %q", err.Error())
			p.errors = append(p.errors, msg)
			return nil, err
		}
		p.nextToken()
		return expr, err
	}
	expression := &PrefixExpression{
		Operator: p.curToken.Literal,
		Token:    p.curToken,
	}

	p.nextToken()

	prefixBP := BPS[PREFIX]
	expr, err := p.parseExpression(prefixBP)
	if err != nil {
		return nil, err
	}

	expression.Right = expr
	return expression, nil
}

func (p *Parser) parseInfixExpression(left Expression) (Expression, error) {
	expression := &InfixExpression{
		Operator: p.curToken.Literal,
		Token:    p.curToken,
		Left:     left,
	}

	curTokenType := p.curToken.Type

	bindingPower := p.currentBindingPower()
	p.nextToken()

	if curTokenType == EXPONENTIAL {
		bindingPower = bindingPower - 1
	}

	expr, err := p.parseExpression(bindingPower)
	if err != nil {
		return nil, err
	}

	expression.Right = expr
	return expression, nil
}

// The core of the Pratt Parsing Algorithm
func (p *Parser) parseExpression(rbp int) (Expression, error) {
	nud := p.NUDS[p.curToken.Type]
	if nud == nil {
		msg := fmt.Sprintf("could not find prefix parser for token type=%q", p.curToken.Type)
		p.errors = append(p.errors, msg)
		return nil, nil
	}

	left, err := nud()
	if err != nil {
		return nil, err
	}

	for p.peekBindingPower() > rbp {
		led := p.LEDS[p.peekToken.Type]
		if led == nil {
			return left, nil
		}
		p.nextToken()

		left, err = led(left)
		if err != nil {
			return nil, err
		}
	}
	return left, nil
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

func (p *Parser) LogErrors() {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	fmt.Printf("parser has %d errors", len(errors))
	for _, msg := range errors {
		fmt.Printf("parser error: %q", msg)
	}

	panic("An error occured while parsing expression")
}
