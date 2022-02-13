package utils

import (
	"bytes"
	"fmt"

	"github.com/lemon-mint/vstruct/compile/backend/dart"
	"github.com/lemon-mint/vstruct/compile/backend/golang"
	"github.com/lemon-mint/vstruct/compile/backend/python"
	"github.com/lemon-mint/vstruct/compile/backend/rust"
	"github.com/lemon-mint/vstruct/compile/frontend"
	"github.com/lemon-mint/vstruct/lexer"
	"github.com/lemon-mint/vstruct/parser"
)

type CompilerOut struct {
	Err  string `json:"err"`
	Code string `json:"code"`
}

func BuildVstructCLI(args []string, fileName string) CompilerOut {
	if len(args) != 3 {
		return CompilerOut{
			Err: "Invalid arguments",
		}
	}

	lang := args[0]
	pkgname := args[1]
	input := args[2] + "\n"

	if pkgname == "" {
		pkgname = "main"
	}

	lex := lexer.NewLexer([]rune(input), fileName)
	p := parser.New(lex)
	file, err := p.Parse()
	if err != nil {
		return CompilerOut{
			Err: err.Error(),
		}
	}
	front := frontend.New(file)
	err = front.Compile()
	if err != nil {
		return CompilerOut{
			Err: err.Error(),
		}
	}

	goir := front.Output()
	goir.Options.UseUnsafe = true

	var buf bytes.Buffer

	switch lang {
	case "go":
		err = golang.Generate(&buf, goir, pkgname)
	case "rust":
		err = rust.Generate(&buf, goir, pkgname)
	case "dart":
		err = dart.Generate(&buf, goir, pkgname)
	case "python":
		err = python.Generate(&buf, goir, pkgname)
	default:
		err = fmt.Errorf("unknown language: %s", lang)
	}

	if err != nil {
		return CompilerOut{
			Err: err.Error(),
		}
	}

	return CompilerOut{
		Code: buf.String(),
		Err:  "",
	}
}
