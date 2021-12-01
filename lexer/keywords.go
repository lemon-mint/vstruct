package lexer

var keywords = map[string]TokenType{
	"struct": TOKEN_STRUCT,
	"enum":   TOKEN_ENUM,
	"alias":  TOKEN_ALIAS,
}

func LookupKeyword(s string) TokenType {
	if t, ok := keywords[s]; ok {
		return t
	}
	return TOKEN_IDENTIFIER
}
