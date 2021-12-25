package dart

import (
	"fmt"
	"io"

	"github.com/lemon-mint/vstruct/ir"
)

func writeStructs(w io.Writer, i *ir.IR) {
	for _, s := range i.Structs {
		fmt.Fprintf(w, "class %s {\n", NameConv(s.Name))
		fmt.Fprintf(w, "  Uint8List vstruct__buf = Uint8List(0);\n\n")
		var allFields []*ir.Field
		allFields = append(allFields, s.FixedFields...)
		allFields = append(allFields, s.DynamicFields...)
		fmt.Fprintf(w, "  %s(", TypeConv(s.Name))
		for i, f := range allFields {
			if i != 0 {
				fmt.Fprintf(w, ", ")
			}
			fmt.Fprintf(w, "%s %s", TypeConv(f.Type), NameConv(f.Name))
		}
		fmt.Fprintf(w, ") {\n")
		fmt.Fprintf(w, "    int __vstruct__size = %d", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
		for _, f := range s.DynamicFields {
			switch f.TypeInfo.FieldType {
			case ir.FieldType_BYTES:
				fmt.Fprintf(w, " + %s.lengthInBytes", NameConv(f.Name))
			case ir.FieldType_STRING:
				fmt.Fprintf(w, " + %s.length", NameConv(f.Name))
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, " + %s.lengthInBytes", NameConv(f.Name))
			}
		}
		fmt.Fprintf(w, ";\n")
		fmt.Fprintf(w, "    vstruct__buf = Uint8List(__vstruct__size);\n")
		fmt.Fprintf(w, "    vstruct__buf = Serialize(vstruct__buf")
		for _, f := range allFields {
			fmt.Fprintf(w, ", %s", NameConv(f.Name))
		}
		fmt.Fprintf(w, ");\n")
		fmt.Fprintf(w, "  }\n\n")
		fmt.Fprintf(w, "  int get lengthInBytes => vstruct__buf.lengthInBytes;\n\n")

		fmt.Fprintf(w, "  Uint8List as_bytes_mut() {\n")
		fmt.Fprintf(w, "    return vstruct__buf;\n")
		fmt.Fprintf(w, "  }\n\n")

		fmt.Fprintf(w, "  %s.fromBytes(Uint8List b) {\n", TypeConv(s.Name))
		fmt.Fprintf(w, "    vstruct__buf = b;\n")
		fmt.Fprintf(w, "  }\n\n")

		fmt.Fprintf(w, "  Uint8List Serialize(Uint8List dst")
		for _, f := range allFields {
			fmt.Fprintf(w, ", %s %s", TypeConv(f.Type), NameConv(f.Name))
		}

		var IsFixed bool = len(s.DynamicFields) == 0

		fmt.Fprintf(w, ") {\n")
		var tmpIdx int = 0
		for _, f := range s.FixedFields {
			switch f.TypeInfo.FieldType {
			case ir.FieldType_STRUCT:
				//fmt.Fprintf(w, "copy(dst[%d:%d], %s)\n", f.Offset, f.Offset+f.TypeInfo.Size, NameConv(f.Name))
				fmt.Fprintf(w, "    Uint8List __tmp_%d = %s.as_bytes_mut();\n", tmpIdx, NameConv(f.Name))
				fmt.Fprintf(w, "    for (int i = 0; i < %s.lengthInBytes; i++) {\n", NameConv(f.Name))
				fmt.Fprintf(w, "      dst[%d + i] = __tmp_%d[i];\n", f.Offset, tmpIdx)
				fmt.Fprintf(w, "    }\n")
			case ir.FieldType_BOOL:
				fmt.Fprintf(w, "	dst[%d] = %s ? 1 : 0;\n", f.Offset, NameConv(f.Name))
			case ir.FieldType_UINT, ir.FieldType_INT, ir.FieldType_FLOAT:
				fmt.Fprintf(w, "    Uint8List __tmp_%d = %s.toBytes();\n", tmpIdx, NameConv(f.Name))
				fmt.Fprintf(w, "    for (int i = 0; i < %d; i++) {\n", f.TypeInfo.Size)
				fmt.Fprintf(w, "      dst[%d + i] = __tmp_%d[i];\n", f.Offset, tmpIdx)
				fmt.Fprintf(w, "    }\n")
			case ir.FieldType_ENUM:
				if f.TypeInfo.Size == 1 {
					fmt.Fprintf(w, "    dst[%d] = %s.index;\n", f.Offset, NameConv(f.Name))
				} else {
					fmt.Fprintf(w, "    Uint8List __tmp_%d = U%d(%s.index).toBytes();\n", tmpIdx, f.TypeInfo.Size, NameConv(f.Name))
					fmt.Fprintf(w, "    for (int i = 0; i < %d; i++) {\n", f.TypeInfo.Size)
					fmt.Fprintf(w, "      dst[%d + i] = __tmp_%d[i];\n", f.Offset, tmpIdx)
					fmt.Fprintf(w, "    }\n")
				}
			}
			tmpIdx++
			fmt.Fprintf(w, "\n")
		}
		if !IsFixed {
			fmt.Fprintf(w, "    U64 __index = U64(%d);\n", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
			for i, f := range s.DynamicFields {
				//fmt.Fprintf(w, "__tmp_%d := uint64(len(%s)) +__index\n", tmpIdx, NameConv(f.Name))
				switch f.TypeInfo.FieldType {
				case ir.FieldType_BYTES:
					fmt.Fprintf(w, "    Uint8List __tmp_%d = (U64(%s.lengthInBytes) + __index).toBytes();\n", tmpIdx, NameConv(f.Name))
				case ir.FieldType_STRING:
					fmt.Fprintf(w, "    Uint8List __tmp_%d = (U64(%s.length) + __index).toBytes();\n", tmpIdx, NameConv(f.Name))
				case ir.FieldType_STRUCT:
					fmt.Fprintf(w, "    Uint8List __tmp_%d = (U64(%s.lengthInBytes) + __index).toBytes();\n", tmpIdx, NameConv(f.Name))
				}
				/*
					for j := 0; j < 8; j++ {
						if j == 0 {
							fmt.Fprintf(w, "dst[%d] = byte(__tmp_%d)\n", s.DynamicFieldHeadOffsets[i]+j, tmpIdx)
						} else {
							fmt.Fprintf(w, "dst[%d] = byte(__tmp_%d >> %d)\n", s.DynamicFieldHeadOffsets[i]+j, tmpIdx, 8*j)
						}
					}
				*/
				fmt.Fprintf(w, "    for (int i = 0; i < 8; i++) {\n")
				fmt.Fprintf(w, "      dst[%d + i] = __tmp_%d[i];\n", s.DynamicFieldHeadOffsets[i], tmpIdx)
				fmt.Fprintf(w, "    }\n")
				/*
					switch f.TypeInfo.FieldType {
					case ir.FieldType_STRUCT:
						fmt.Fprintf(w, "copy(dst[__index:__tmp_%d], %s)\n", tmpIdx, NameConv(f.Name))
					case ir.FieldType_BYTES:
						fmt.Fprintf(w, "copy(dst[__index:__tmp_%d], %s)\n", tmpIdx, NameConv(f.Name))
					case ir.FieldType_STRING:
						fmt.Fprintf(w, "copy(dst[__index:__tmp_%d], %s)\n", tmpIdx, NameConv(f.Name))
					}
				*/
				tmpIdx++
				switch f.TypeInfo.FieldType {
				case ir.FieldType_STRUCT:
					fmt.Fprintf(w, "    Uint8List __tmp_%d = %s.as_bytes_mut();\n", tmpIdx, NameConv(f.Name))
					fmt.Fprintf(w, "    for (int i = 0; i < %s.lengthInBytes; i++) {\n", NameConv(f.Name))
					fmt.Fprintf(w, "      dst[(__index + U64(i)).value.toInt()] = __tmp_%d[i];\n", tmpIdx)
					fmt.Fprintf(w, "    }\n")
				case ir.FieldType_BYTES:
					fmt.Fprintf(w, "    for (int i = 0; i < %s.lengthInBytes; i++) {\n", NameConv(f.Name))
					fmt.Fprintf(w, "      dst[(__index + U64(i)).value.toInt()] = %s[i];\n", NameConv(f.Name))
					fmt.Fprintf(w, "    }\n")
				case ir.FieldType_STRING:
					fmt.Fprintf(w, "    List<int> __tmp_%d = utf8.encode(%s);\n", tmpIdx, NameConv(f.Name))
					tmpIdx++
					fmt.Fprintf(w, "    Uint8List __tmp_%d = Uint8List.fromList(__tmp_%d);\n", tmpIdx, tmpIdx-1)
					fmt.Fprintf(w, "    for (int i = 0; i < %s.length; i++) {\n", NameConv(f.Name))
					fmt.Fprintf(w, "      dst[(__index + U64(i)).value.toInt()] = __tmp_%d[i];\n", tmpIdx)
					fmt.Fprintf(w, "    }\n")
				}

				if i != len(s.DynamicFields)-1 {
					switch f.TypeInfo.FieldType {
					case ir.FieldType_STRUCT:
						fmt.Fprintf(w, "    __index = __index + U64(%s.lengthInBytes);\n", NameConv(f.Name))
					case ir.FieldType_BYTES:
						fmt.Fprintf(w, "	__index = __index + U64(%s.lengthInBytes);\n", NameConv(f.Name))
					case ir.FieldType_STRING:
						fmt.Fprintf(w, "    __index = __index + U64(%s.length);\n", NameConv(f.Name))
					}
				}
				tmpIdx++
			}
		}
		fmt.Fprintf(w, "    return dst;\n")
		fmt.Fprintf(w, "  }\n\n")
		/*
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
		*/
		for _, f := range s.FixedFields {
			fmt.Fprintf(w, "  %s %s() {\n", TypeConv(f.Type), NameConv(f.Name))
			switch f.TypeInfo.FieldType {
			case ir.FieldType_BOOL:
				fmt.Fprintf(w, "    return s[%d] != 0;\n", f.Offset)
			case ir.FieldType_UINT:
				fmt.Fprintf(w, "    U%d _value = U%d.fromBytes(vstruct__buf.sublist(%d, %d));\n", f.TypeInfo.Size*8, f.TypeInfo.Size*8, f.Offset, f.Offset+f.TypeInfo.Size)
				fmt.Fprintf(w, "    return _value;\n")
			case ir.FieldType_INT:
				fmt.Fprintf(w, "    I%d _value = I%d.fromBytes(vstruct__buf.sublist(%d, %d));\n", f.TypeInfo.Size*8, f.TypeInfo.Size*8, f.Offset, f.Offset+f.TypeInfo.Size)
				fmt.Fprintf(w, "    return _value;\n")
			case ir.FieldType_FLOAT:
				fmt.Fprintf(w, "    F%d _value = F%d.fromBytes(vstruct__buf.sublist(%d, %d));\n", f.TypeInfo.Size*8, f.TypeInfo.Size*8, f.Offset, f.Offset+f.TypeInfo.Size)
				fmt.Fprintf(w, "    return _value;\n")
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, "    return %s.fromBytes(vstruct__buf.sublist(%d, %d));\n", TypeConv(f.Type), f.Offset, f.Offset+f.TypeInfo.Size)
			case ir.FieldType_ENUM:
				if f.TypeInfo.Size == 1 {
					fmt.Fprintf(w, "    return %s.values[vstruct__buf[%d]];\n", TypeConv(f.Type), f.Offset)
				} else {
					fmt.Fprintf(w, "    U%d _value = U%d.fromBytes(vstruct__buf.sublist(%d, %d));\n", f.TypeInfo.Size*8, f.TypeInfo.Size*8, f.Offset, f.Offset+f.TypeInfo.Size)
					fmt.Fprintf(w, "    return %s.values[_value.value];\n", TypeConv(f.Type))
				}
			}
			fmt.Fprintf(w, "  }\n\n")
		}

		/*
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
		*/

		fmt.Fprintf(w, "}\n\n")
	}
}
