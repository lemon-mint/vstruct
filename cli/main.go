package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

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

func CreateDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vstruct <file>")
		os.Exit(1)
	}
	for _, path := range os.Args[1:] {
		if !strings.HasSuffix(path, ".vstruct") {
			path = path + ".vstruct"
		}

		filename, err := filepath.Abs(path)
		if err != nil {
			log.Println(err)
			continue
		}

		input := ReadFileAsString(filename)
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
		goir.Options.UseUnsafe = true

		packageName := filepath.Base(path)
		packageName = strings.TrimSuffix(packageName, ".vstruct")

		var buf bytes.Buffer
		err = golang.Generate(&buf, goir, packageName)
		if err != nil {
			fmt.Println(buf.String())
			panic(err)
		}
		out := buf.String()
		fmt.Println(out)

		if err := CreateDirectory(filepath.Join(".", "vstruct", "model", packageName)); err != nil {
			panic(err)
		}
		f, err := os.Create(filepath.Join(".", "vstruct", "model", packageName, packageName+".go"))
		if err != nil {
			panic(err)
		}
		f.WriteString(out)
		f.Close()
	}
}
