package ast

import (
	"Ahmadi/token"
	"bytes"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, statement := range p.Statements {
		out.WriteString(statement.String())
	}

	return out.String()
}

type DefStatement struct {
	Token token.Token // def
	Name  *Identifier
	Value Expression
}

func (ds *DefStatement) statementNode()       {}
func (ds *DefStatement) TokenLiteral() string { return ds.Token.Literal }
func (ds *DefStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ds.TokenLiteral() + " ")
	out.WriteString(ds.Name.String())
	out.WriteString(" = ")

	if ds.Value != nil {
		out.WriteString(ds.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
	return i.Value
}

type ReturnStatement struct {
	Token       token.Token // return
	ReturnValue Expression
}

func (returnStatement *ReturnStatement) statementNode()       {}
func (returnStatement *ReturnStatement) TokenLiteral() string { return returnStatement.Token.Literal }
func (returnStatement *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(returnStatement.TokenLiteral() + " ")

	if returnStatement.ReturnValue != nil {
		out.WriteString(returnStatement.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (expressionStatement *ExpressionStatement) statementNode() {}
func (expressionStatement *ExpressionStatement) TokenLiteral() string {
	return expressionStatement.Token.Literal
}
func (expressionStatement *ExpressionStatement) String() string {
	if expressionStatement.Expression != nil {
		return expressionStatement.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (integerLiteral *IntegerLiteral) expressionNode()      {}
func (integerLiteral *IntegerLiteral) TokenLiteral() string { return integerLiteral.Token.Literal }
func (integerLiteral *IntegerLiteral) String() string       { return integerLiteral.TokenLiteral() }

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (prefixExpression *PrefixExpression) expressionNode() {}
func (prefixExpression *PrefixExpression) TokenLiteral() string {
	return prefixExpression.Token.Literal
}
func (prefixExpression *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(prefixExpression.Operator)
	out.WriteString(prefixExpression.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (infixExpression *InfixExpression) expressionNode() {}
func (infixExpression *InfixExpression) TokenLiteral() string {
	return infixExpression.Token.Literal
}
func (infixExpression *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(infixExpression.Left.String())
	out.WriteString(" " + infixExpression.Operator + " ")
	out.WriteString(infixExpression.Right.String())
	out.WriteString(")")

	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (bl *Boolean) expressionNode()      {}
func (bl *Boolean) TokenLiteral() string { return bl.Token.Literal }
func (bl *Boolean) String() string       { return bl.TokenLiteral() }
