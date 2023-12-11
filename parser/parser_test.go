package parser

import (
	"Ahmadi/ast"
	"Ahmadi/lexer"
	"fmt"
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

func TestIdentifierExpression(t *testing.T) {
	input := "test;"

	lexer := lexer.New(input)
	parser := New(lexer)
	program := parser.ParseProgram()
	checkParseErrors(t, parser)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	identifier, ok := statement.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression not *ast.Identifier. got=%T", statement.Expression)
	}

	if identifier.Value != "test" {
		t.Fatalf("identifier.Value not %s. got='%s'", "test", identifier.Value)
	}

	if identifier.TokenLiteral() != "test" {
		t.Fatalf("identifier.TokenLiteral() not %s. got='%s'", "test", identifier.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "100;"

	lexer := lexer.New(input)
	parser := New(lexer)
	program := parser.ParseProgram()
	checkParseErrors(t, parser)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] not *ast.ExpressionStatement. got='%T'", program.Statements[0])
	}

	literal, ok := statement.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got='%T'", statement.Expression)
	}

	if literal.Value != 100 {
		t.Fatalf("literal.Value not %d. got=%d", 100, literal.Value)
	}

	if literal.TokenLiteral() != "100" {
		t.Fatalf("literal.TokenLiteral not %s. got'%s'", "100", literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{"!100;", "!", 100},
		{"-20;", "-", 20},
		{"!true;", "!", true},
		{"!false;", "!", false},
	}

	for _, test := range prefixTests {
		lexer := lexer.New(test.input)
		parser := New(lexer)
		program := parser.ParseProgram()
		checkParseErrors(t, parser)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statement. got=%d", 1, len(program.Statements))
		}

		statement, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got'%T'", program.Statements[0])
		}

		exp, ok := statement.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("statement is not ast.PrefixExpression. got='%T'", statement.Expression)
		}

		if exp.Operator != test.operator {
			t.Fatalf("exp.Operator is not '%s'. got='%s'", test.operator, exp.Operator)
		}

		if value, ok := test.value.(bool); ok {
			if !testBooleanLiteral(t, exp.Right, value) {
				return
			}
		} else {
			if !testIntegerLiteral(t, exp.Right, test.value) {
				return
			}
		}
	}
}

func testIntegerLiteral(t *testing.T, integerLiteral ast.Expression, value interface{}) bool {
	integ, ok := integerLiteral.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("integerLiteral not *ast.IntegerLiteral. got='%T'", integerLiteral)
		return false
	}

	if val, ok := value.(int); ok {
		if integ.Value != int64(val) {
			t.Errorf("integ.Value not %d. got='%d'", value, integ.Value)
			return false
		}
	} else {
		if integ.Value != value.(int64) {
			t.Errorf("integ.Value not %d. got='%d'", value, integ.Value)
			return false
		}
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got='%s'", value, integ.TokenLiteral())
		return false
	}

	return true
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"1 + 2;", 1, "+", 2},
		{"5 - 5;", 5, "-", 5},
		{"10 * 1000;", 10, "*", 1000},
		{"120 > 121;", 120, ">", 121},
		{"0 < 10000;", 0, "<", 10000},
		{"9 == 9;", 9, "==", 9},
		{"12 != 11;", 12, "!=", 11},
		{"12 / 3;", 12, "/", 3},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}

	for _, test := range infixTests {
		lexer := lexer.New(test.input)
		parser := New(lexer)
		program := parser.ParseProgram()
		checkParseErrors(t, parser)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, len(program.Statements))
		}

		statement, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got='%T'", program.Statements[0])
		}

		exp, ok := statement.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp is not ast.InfixExpression. got='%T'", statement.Expression)
		}

		if exp.Operator != test.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", test.operator, exp.Operator)
		}

		if value, ok := test.rightValue.(bool); ok {
			if !testBooleanLiteral(t, exp.Right, value) {
				return
			}
		} else {
			if !testIntegerLiteral(t, exp.Right, test.rightValue.(int)) {
				return
			}
		}

		if value, ok := test.leftValue.(bool); ok {
			if !testBooleanLiteral(t, exp.Left, value) {
				return
			}
		} else {
			if !testIntegerLiteral(t, exp.Left, test.leftValue.(int)) {
				return
			}
		}

		if !testInfixExpression(t, statement.Expression, test.leftValue, test.operator, test.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == true",
			"((3 > 5) == true)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
		{
			"a + add(b * c) + d",
			"((a + add((b * c))) + d)",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
		},
		{
			"a * [1, 2, 3, 4][b * c] * d",
			"((a * ([1, 2, 3, 4][(b * c)])) * d)",
		},
		{
			"add(a * b[2], b[1], 2 * [1, 2][1])",
			"add((a * (b[2])), (b[1]), (2 * ([1, 2][1])))",
		},
	}

	for _, test := range tests {
		lexer := lexer.New(test.input)
		parser := New(lexer)
		program := parser.ParseProgram()
		checkParseErrors(t, parser)

		if program.String() != test.expected {
			t.Errorf("expected=%q, got=%q", test.expected, program.String())
		}
	}
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	identifier, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}

	if identifier.Value != value {
		t.Errorf("identifier.Value not %s. got='%s'", value, identifier.Value)
		return false
	}

	if identifier.TokenLiteral() != value {
		t.Errorf("identifier.TokenLiteral not %s. got='%s'", value, identifier.TokenLiteral())
		return false
	}

	return true
}

