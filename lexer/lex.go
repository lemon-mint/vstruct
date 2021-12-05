package lexer

type Lexer struct {
	Filename string

	input []rune

	position int
	cursor   int

	line int
	col  int

	// The current char
	char rune

	lastToken Token
}

// NewLexer creates a new Lexer instance
func NewLexer(input []rune, filename string) *Lexer {
	l := &Lexer{
		Filename: filename,
		input:    input,
		position: 0,
		cursor:   0,

		line: 1,
		col:  0,
	}
	l.ReadChar()
	return l
}

func (l *Lexer) ReadChar() bool {
	if l.cursor >= len(l.input) {
		l.char = 0
		l.lastToken = l.NewToken(TOKEN_EOF, "")
		return false
	} else {
		l.char = l.input[l.cursor]
		if l.char == '\n' {
			l.line++
			l.col = 0
		}
	}
	l.position = l.cursor
	l.cursor++
	l.col++
	return true
}

func (l *Lexer) skipWhitespace() bool {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		if !l.ReadChar() {
			return false
		}
	}
	return true
}

func (l *Lexer) isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_' || c >= '0' && c <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for l.isLetter(l.char) {
		if !l.ReadChar() {
			break
		}
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) NewToken(tokenType TokenType, literal string) Token {
	return Token{
		File:    l.Filename,
		Type:    tokenType,
		Literal: literal,
		Line:    l.line,
		Col:     l.col,
	}
}

func (l *Lexer) NextToken() Token {
	var t Token

	if !l.skipWhitespace() {
		l.lastToken = l.NewToken(TOKEN_EOF, "")
		return l.lastToken
	}
	if l.lastToken.Type == TOKEN_EOF {
		return l.lastToken
	}

	switch l.char {
	case '{':
		t = l.NewToken(TOKEN_OPEN_BRACE, "{")
		l.ReadChar()
	case '}':
		t = l.NewToken(TOKEN_CLOSE_BRACE, "}")
		l.ReadChar()
	case ',':
		t = l.NewToken(TOKEN_COMMA, ",")
		l.ReadChar()
	case ';':
		t = l.NewToken(TOKEN_SEMICOLON, ";")
		l.ReadChar()
	case '=':
		t = l.NewToken(TOKEN_EQUAL, "=")
		l.ReadChar()
	default:
		if l.isLetter(l.char) {
			ident := l.readIdentifier()
			t = l.NewToken(LookupKeyword(ident), ident)
		} else {
			t = l.NewToken(TOKEN_UNKNOWN, string(l.char))
		}
	}
	l.lastToken = t
	return t
}
