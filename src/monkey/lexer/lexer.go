package lexer

import (
	"fmt"
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


	l.skipWhitespace()

	switch l.ch {

	case '=' :
	
		if l.peekChar() == '=' {

			fmt.Println("this is an equal")
			char := l.ch

			l.readChar()
			literal := string(char) + string(l.ch)
			tok.Literal = literal
			tok.Type = token.EQUAL
			
		} else {
			tok = token.NewToken(token.ASSIGN, l.ch)
		}


	case '+' :
		tok = token.NewToken(token.PLUS, l.ch)

	case '(':

		 tok = token.NewToken(token.LPAREN, l.ch)

	case ')':

		tok = token.NewToken(token.RPAREN, l.ch)
	
	case '{':

		tok = token.NewToken(token.LBRACE, l.ch)

	case '}':

		tok = token.NewToken(token.RBRACE,l.ch)
	
	case ';':

		tok = token.NewToken(token.SEMICOLON, l.ch)
	
	case ',':
		tok = token.NewToken(token.COMMA,l.ch)

	case '-':
		tok = token.NewToken(token.MINUS, l.ch)

	case '*':
		tok = token.NewToken(token.ASTERISK, l.ch)

	case '/':
		tok = token.NewToken(token.SLASH, l.ch)

	case '!':
		if l.input[l.readPosition] == '=' {

			literal := string(l.ch) + string(l.input[l.readPosition])
			tok.Type = token.NOT_EQUAL
			tok.Literal = literal
			l.readChar()
		} else {
			tok = token.NewToken(token.BANG, l.ch)
		}

	case '<':

		tok = token.NewToken(token.LT, l.ch)

	case '>': 
		tok = token.NewToken(token.GT, l.ch)

	case 0:

		tok.Type = token.EOF
		tok.Literal = ""

	// If the current char is not one of the symbols, check if its a letter
	default: 
	 if isLetter(l.ch){
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)
		//Return here since we already advanced in the input when identifying the identifier
		return tok
	  } else if isDigit(l.ch) {
		tok.Type = token.INT
		tok.Literal = l.readNumber()
		return  tok
	  } else {
		tok = token.NewToken(token.ILLEGAL, l.ch)
	  }

	}

	l.readChar()
	return tok
}


func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(ch byte) bool {
return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
 position := l.position

 for isLetter(l.ch) {
	l.readChar()
 }

 return l.input[position: l.position]
}


func (l *Lexer) readNumber() string {
	position:= l.position

	for isDigit(l.ch){
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
	l.readChar()
}
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

func (l *Lexer) peekChar() byte {

	if l.readPosition >= len(l.input) {
		return 0
	} 
	return l.input[l.readPosition]
}