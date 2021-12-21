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

		fmt.Fprintf(w, "impl %s {\n", NameConv(e.Name))
		fmt.Fprintf(w, "    pub fn from_u%d(value: u%d) -> %s {\n", e.Size*8, e.Size*8, NameConv(e.Name))
		fmt.Fprintf(w, "        match value {\n")
		for i, v := range e.Options {
			fmt.Fprintf(w, "            %d => %s::%s,\n", i, NameConv(e.Name), NameConv(v))
		}
		fmt.Fprintf(w, "            _ => panic!(\"invalid value for %s: {}\", value),\n", NameConv(e.Name))
		fmt.Fprintf(w, "        }\n")
		fmt.Fprintf(w, "    }\n")
		fmt.Fprintf(w, "}\n\n")
	}
}
