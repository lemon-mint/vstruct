package main

import (
	"io"
	"log"
	"os"

	"github.com/kr/pretty"
	"github.com/lemon-mint/vstruct/lexer"
	"github.com/lemon-mint/vstruct/parser"
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
	input := ReadFileAsString("./test.vstruct")
	lex := lexer.NewLexer([]rune(input), "./test.vstruct")
	p := parser.New(lex)
	file, err := p.Parse()
	if err != nil {
		panic(err)
	}
	log.Println(file)
	pretty.Println(file)
}
