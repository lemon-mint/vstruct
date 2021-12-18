package golang

import (
	"fmt"
	"go/format"
	"io"
	"strings"

	"github.com/lemon-mint/vstruct/ir"
)

func Generate(w io.Writer, i *ir.IR, packageName string) error {
	var codedataBuf strings.Builder
	writeEnums(&codedataBuf, i)
	writeStructs(&codedataBuf, i)
	output := fmt.Sprintf(
		`package %s

%s
`,
		packageName,
		codedataBuf.String(),
	)
	fmted, err := format.Source([]byte(output))
	if err != nil {
		fmt.Println(output)
		return err
	}
	_, err = w.Write(fmted)
	return err
}

func writeStructs(w io.Writer, i *ir.IR) {
	for _, s := range i.Structs {
		fmt.Fprintf(w, "type %s []byte\n\n", NameConv(s.Name))

		for _, f := range s.FixedFields {
			fmt.Fprintf(w, "func (s %s) %s() %s {\n", NameConv(s.Name), NameConv(f.Name), TypeConv(f.Type))
			switch f.TypeInfo.FieldType {
			case ir.FieldType_UINT, ir.FieldType_INT:
				fmt.Fprintf(w, "_ = s[%d]\n", f.Offset+f.TypeInfo.Size-1)
				fmt.Fprintf(w, "var __v uint%d = ", f.TypeInfo.Size*8)
				for i := 0; i < f.TypeInfo.Size; i++ {
					if i == 0 {
						fmt.Fprintf(w, "uint%d(s[%d])", f.TypeInfo.Size*8, f.Offset+i)
					} else {
						fmt.Fprintf(w, "|\nuint%d(s[%d])<<%d", f.TypeInfo.Size*8, f.Offset+i, i*8)
					}
				}
				fmt.Fprintf(w, "\nreturn %s(__v)\n", TypeConv(f.Type))
			case ir.FieldType_STRING:
				fmt.Fprintf(w, "return string(s[%d:%d])\n", f.Offset, f.Offset+f.TypeInfo.Size)
			case ir.FieldType_BYTES:
				fmt.Fprintf(w, "return []byte(s[%d:%d])\n", f.Offset, f.Offset+f.TypeInfo.Size)
			case ir.FieldType_ENUM:
				if f.TypeInfo.Size != 1 {
					fmt.Printf("_ = s[%d]\n", f.Offset+f.TypeInfo.Size-1)
					fmt.Fprintf(w, "var __v uint%d = ", f.TypeInfo.Size*8)
					for i := 0; i < f.TypeInfo.Size; i++ {
						if i == 0 {
							fmt.Fprintf(w, "uint%d(s[%d])", f.TypeInfo.Size*8, f.Offset+i)
						} else {
							fmt.Fprintf(w, "|\nuint%d(s[%d])<<%d", f.TypeInfo.Size*8, f.Offset+i, i*8)
						}
					}
					fmt.Fprintf(w, "\nreturn %s(__v)\n", TypeConv(f.Type))
				} else {
					fmt.Fprintf(w, "return %s(s[%d])\n", TypeConv(f.Type), f.Offset)
				}
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, "return %s(s[%d:%d])\n", TypeConv(f.Type), f.Offset, f.Offset+f.TypeInfo.Size)
			}
			fmt.Fprintf(w, "}\n\n")
		}
	}
}
