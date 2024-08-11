package parser

import (
	"testing"
	"tarbuj/ast"
	"tarbuj/lexer"
)


func TestLetStatements(t *testing.T) {
	var input = `
let x = 5;
let y = 10;
let number = 56;
	`

	var l = lexer.New(input)
	var p = New(l)

	var program = p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	var tests = []struct{
		expectedIdentifier string
	} {
		{"x"},
		{"y"},
		{"number"},
	}

	for i, tt := range tests {
		var stmt = program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	var letStmt, ok = s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmtName.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	return true
}

fucn checkParserErrors(t *testing.T, p *Parser) {
	var errors = p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
