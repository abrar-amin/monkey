
package parser

import (
	"testing"
	"monkey/ast"
	"monkey/lexer"
)

func TestLetStatements(t *testing.T){
	input := `
let x = 67;
let y = 65;
let multichar = 8343434343;	
`

	l := lexer.New(input)
	//Scoping rules
	p := New(l)

	program := p.ParseProgram()
	if program == nil{
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. instead got: %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string 
	}{
		{"x"},
		{"y"},
		{"multichar"},
	}

	// go syntax is very interesting...
	for i, tt := range tests{

		//program.Statements must correspond to the testing order
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier){
			//error printing is handled by testLetStatement
			return
		}
	}


}

func testLetStatement(t *testing.T, s ast.Statement,name string) bool{
	if s.TokenLiteral() != "let"{
		t.Errorf("s.TokenLiteral not 'let'. It was %q", s.TokenLiteral())
		return false 
	}

	// this is a Go type assertion -> checks to see if s is a *ast.LetStatement
	letStmt, ok := s.(*ast.LetStatement)
	if !ok{
		t.Errorf("S not *ast.LetStatement. got %T", s)
		return false 
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false 
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}
	return true
}