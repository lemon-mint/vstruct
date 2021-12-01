package lexer

var keywords = map[string]TokenType{
	"struct": TOKEN_STRUCT,
	"enum":   TOKEN_ENUM,
	"alias":  TOKEN_ALIAS,
}

var reserved = map[string]TokenType{
	"uint8":   TOKEN_UINT8,
	"uint16":  TOKEN_UINT16,
	"uint32":  TOKEN_UINT32,
	"uint64":  TOKEN_UINT64,
	"int8":    TOKEN_INT8,
	"int16":   TOKEN_INT16,
	"int32":   TOKEN_INT32,
	"int64":   TOKEN_INT64,
	"float32": TOKEN_FLOAT32,
	"float64": TOKEN_FLOAT64,
	"bytes":   TOKEN_BYTES,
	"string":  TOKEN_STRING,
	"bool":    TOKEN_BOOL,
}

func LookupKeyword(s string) TokenType {
	if t, ok := keywords[s]; ok {
		return t
	}
	if t, ok := reserved[s]; ok {
		return t
	}
	return TOKEN_IDENTIFIER
}
