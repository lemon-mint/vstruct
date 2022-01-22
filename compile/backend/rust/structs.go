package rust

import (
	"fmt"
	"io"

	"github.com/lemon-mint/vstruct/ir"
)

func writeStructs(w io.Writer, i *ir.IR) {
	for _, s := range i.Structs {
		fmt.Fprintf(w, "pub struct %s {\n", NameConv(s.Name))
		fmt.Fprintf(w, "    buffer: Vec<u8>,\n")
		fmt.Fprintf(w, "}\n\n")

		fmt.Fprintf(w, "impl %s {\n", NameConv(s.Name))
		fmt.Fprintf(w, "    pub fn new(")
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
		if len(s.DynamicFields) > 0 && len(s.FixedFields) > 0 {
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
		var IsFixed bool = len(s.DynamicFields) == 0
		for _, f := range s.FixedFields {
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
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, "        let __struct_bytes_%s = %s.as_bytes();\n", NameConv(f.Name), NameConv(f.Name))
				fmt.Fprintf(w, "        buffer.extend_from_slice(__struct_bytes_%s);\n", NameConv(f.Name))
			}
			fmt.Fprintf(w, "\n")
		}

		if !IsFixed {
			fmt.Fprintf(w, "        let mut __dyn_index = %d;\n", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
		}

		for _, f := range s.DynamicFields {
			fmt.Fprintf(w, "        __dyn_index += %s.len();\n", NameConv(f.Name))

			for i := 0; i < 8; i++ {
				if i == 0 {
					fmt.Fprintf(w, "        buffer.push(__dyn_index as u8);\n")
				} else {
					fmt.Fprintf(w, "        buffer.push((__dyn_index >> %d) as u8);\n", 8*i)
				}
			}

			switch f.TypeInfo.FieldType {
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
		fmt.Fprintf(w, "    pub fn as_bytes(&self) -> &[u8] {\n")
		fmt.Fprintf(w, "        &self.buffer[..]\n")
		fmt.Fprintf(w, "    }\n\n")
		fmt.Fprintf(w, "    pub fn as_bytes_mut(&mut self) -> &mut [u8] {\n")
		fmt.Fprintf(w, "        &mut self.buffer[..]\n")
		fmt.Fprintf(w, "    }\n\n")
		fmt.Fprintf(w, "    pub fn len(&self) -> usize {\n")
		fmt.Fprintf(w, "        self.buffer.len()\n")
		fmt.Fprintf(w, "    }\n\n")
		fmt.Fprintf(w, "    pub fn from_bytes(bytes: &[u8]) -> %s {\n", NameConv(s.Name))
		fmt.Fprintf(w, "        let mut buffer = Vec::new();\n")
		fmt.Fprintf(w, "        buffer.extend_from_slice(bytes);\n")
		fmt.Fprintf(w, "        %s { buffer: buffer }\n", NameConv(s.Name))
		fmt.Fprintf(w, "    }\n\n")

		for _, f := range s.FixedFields {
			switch f.TypeInfo.FieldType {
			case ir.FieldType_BOOL:
				fmt.Fprintf(w, "    pub fn %s(&self) -> bool {\n", NameConv(f.Name))
				fmt.Fprintf(w, "        self.buffer[%d] != 0\n", f.Offset)
				fmt.Fprintf(w, "    }\n\n")
			case ir.FieldType_UINT:
				fmt.Fprintf(w, "    pub fn %s(&self) -> u%d {\n", NameConv(f.Name), f.TypeInfo.Size*8)
				fmt.Fprintf(w, "        let mut __result = ")
				for i := 0; i < f.TypeInfo.Size; i++ {
					if i == 0 {
						fmt.Fprintf(w, "self.buffer[%d] as u%d", f.Offset+i, f.TypeInfo.Size*8)
					} else {
						fmt.Fprintf(w, " \n            | (self.buffer[%d] as u%d) << %d", f.Offset+i, f.TypeInfo.Size*8, 8*i)
					}
				}
				fmt.Fprintf(w, ";\n")
				fmt.Fprintf(w, "        __result\n")
				fmt.Fprintf(w, "    }\n\n")
			case ir.FieldType_INT:
				fmt.Fprintf(w, "    pub fn %s(&self) -> i%d {\n", NameConv(f.Name), f.TypeInfo.Size*8)
				fmt.Fprintf(w, "        let mut __result = ")
				for i := 0; i < f.TypeInfo.Size; i++ {
					if i == 0 {
						fmt.Fprintf(w, "self.buffer[%d] as i%d", f.Offset+i, f.TypeInfo.Size*8)
					} else {
						fmt.Fprintf(w, " \n            | (self.buffer[%d] as i%d) << %d", f.Offset+i, f.TypeInfo.Size*8, 8*i)
					}
				}
				fmt.Fprintf(w, ";\n")
				fmt.Fprintf(w, "        __result\n")
				fmt.Fprintf(w, "    }\n\n")
			case ir.FieldType_FLOAT:
				fmt.Fprintf(w, "    pub fn %s(&self) -> f%d {\n", NameConv(f.Name), f.TypeInfo.Size*8)
				fmt.Fprintf(w, "        f%d::from_le_bytes(self.buffer[%d..%d].try_into().unwrap(\"%s\")).unwrap()\n",
					f.TypeInfo.Size*8,
					f.Offset,
					f.Offset+f.TypeInfo.Size,
					fmt.Sprintf("ValueError: %s", f.Name),
				)
				fmt.Fprintf(w, "    }\n\n")
			case ir.FieldType_ENUM:
				fmt.Fprintf(w, "    pub fn %s(&self) -> %s {\n", NameConv(f.Name), NameConv(f.Type))
				if f.TypeInfo.Size == 1 {
					fmt.Fprintf(w, "        %s::from_u8(self.buffer[%d])\n", NameConv(f.Type), f.Offset)
				} else {
					fmt.Fprintf(w, "        let mut __result_u%d = ", f.TypeInfo.Size*8)
					for i := 0; i < f.TypeInfo.Size; i++ {
						if i == 0 {
							fmt.Fprintf(w, "self.buffer[%d] as u%d", f.Offset+i, f.TypeInfo.Size*8)
						} else {
							fmt.Fprintf(w, " \n            | (self.buffer[%d] as u%d) << %d", f.Offset+i, f.TypeInfo.Size*8, 8*i)
						}
					}
					fmt.Fprintf(w, ";\n")
					fmt.Fprintf(w, "        %s::from_u%d(__result_u%d)\n", NameConv(f.Type), f.TypeInfo.Size*8, f.TypeInfo.Size*8)
				}
				fmt.Fprintf(w, "    }\n\n")
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, "    pub fn %s(&self) -> %s {\n", NameConv(f.Name), NameConv(f.Type))
				fmt.Fprintf(w, "        // TODO: Think about rust's borrowing rules\n")
				fmt.Fprintf(w, "        %s::from_bytes(&self.buffer[%d..%d])\n", NameConv(f.Type), f.Offset, f.Offset+f.TypeInfo.Size)
				fmt.Fprintf(w, "    }\n\n")
			}
		}

		for i, f := range s.DynamicFields {
			fmt.Fprintf(w, "    pub fn %s(&self) -> %s {\n", NameConv(f.Name), NameConv(f.Type))
			fmt.Fprintf(w, "        let __off0: u64 = ")
			if i == 0 {
				fmt.Fprintf(w, "%d", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
			} else {
				for j := 0; j < 8; j++ {
					if j == 0 {
						fmt.Fprintf(w, "self.buffer[%d] as u64", s.DynamicFieldHeadOffsets[i]-8+j)
					} else {
						fmt.Fprintf(w, " \n            | (self.buffer[%d] as u64) << %d", s.DynamicFieldHeadOffsets[i]-8+j, j*8)
					}
				}
			}
			fmt.Fprintf(w, ";\n")
			fmt.Fprintf(w, "        let __off1: u64 = ")
			for j := 0; j < 8; j++ {
				if j == 0 {
					fmt.Fprintf(w, "self.buffer[%d] as u64", s.DynamicFieldHeadOffsets[i]+j)
				} else {
					fmt.Fprintf(w, " \n            | (self.buffer[%d] as u64) << %d", s.DynamicFieldHeadOffsets[i]+j, j*8)
				}
			}
			fmt.Fprintf(w, ";\n")

			fmt.Fprintf(w, "        // TODO: Think about rust's borrowing rules\n\n")
			switch f.TypeInfo.FieldType {
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, "        %s::from_bytes(&self.buffer[__off0 as usize..__off1 as usize])\n", NameConv(f.Type))
			case ir.FieldType_STRING:
				fmt.Fprintf(w, "        String::from_utf8(self.buffer[__off0 as usize..__off1 as usize].to_vec()).unwrap()\n")
			case ir.FieldType_BYTES:
				fmt.Fprintf(w, "        self.buffer[__off0 as usize..__off1 as usize].to_vec()\n")
			}
			fmt.Fprintf(w, "    }\n\n")
		}

		fmt.Fprintf(w, "}\n\n")
	}
}
