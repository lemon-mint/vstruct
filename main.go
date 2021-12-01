package main

import (
	"io"
	"log"
	"os"

	"github.com/lemon-mint/vstruct/lexer"
)

func ReadFileAsString(fileName string) string {
	os.Open(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func main() {
	input := ReadFileAsString("test.vstruct")
	lex := lexer.NewLexer([]rune(input))
TokenizeLoop:
	for {
		token := lex.NextToken()
		log.Printf("%v", token)
		if token.Type == lexer.TOKEN_UNKNOWN || token.Type == lexer.TOKEN_EOF {
			switch token.Type {
			case lexer.TOKEN_UNKNOWN:
				log.Printf("Unknown token: %v", token)
			case lexer.TOKEN_EOF:
				break TokenizeLoop
			}
		}
	}
}
