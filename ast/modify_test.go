package ast

import (
	"reflect"
	"testing"
)

func turnOneIntoTwo(node Node) Node {
	integer, ok := node.(*IntegerLiteral)
	if !ok {
		return node
	}

	if integer.Value != 1 {
		return node
	}

	integer.Value = 2
	return integer
}

func one() Expression { return &IntegerLiteral{Value: 1} }
func two() Expression { return &IntegerLiteral{Value: 2} }

func TestModify(t *testing.T) {

	tests := []struct {
		input    Node
		expected Node
	}{
		{one(), two()},
		{
			&Program{
				Statements: []Statement{
					&ExpressionStatement{Expression: one()},
				},
			},
			&Program{
				Statements: []Statement{
					&ExpressionStatement{Expression: two()},
				},
			},
		},
		{
			&InfixExpression{Left: one(), Operator: "+", Right: two()},
			&InfixExpression{Left: two(), Operator: "+", Right: two()},
		},
		{
			&InfixExpression{Left: two(), Operator: "+", Right: one()},
			&InfixExpression{Left: two(), Operator: "+", Right: two()},
		},
		{
			&PrefixExpression{Operator: "-", Right: one()},
			&PrefixExpression{Operator: "-", Right: two()},
		},
		{
			&IndexExpression{Left: one(), Index: one()},
			&IndexExpression{Left: two(), Index: two()},
		},
		{
			&IfExpression{
				Condition: one(),
				Consequence: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: one()},
					},
				},
				Alternative: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: one()},
					},
				},
			},
			&IfExpression{
				Condition: two(),
				Consequence: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: two()},
					},
				},
				Alternative: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: two()},
					},
				},
			},
		},
		{
			&ReturnStatement{ReturnValue: one()},
			&ReturnStatement{ReturnValue: two()},
		},
		{
			&LetStatement{Value: one()},
			&LetStatement{Value: two()},
		},
		{
			&FunctionLiteral{
				Parameters: []*Identifier{},
				Body: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: one()},
					},
				},
			},
			&FunctionLiteral{
				Parameters: []*Identifier{},
				Body: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: two()},
					},
				},
			},
		},
		{
			&ArrayLiteral{Elements: []Expression{one(), two()}},
			&ArrayLiteral{Elements: []Expression{two(), two()}},
		},
	}

	for _, tt := range tests {
		modified := Modify(tt.input, turnOneIntoTwo)

		equal := reflect.DeepEqual(modified, tt.expected)

		if !equal {
			t.Errorf("not equal. got=%#v, want=%#v", modified, tt.expected)
		}
	}
}

func TestModifyHash(t *testing.T) {
	hashLiteral := &HashLiteral{
		Pairs: map[Expression]Expression{
			one(): one(),
			one(): one(),
		},
	}

	Modify(hashLiteral, turnOneIntoTwo)

	for key, val := range hashLiteral.Pairs {
		key, _ := key.(*IntegerLiteral)
		if key.Value != 2 {
			t.Errorf("value is not %d, got=%d", 2, key.Value)
		}
		val, _ := val.(*IntegerLiteral)
		if val.Value != 2 {
			t.Errorf("value is not %d, got=%d", 2, val.Value)
		}
	}
}
