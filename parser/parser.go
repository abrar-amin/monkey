package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct{
	l *lexer.lexer

	curToken token.Token
	peektoken token.Token

}

func New(l *lexer.Lexer) *Parser{

	//Shorthand for creating a new parser, storing a reference in p.
	p := &Parser{l : l}

	//Make sure curToken and peekToken have values
	p.nextToken()
	p.nextToken()


	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken

	//Use the lexer to get the next token
	p.peekToken = p.l.nextToken()
}