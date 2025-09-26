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

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 //ASCII code for "NUL"
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
}