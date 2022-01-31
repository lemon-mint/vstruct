package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/lemon-mint/vstruct/compile/backend/dart"
	"github.com/lemon-mint/vstruct/compile/backend/golang"
	"github.com/lemon-mint/vstruct/compile/backend/rust"
	"github.com/lemon-mint/vstruct/compile/backend/python"
	"github.com/lemon-mint/vstruct/compile/frontend"
	"github.com/lemon-mint/vstruct/ir"
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
	if len(os.Args) < 3 {
		fmt.Println("Usage: vstruct <language: go, rust, dart, py> <file>...")
		os.Exit(1)
	}
	generate := (func(w io.Writer, i *ir.IR, packageName string) error)(nil)
	modelPath := filepath.Join("vstruct", "model")
	ext := ""

	language := os.Args[1]
	switch language {
	case "go":
		generate = golang.Generate
		modelPath = filepath.Join(".", modelPath)
		ext = ".go"
	case "rust":
		generate = rust.Generate
		modelPath = filepath.Join("src", modelPath)
		ext = ".rs"
	case "dart":
		generate = dart.Generate
		modelPath = filepath.Join("lib", modelPath)
		ext = ".dart"
	case "py":
		generate = python.Generate
		modelPath = filepath.Join(".", modelPath)
		ext = ".py"
	default:
		fmt.Printf("Unknown language: %s\n", language)
		os.Exit(1)
	}

	for _, path := range os.Args[2:] {
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
		err = generate(&buf, goir, packageName)
		if err != nil {
			fmt.Println(buf.String())
			panic(err)
		}
		out := buf.String()
		fmt.Println(out)

		if err := CreateDirectory(filepath.Join(modelPath, packageName)); err != nil {
			panic(err)
		}
		f, err := os.Create(filepath.Join(modelPath, packageName, packageName+ext))
		if err != nil {
			panic(err)
		}
		f.WriteString(out)
		f.Close()
	}
}
