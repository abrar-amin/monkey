package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct{
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
	error []string

}

func New(l *lexer.Lexer) *Parser{

	//Shorthand for creating a new parser, storing a reference in p.
	p := &Parser{l : l, errors : []string{}}
	//Make sure curToken and peekToken have values
	p.nextToken()
	p.nextToken()


	return p
}

func (p *Parser) peekError(t token.TokenType){
	msg := fmt.Sprintf("expected next token has to be %s, got %s instead", t,p.peekToken.Type)
	p.errors = append(p.errors,msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken

	//Use the lexer to get the next token like in the 
	//case with an arithmetic expression
	p.peekToken = p.l.NextToken()
}


//The recursive part (recursive descent)
func (p *Parser) ParseProgram() *ast.Program{
	//This is the top level node
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		//Recursive part
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}	
		p.nextToken()
	}
	return program
}


func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
		case token.LET:
			return p.parseLetStatement()

		//TODO: implement other statements
		default:
			return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement{
	stmt := &ast.LetStatement{Token: p.curToken}

	//the variable should have a name
	if !p.expectPeek(token.IDENT){
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Must be followed by an equal sign; the expect methods here provide a way to enforce types  
	if !p.expectPeek(token.ASSIGN){
		return nil
	}

	// TODO: Expressions
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}