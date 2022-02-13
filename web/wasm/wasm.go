//go:build wasm
// +build wasm

package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/lemon-mint/vstruct/utils"
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
		stringArgs := make([]string, len(args))
		for i, arg := range args {
			stringArgs[i] = arg.String()
		}
		output := utils.BuildVstructCLI(stringArgs)
		if output.Err != "" {
			js.Global().Get("console").Call("error", output.Err)
		}

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
