package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/lemon-mint/vstruct/compile/backend/dart"
	"github.com/lemon-mint/vstruct/compile/backend/golang"
	"github.com/lemon-mint/vstruct/compile/backend/python"
	"github.com/lemon-mint/vstruct/compile/backend/rust"
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
	IRData := front.Output()
	IRData.Options.UseUnsafe = true

	var buf bytes.Buffer
	err = golang.Generate(&buf, IRData, "main")
	if err != nil {
		fmt.Println(buf.String())
		panic(err)
	}
	out := buf.String()
	fmt.Println(out)
	f, err := os.Create("./_out/main.go")
	if err != nil {
		panic(err)
	}
	f.WriteString(out)
	f.Close()

	buf.Reset()
	err = rust.Generate(&buf, IRData, "main")
	if err != nil {
		fmt.Println(buf.String())
		panic(err)
	}
	out = buf.String()
	fmt.Println(out)
	f, err = os.Create("./_out/main.rs")
	if err != nil {
		panic(err)
	}
	f.WriteString(out)
	f.WriteString("fn main() {}\n")
	f.Close()

	buf.Reset()
	err = dart.Generate(&buf, IRData, "main")
	if err != nil {
		fmt.Println(buf.String())
		panic(err)
	}
	out = buf.String()
	fmt.Println(out)
	f, err = os.Create("./_out/bin/main.dart")
	if err != nil {
		panic(err)
	}
	f.WriteString(out)
	f.Close()

	buf.Reset()
	err = python.Generate(&buf, IRData, "main")
	if err != nil {
		fmt.Println(buf.String())
		panic(err)
	}
	out = buf.String()
	fmt.Println(out)
	f, err = os.Create("./_out/main.py")
	if err != nil {
		panic(err)
	}
	f.WriteString(out)
	f.Close()
}
