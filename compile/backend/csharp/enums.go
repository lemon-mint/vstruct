package csharp

import (
	"fmt"
	"io"

	"github.com/lemon-mint/vstruct/ir"
)

func writeEnums(w io.Writer, i *ir.IR) {
	for _, e := range i.Enums {
		fmt.Fprintf(w, "\tpublic enum %s : ulong\n", NameConv(e.Name))
		fmt.Fprintf(w, "\t{\n")
		for _, o := range e.Options {
			fmt.Fprintf(w, "\t\t%s,\n", NameConv(o))
		}
		fmt.Fprintf(w, "\t}\n\n")
	}
}
