package token


const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT" // 1343456
	// Operators
	ASSIGN = "="
	PLUS = "+"
	MINUS= "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	LT="<"
	GT=">"
	EQUAL = "=="
	NOT_EQUAL = "!="
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
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
	TRUE = "TRUE"

	)

type TokenType string


type Token struct {
	Type    TokenType
	Literal string
}

var keywords =  map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"false": FALSE,
	"return": RETURN,
	"true": TRUE,
	"if": IF,
	"else": ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok:= keywords[ident]; ok {
		return tok
	}
	return IDENT
}
func NewToken(tokenType TokenType, literal byte) Token{

	return Token{
		Type: tokenType,
		Literal: string(literal),
	}

}