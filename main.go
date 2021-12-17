package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/kr/pretty"
	"github.com/lemon-mint/vstruct/compile/backend/golang"
	"github.com/lemon-mint/vstruct/compile/frontend"
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
	path, err := filepath.Abs("./test.vstruct")
	if err != nil {
		panic(err)
	}
	input := ReadFileAsString(path)
	lex := lexer.NewLexer([]rune(input), path)
	p := parser.New(lex)
	file, err := p.Parse()
	if err != nil {
		panic(err)
	}
	front := frontend.New(file)
	err = front.Compile()
	if err != nil {
		panic(err)
	}
	goir := front.Output()
	pretty.Println(goir)
	var buf bytes.Buffer
	err = golang.Generate(&buf, goir, "main")
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}
