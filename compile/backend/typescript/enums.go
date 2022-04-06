package typescript

import (
	"fmt"
	"io"

	"github.com/lemon-mint/vstruct/ir"
)

func writeEnums(w io.Writer, i *ir.IR) {
	for _, e := range i.Enums {
		fmt.Fprintf(w, "export enum %s {\n", NameConv(e.Name))
		for _, o := range e.Options {
			fmt.Fprintf(w, "%s,\n", NameConv(o))
		}
		fmt.Fprintf(w, "}\n\n")
	}
}
