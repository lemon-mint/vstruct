package lexer

type Lexer struct {
	input []rune

	position int
	cursor   int

	line int
	col  int

	// The current char
	char rune
}

// NewLexer creates a new Lexer instance
func NewLexer(input []rune) *Lexer {
	l := &Lexer{
		input:    input,
		position: 0,
		cursor:   0,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() bool {
	if l.cursor >= len(l.input) {
		l.char = 0
		return false
	} else {
		l.char = l.input[l.cursor]
	}
	l.position = l.cursor
	l.cursor++
	l.col++
	return true
}

func (l *Lexer) skipWhitespace() bool {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		if l.char == '\n' {
			l.line++
			l.col = 0
		}

		if !l.readChar() {
			return false
		}
	}
	return true
}

func (l *Lexer) isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for l.isLetter(l.char) {
		if !l.readChar() {
			break
		}
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) NewToken(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: string(l.char),
		File:    "",
		Line:    l.line,
		Col:     l.col,
	}
}

func (l *Lexer) NextToken() Token {
	var t Token

	l.skipWhitespace()

	switch l.char {
	case '{':
		t = l.NewToken(TOKEN_OPEN_BRACE, "{")
	case '}':
		t = l.NewToken(TOKEN_CLOSE_BRACE, "}")
	case ',':
		t = l.NewToken(TOKEN_COMMA, ",")
	case ';':
		t = l.NewToken(TOKEN_SEMICOLON, ";")

	default:
		if l.isLetter(l.char) {
			ident := l.readIdentifier()
			t = l.NewToken(LookupKeyword(ident), ident)
		} else {
			t = l.NewToken(TOKEN_UNKNOWN, string(l.char))
		}
	}

	return t
}
