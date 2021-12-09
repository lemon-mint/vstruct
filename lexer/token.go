package lexer

import "fmt"

type TokenType int

func (t TokenType) String() string {
	return tokenNames[t]
}

const (
	TOKEN_UNKNOWN TokenType = iota
	TOKEN_EOF
	TOKEN_RESERVED
	// Keywords
	TOKEN_STRUCT
	TOKEN_ENUM
	TOKEN_ALIAS
	// Symbols
	TOKEN_OPEN_BRACE
	TOKEN_CLOSE_BRACE
	// Separators
	TOKEN_COMMA
	TOKEN_SEMICOLON
	// Operators
	TOKEN_EQUAL
	// Literals
	TOKEN_IDENTIFIER

	// built-in types
	TOKEN_RAWTYPE
)

var tokenNames = map[TokenType]string{
	TOKEN_UNKNOWN:     "UNKNOWN",
	TOKEN_EOF:         "EOF",
	TOKEN_STRUCT:      "STRUCT",
	TOKEN_ENUM:        "ENUM",
	TOKEN_ALIAS:       "ALIAS",
	TOKEN_OPEN_BRACE:  "OPEN_BRACE",
	TOKEN_CLOSE_BRACE: "CLOSE_BRACE",
	TOKEN_COMMA:       "COMMA",
	TOKEN_SEMICOLON:   "SEMICOLON",
	TOKEN_EQUAL:       "EQUAL",
	TOKEN_IDENTIFIER:  "TOKEN_IDENTIFIER",
	TOKEN_RAWTYPE:     "RAWTYPE",
}

type Token struct {
	Type    TokenType
	Literal string

	File string
	Line int
	Col  int
}

func (t Token) String() string {
	return fmt.Sprintf("Token<%s>[%s:%d:%d: %s]", tokenNames[t.Type], t.File, t.Line, t.Col, t.Literal)
}
