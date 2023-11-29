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

	if program == nil {
		t.Fatalf("Nil returned from ParseProgram()")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements contain %d, expected %d", len(program.Statements), 1)
	}

	tests := []struct {
		expectedIndentifier string
	}{
		{"x"},
	}

	for index, test := range tests {
		statement := program.Statements[index]
		if !testDefStatement(t, statement, test.expectedIndentifier) {
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
