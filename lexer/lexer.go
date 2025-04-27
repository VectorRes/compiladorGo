package lexer

import (
	"mi-proyecto/token"
)

type Lexer struct {
	source       string
	character    byte
	position     int
	readPosition int
}

func New(input string) *Lexer {
	l := &Lexer{source: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.character {
	case '=':
		if l.peekChar() == '=' {
			ch := l.character
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.character)}
		} else {
			tok = token.Token{Type: token.ASSIGN, Literal: string(l.character)}
		}
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: string(l.character)}
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		if isLetter(l.character) {
			literal := l.readLiteral()
			tokType := token.LookupTokenType(literal)
			return token.Token{Type: tokType, Literal: literal}
		} else if isDigit(l.character) {
			literal := l.readNumber()
			tok = token.Token{Type: token.INT, Literal: literal}
		} else {
			tok = token.Token{Type: token.ILLEGAL, Literal: string(l.character)}
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.character == ' ' || l.character == '\t' || l.character == '\n' || l.character == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.source) {
		l.character = 0
	} else {
		l.character = l.source[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.source) {
		return 0
	}
	return l.source[l.readPosition]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.character) {
		l.readChar()
	}
	return l.source[position:l.position]
}

func (l *Lexer) readLiteral() string {
	position := l.position
	for isLetter(l.character) || isDigit(l.character) {
		l.readChar()
	}
	return l.source[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
