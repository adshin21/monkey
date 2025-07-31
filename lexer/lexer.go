// Package lexer takes source code and outputs tokens
package lexer

import (
	"github.com/adshin21/monkey/token"
)

type Lexer struct {
	input        string
	position     int  // Current character position (where l.ch points)
	readPosition int  // Next character to read (lookahead)
	ch           byte // Current character being examined
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOTEQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	startPos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

func (l *Lexer) skipWhitespace() {
	for isWhiteSpace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	startPos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
