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

	return Eval(program, object.NewEnvironment())
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
		{
			"nested if",
			`
			if (10 > 1) {
				if(10 > 2) {
					return 10;
				}

				return 11;
			}

			return 12;
			`,
			10,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			evaluated := testEval(test.input)
			testIntegerObject(t, evaluated, test.expected)
		})
	}
}

func TestErrorHandling(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
			if (10 > 1) {
			if (10 > 1) {
			return true + false;
			}
			return 1;
			}
			`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"identifier not found: foobar",
		},
		{
			`"Hello" - "World"`,
			"unknown operator: STRING - STRING",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			evaluated := testEval(test.input)

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("no error object returned. got=%T", evaluated)
				return
			}

			if errObj.Message != test.expectedMessage {
				t.Errorf("wrong error message. expected=%q, got=%q", test.expectedMessage, errObj.Message)
				return
			}
		})
	}
}

func TestDefStatements(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected int64
	}{
		{"def a = 5; a;", 5},
		{"def a = 5 * 5; a;", 25},
		{"def a = 5; def b = a; b;", 5},
		{"def a = 5; def b = a; def c = a + b + 5; c;", 15},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			testIntegerObject(t, testEval(test.input), test.expected)
		})
	}
}

func TestFunctionObject(t *testing.T) {
	t.Parallel()
	input := "fun(x) {x + 2; };"

	evaluated := testEval(input)
	fun, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T", evaluated)
	}

	if len(fun.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v", fun.Parameters)
	}

	if fun.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fun.Parameters[0])
	}

	expectedBody := "(x + 2)"
	if fun.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got %q", expectedBody, fun.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected int64
	}{
		{"def identity = fun(x) { x; }; identity(5);", 5},
		{"def identity = fun(x) { return x; }; identity(5);", 5},
		{"def double = fun(x) { x * 2; }; double(5);", 10},
		{"def add = fun(x, y) { x + y; }; add(5, 5);", 10},
		{"def add = fun(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fun(x) { x; }(5)", 5},
	}

	for _, test := range tests {
		t.Run(test.input[:10]+"...", func(t *testing.T) {
			testIntegerObject(t, testEval(test.input), test.expected)
		})
	}
}

func TestClosures(t *testing.T) {
	t.Parallel()

	input := `
	def adder = fun(x) {
		fun(y) {
			x + y;
		};
	};

	def t = adder(2);
	t(3);
	`

	testIntegerObject(t, testEval(input), 5)
}

func TestStringLiteral(t *testing.T) {
	input := `"Ali Ahmadi!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T", evaluated)
	}

	if str.Value != "Ali Ahmadi!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	t.Parallel()

	input := `"Ali" + " " + "Ahmadi"`
	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T", evaluated)
	}

	if str.Value != "Ali Ahmadi" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestBuiltinFunctions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{`len(1)`, "argument to `len` not supported, got INTEGER"},
		{`len("one", "two")`, "wrong number of arguments. got=2, want=1"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			evaluated := testEval(test.input)

			switch expected := test.expected.(type) {
			case int:
				testIntegerObject(t, evaluated, int64(expected))

			case string:
				errObj, ok := evaluated.(*object.Error)
				if !ok {
					t.Errorf("object is not Error. got=%T", evaluated)
					return
				}

				if errObj.Message != expected {
					t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Message)
				}
			}
		})
	}
}

func TestArrayLiterals(t *testing.T) {
	t.Parallel()
	input := `[1 * 1, 2 * 2, 3 * 3]`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T", evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong number of elements. got=%d", len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 9)
}

func TestArrayIndexExpressions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"[1, 2, 3][0]",
			1,
		},
		{
			"[1, 2, 3][1]",
			2,
		},
		{
			"[1, 2, 3][2]",
			3,
		},
		{
			"def i = 0; [1][i];",
			1,
		},
		{
			"[1, 2, 3][1 + 1];",
			3,
		},
		{
			"def arr = [1, 2, 3]; arr[2];",
			3,
		},
		{
			"def arr = [1, 2, 3]; arr[0] + arr[1] + arr[2];",
			6,
		},
		{
			"def arr = [1, 2, 3]; def i = arr[0]; arr[i]",
			2,
		},
		{
			"[1, 2, 3][3]",
			nil,
		},
		{
			"[1, 2, 3][-1]",
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
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
