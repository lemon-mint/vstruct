package lexer

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

	// user-defined types
	TOKEN_USER_TYPE
)

type Token struct {
	Type    TokenType
	Literal string

	File string
	Line int
	Col  int
}
