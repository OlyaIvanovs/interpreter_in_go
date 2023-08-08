package parser

import (
	"testing"
	"strconv"
	
	"github.com/OlyaIvanovs/interpreter_in_go/lexer"
	"github.com/OlyaIvanovs/interpreter_in_go/ast"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	
	l := lexer.New(input)
	p := New(l)
	
	program := p.ParseProgram()	
	checkParseErrors(t, p)
	
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}
	
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	
	if len(errors) == 0 {
		return
	}
	
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parse error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral is not 'let', got=%q",s)
		return false
	}
	
	letSmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}
	
	if letSmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %s, got=%s", name, letSmt.Name.Value)
		return false
	}
	
	if letSmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not %s. got=%s", name, letSmt.Name)
		return false
	}
	
	return true
}

func TestReturnStatement(t *testing.T) {
	input := `
return 5; 
return 10; 
return 99932;
`
 l := lexer.New(input)
 p := New(l)
	
 program := p.ParseProgram()
 checkParseErrors(t, p)
	
 if len(program.Statements) != 3 {
 	t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
 }
 
 	for _, stmt := range program.Statements {
 		returnStmt, ok := stmt.(*ast.ReturnStatement)
 		if !ok {
 			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
 			continue
 		}
 		
 		if returnStmt.TokenLiteral() != "return" {
 			t.Errorf("returnStmt.Token is not 'return', got %q", returnStmt.TokenLiteral())
 		}
 	}
}

func TestBoolenExpression(t *testing.T) {
	input := "true"
	
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)
		
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements, got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	
	b, ok := stmt.Expression.(*ast.Boolean)
	if !ok {
		t.Fatalf("exp is not *ast.Boolean, got=%T", stmt.Expression)
	}
	
	if b.Value != true {
		t.Errorf("ident.Value not 'foobar'.got=%t", b.Value)
	}
	if b.TokenLiteral() != "true" {
		t.Errorf("b.TokenLiteral not bool. got=%s", b.TokenLiteral())
	}
}

func TestIndentifierExpression(t *testing.T) {
	input := "foobar;"
	
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)
		
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements, got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp is not *ast.Identifier, got=%T", stmt.Expression)
	}
	
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not 'foobar'.got=%s", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not foobar. got=%s", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"
	
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)
		
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements, got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp is not *ast.IntegerLiteral, got=%T", stmt.Expression)
	}
	
	if literal.Value != 5 {
		t.Errorf("literal.Value not '5'.got=%d", literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not 5. got=%s", literal.TokenLiteral())
	}
}


func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier, got=%T", exp)
		return false
	}
	
	if ident.Value != value {
		t.Errorf("ident.Value not %s, got=%s", value, ident.Value)
		return false
	}
	
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value, ident)
		return false
	}
	
	return true
} 

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
		case int:
			return testIntegerLiteral(t, exp, int64(v))
		case int64:
			return testIntegerLiteral(t, exp, v)
		case string:
			return testIdentifier(t, exp, v)
		case bool:
			return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("type of exp not handler. got=%T", exp)
	return false
}

func TestParsingPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input string
		operator string
		integerValue interface{}
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
		{"!true", "!", true},
		{"!false", "!", false},
	}
	
	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)
			
		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements, got=%d", len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}
		
		  exp, ok := stmt.Expression.(*ast.PrefixExpression)
		  if !ok {
		  	t.Fatalf("exp is not *ast.PrefixExpression, got=%T", stmt.Expression)
		  }
				
		  if exp.Operator != tt.operator {
		  	t.Errorf("exp.Operator not '%s'.got=%s", tt.operator, exp.Operator)
		  }
		  
		  if !testLiteralExpression(t, exp.Right, tt.integerValue) {
		  	return
		  }
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral, got=%T", il)
		return false
	}
	
	if integ.Value != value {
		t.Errorf("integ.Value not %d. got =%d", value, integ.Value)
		return false	
	}
	
	if integ.TokenLiteral() != strconv.FormatInt(value, 10) {
		t.Errorf("integ.TokenLiteral not %d. got =%s", value, integ.TokenLiteral())
		return false
	}
	
	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	bo, ok := exp.(*ast.Boolean)
	if !ok {
		t.Errorf("exp not *ast.Boolean. got=%T", exp)
		return false
	}
	
	if bo.Value != value {
		t.Errorf("bo.Value not %t. got=%t", value, bo.Value)
		return false
	}
	
	if bo.TokenLiteral() != strconv.FormatBool(value) {
		t.Errorf("bo.TokenLiteral not %t. got=%s", value, bo.TokenLiteral())
		return false
	}
	
	return true
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.InfixExpression, got=%T(%s)", exp, exp)
		return false
	}
	
	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}
	
	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
		return false
	}

	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}

	return true
}

func TestParsingInfixExpression(t *testing.T) {
	prefixTests := []struct {
		input string
		leftValue interface {}
		operator string
		rightValue interface {}
	}{
		{"5 + 5", 5, "+" ,5},
		{"5 - 5", 5, "-" ,5},
		{"5 * 5", 5, "*" ,5},
		{"5 / 5", 5, "/" ,5},
		{"5 > 5", 5, ">" ,5},
		{"5 < 5", 5, "<" ,5},
		{"5 == 5", 5,"==" ,5},
		{"5 != 5", 5, "!=" ,5},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}
	
	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)
			
		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements, got=%d", len(program.Statements))
		}
		
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}
		
		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp is not *ast.InfixExpression, got=%T", stmt.Expression)
		}
			
		if !testLiteralExpression(t, exp.Left, tt.leftValue) {
		  	return
		}
		  
		if exp.Operator != tt.operator {
			t.Errorf("exp.Operator not '%s'.got=%s", tt.operator, exp.Operator)
		}
		  
		if !testLiteralExpression(t, exp.Right, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct{
		input string
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
			"3 > 5 == false",	
			"((3 > 5) == false)",
		},	
		{
			"3 < 5 == true",	
			"((3 < 5) == true)",
		},	
	}
	
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)
		
		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q got=%q", tt.expected, actual)
		}
	}
}


