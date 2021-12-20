package golang

import (
	"fmt"
	"io"

	"github.com/lemon-mint/vstruct/ir"
)

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
			fmt.Fprintf(w, "_ = s[%d]\n", s.DynamicFieldHeadOffsets[i+1]-1)
			fmt.Fprintf(w, "var __off0 uint64 = ")
			if i == 0 {
				fmt.Fprintf(w, "%d", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
			} else {
				for j := 0; j < 8; j++ {
					if j == 0 {
						fmt.Fprintf(w, "uint64(s[%d])", s.DynamicFieldHeadOffsets[i]-8+j)
					} else {
						fmt.Fprintf(w, "|\nuint64(s[%d])<<%d", s.DynamicFieldHeadOffsets[i]-8+j, j*8)
					}
				}
			}
			fmt.Fprintf(w, "\nvar __off1 uint64 = ")
			for j := 0; j < 8; j++ {
				if j == 0 {
					fmt.Fprintf(w, "uint64(s[%d])", s.DynamicFieldHeadOffsets[i]+j)
				} else {
					fmt.Fprintf(w, "|\nuint64(s[%d])<<%d", s.DynamicFieldHeadOffsets[i]+j, j*8)
				}
			}
			switch f.TypeInfo.FieldType {
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, "\nreturn %s(s[__off0:__off1])\n", TypeConv(f.Type))
			case ir.FieldType_STRING:
				fmt.Fprintf(w, "\nvar __v = s[__off0:__off1]\n")
				if f.Type == "string" {
					fmt.Fprintf(w, "\nreturn *(*string)(unsafe.Pointer(&__v))\n")
				} else {
					fmt.Fprintf(w, "\nreturn *(*%s)(unsafe.Pointer(&__v))\n", TypeConv(f.Type))
				}
			case ir.FieldType_BYTES:
				fmt.Fprintf(w, "\nreturn %s(s[__off0:__off1])\n", TypeConv(f.Type))
			}
			fmt.Fprintf(w, "}\n\n")
		}

		fmt.Fprintf(w, "func (s %s) Vstruct_Validate() bool {\n", NameConv(s.Name))
		if s.IsFixed && len(s.DynamicFields) == 0 {
			fmt.Fprintf(w, "return len(s) >= %d\n", s.TotalFixedFieldSize)
		} else {
			fmt.Fprintf(w, "if len(s) < %d {\n", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
			fmt.Fprintf(w, "return false\n")
			fmt.Fprintf(w, "}\n")
			for i, f := range s.DynamicFieldHeadOffsets {
				_ = f
				fmt.Fprintf(w, "\nvar __off%d uint64 = ", i)
				if i == 0 {
					fmt.Fprintf(w, "%d", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
				} else {
					for j := 0; j < 8; j++ {
						if j == 0 {
							fmt.Fprintf(w, "uint64(s[%d])", s.DynamicFieldHeadOffsets[i]-8+j)
						} else {
							fmt.Fprintf(w, "|\nuint64(s[%d])<<%d", s.DynamicFieldHeadOffsets[i]-8+j, j*8)
						}
					}
				}
			}
			fmt.Fprintf(w, "\nvar __off%d uint64 = uint64(len(s))", len(s.DynamicFieldHeadOffsets))
			var dynStructFields []*ir.Field
			for _, f := range s.DynamicFields {
				if f.TypeInfo.FieldType == ir.FieldType_STRUCT {
					dynStructFields = append(dynStructFields, f)
				}
			}
			if len(dynStructFields) == 0 {
				fmt.Fprintf(w, "\nreturn ")
				for i, f := range s.DynamicFieldHeadOffsets {
					fmt.Fprintf(w, "__off%d <= __off%d ", i, i+1)
					if i != len(s.DynamicFieldHeadOffsets)-1 {
						fmt.Fprintf(w, "&& ")
					}
					_ = f
				}
			} else {
				fmt.Fprintf(w, "\nif ")
				for i, f := range s.DynamicFieldHeadOffsets {
					fmt.Fprintf(w, "__off%d <= __off%d ", i, i+1)
					if i != len(s.DynamicFieldHeadOffsets)-1 {
						fmt.Fprintf(w, "&& ")
					}
					_ = f
				}
				fmt.Fprintf(w, "{\n")
				fmt.Fprintf(w, "return ")
				for i, f := range dynStructFields {
					fmt.Fprintf(w, "s.%s().Vstruct_Validate()", NameConv(f.Name))
					if i != len(dynStructFields)-1 {
						fmt.Fprintf(w, " && ")
					}
				}
				fmt.Fprintf(w, "\n}\n")
				fmt.Fprintf(w, "\nreturn false\n")
			}
		}
		fmt.Fprintf(w, "}\n\n")

		fmt.Fprintf(w, "func (s %s) String() string {\n", NameConv(s.Name))
		fmt.Fprintf(w, "if !s.Vstruct_Validate() {\n")
		fmt.Fprintf(w, "return \"%s (invalid)\"\n", NameConv(s.Name))
		fmt.Fprintf(w, "}\n")
		fmt.Fprintf(w, "var __b strings.Builder\n")
		fmt.Fprintf(w, "__b.WriteString(\"%s {\")\n", NameConv(s.Name))
		var allFields []*ir.Field
		allFields = append(allFields, s.FixedFields...)
		allFields = append(allFields, s.DynamicFields...)
		for i, f := range allFields {
			if i != 0 {
				fmt.Fprintf(w, "__b.WriteString(\", \")\n")
			}
			fmt.Fprintf(w, "__b.WriteString(\"%s: \")\n", NameConv(f.Name))
			switch f.TypeInfo.FieldType {
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, "__b.WriteString(s.%s().String())\n", NameConv(f.Name))
			case ir.FieldType_STRING:
				if f.Type == "string" {
					fmt.Fprintf(w, "__b.WriteString(strconv.Quote(s.%s()))\n", NameConv(f.Name))
				} else {
					fmt.Fprintf(w, "__b.WriteString(strconv.Quote(string(s.%s())))\n", NameConv(f.Name))
				}
			case ir.FieldType_BYTES:
				fmt.Fprintf(w, "__b.WriteString(fmt.Sprint(s.%s()))\n", NameConv(f.Name))
			case ir.FieldType_BOOL:
				fmt.Fprintf(w, "__b.WriteString(strconv.FormatBool(s.%s()))\n", NameConv(f.Name))
			case ir.FieldType_INT:
				fmt.Fprintf(w, "__b.WriteString(strconv.FormatInt(int64(s.%s()), 10))\n", NameConv(f.Name))
			case ir.FieldType_UINT:
				fmt.Fprintf(w, "__b.WriteString(strconv.FormatUint(uint64(s.%s()), 10))\n", NameConv(f.Name))
			case ir.FieldType_FLOAT:
				fmt.Fprintf(w, "__b.WriteString(strconv.FormatFloat(float64(s.%s()), 'g', -1, %d))\n", NameConv(f.Name), f.TypeInfo.Size)
			case ir.FieldType_ENUM:
				fmt.Fprintf(w, "__b.WriteString(s.%s().String())\n", NameConv(f.Name))
			}
		}
		fmt.Fprintf(w, "__b.WriteString(\"}\")\n")
		fmt.Fprintf(w, "return __b.String()\n")
		fmt.Fprintf(w, "}\n\n")
	}

	for _, s := range i.Structs {
		fmt.Fprintf(w, "func Serialize_%s(dst %s", TypeConv(s.Name), TypeConv(s.Name))
		var allFields []*ir.Field
		allFields = append(allFields, s.FixedFields...)
		allFields = append(allFields, s.DynamicFields...)
		for _, f := range allFields {
			fmt.Fprintf(w, ", %s %s", NameConv(f.Name), TypeConv(f.Type))
		}

		var IsFixed bool = len(s.DynamicFields) == 0

		fmt.Fprintf(w, ") %s {\n", TypeConv(s.Name))
		if IsFixed && len(s.FixedFields) > 0 {
			fmt.Fprintf(w, "_ = dst[%d]\n", s.TotalFixedFieldSize-1)
		} else if !IsFixed {
			fmt.Fprintf(w, "_ = dst[%d]\n", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1]-1)
		}
		var tmpIdx int = 0
		for _, f := range s.FixedFields {
			switch f.TypeInfo.FieldType {
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, "copy(dst[%d:%d], %s)\n", f.Offset, f.Offset+f.TypeInfo.Size, NameConv(f.Name))
			case ir.FieldType_BOOL:
				fmt.Fprintf(w, "dst[%d] = *(*byte)(unsafe.Pointer(&%s))\n", f.Offset, NameConv(f.Name))
			case ir.FieldType_INT:
				fmt.Fprintf(w, "var __tmp_%d = uint%d(%s)\n", tmpIdx, f.TypeInfo.Size*8, NameConv(f.Name))
				for i := 0; i < f.TypeInfo.Size; i++ {
					if i == 0 {
						fmt.Fprintf(w, "dst[%d] = byte(__tmp_%d)\n", f.Offset+i, tmpIdx)
					} else {
						fmt.Fprintf(w, "dst[%d] = byte(__tmp_%d >> %d)\n", f.Offset+i, tmpIdx, 8*i)
					}
				}
			case ir.FieldType_UINT:
				for i := 0; i < f.TypeInfo.Size; i++ {
					if i == 0 {
						fmt.Fprintf(w, "dst[%d] = byte(%s)\n", f.Offset+i, NameConv(f.Name))
					} else {
						fmt.Fprintf(w, "dst[%d] = byte(%s >> %d)\n", f.Offset+i, NameConv(f.Name), 8*i)
					}
				}
			case ir.FieldType_FLOAT:
				fmt.Fprintf(w, "var __tmp_%d = math.Float%dbits(%s)\n", tmpIdx, f.TypeInfo.Size*8, NameConv(f.Name))
				for i := 0; i < f.TypeInfo.Size; i++ {
					if i == 0 {
						fmt.Fprintf(w, "dst[%d] = byte(__tmp_%d)\n", f.Offset+i, tmpIdx)
					} else {
						fmt.Fprintf(w, "dst[%d] = byte(__tmp_%d >> %d)\n", f.Offset+i, tmpIdx, 8*i)
					}
				}
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
				fmt.Fprintf(w, "__tmp_%d := uint64(len(%s))\n", tmpIdx, NameConv(f.Name))
				for j := 0; j < 8; j++ {
					if j == 0 {
						fmt.Fprintf(w, "dst[%d] = byte(__tmp_%d)\n", s.DynamicFieldHeadOffsets[i]+j, tmpIdx)
					} else {
						fmt.Fprintf(w, "dst[%d] = byte(__tmp_%d >> %d)\n", s.DynamicFieldHeadOffsets[i]+j, tmpIdx, 8*j)
					}
				}
				switch f.TypeInfo.FieldType {
				case ir.FieldType_STRUCT:
					fmt.Fprintf(w, "copy(dst[__index:__index+__tmp_%d], %s)\n", tmpIdx, NameConv(f.Name))
				case ir.FieldType_BYTES:
					fmt.Fprintf(w, "copy(dst[__index:__index+__tmp_%d], %s)\n", tmpIdx, NameConv(f.Name))
				case ir.FieldType_STRING:
					fmt.Fprintf(w, "copy(dst[__index:__index+__tmp_%d], %s)\n", tmpIdx, NameConv(f.Name))
				}
				if i != len(s.DynamicFields)-1 {
					fmt.Fprintf(w, "__index += __tmp_%d\n", tmpIdx)
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
