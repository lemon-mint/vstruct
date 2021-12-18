package golang

import (
	"fmt"
	"io"

	"github.com/lemon-mint/vstruct/ir"
)

func writeEnums(w io.Writer, i *ir.IR) {
	for _, e := range i.Enums {
		fmt.Fprintf(w, "type %s uint%d\n", NameConv(e.Name), e.Size*8)
		fmt.Fprintf(w, "const (\n")
		for i, o := range e.Options {
			fmt.Fprintf(w, "%s_%s %s = %d\n", NameConv(e.Name), NameConv(o), NameConv(e.Name), i)
		}
		fmt.Fprintf(w, ")\n\n")

		fmt.Fprintf(w, "func (e %s) String() string {\n", NameConv(e.Name))
		fmt.Fprintf(w, "switch e {\n")
		for _, o := range e.Options {
			fmt.Fprintf(w, "case %s_%s:\n", NameConv(e.Name), NameConv(o))
			fmt.Fprintf(w, "return \"%s\"\n", NameConv(o))
		}
		fmt.Fprintf(w, "}\n")
		fmt.Fprintf(w, "return \"\"\n")
		fmt.Fprintf(w, "}\n\n")

		fmt.Fprintf(w, "func (e %s) Match(\n", NameConv(e.Name))
		for _, o := range e.Options {
			fmt.Fprintf(w, "on%s func(),\n", NameConv(o))
		}
		fmt.Fprintf(w, ") {\n")
		fmt.Fprintf(w, "switch e {\n")
		for _, o := range e.Options {
			fmt.Fprintf(w, "case %s_%s:\n", NameConv(e.Name), NameConv(o))
			fmt.Fprintf(w, "on%s()\n", NameConv(o))
		}
		fmt.Fprintf(w, "}\n")
		fmt.Fprintf(w, "}\n\n")
	}
}
