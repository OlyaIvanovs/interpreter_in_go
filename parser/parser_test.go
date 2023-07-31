package parser

import (
	"testing"
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
	p := New(l);
	
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram returns nil")
	}
	
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