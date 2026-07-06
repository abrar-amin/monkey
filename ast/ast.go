package ast 

import "monkey/token"


// More or less OOP subclasses 
type Node interface {
	TokenLiteral() string
}


type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}


// Program is the root of every AST in a Monkey program .
type Program struct{
	Statements []Statement
}

func(p *Program) tokenLiteral() {
	if(len(p.statements) > 0)
	{
		return p.Statements[0].TokenLiteral()
	}
	else{
		return ""
	}
}

type LetStatement struct{
	Token token.Token // token.LET token
	Name *Identifier
	Value Expression 
}


//Empty bodies are known as marker methods.
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {return ls.Token.Literal}

type Identifier struct{
	Token token.Token //token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {return i.Token.Literal}