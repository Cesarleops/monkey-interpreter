package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input string
	position int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	ch byte // current char under examination
}
	
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}


func (l *Lexer) NextToken() token.Token {

	var tok token.Token

	switch l.ch {
	case '=' :

		tok = *token.NewToken(token.ASSIGN, "=")

	case '+' :
		tok = *token.NewToken(token.PLUS, "+")

	case '(':

		 tok = *token.NewToken(token.LPAREN, "(")

	case ')':

		tok = *token.NewToken(token.RPAREN, ")")
	
	case '{':

		tok = *token.NewToken(token.LBRACE, "{")

	case '}':

		tok = *token.NewToken(token.RBRACE, "}")
	
	case ';':

		tok = *token.NewToken(token.SEMICOLON, ";")
	
	case 0:

		tok = *token.NewToken(token.EOF , "")

	case ',':

		tok = *token.NewToken(token.COMMA, ",")

	}

	l.readChar()
	return tok
}


func (l *Lexer) readChar(){

	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition+=1

}