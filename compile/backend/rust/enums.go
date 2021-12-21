package rust

import (
	"fmt"
	"io"

	"github.com/lemon-mint/vstruct/ir"
)

func writeEnums(w io.Writer, i *ir.IR) {
	for _, e := range i.Enums {
		fmt.Fprintf(w, "enum %s {\n", NameConv(e.Name))
		for i, v := range e.Options {
			if i > 0 {
				fmt.Fprintf(w, ",\n")
			}
			fmt.Fprintf(w, "    %s", NameConv(v))
		}
		fmt.Fprintf(w, "\n}\n\n")
	}
}
