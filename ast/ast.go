package ast 

import ("monkey/token", "bytes")


// More or less OOP subclasses 
type Node interface {
	TokenLiteral() string
	String() string
}


type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

//essentially jsut an expression wrapper for cases like `x + 5;`
type ExpressionStatement interface{
	Token token.Token 
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// Program is the root of every AST in a Monkey program .
type Program struct{
	Statements []Statement
}

func (p *Program) String() string {
var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}



func (ls *LetStatement) String() string {
var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}


func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}


func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func(p *Program) tokenLiteral() string {
	if(len(p.Statements) > 0){
		return p.Statements[0].TokenLiteral()
	}else{
		return ""
	}
}

// To satisfy the Expression interface.
func (i *Identifier) String() string { return i.Value }

type LetStatement struct{
	Token token.Token // token.LET token
	Name *Identifier
	Value Expression 
}


//Empty bodies are known as marker methods.
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {return ls.Token.Literal}

// For variable naemes
type Identifier struct{
	Token token.Token //token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {return i.Token.Literal}