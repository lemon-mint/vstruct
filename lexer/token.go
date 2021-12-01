package lexer

import "fmt"

type TokenType int

const (
	TOKEN_UNKNOWN TokenType = iota
	TOKEN_EOF
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

	// Types
	TOKEN_UINT8
	TOKEN_UINT16
	TOKEN_UINT32
	TOKEN_UINT64
	TOKEN_INT8
	TOKEN_INT16
	TOKEN_INT32
	TOKEN_INT64
	TOKEN_FLOAT32
	TOKEN_FLOAT64
	TOKEN_BYTES
	TOKEN_STRING
	TOKEN_BOOL
	// Literals
	TOKEN_IDENTIFIER
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
	TOKEN_UINT8:       "UINT8",
	TOKEN_UINT16:      "UINT16",
	TOKEN_UINT32:      "UINT32",
	TOKEN_UINT64:      "UINT64",
	TOKEN_INT8:        "INT8",
	TOKEN_INT16:       "INT16",
	TOKEN_INT32:       "INT32",
	TOKEN_INT64:       "INT64",
	TOKEN_FLOAT32:     "FLOAT32",
	TOKEN_FLOAT64:     "FLOAT64",
	TOKEN_BYTES:       "BYTES",
	TOKEN_STRING:      "STRING",
	TOKEN_BOOL:        "BOOL",
	TOKEN_IDENTIFIER:  "TOKEN_IDENTIFIER",
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
