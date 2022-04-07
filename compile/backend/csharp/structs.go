package csharp

import (
	"fmt"
	"io"

	"github.com/lemon-mint/vstruct/ir"
)

func writeStructs(w io.Writer, i *ir.IR) {
	for _, s := range i.Structs {
		fmt.Fprintf(w, "\tclass %s\n\t{\n", NameConv(s.Name))

		fmt.Fprintf(w, "\t\tpublic %s(int size)\n\t\t{\n", NameConv(s.Name))
		fmt.Fprintf(w, "\t\t\tthis.value = new byte[size];\n")
		fmt.Fprint(w, "\t\t}\n\n")

		fmt.Fprintf(w, "\t\tpublic static %s FromBytes(byte[] bytes)\n\t\t{\n", NameConv(s.Name))
		fmt.Fprintf(w, "\t\t\tvar s = new %s(bytes.Length);\n", NameConv(s.Name))
		fmt.Fprintf(w, "\t\t\tArray.Copy(bytes, 0, s.value, 0, bytes.Length);\n")
		fmt.Fprint(w, "\t\t\treturn s;\n")
		fmt.Fprint(w, "\t\t}\n\n")

		fmt.Fprintf(w, "\t\tpublic static %s Serialize(%s dst", TypeConv(s.Name), TypeConv(s.Name))
		var allFields []*ir.Field
		allFields = append(allFields, s.FixedFields...)
		allFields = append(allFields, s.DynamicFields...)
		for _, f := range allFields {
			fmt.Fprintf(w, ", %s %s", TypeConv(f.Type), NameConv(f.Name))
		}

		var IsFixed bool = len(s.DynamicFields) == 0

		fmt.Fprintf(w, ")\n\t\t{\n")
		// if IsFixed && len(s.FixedFields) > 0 {
		// 	fmt.Fprintf(w, "_ = dst[%d]\n", s.TotalFixedFieldSize-1)
		// } else if !IsFixed {
		// 	fmt.Fprintf(w, "_ = dst[%d]\n", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1]-1)
		// }
		var tmpIdx int = 0
		for _, f := range s.FixedFields {
			switch f.TypeInfo.FieldType {
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, "\t\t\tArray.Copy(%s.value, %d, dst.value, %d, %d);\n", NameConv(f.Name), 0, f.Offset, f.TypeInfo.Size)
			case ir.FieldType_BOOL:
				fmt.Fprintf(w, "\t\t\tdst.value[%d] = (byte)%s;\n", f.Offset, NameConv(f.Name))
			case ir.FieldType_INT:
				fmt.Fprintf(w, "\t\t\tvar __tmp_%d = BitConverter.GetBytes(%s);\n", tmpIdx, NameConv(f.Name))
				fmt.Fprintf(w, "\t\t\tif (!BitConverter.IsLittleEndian)\n\t\t\t{\n\t\t\t\tArray.Reverse(__tmp_%d);\n\t\t\t}\n", tmpIdx)
				fmt.Fprintf(w, "\t\t\tArray.Copy(__tmp_%d, 0, dst.value, %d, %d);\n", tmpIdx, f.Offset, f.TypeInfo.Size)
			case ir.FieldType_UINT:
				fmt.Fprintf(w, "\t\t\tvar __tmp_%d = BitConverter.GetBytes(%s);\n", tmpIdx, NameConv(f.Name))
				fmt.Fprintf(w, "\t\t\tif (!BitConverter.IsLittleEndian)\n\t\t\t{\n\t\t\t\tArray.Reverse(__tmp_%d);\n\t\t\t}\n", tmpIdx)
				fmt.Fprintf(w, "\t\t\tArray.Copy(__tmp_%d, 0, dst.value, %d, %d);\n", tmpIdx, f.Offset, f.TypeInfo.Size)
			case ir.FieldType_FLOAT:
				fmt.Fprintf(w, "\t\t\tvar __tmp_%d = BitConverter.GetBytes(%s);\n", tmpIdx, NameConv(f.Name))
				fmt.Fprintf(w, "\t\t\tif (!BitConverter.IsLittleEndian)\n\t\t\t{\n\t\t\t\tArray.Reverse(__tmp_%d);\n\t\t\t}\n", tmpIdx)
				fmt.Fprintf(w, "\t\t\tArray.Copy(__tmp_%d, 0, dst.value, %d, %d);\n", tmpIdx, f.Offset, f.TypeInfo.Size)
			case ir.FieldType_ENUM:
				if f.TypeInfo.Size == 1 {
					fmt.Fprintf(w, "\t\t\tdst.value[%d] = (byte)%s;\n", f.Offset, NameConv(f.Name))
				} else {
					fmt.Fprintf(w, "\t\t\tvar __tmp_%d = BitConverter.GetBytes(%s);\n", tmpIdx, NameConv(f.Name))
					fmt.Fprintf(w, "\t\t\tif (!BitConverter.IsLittleEndian)\n\t\t\t{\n\t\t\t\tArray.Reverse(__tmp_%d);\n\t\t\t}\n", tmpIdx)
					fmt.Fprintf(w, "\t\t\tArray.Copy(__tmp_%d, 0, dst.value, %d, %d);\n", tmpIdx, f.Offset, f.TypeInfo.Size)
				}
			}
			tmpIdx++
		}
		fmt.Fprintf(w, "\n")
		if !IsFixed {
			fmt.Fprintf(w, "\t\t\tvar __index = (ulong)%d;\n", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
			for i, f := range s.DynamicFields {
				switch f.TypeInfo.FieldType {
				case ir.FieldType_STRING:
					fmt.Fprintf(w, "\t\t\tvar __tmp_%d = (ulong)System.Text.Encoding.UTF8.GetBytes(%s).Length +__index;\n", tmpIdx, NameConv(f.Name))
				case ir.FieldType_BYTES:
					fmt.Fprintf(w, "\t\t\tvar __tmp_%d = (ulong)%s.Length +__index;\n", tmpIdx, NameConv(f.Name))
				default:
					fmt.Fprintf(w, "\t\t\tvar __tmp_%d = (ulong)%s.value.Length +__index;\n", tmpIdx, NameConv(f.Name))
				}
				fmt.Fprintf(w, "\t\t\tvar __tmp_%d_len = BitConverter.GetBytes(__tmp_%d);\n", tmpIdx, tmpIdx)
				fmt.Fprintf(w, "\t\t\tif (!BitConverter.IsLittleEndian)\n\t\t\t{\n\t\t\t\tArray.Reverse(__tmp_%d_len);\n\t\t\t}\n", tmpIdx)
				fmt.Fprintf(w, "\t\t\tArray.Copy(__tmp_%d_len, 0, dst.value, %d, 8);\n", tmpIdx, s.DynamicFieldHeadOffsets[i])

				switch f.TypeInfo.FieldType {
				case ir.FieldType_STRING:
					fmt.Fprintf(w, "\t\t\tvar __buf_%d = System.Text.Encoding.UTF8.GetBytes(%s);\n", tmpIdx, NameConv(f.Name))
				case ir.FieldType_BYTES:
					fallthrough
				default:
					fmt.Fprintf(w, "\t\t\tvar __buf_%d = %s;\n", tmpIdx, NameConv(f.Name))
				}
				switch f.TypeInfo.FieldType {
				case ir.FieldType_STRUCT:
					fmt.Fprintf(w, "\t\t\tif (!BitConverter.IsLittleEndian)\n\t\t\t{\n\t\t\t\tArray.Reverse(__buf_%d.value);\n\t\t\t}\n", tmpIdx)
				default:
					fmt.Fprintf(w, "\t\t\tif (!BitConverter.IsLittleEndian)\n\t\t\t{\n\t\t\t\tArray.Reverse(__buf_%d);\n\t\t\t}\n", tmpIdx)
				}
				switch f.TypeInfo.FieldType {
				case ir.FieldType_STRUCT:
					fmt.Fprintf(w, "\t\t\tArray.Copy(__buf_%d.value, 0, dst.value, (long)__index, (long)__tmp_%d-(long)__index);\n", tmpIdx, tmpIdx)
				case ir.FieldType_BYTES:
					fallthrough
				case ir.FieldType_STRING:
					fmt.Fprintf(w, "\t\t\tArray.Copy(__buf_%d, 0, dst.value, (long)__index, (long)__tmp_%d-(long)__index);\n", tmpIdx, tmpIdx)
				}
				if i != len(s.DynamicFields)-1 {
					switch f.TypeInfo.FieldType {
					case ir.FieldType_STRING:
						fmt.Fprintf(w, "\t\t\t__index += (ulong)(ulong)System.Text.Encoding.UTF8.GetBytes(%s).Length;\n", NameConv(f.Name))
					case ir.FieldType_BYTES:
						fmt.Fprintf(w, "\t\t\t__index += (ulong)%s.Length;\n", NameConv(f.Name))
					default:
						fmt.Fprintf(w, "\t\t\t__index += (ulong)%s.value.Length;\n", NameConv(f.Name))
					}
				}
				tmpIdx++
			}
		}
		fmt.Fprintf(w, "\t\t\treturn dst;\n")
		fmt.Fprintf(w, "\t\t}\n\n")

		fmt.Fprintf(w, "\t\tpublic static %s New(", TypeConv(s.Name))
		for i, f := range allFields {
			if i != 0 {
				fmt.Fprintf(w, ", ")
			}
			fmt.Fprintf(w, "%s %s", TypeConv(f.Type), NameConv(f.Name))
		}
		fmt.Fprintf(w, ")\n\t\t{\n")
		fmt.Fprintf(w, "\t\t\tvar __vstruct__size = (ulong)%d ", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
		for _, f := range s.DynamicFields {
			switch f.TypeInfo.FieldType {
			case ir.FieldType_STRING:
				fmt.Fprintf(w, "+ (ulong)System.Text.Encoding.UTF8.GetBytes(%s).Length ", NameConv(f.Name))
			case ir.FieldType_BYTES:
				fmt.Fprintf(w, "+ (ulong)%s.Length ", NameConv(f.Name))
			default:
				fmt.Fprintf(w, "+ (ulong)%s.value.Length ", NameConv(f.Name))
			}
		}
		fmt.Fprintf(w, ";\n")
		fmt.Fprintf(w, "\t\t\tvar __vstruct__buf = %s.FromBytes(new byte[__vstruct__size]);\n", TypeConv(s.Name))
		fmt.Fprintf(w, "\t\t\t__vstruct__buf = %s.Serialize(__vstruct__buf", TypeConv(s.Name))
		for _, f := range allFields {
			fmt.Fprintf(w, ", %s", NameConv(f.Name))
		}
		fmt.Fprintf(w, ");\n")
		fmt.Fprintf(w, "\t\t\treturn __vstruct__buf;\n")
		fmt.Fprintf(w, "\t\t}\n\n")

		fmt.Fprintf(w, "\t\tpublic byte[] value = new byte[0];\n")

		for _, f := range s.FixedFields {
			fmt.Fprintf(w, "\t\tpublic %s Get%s()\n\t\t{\n", TypeConv(f.Type), NameConv(f.Name))
			switch f.TypeInfo.FieldType {
			case ir.FieldType_BOOL:
				fmt.Fprintf(w, "\t\t\treturn s[%d] != 0;\n", f.Offset)
			case ir.FieldType_UINT:
				fmt.Fprintf(w, "\t\t\tvar buf = new byte[%d];\n", f.TypeInfo.Size)
				fmt.Fprintf(w, "\t\t\tArray.Copy(this.value, %d, buf, 0, %d);\n", f.Offset, f.TypeInfo.Size)
				fmt.Fprint(w, "\t\t\tif (!BitConverter.IsLittleEndian)\n\t\t\t{\n\t\t\t\tArray.Reverse(buf);\n\t\t\t}\n")
				switch f.TypeInfo.Size {
				case 1:
					fmt.Fprintf(w, "\t\t\treturn ((byte)buf[0]);\n")
				case 2:
					fmt.Fprintf(w, "\t\t\treturn BitConverter.ToUInt16(buf, 0);\n")
				case 4:
					fmt.Fprintf(w, "\t\t\treturn BitConverter.ToUInt32(buf, 0);\n")
				case 8:
					fmt.Fprintf(w, "\t\t\treturn BitConverter.ToUInt64(buf, 0);\n")
				}
			case ir.FieldType_INT:
				fmt.Fprintf(w, "\t\t\tvar buf = new byte[%d];\n", f.TypeInfo.Size)
				fmt.Fprintf(w, "\t\t\tArray.Copy(this.value, %d, buf, 0, %d);\n", f.Offset, f.TypeInfo.Size)
				fmt.Fprint(w, "\t\t\tif (!BitConverter.IsLittleEndian)\n\t\t\t{\n\t\t\t\tArray.Reverse(buf);\n\t\t\t}\n")
				switch f.TypeInfo.Size {
				case 1:
					fmt.Fprintf(w, "\t\t\treturn ((sbyte)buf[0]);\n")
				case 2:
					fmt.Fprintf(w, "\t\t\treturn BitConverter.ToInt16(buf, 0);\n")
				case 4:
					fmt.Fprintf(w, "\t\t\treturn BitConverter.ToInt32(buf, 0);\n")
				case 8:
					fmt.Fprintf(w, "\t\t\treturn BitConverter.ToInt64(buf, 0);\n")
				}
			case ir.FieldType_FLOAT:
				fmt.Fprintf(w, "\t\t\tvar buf = new byte[%d];\n", f.TypeInfo.Size)
				fmt.Fprintf(w, "\t\t\tArray.Copy(this.value, %d, buf, 0, %d);\n", f.Offset, f.TypeInfo.Size)
				fmt.Fprint(w, "\t\t\tif (!BitConverter.IsLittleEndian)\n\t\t\t{\n\t\t\t\tArray.Reverse(buf);\n\t\t\t}\n")
				switch f.TypeInfo.Size {
				case 4:
					fmt.Fprintf(w, "\t\t\treturn BitConverter.ToSingle(buf, 0);\n")
				case 8:
					fmt.Fprintf(w, "\t\t\treturn BitConverter.ToDouble(buf, 0);\n")
				}
			case ir.FieldType_ENUM:
				if f.TypeInfo.Size != 1 {
					fmt.Fprintf(w, "\t\t\tvar buf = new Byte[%d];\n", f.TypeInfo.Size)
					fmt.Fprintf(w, "\t\t\tArray.Copy(this.value, %d, buf, 0, %d);\n", f.Offset, f.TypeInfo.Size)
					fmt.Fprint(w, "\t\t\tif (!BitConverter.IsLittleEndian)\n\t\t\t{\n\t\t\t\tArray.Reverse(buf);\n\t\t\t}\n")
					switch f.TypeInfo.Size {
					case 2:
						fmt.Fprintf(w, "\t\t\treturn (%s)((ushort)BitConverter.ToUInt16(buf, 0));\n", NameConv(f.Type))
					case 4:
						fmt.Fprintf(w, "\t\t\treturn (%s)((uint)BitConverter.ToUInt32(buf, 0));\n", NameConv(f.Type))
					case 8:
						fmt.Fprintf(w, "\t\t\treturn (%s)((ulong)BitConverter.ToUInt64(buf, 0));\n", NameConv(f.Type))
					}
				} else {
					fmt.Fprintf(w, "\t\t\treturn (%s)(this.value[%d]);\n", TypeConv(f.Type), f.Offset)
				}
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, "\t\t\tvar buf = new Byte[%d];\n", f.TypeInfo.Size)
				fmt.Fprintf(w, "\t\t\tArray.Copy(this.value, %d, buf, 0, %d);\n", f.Offset, f.TypeInfo.Size)
				fmt.Fprintf(w, "\t\t\treturn %s.FromBytes(buf);\n", TypeConv(f.Type))
			}
			fmt.Fprintf(w, "\t\t}\n\n")
		}

		for i, f := range s.DynamicFields {
			fmt.Fprintf(w, "\t\tpublic %s Get%s()\n\t\t{\n", TypeConv(f.Type), NameConv(f.Name))
			if i == 0 {
				fmt.Fprintf(w, "\t\t\tulong __off0 = ")
				fmt.Fprintf(w, "%d;\n", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
			} else {
				fmt.Fprintf(w, "\t\t\tvar buf_%d_1 = new byte[8];\n", i)
				fmt.Fprintf(w, "\t\t\tArray.Copy(this.value, %d, buf_%d_1, 0, 8);\n", s.DynamicFieldHeadOffsets[i]-8, i)
				fmt.Fprintf(w, "\t\t\tif (!BitConverter.IsLittleEndian)\n\t\t\t{\n\t\t\t\tArray.Reverse(buf_%d_1);\n\t\t\t}\n", i)
				fmt.Fprintf(w, "\t\t\tulong __off0 = ")
				fmt.Fprintf(w, "((ulong)BitConverter.ToUInt64(buf_%d_1, 0));\n", i)
			}
			fmt.Fprintf(w, "\t\t\tvar buf_%d_2 = new byte[8];\n", i)
			fmt.Fprintf(w, "\t\t\tArray.Copy(this.value, %d, buf_%d_2, 0, 8);\n", s.DynamicFieldHeadOffsets[i], i)
			fmt.Fprintf(w, "\t\t\tif (!BitConverter.IsLittleEndian)\n\t\t\t{\n\t\t\t\tArray.Reverse(buf_%d_2);\n\t\t\t}\n", i)
			fmt.Fprintf(w, "\t\t\tulong __off1 = ")
			fmt.Fprintf(w, "((ulong)BitConverter.ToUInt64(buf_%d_2, 0));\n", i)
			switch f.TypeInfo.FieldType {
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, "\t\t\tvar buf_%d_3 = new Byte[__off1-__off0];\n", i)
				fmt.Fprintf(w, "\t\t\tArray.Copy(this.value, (long)__off0, buf_%d_3, 0, (long)__off1-(long)__off0);\n", i)
				fmt.Fprintf(w, "\t\t\treturn %s.FromBytes(buf_%d_3);\n", TypeConv(f.Type), i)
			case ir.FieldType_STRING:
				fmt.Fprintf(w, "\t\t\tvar buf_%d_3 = new Byte[__off1-__off0];\n", i)
				fmt.Fprintf(w, "\t\t\tArray.Copy(this.value, (long)__off0, buf_%d_3, 0, (long)__off1-(long)__off0);\n", i)
				fmt.Fprintf(w, "\t\t\treturn ((%s)System.Text.Encoding.UTF8.GetString(buf_%d_3));\n", TypeConv(f.Type), i)
			case ir.FieldType_BYTES:
				fmt.Fprintf(w, "\t\t\tvar buf_%d_3 = new Byte[__off1-__off0];\n", i)
				fmt.Fprintf(w, "\t\t\tArray.Copy(this.value, (long)__off0, buf_%d_3, 0, (long)__off1-(long)__off0);\n", i)
				fmt.Fprintf(w, "\t\t\treturn (%s)buf_%d_3;\n", TypeConv(f.Type), i)
			}
			fmt.Fprintf(w, "\t\t}\n\n")
		}
		fmt.Fprintf(w, "\t}\n\n")

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

		// fmt.Fprintf(w, "func (s %s) String() string {\n", NameConv(s.Name))
		// fmt.Fprintf(w, "if !s.Vstruct_Validate() {\n")
		// fmt.Fprintf(w, "return \"%s (invalid)\"\n", NameConv(s.Name))
		// fmt.Fprintf(w, "}\n")
		// fmt.Fprintf(w, "var __b strings.Builder\n")
		// fmt.Fprintf(w, "__b.WriteString(\"%s {\")\n", NameConv(s.Name))
		// var allFields []*ir.Field
		// allFields = append(allFields, s.FixedFields...)
		// allFields = append(allFields, s.DynamicFields...)
		// for i, f := range allFields {
		// 	if i != 0 {
		// 		fmt.Fprintf(w, "__b.WriteString(\", \")\n")
		// 	}
		// 	fmt.Fprintf(w, "__b.WriteString(\"%s: \")\n", NameConv(f.Name))
		// 	switch f.TypeInfo.FieldType {
		// 	case ir.FieldType_STRUCT:
		// 		fmt.Fprintf(w, "__b.WriteString(s.%s().String())\n", NameConv(f.Name))
		// 	case ir.FieldType_STRING:
		// 		if f.Type == "string" {
		// 			fmt.Fprintf(w, "__b.WriteString(strconv.Quote(s.%s()))\n", NameConv(f.Name))
		// 		} else {
		// 			fmt.Fprintf(w, "__b.WriteString(strconv.Quote(string(s.%s())))\n", NameConv(f.Name))
		// 		}
		// 	case ir.FieldType_BYTES:
		// 		fmt.Fprintf(w, "__b.WriteString(fmt.Sprint(s.%s()))\n", NameConv(f.Name))
		// 	case ir.FieldType_BOOL:
		// 		fmt.Fprintf(w, "__b.WriteString(strconv.FormatBool(s.%s()))\n", NameConv(f.Name))
		// 	case ir.FieldType_INT:
		// 		fmt.Fprintf(w, "__b.WriteString(strconv.FormatInt(int64(s.%s()), 10))\n", NameConv(f.Name))
		// 	case ir.FieldType_UINT:
		// 		fmt.Fprintf(w, "__b.WriteString(strconv.FormatUint(uint64(s.%s()), 10))\n", NameConv(f.Name))
		// 	case ir.FieldType_FLOAT:
		// 		fmt.Fprintf(w, "__b.WriteString(strconv.FormatFloat(float64(s.%s()), 'g', -1, %d))\n", NameConv(f.Name), f.TypeInfo.Size)
		// 	case ir.FieldType_ENUM:
		// 		fmt.Fprintf(w, "__b.WriteString(s.%s().String())\n", NameConv(f.Name))
		// 	}
		// }
		// fmt.Fprintf(w, "__b.WriteString(\"}\")\n")
		// fmt.Fprintf(w, "return __b.String()\n")
		// fmt.Fprintf(w, "}\n\n")
	}
}
