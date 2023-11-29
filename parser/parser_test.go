package parser

import (
	"Ahmadi/ast"
	"Ahmadi/lexer"
	"testing"
)

func TestDefStatements(t *testing.T) {
	input := `
	def x = 12;
	`

	lexer := lexer.New(input)
	parser := New(lexer)
	program := parser.ParseProgram()
	checkParseErrors(t, parser)

	if program == nil {
		nilProgram(t)
	}

	if len(program.Statements) != 1 {
		statementCountError(t, 1, len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
	}

	for index, test := range tests {
		statement := program.Statements[index]
		if !testDefStatement(t, statement, test.expectedIdentifier) {
			return
		}
	}
}

func TestIncorrectDefStatements(t *testing.T) {
	input := `
	def first_age = 25;
	def second_age = 23;
	`

	lexer := lexer.New(input)
	parser := New(lexer)

	program := parser.ParseProgram()
	checkParseErrors(t, parser)

	if program == nil {
		nilProgram(t)
	}

	if len(program.Statements) != 2 {
		statementCountError(t, 2, len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"first_age"},
		{"second_age"},
	}

	for index, test := range tests {
		staement := program.Statements[index]
		if !testDefStatement(t, staement, test.expectedIdentifier) {
			return
		}
	}

}

func testDefStatement(t *testing.T, statement ast.Statement, identifier string) bool {
	if statement.TokenLiteral() != "def" {
		t.Errorf("statement.TokenLiteral not 'def', got=%q", statement.TokenLiteral())
		return false
	}

	defStatement, ok := statement.(*ast.DefStatement)

	if !ok {
		t.Errorf("statement not *ast.DefStatement, got=%T", statement)
		return false
	}

	if defStatement.Name.Value != identifier {
		t.Errorf("defStatement.Name.Value not '%s' got=%s", identifier, defStatement.Name.Value)
		return false
	}

	if defStatement.Name.TokenLiteral() != identifier {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s", identifier, defStatement.Name.TokenLiteral())
		return false
	}

	if defStatement.Token.Literal != "def" {
		t.Errorf("defStatement.Token.Literal not 'def' got='%s'", defStatement.Token.Literal)
	}

	return true
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, message := range errors {
		t.Errorf("parser error: %q", message)
	}
	t.FailNow()
}

func nilProgram(t *testing.T) {
	t.Fatalf("Nil returned from ParseProgram()")
}

func statementCountError(t *testing.T, expected int, found int) {
	t.Fatalf("program.Statements contain %d, expected %d", found, expected)
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return -12;
	`

	lexer := lexer.New(input)
	parser := New(lexer)

	program := parser.ParseProgram()
	checkParseErrors(t, parser)

	if len(program.Statements) != 3 {
		statementCountError(t, 3, len(program.Statements))
	}

	for _, statement := range program.Statements {
		returnStmt, ok := statement.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("statement not *ast.ReturnStatement, got='%T'", statement)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got='%q'", returnStmt.TokenLiteral())
		}
	}
}
