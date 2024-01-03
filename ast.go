package main

import (
	"bytes"
	"math"
)

// The basic form of representation in the program, represents a parsed token and a Node of the AST
type Expression interface {
	TokenLiteral() string
	String() string
	Evaluate() float64
}

// Calculator represents the 'program'
type Calculator struct {
	Expression Expression
}

func (c *Calculator) TokenLiteral() string { return c.Expression.TokenLiteral() }
func (c *Calculator) String() string       { return c.Expression.String() }
func (c *Calculator) Evaluate() float64 {
	return c.Expression.Evaluate()
}

type NumberLiteral struct {
	Token Token
	Value float64
}

func (nl *NumberLiteral) TokenLiteral() string { return nl.Token.Literal }
func (nl *NumberLiteral) String() string       { return nl.Token.Literal }
func (nl *NumberLiteral) Evaluate() float64 {
	return nl.Value
}

type InfixExpression struct {
	Token    Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

func (ie *InfixExpression) Evaluate() float64 {
	if ie.Operator == PLUS {
		return ie.Left.Evaluate() + ie.Right.Evaluate()
	}

	if ie.Operator == MINUS {
		return ie.Left.Evaluate() - ie.Right.Evaluate()
	}

	if ie.Operator == MULT {
		return ie.Left.Evaluate() * ie.Right.Evaluate()
	}

	if ie.Operator == DIVIDE {
		return ie.Left.Evaluate() / ie.Right.Evaluate()
	}

	if ie.Operator == EXPONENTIAL {
		return math.Pow(ie.Left.Evaluate(), ie.Right.Evaluate())
	}
	return 0
}

type PrefixExpression struct {
	Token    Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (pe *PrefixExpression) Evaluate() float64 {
	if pe.Operator == MINUS {
		return -pe.Right.Evaluate()
	}

	return 0
}
