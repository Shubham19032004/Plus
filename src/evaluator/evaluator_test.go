package evaluator

import (
	"testing"

	"github.com/Shubham19032004/plus/src/lexer"
	"github.com/Shubham19032004/plus/src/object"
	"github.com/Shubham19032004/plus/src/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5+5+5+5-10", 10},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}
func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1<2", true},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!!true", true},
		{"!!5", true},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}
func TestIfElseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if(true){10}", 10},
		{"if(false){10}", nil},
		{"if(1){10}", 10},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}
func TestReturnStatements(t *testing.T){
	tests:=[]struct{
		input string
		expected int64
	}{
		{"return 10;",10},
		{"return 10; 9",10},
	}
	for _,tt :=range tests{
		evaluated:=testEval(tt.input)
		testIntegerObject(t,evaluated,tt.expected)
	}
}
func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not Null. got=%T (%+v)", obj, obj)
		return false
	}
	return true

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
		t.Errorf("object is not Integer. got=%T (%+V)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value.got=%d,want=%d", result.Value, expected)
		return false
	}
	return true
}
func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got =%T (%+V)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has worng value. got=%t,want=%t", result.Value, expected)
		return false
	}
	return true
}
