package rust

import (
	"fmt"
	"io"

	"github.com/lemon-mint/vstruct/ir"
)

func writeStructs(w io.Writer, i *ir.IR) {
	for _, s := range i.Structs {
		fmt.Fprintf(w, "struct %s {\n", NameConv(s.Name))
		fmt.Fprintf(w, "    buffer: Vec<u8>,\n")
		fmt.Fprintf(w, "}\n\n")

		fmt.Fprintf(w, "impl %s {\n", NameConv(s.Name))
		fmt.Fprintf(w, "    fn new(")
		for i, f := range s.FixedFields {
			if i > 0 {
				fmt.Fprintf(w, ", ")
			}
			if NeedRef(f.TypeInfo.FieldType) {
				fmt.Fprintf(w, "%s: &%s", NameConv(f.Name), TypeConv(f.Type))
			} else {
				fmt.Fprintf(w, "%s: %s", NameConv(f.Name), TypeConv(f.Type))
			}
		}
		if len(s.FixedFields) > 0 {
			fmt.Fprintf(w, ", ")
		}
		for i, f := range s.DynamicFields {
			if i > 0 {
				fmt.Fprintf(w, ", ")
			}
			if NeedRef(f.TypeInfo.FieldType) {
				fmt.Fprintf(w, "%s: &%s", NameConv(f.Name), TypeConv(f.Type))
			} else {
				fmt.Fprintf(w, "%s: %s", NameConv(f.Name), TypeConv(f.Type))
			}
		}
		fmt.Fprintf(w, ") -> %s {\n", NameConv(s.Name))
		fmt.Fprintf(w, "        let mut buffer = Vec::new();\n\n")
		var allFields []*ir.Field
		allFields = append(allFields, s.FixedFields...)
		allFields = append(allFields, s.DynamicFields...)
		for _, f := range allFields {
			switch f.TypeInfo.FieldType {
			case ir.FieldType_BOOL:
				fmt.Fprintf(w, "        buffer.push(if %s { 1 } else { 0 });\n", NameConv(f.Name))
			case ir.FieldType_UINT:
				for i := 0; i < f.TypeInfo.Size; i++ {
					if i == 0 {
						fmt.Fprintf(w, "        buffer.push(%s as u8);\n", NameConv(f.Name))
					} else {
						fmt.Fprintf(w, "        buffer.push((%s >> %d) as u8);\n", NameConv(f.Name), 8*i)
					}
				}
			case ir.FieldType_INT:
				fmt.Fprintf(w, "        let __unsigned_%s = %s as u%d;\n", NameConv(f.Name), NameConv(f.Name), f.TypeInfo.Size*8)
				for i := 0; i < f.TypeInfo.Size; i++ {
					if i == 0 {
						fmt.Fprintf(w, "        buffer.push(__unsigned_%s as u8);\n", NameConv(f.Name))
					} else {
						fmt.Fprintf(w, "        buffer.push((__unsigned_%s >> %d) as u8);\n", NameConv(f.Name), 8*i)
					}
				}
			case ir.FieldType_ENUM:
				fmt.Fprintf(w, "        buffer.push(%s as u%d);\n", NameConv(f.Name), f.TypeInfo.Size*8)
			case ir.FieldType_FLOAT:
				fmt.Fprintf(w, "        let __float_bytes_%s = %s.to_le_bytes();\n", NameConv(f.Name), NameConv(f.Name))
				for i := 0; i < f.TypeInfo.Size; i++ {
					fmt.Fprintf(w, "        buffer.push(__float_bytes_%s[%d]);\n", NameConv(f.Name), i)
				}
			case ir.FieldType_BYTES:
				fmt.Fprintf(w, "        buffer.extend_from_slice(&%s[..]);\n", NameConv(f.Name))
			case ir.FieldType_STRING:
				fmt.Fprintf(w, "        let __string_bytes_%s = %s.as_bytes();\n", NameConv(f.Name), NameConv(f.Name))
				fmt.Fprintf(w, "        buffer.extend_from_slice(__string_bytes_%s);\n", NameConv(f.Name))
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, "        let __struct_bytes_%s = %s.as_bytes();\n", NameConv(f.Name), NameConv(f.Name))
				fmt.Fprintf(w, "        buffer.extend_from_slice(__struct_bytes_%s);\n", NameConv(f.Name))
			}
			fmt.Fprintf(w, "\n")
		}
		fmt.Fprintf(w, "        %s { buffer: buffer }\n", NameConv(s.Name))
		fmt.Fprintf(w, "    }\n\n")
		fmt.Fprintf(w, "    fn as_bytes(&self) -> &[u8] {\n")
		fmt.Fprintf(w, "        &self.buffer[..]\n")
		fmt.Fprintf(w, "    }\n")
		fmt.Fprintf(w, "}\n\n")
	}
}
