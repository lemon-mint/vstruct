package typescript

import (
	"fmt"
	"io"

	"github.com/lemon-mint/vstruct/ir"
)

func writeStructs(w io.Writer, i *ir.IR) {
	for _, s := range i.Structs {
		fmt.Fprintf(w, "class %s {\n", NameConv(s.Name))
		fmt.Fprintf(w, "	public value: Uint8Array = new Uint8Array();\n")
		fmt.Fprintf(w, `	constructor(value: Uint8Array) {
		this.value = value;
	}`)
		fmt.Fprint(w, "\n\n")

		for _, f := range s.FixedFields {
			fmt.Fprintf(w, "public %s(): %s {\n", NameConv(f.Name), TypeConv(f.Type))
			switch f.TypeInfo.FieldType {
			case ir.FieldType_BOOL:
				fmt.Fprintf(w, "return this.value[%d] != 0\n", f.Offset)
			case ir.FieldType_UINT, ir.FieldType_INT:
				fmt.Fprintf(w, "let __v: bigint = ")
				for i := 0; i < f.TypeInfo.Size; i++ {
					if i == 0 {
						fmt.Fprintf(w, "BigInt(this.value[%d])", f.Offset+i)
					} else {
						fmt.Fprintf(w, "|\nBigInt(this.value[%d]) << %dn", f.Offset+i, i*8)
					}
				}
				fmt.Fprintf(w, "\nreturn __v\n")
			case ir.FieldType_FLOAT:
				switch f.TypeInfo.Size {
				case 4:
					fmt.Fprintf(w, "\nreturn %s(new Float32Array(this.value.slice(%d, %d+4))[0])\n", TypeConv(f.Type), f.Offset, f.Offset)
				case 8:
					fmt.Fprintf(w, "\nreturn %s(new Float64Array(this.value.slice(%d, %d+8))[0])\n", TypeConv(f.Type), f.Offset, f.Offset)
				default:
					panic("unsupported float size")
				}
			case ir.FieldType_ENUM:
				if f.TypeInfo.Size != 1 {
					fmt.Fprintf(w, "let __v: number = ")
					for i := 0; i < f.TypeInfo.Size; i++ {
						if i == 0 {
							fmt.Fprintf(w, "s[%d]", f.Offset+i)
						} else {
							fmt.Fprintf(w, "|\nthis.value[%d] << %d", f.Offset+i, i*8)
						}
					}
					fmt.Fprintf(w, "\nreturn <%s><unknown>%s[__v]\n", TypeConv(f.Type), TypeConv(f.Type))
				} else {
					fmt.Fprintf(w, "return <%s><unknown>%s[this.value[%d]]\n", TypeConv(f.Type), TypeConv(f.Type), f.Offset)
				}
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, "return new %s(this.value.slice(%d, %d))\n", TypeConv(f.Type), f.Offset, f.Offset+f.TypeInfo.Size)
			}
			fmt.Fprintf(w, "}\n\n")
		}

		for i, f := range s.DynamicFields {
			fmt.Fprintf(w, "public %s(): %s {\n", NameConv(f.Name), TypeConv(f.Type))
			fmt.Fprintf(w, "let __off0 = ")
			if i == 0 {
				fmt.Fprintf(w, "%d", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
			} else {
				for j := 0; j < 8; j++ {
					if j == 0 {
						fmt.Fprintf(w, "this.value[%d]", s.DynamicFieldHeadOffsets[i]-8+j)
					} else {
						fmt.Fprintf(w, "|\nthis.value[%d] << %d", s.DynamicFieldHeadOffsets[i]-8+j, j*8)
					}
				}
			}
			fmt.Fprintf(w, "\nlet __off1 = ")
			for j := 0; j < 8; j++ {
				if j == 0 {
					fmt.Fprintf(w, "this.value[%d]", s.DynamicFieldHeadOffsets[i]+j)
				} else {
					fmt.Fprintf(w, "|\nthis.value[%d] << %d", s.DynamicFieldHeadOffsets[i]+j, j*8)
				}
			}
			switch f.TypeInfo.FieldType {
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, "\nreturn new %s(this.value.slice(__off0, __off1))\n", TypeConv(f.Type))
			case ir.FieldType_STRING:
				fmt.Fprintf(w, "\nlet __v = this.value.slice(__off0, __off1)\n")
				fmt.Fprintf(w, "\nreturn new TextDecoder(\"utf-8\").decode(__v)\n")
			case ir.FieldType_BYTES:
				fmt.Fprintf(w, "\nreturn this.value.slice(__off0, __off1)\n")
			}
			fmt.Fprintf(w, "}\n\n")
		}
		fmt.Fprintf(w, "}\n\n")

		// fmt.Fprintf(w, "func (s %s) Vstruct_Validate() bool {\n", NameConv(s.Name))
		// if s.IsFixed && len(s.DynamicFields) == 0 {
		// 	fmt.Fprintf(w, "return len(s) >= %d\n", s.TotalFixedFieldSize)
		// } else {
		// 	fmt.Fprintf(w, "if len(s) < %d {\n", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
		// 	fmt.Fprintf(w, "return false\n")
		// 	fmt.Fprintf(w, "}\n")
		// 	if s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1] > 0 {
		// 		// Add Bounds Check Elimination
		// 		fmt.Fprintf(w, "\n_ = s[%d]\n", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1]-1)
		// 	}

		// 	for i, f := range s.DynamicFieldHeadOffsets {
		// 		_ = f
		// 		fmt.Fprintf(w, "\nvar __off%d uint64 = ", i)
		// 		if i == 0 {
		// 			fmt.Fprintf(w, "%d", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
		// 		} else {
		// 			for j := 0; j < 8; j++ {
		// 				if j == 0 {
		// 					fmt.Fprintf(w, "uint64(s[%d])", s.DynamicFieldHeadOffsets[i]-8+j)
		// 				} else {
		// 					fmt.Fprintf(w, "|\nuint64(s[%d])<<%d", s.DynamicFieldHeadOffsets[i]-8+j, j*8)
		// 				}
		// 			}
		// 		}
		// 	}
		// 	fmt.Fprintf(w, "\nvar __off%d uint64 = uint64(len(s))", len(s.DynamicFieldHeadOffsets))
		// 	var dynStructFields []*ir.Field
		// 	for _, f := range s.DynamicFields {
		// 		if f.TypeInfo.FieldType == ir.FieldType_STRUCT {
		// 			dynStructFields = append(dynStructFields, f)
		// 		}
		// 	}
		// 	if len(dynStructFields) == 0 {
		// 		fmt.Fprintf(w, "\nreturn ")
		// 		for i, f := range s.DynamicFieldHeadOffsets {
		// 			fmt.Fprintf(w, "__off%d <= __off%d ", i, i+1)
		// 			if i != len(s.DynamicFieldHeadOffsets)-1 {
		// 				fmt.Fprintf(w, "&& ")
		// 			}
		// 			_ = f
		// 		}
		// 	} else {
		// 		fmt.Fprintf(w, "\nif ")
		// 		for i, f := range s.DynamicFieldHeadOffsets {
		// 			fmt.Fprintf(w, "__off%d <= __off%d ", i, i+1)
		// 			if i != len(s.DynamicFieldHeadOffsets)-1 {
		// 				fmt.Fprintf(w, "&& ")
		// 			}
		// 			_ = f
		// 		}
		// 		fmt.Fprintf(w, "{\n")
		// 		fmt.Fprintf(w, "return ")
		// 		for i, f := range dynStructFields {
		// 			fmt.Fprintf(w, "s.%s().Vstruct_Validate()", NameConv(f.Name))
		// 			if i != len(dynStructFields)-1 {
		// 				fmt.Fprintf(w, " && ")
		// 			}
		// 		}
		// 		fmt.Fprintf(w, "\n}\n")
		// 		fmt.Fprintf(w, "\nreturn false\n")
		// 	}
		// }
		// fmt.Fprintf(w, "}\n\n")
	}

	for _, s := range i.Structs {
		fmt.Fprintf(w, "function Serialize_%s(dst: %s", TypeConv(s.Name), TypeConv(s.Name))
		var allFields []*ir.Field
		allFields = append(allFields, s.FixedFields...)
		allFields = append(allFields, s.DynamicFields...)
		for _, f := range allFields {
			fmt.Fprintf(w, ", %s: %s", NameConv(f.Name), TypeConv(f.Type))
		}

		var IsFixed bool = len(s.DynamicFields) == 0

		fmt.Fprintf(w, "): %s {\n", TypeConv(s.Name))
		// if IsFixed && len(s.FixedFields) > 0 {
		// 	fmt.Fprintf(w, "_ = dst[%d]\n", s.TotalFixedFieldSize-1)
		// } else if !IsFixed {
		// 	fmt.Fprintf(w, "_ = dst[%d]\n", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1]-1)
		// }
		var tmpIdx int = 0
		for _, f := range s.FixedFields {
			switch f.TypeInfo.FieldType {
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, "dst.value = new Uint8Array(...dst.value, ...%s)\n", NameConv(f.Name))
			case ir.FieldType_BOOL:
				fmt.Fprintf(w, "dst.value = new Uint8Array(...dst.value, %s==true)\n", NameConv(f.Name))
			case ir.FieldType_INT:
				fmt.Fprintf(w, "var __tmp_%d = BigInt(%s)\n", tmpIdx, NameConv(f.Name))
				for i := 0; i < f.TypeInfo.Size; i++ {
					if i == 0 {
						fmt.Fprintf(w, "dst.value = new Uint8Array(...dst.value, Number(__tmp_%d))\n", tmpIdx)
					} else {
						fmt.Fprintf(w, "dst.value = new Uint8Array(...dst.value, Number(__tmp_%d >> %dn))\n", tmpIdx, 8*i)
					}
				}
			case ir.FieldType_UINT:
				for i := 0; i < f.TypeInfo.Size; i++ {
					if i == 0 {
						fmt.Fprintf(w, "dst.value = new Uint8Array(...dst.value, Number(%s))\n", NameConv(f.Name))
					} else {
						fmt.Fprintf(w, "dst.value = new Uint8Array(...dst.value, Number(%s >> %d))\n", NameConv(f.Name), 8*i)
					}
				}
			case ir.FieldType_FLOAT:
				fmt.Fprintf(w, "var __tmp_%d = new Uint8Array(%d)\n", tmpIdx, f.TypeInfo.Size)
				fmt.Fprintf(w, "var __tmp_%d_f = new Float%dArray(__temp_%d.buffer)\n", tmpIdx, f.TypeInfo.Size*8, tmpIdx)
				fmt.Fprintf(w, "__tmp_%d_f[0] = %s\n", tmpIdx, NameConv(f.Name))
				fmt.Fprintf(w, "dst.value = new Uint8Array(...dst.value, ...__tmp_%d)\n", tmpIdx)
			case ir.FieldType_ENUM:
				if f.TypeInfo.Size == 1 {
					fmt.Fprintf(w, "dst[%d] = byte(%s)\n", f.Offset, NameConv(f.Name))
				} else {
					fmt.Fprintf(w, "var __tmp_%d = uint%d(%s)\n", tmpIdx, f.TypeInfo.Size*8, NameConv(f.Name))
					for i := 0; i < f.TypeInfo.Size; i++ {
						if i == 0 {
							fmt.Fprintf(w, "dst[%d] = byte(__tmp_%d)\n", f.Offset+i, tmpIdx)
						} else {
							fmt.Fprintf(w, "dst[%d] = byte(__tmp_%d >> %d)\n", f.Offset+i, tmpIdx, 8*i)
						}
					}
				}
			}
			tmpIdx++
		}
		fmt.Fprintf(w, "\n")
		if !IsFixed {
			fmt.Fprintf(w, "var __index = uint64(%d)\n", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
			for i, f := range s.DynamicFields {
				fmt.Fprintf(w, "__tmp_%d := uint64(len(%s)) +__index\n", tmpIdx, NameConv(f.Name))
				for j := 0; j < 8; j++ {
					if j == 0 {
						fmt.Fprintf(w, "dst[%d] = byte(__tmp_%d)\n", s.DynamicFieldHeadOffsets[i]+j, tmpIdx)
					} else {
						fmt.Fprintf(w, "dst[%d] = byte(__tmp_%d >> %d)\n", s.DynamicFieldHeadOffsets[i]+j, tmpIdx, 8*j)
					}
				}
				switch f.TypeInfo.FieldType {
				case ir.FieldType_STRUCT:
					fmt.Fprintf(w, "copy(dst[__index:__tmp_%d], %s)\n", tmpIdx, NameConv(f.Name))
				case ir.FieldType_BYTES:
					fmt.Fprintf(w, "copy(dst[__index:__tmp_%d], %s)\n", tmpIdx, NameConv(f.Name))
				case ir.FieldType_STRING:
					fmt.Fprintf(w, "copy(dst[__index:__tmp_%d], %s)\n", tmpIdx, NameConv(f.Name))
				}
				if i != len(s.DynamicFields)-1 {
					fmt.Fprintf(w, "__index += uint64(len(%s))\n", NameConv(f.Name))
				}
				tmpIdx++
			}
		}
		fmt.Fprintf(w, "return dst\n")
		fmt.Fprintf(w, "}\n\n")

		fmt.Fprintf(w, "func New_%s(", TypeConv(s.Name))
		for i, f := range allFields {
			if i != 0 {
				fmt.Fprintf(w, ", ")
			}
			fmt.Fprintf(w, "%s %s", NameConv(f.Name), TypeConv(f.Type))
		}
		fmt.Fprintf(w, ") %s {\n", TypeConv(s.Name))
		fmt.Fprintf(w, "var __vstruct__size = %d", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
		for _, f := range s.DynamicFields {
			fmt.Fprintf(w, "+len(%s)", NameConv(f.Name))
		}
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "var __vstruct__buf = make(%s, __vstruct__size)\n", TypeConv(s.Name))
		fmt.Fprintf(w, "__vstruct__buf = Serialize_%s(__vstruct__buf", TypeConv(s.Name))
		for _, f := range allFields {
			fmt.Fprintf(w, ", %s", NameConv(f.Name))
		}
		fmt.Fprintf(w, ")\n")
		fmt.Fprintf(w, "return __vstruct__buf\n")
		fmt.Fprintf(w, "}\n\n")
	}
}
