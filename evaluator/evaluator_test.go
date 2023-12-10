package evaluator

import (
	"Ahmadi/lexer"
	"Ahmadi/object"
	"Ahmadi/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"120", 120},
		{"-2", -2},
		{"-39", -39},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		testIntegerObject(t, evaluated, test.expected)
	}
}

func testEval(input string) object.Object {
	lexer := lexer.New(input)
	parser := parser.New(lexer)
	program := parser.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

func TestEvalBooleanExpression(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		testBoolObject(t, evaluated, test.expected)
	}
}

func testBoolObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}

	return true
}

func TestBangOperator(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		testBoolObject(t, evaluated, test.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", "if (true) { 10 }", 10},
		{"if (false) { 10 }", "if (false) { 10 }", nil},
		{"if (1) { 10 }", "if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", "if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", "if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", "if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", "if (1 < 2) { 10 } else { 20 }", 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			evaluated := testEval(test.input)
			integer, ok := test.expected.(int)

			if ok {
				testIntegerObject(t, evaluated, int64(integer))
			} else {
				testNullObject(t, evaluated)
			}
		})
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    string
		expected int64
	}{
		{"return 10;", "return 10;", 10},
		{"return 10; 9;", "return 10; 9;", 10},
		{"return 2 * 5; 9;", "return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", "9; return 2 * 5; 9;", 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			evaluated := testEval(test.input)
			testIntegerObject(t, evaluated, test.expected)
		})
	}
}
