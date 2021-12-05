package lexer

var keywords = map[string]TokenType{
	"struct": TOKEN_STRUCT,
	"enum":   TOKEN_ENUM,
	"alias":  TOKEN_ALIAS,
}

var reserved = map[string]TokenType{
	"import": TOKEN_RESERVED,
	"as":     TOKEN_RESERVED,
	"rpc":    TOKEN_RESERVED,
	"func":   TOKEN_RESERVED,

	"union": TOKEN_RESERVED,
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
