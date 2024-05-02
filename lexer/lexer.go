package lexer

import (
	"github.com/haitammabrouk/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// read the current char under examination
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// we initialize a token based on the the current char
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=' :
		tok = newToken(token.ASSIGN, l.ch)
	case '+' :
		tok = newToken(token.PLUS, l.ch)
	case '(' :
		tok = newToken(token.LPAREN, l.ch)
	case ')' :
		tok = newToken(token.RPAREN, l.ch)
	case ';' :
		tok = newToken(token.SEMICOLON, l.ch)
	case ',' :
		tok = newToken(token.COMMA, l.ch)
	case '{' :
		tok = newToken(token.LBRACE, l.ch)
	case '}' :
		tok = newToken(token.RBRACE, l.ch)
	case 0 :
		tok.Literal = ""
		tok.Type = token.EOF
	default :
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Read the rest of the identifier or keyword until we encounter a non letter character
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// check if the current char is a letter
func isLetter(ch byte) bool {
	return 'a' <= ch && ch >= 'z' || 'A' <= ch && ch >= 'Z' || ch == '_'
}

// Read the rest of the number and turn into token
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// check if the current char is a digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// whitespaces are considered just delimiters so they dont have any meaning
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}