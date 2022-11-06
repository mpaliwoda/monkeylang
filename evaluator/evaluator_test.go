package evaluator

import (
	"testing"

	"github.com/mpaliwoda/monkeylang/lexer"
	"github.com/mpaliwoda/monkeylang/object"
	"github.com/mpaliwoda/monkeylang/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		// just regular stuff
		{"5", 5},
		{"10", 10},

		// prefix
		{"-5", -5},
		{"-10", -10},

		// infix
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

	for _, testCase := range tests {
		evaluated := testEval(testCase.input)
		testIntegerObject(t, evaluated, testCase.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		// raw
		{"true", true},
		{"false", false},

		// integer comparisons
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},

		// boolean comparisons
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},

		// mixed comparisons
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
	}

	for _, testCase := range tests {
		evaluated := testEval(testCase.input)
		testBooleanObject(t, evaluated, testCase.expected)
	}
}

func TestEvalBangOperator(t *testing.T) {
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

	for _, testCase := range tests {
		evaluated := testEval(testCase.input)
		testBooleanObject(t, evaluated, testCase.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)

	if !ok {
		t.Errorf("obj is not an Integer, got %T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("result.Value expected to be %d, got %d", expected, result.Value)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)

	if !ok {
		t.Errorf("obj is not an Boolean, got %T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("result.Value expected to be %t, got %t", expected, result.Value)
		return false
	}

	return true
}