func testLiteralExpression(
	t *testing.T,
	exp ast.Expression,
	expected interface{},
) bool {
	switch value := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(value))
	case int64:
		return testIntegerLiteral(t, exp, value)
	case string:
		return testIdentifier(t, exp, value)
	case bool:
		return testBooleanLiteral(t, exp, value)
	}

	t.Errorf("type of exp not handled. got=%T", expected)
	return false
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {
	operatorExpression, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.InfixExpression. got=%T(%s)", exp, exp)
		return false
	}

	if !testLiteralExpression(t, operatorExpression.Left, left) {
		return false
	}

	if operatorExpression.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%q", operator, operatorExpression.Operator)
		return false
	}

	if !testLiteralExpression(t, operatorExpression.Right, right) {
		return false
	}

	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	bol, ok := exp.(*ast.Boolean)
	if !ok {
		t.Errorf("exp not *ast>Boolean. got=%T", exp)
		return false
	}

	if bol.Value != value {
		t.Errorf("bol.Value is not %v. got'%v'", value, bol.Value)
		return false
	}

	if bol.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("bol.TokenLiteral not %t. got=%s", value, bol.TokenLiteral())
		return false
	}

	return true
}

func TestIfExpression(t *testing.T) {
	input := `if (x < y) { x }`

	lexer := lexer.New(input)
	parser := New(lexer)
	program := parser.ParseProgram()
	checkParseErrors(t, parser)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statement. got=%d", 1, len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	exp, ok := statement.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("statement.Expression is not ast.IfExpression. got=%T", statement.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}

	if len(exp.Consequence.Statements) != 1 {
		t.Errorf("consequence is not %d statements. got=%d", 1, len(exp.Consequence.Statements))
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("exp.Consequence.Statements[0] is not ast.ExpressionStatement. got=%T", exp.Consequence.Statements[0])
	}

	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}

	if exp.Alternative != nil {
		t.Errorf("exp.Alternative.Statements was not nil. got=%+v", exp.Alternative)
	}
}

func TestIfElseExpression(t *testing.T) {
	input := `if (x < y) { x } else { y }`

	lexer := lexer.New(input)
	parser := New(lexer)
	program := parser.ParseProgram()
	checkParseErrors(t, parser)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statement. got=%d", 1, len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	exp, ok := statement.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("statement.Expression is not ast.IfExpression. got=%T", statement.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}

	if len(exp.Consequence.Statements) != 1 {
		t.Errorf("consequence is not %d statements. got=%d", 1, len(exp.Consequence.Statements))
	}

	if len(exp.Alternative.Statements) != 1 {
		t.Errorf("alternative is not %d statement. got=%d", 1, len(exp.Alternative.Statements))
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("exp.Consequence.Statements[0] is not ast.ExpressionStatement. got=%T", exp.Consequence.Statements[0])
	}

	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}

	alternative, ok := exp.Alternative.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("exp.Alternative.Statements[0] is not ast.ExpressionStatement. got=%T", exp.Alternative.Statements[0])
	}

	if !testIdentifier(t, alternative.Expression, "y") {
		return
	}
}

