package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/lemon-mint/vstruct/compile/backend/dart"
	"github.com/lemon-mint/vstruct/compile/backend/golang"
	"github.com/lemon-mint/vstruct/compile/backend/rust"
	"github.com/lemon-mint/vstruct/compile/frontend"
	"github.com/lemon-mint/vstruct/lexer"
	"github.com/lemon-mint/vstruct/parser"
)

type CompilerOut struct {
	Err  string `json:"err"`
	Code string `json:"code"`
}

func main() {
	compile := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// args[0] : language (go|rust|dart)
		// args[1] : pkgname (string)
		// args[2] : input vstruct file (string)
		output := func() CompilerOut {
			if len(args) != 3 {
				return CompilerOut{
					Err: "Invalid arguments",
				}
			}

			lang := args[0].String()
			pkgname := args[1].String()
			input := args[2].String()

			if pkgname == "" {
				pkgname = "main"
			}

			lex := lexer.NewLexer([]rune(input), "/usr/local/vstruct/playground.vstruct")
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
			default:
				err = fmt.Errorf("Unknown language: %s", lang)
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
		}()

		jsonout, err := json.Marshal(output)
		if err != nil {
			return ":ERROR:" + err.Error()
		}
		return string(jsonout)
	})

	js.Global().Set("v_compile", compile)

	ch := make(chan struct{})
	js.Global().Set("v_loaded", true)
	<-ch
}
