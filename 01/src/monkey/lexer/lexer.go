package lexer

import token "monkey/token"

type Lexer struct {
	input        string
	position     int  //current char's positon
	readPosition int  //current reading position(position + 1)
	ch           byte //curretn char
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	l.readChar()
	return l
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 //ASCII code for "NUL"
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	l.eatWhitespace()

	var t token.Token

	switch l.ch {
	case '=':
		t = newToken(token.ASSIGN, l.ch)
	case '+':
		t = newToken(token.PLUS, l.ch)
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case '{':
		t = newToken(token.LBRACE, l.ch)
	case '}':
		t = newToken(token.RBRACE, l.ch)
	case ';':
		t = newToken(token.SEMICOLON, l.ch)
	case ',':
		t = newToken(token.COMMA, l.ch)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isDigit(l.ch) {
			t.Literal = l.readNumber()
			t.Type = token.INT
			return t
		} else if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)
			return t
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return t

}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
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

func (l *Lexer) readIdentifier() string {
	startPos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && '9' >= ch
}

func isLetter(ch byte) bool {
	return ('a' <= ch && 'z' >= ch) || ('A' <= ch && 'Z' >= ch)
}
