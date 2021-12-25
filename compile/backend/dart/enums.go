package dart

import (
	"fmt"
	"io"

	"github.com/lemon-mint/vstruct/ir"
)

func writeEnums(w io.Writer, i *ir.IR) {
	for _, e := range i.Enums {
		fmt.Fprintf(w, "enum %s {\n", NameConv(e.Name))
		for _, v := range e.Options {
			fmt.Fprintf(w, "\t%s,\n", NameConv(v))
		}
		fmt.Fprintf(w, "}\n\n")
	}
}