func TestFunctionLiteralParsing(t *testing.T) {
	input := `fun(x, y) { x + y; }`

	lexer := lexer.New(input)
	parser := New(lexer)
	program := parser.ParseProgram()
	checkParseErrors(t, parser)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	function, ok := statement.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("stmtement.Expression is not ast.FunctionLiteral. got=%T", statement.Expression)
	}

	if len(function.Parameters) != 2 {
		t.Fatalf("function literal parameters wrong. want %d, got=%d", 2, len(function.Parameters))
	}

	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "y")

	if len(function.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements has not %d statements. got=%d", 1, len(function.Body.Statements))
	}

	bodyStatement, ok := function.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("function body stmt is not ast.ExpressionStatement. got=%T", function.Body.Statements[0])
	}

	testInfixExpression(t, bodyStatement.Expression, "x", "+", "y")
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{input: "fun() {};", expectedParams: []string{}},
		{input: "fun(x) {};", expectedParams: []string{"x"}},
		{input: "fun(x, y, z) {};", expectedParams: []string{"x", "y", "z"}},
	}

	for _, test := range tests {
		lexer := lexer.New(test.input)
		parser := New(lexer)
		program := parser.ParseProgram()
		checkParseErrors(t, parser)

		statement := program.Statements[0].(*ast.ExpressionStatement)
		function := statement.Expression.(*ast.FunctionLiteral)

		if len(function.Parameters) != len(test.expectedParams) {
			t.Errorf("length parameters wrong. want %d, got=%d", len(test.expectedParams), len(function.Parameters))
		}

		for i, ident := range test.expectedParams {
			testLiteralExpression(t, function.Parameters[i], ident)
		}
	}
}

func TestCallExpressionParsing(t *testing.T) {
	input := `subtract(10 * 2, 12)`

	lexer := lexer.New(input)
	parser := New(lexer)
	program := parser.ParseProgram()
	checkParseErrors(t, parser)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	exp, ok := statement.Expression.(*ast.CallExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.CallExpression. got=%T", statement.Expression)
	}

	if !testIdentifier(t, exp.Function, "subtract") {
		return
	}

	if len(exp.Arguments) != 2 {
		t.Fatalf("wrong length of arguments. got=%d", len(exp.Arguments))
	}

	testInfixExpression(t, exp.Arguments[0], 10, "*", 2)
	testLiteralExpression(t, exp.Arguments[1], 12)
}

func TestMultipleDefStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"def x = 10;", "x", 10},
		{"def y = true;", "y", true},
		{"def foobar = y;", "foobar", "y"},
	}

	for _, test := range tests {
		lexer := lexer.New(test.input)
		parser := New(lexer)
		program := parser.ParseProgram()
		checkParseErrors(t, parser)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, len(program.Statements))
		}

		statement := program.Statements[0]
		if !testDefStatement(t, statement, test.expectedIdentifier) {
			return
		}

		value := statement.(*ast.DefStatement).Value
		if !testLiteralExpression(t, value, test.expectedValue) {
			return
		}
	}
}

func TestStringLiteralExpression(t *testing.T) {
	input := `"Hello Ali Ahmadi!"`

	lexer := lexer.New(input)
	parser := New(lexer)
	program := parser.ParseProgram()
	checkParseErrors(t, parser)

	statement := program.Statements[0].(*ast.ExpressionStatement)
	literal, ok := statement.Expression.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("exp not *ast.StringLiteral. got=%T", statement.Expression)
	}

	if literal.Value != "Hello Ali Ahmadi!" {
		t.Errorf("literal.Value not %q. got=%q", input, literal.Value)
	}
}

func TestParsingArrayLiterals(t *testing.T) {
	t.Parallel()
	input := `[1, 2 * 2, 12 / 1, 45]`

	lexer := lexer.New(input)
	parser := New(lexer)
	program := parser.ParseProgram()
	checkParseErrors(t, parser)

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	array, ok := statement.Expression.(*ast.ArrayLiteral)

	if !ok {
		t.Fatalf("exp not ast.ArrayLiteral. got=%T", statement.Expression)
	}

	if len(array.Elements) != 4 {
		t.Fatalf("len(array.Elements) not %d. got=%d", 4, len(array.Elements))
	}

	testIntegerLiteral(t, array.Elements[0], 1)
	testInfixExpression(t, array.Elements[1], 2, "*", 2)
	testInfixExpression(t, array.Elements[2], 12, "/", 1)
	testIntegerLiteral(t, array.Elements[3], 45)
}

func TestParsingIndexExpressions(t *testing.T) {
	input := `arr[1 + 2]`

	lexer := lexer.New(input)
	parser := New(lexer)
	program := parser.ParseProgram()
	checkParseErrors(t, parser)

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	indexExp, ok := statement.Expression.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("exp not *ast.IndexExpression. got=%T", statement.Expression)
	}

	if !testIdentifier(t, indexExp.Left, "arr") {
		return
	}

	if !testInfixExpression(t, indexExp.Index, 1, "+", 2) {
		return
	}
}
