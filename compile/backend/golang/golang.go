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

import (
	"math"
)

var _ = math.Float32frombits
var _ = math.Float64frombits

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
			case ir.FieldType_BOOL:
				fmt.Fprintf(w, "return %s(s[%d] != 0)\n", TypeConv(f.Type), f.Offset)
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
			case ir.FieldType_FLOAT:
				fmt.Fprintf(w, "_ = s[%d]\n", f.Offset+f.TypeInfo.Size-1)
				fmt.Fprintf(w, "var __v uint%d = ", f.TypeInfo.Size*8)
				for i := 0; i < f.TypeInfo.Size; i++ {
					if i == 0 {
						fmt.Fprintf(w, "uint%d(s[%d])", f.TypeInfo.Size*8, f.Offset+i)
					} else {
						fmt.Fprintf(w, "|\nuint%d(s[%d])<<%d", f.TypeInfo.Size*8, f.Offset+i, i*8)
					}
				}
				fmt.Fprintf(w, "\nreturn %s(math.Float%dfrombits(__v))\n", TypeConv(f.Type), f.TypeInfo.Size*8)
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

		for i, f := range s.DynamicFields {
			fmt.Fprintf(w, "func (s %s) %s() %s {\n", NameConv(s.Name), NameConv(f.Name), TypeConv(f.Type))
			fmt.Fprintf(w, "_ = s[%d]\n", s.DynamicFieldHeadOffsets[i]+15)
			fmt.Fprintf(w, "var __off0 uint64 = ")
			for j := 0; j < 8; j++ {
				if j == 0 {
					fmt.Fprintf(w, "uint64(s[%d])", s.DynamicFieldHeadOffsets[i]+j)
				} else {
					fmt.Fprintf(w, "|\nuint64(s[%d])<<%d", s.DynamicFieldHeadOffsets[i]+j, j*8)
				}
			}
			fmt.Fprintf(w, "\nvar __off1 uint64 = ")
			for j := 0; j < 8; j++ {
				if j == 0 {
					fmt.Fprintf(w, "uint64(s[%d])", s.DynamicFieldHeadOffsets[i]+8+j)
				} else {
					fmt.Fprintf(w, "|\nuint64(s[%d])<<%d", s.DynamicFieldHeadOffsets[i]+8+j, j*8)
				}
			}
			switch f.TypeInfo.FieldType {
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, "\nreturn %s(s[__off0:__off1])\n", TypeConv(f.Type))
			case ir.FieldType_STRING:
				fmt.Fprintf(w, "\nreturn %s(s[__off0:__off1])\n", TypeConv(f.Type))
			case ir.FieldType_BYTES:
				fmt.Fprintf(w, "\nreturn %s(s[__off0:__off1])\n", TypeConv(f.Type))
			}
			fmt.Fprintf(w, "}\n\n")
		}
	}
}
