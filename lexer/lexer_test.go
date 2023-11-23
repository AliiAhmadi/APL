package lexer

import (
	"Ahmadi/token"
	"testing"
)

func TestNextToken(t *testing.T) {

	input := `+=+=(){};,`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.ASSIGN, "="},
		{token.LPARENTHESES, "("},
		{token.RPARENTHESES, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.COMMA, ","},
	}

	lexer := New(input)

	for i, test := range tests {
		tok := lexer.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenWiderWithMoreDetails(t *testing.T) {
	input := `
	def first_variable = 10;
	def second_variable = 20;

	def adder = fun(first, second) {
		return first + second;
	}

	def result = adder(first_variable, second_variable);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.DEF, "def"},
		{token.ID, "first_variable"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.DEF, "def"},
		{token.ID, "second_variable"},
		{token.ASSIGN, "="},
		{token.INT, "20"},
		{token.SEMICOLON, ";"},
		{token.DEF, "def"},
		{token.ID, "adder"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fun"},
		{token.LPARENTHESES, "("},
		{token.ID, "first"},
		{token.COMMA, ","},
		{token.ID, "second"},
		{token.RPARENTHESES, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.ID, "first"},
		{token.PLUS, "+"},
		{token.ID, "second"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.DEF, "def"},
		{token.ID, "result"},
		{token.ASSIGN, "="},
		{token.ID, "adder"},
		{token.LPARENTHESES, "("},
		{token.ID, "first_variable"},
		{token.COMMA, ","},
		{token.ID, "second_variable"},
		{token.RPARENTHESES, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, test := range tests {
		tok := lexer.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestIfAndElseKeywords(t *testing.T) {
	input := `
	if(first_var > second_var) {
		return first_var
	}else {
		return second_var
	}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IF, "if"},
		{token.LPARENTHESES, "("},
		{token.ID, "first_var"},
		{token.GREATER, ">"},
		{token.ID, "second_var"},
		{token.RPARENTHESES, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.ID, "first_var"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.ID, "second_var"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, test := range tests {
		tok := lexer.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}
