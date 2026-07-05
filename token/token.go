package token 

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT" // 1343456
	// Operators
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
	
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"

	LT = "<"
	GT = ">"
	IF = "if"
	ELSE = "else"
	RETURN = "return"
	TRUE = "true"
	FALSE = "false"
	EQ = "=="
	NOT_EQ = "!="
)

var keywords = map[string]TokenType{
	"fn" : FUNCTION,
	"let" : LET,
	"true" : TRUE,
	"false" : FALSE,
	"if" : IF,
	"else" : ELSE,
	"return" : RETURN,
}

func LookupIdent(ident string) TokenType{
	if tok, ok := keywords[ident]; ok{
		return tok
	}
	return IDENT
}