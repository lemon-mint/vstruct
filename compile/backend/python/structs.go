package python

import (
	"fmt"
	"io"

	"github.com/lemon-mint/vstruct/ir"
)

func writeStructs(w io.Writer, i *ir.IR) {
	for _, s := range i.Structs {
		fmt.Fprintf(w, "class %s():\n", NameConv(s.Name))
		var allFields []*ir.Field
		allFields = append(allFields, s.FixedFields...)
		allFields = append(allFields, s.DynamicFields...)
		fmt.Fprintf(w, "    def __init__(self, ")
		for i, f := range allFields {
			if i != 0 {
				fmt.Fprintf(w, ", ")
			}
			//fmt.Fprintf(w, "%s _%s", TypeConv(f.Type), NameConv(f.Name))
			fmt.Fprintf(w, "%s: %s", NameConv(f.Name), TypeConv(f.Type))
		}
		fmt.Fprintf(w, ") {\n")
		//fmt.Fprintf(w, "  Uint8List vData = Uint8List(0);\n\n")
		//fmt.Fprintf(w, "        self.vData = bytearray(0)\n")
		//fmt.Fprintf(w, "    int vSize = %d", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
		fmt.Fprintf(w, "        vSize = %d", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
		for _, f := range s.DynamicFields {
			switch f.TypeInfo.FieldType {
			case ir.FieldType_BYTES:
				//fmt.Fprintf(w, " + _%s.lengthInBytes", NameConv(f.Name))
				fmt.Fprintf(w, " + len(%s)", NameConv(f.Name))
			case ir.FieldType_STRING:
				//fmt.Fprintf(w, " + _%s.length", NameConv(f.Name))
				fmt.Fprintf(w, " + len(%s)", NameConv(f.Name))
			case ir.FieldType_STRUCT:
				//fmt.Fprintf(w, " + _%s.lengthInBytes", NameConv(f.Name))
				fmt.Fprintf(w, " + len(%s)", NameConv(f.Name))
			}
		}
		//fmt.Fprintf(w, ";\n")
		fmt.Fprintf(w, "\n")
		//fmt.Fprintf(w, "        vData = Uint8List(vSize);\n")
		fmt.Fprintf(w, "        vData = bytearray(vSize)\n")
		//fmt.Fprintf(w, "        vData = vSerialize(vData")
		fmt.Fprintf(w, "        vData = self.vSerialize(self.vData")
		for _, f := range allFields {
			//fmt.Fprintf(w, ", _%s", NameConv(f.Name))
			fmt.Fprintf(w, ", %s", NameConv(f.Name))
		}
		fmt.Fprintf(w, ")\n")
		fmt.Fprintf(w, "\n\n")
		//fmt.Fprintf(w, "  int get lengthInBytes => vData.lengthInBytes;\n\n")
		fmt.Fprintf(w, "    def __len__(self) -> int:\n")
		fmt.Fprintf(w, "        return len(self.vData)\n\n")

		//fmt.Fprintf(w, "  Uint8List toBytes() {\n")
		//fmt.Fprintf(w, "    return vData;\n")
		//fmt.Fprintf(w, "  }\n\n")
		fmt.Fprintf(w, "    def toBytes(self) -> bytearray:\n")
		fmt.Fprintf(w, "        return self.vData\n\n")

		//fmt.Fprintf(w, "  %s.fromBytes(Uint8List b) {\n", TypeConv(s.Name))
		//fmt.Fprintf(w, "    vData = b;\n")
		//fmt.Fprintf(w, "  }\n\n")
		fmt.Fprintf(w, "    @staticmethod\n")
		fmt.Fprintf(w, "    def fromBytes(b: bytearray) -> %s:\n", TypeConv(s.Name))
		fmt.Fprintf(w, "        self = %s.__new__()\n", TypeConv(s.Name))
		fmt.Fprintf(w, "        self.vData = b\n")
		fmt.Fprintf(w, "        return self\n\n")

		//fmt.Fprintf(w, "  Uint8List vSerialize(Uint8List dst")
		fmt.Fprintf(w, "    def vSerialize(self, dst: bytearray")
		for _, f := range allFields {
			//fmt.Fprintf(w, ", %s %s", TypeConv(f.Type), NameConv(f.Name))
			fmt.Fprintf(w, ", %s: %s", NameConv(f.Name), TypeConv(f.Type))
		}
		//fmt.Fprintf(w, ") {\n")
		fmt.Fprintf(w, ") -> bytearray:\n")

		var IsFixed bool = len(s.DynamicFields) == 0
		var tmpIdx int = 0
		for _, f := range s.FixedFields {
			switch f.TypeInfo.FieldType {
			case ir.FieldType_STRUCT:
				//fmt.Fprintf(w, "    Uint8List __tmp_%d = %s.toBytes();\n", tmpIdx, NameConv(f.Name))
				//fmt.Fprintf(w, "    for (int i = 0; i < %s.lengthInBytes; i++) {\n", NameConv(f.Name))
				//fmt.Fprintf(w, "      dst[%d + i] = __tmp_%d[i];\n", f.Offset, tmpIdx)
				//fmt.Fprintf(w, "    }\n")
				fmt.Fprintf(w, "        __tmp_%d = %s.toBytes()\n", tmpIdx, NameConv(f.Name))
				fmt.Fprintf(w, "        dst[%d:%d+len(__tmp_%d)] = __tmp_%d\n", f.Offset, f.Offset, tmpIdx, tmpIdx)
			case ir.FieldType_BOOL:
				//fmt.Fprintf(w, "	dst[%d] = %s ? 1 : 0;\n", f.Offset, NameConv(f.Name))
				fmt.Fprintf(w, "        dst[%d] = 1 if %s else 0\n", f.Offset, NameConv(f.Name))
			case ir.FieldType_UINT, ir.FieldType_INT, ir.FieldType_FLOAT:
				//fmt.Fprintf(w, "    Uint8List __tmp_%d = %s.toBytes();\n", tmpIdx, NameConv(f.Name))
				//fmt.Fprintf(w, "    for (int i = 0; i < %d; i++) {\n", f.TypeInfo.Size)
				//fmt.Fprintf(w, "      dst[%d + i] = __tmp_%d[i];\n", f.Offset, tmpIdx)
				//fmt.Fprintf(w, "    }\n")
				fmt.Fprintf(w, "        __tmp_%d = %s.to_bytes(%d, 'little')\n", tmpIdx, NameConv(f.Name), f.TypeInfo.Size)
				fmt.Fprintf(w, "        dst[%d:%d] = __tmp_%d\n", f.Offset, f.Offset+f.TypeInfo.Size, tmpIdx)
			case ir.FieldType_ENUM:
				if f.TypeInfo.Size == 1 {
					//fmt.Fprintf(w, "    dst[%d] = %s.index;\n", f.Offset, NameConv(f.Name))
					fmt.Fprintf(w, "        dst[%d] = %s.value\n", f.Offset, NameConv(f.Name))
				} else {
					//fmt.Fprintf(w, "    Uint8List __tmp_%d = U%d(%s.index).toBytes();\n", tmpIdx, f.TypeInfo.Size, NameConv(f.Name))
					//fmt.Fprintf(w, "    for (int i = 0; i < %d; i++) {\n", f.TypeInfo.Size)
					//fmt.Fprintf(w, "      dst[%d + i] = __tmp_%d[i];\n", f.Offset, tmpIdx)
					//fmt.Fprintf(w, "    }\n")
					fmt.Fprintf(w, "        __tmp_%d = %s.value.to_bytes(%d, 'little')\n", tmpIdx, NameConv(f.Name), f.TypeInfo.Size)
					fmt.Fprintf(w, "        dst[%d:%d] = __tmp_%d\n", f.Offset, f.Offset+f.TypeInfo.Size, tmpIdx)
				}
			}
			tmpIdx++
			fmt.Fprintf(w, "\n")
		}
		if !IsFixed {
			//fmt.Fprintf(w, "    U64 __index = U64(%d);\n", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
			fmt.Fprintf(w, "        __index = %d\n", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
			for i, f := range s.DynamicFields {
				switch f.TypeInfo.FieldType {
				case ir.FieldType_BYTES:
					//fmt.Fprintf(w, "    Uint8List __tmp_%d = (U64(%s.lengthInBytes) + __index).toBytes();\n", tmpIdx, NameConv(f.Name))
					fmt.Fprintf(w, "        __tmp_%d = (len(%s) + __index).to_bytes(8, 'little')\n", tmpIdx, NameConv(f.Name))
				case ir.FieldType_STRING:
					//fmt.Fprintf(w, "    Uint8List __tmp_%d = (U64(%s.length) + __index).toBytes();\n", tmpIdx, NameConv(f.Name))
					fmt.Fprintf(w, "    	__tmp_%d = (len(%s) + __index).to_bytes(8, 'little')\n", tmpIdx, NameConv(f.Name))
				case ir.FieldType_STRUCT:
					//fmt.Fprintf(w, "    Uint8List __tmp_%d = (U64(%s.lengthInBytes) + __index).toBytes();\n", tmpIdx, NameConv(f.Name))
					fmt.Fprintf(w, "    	__tmp_%d = (%s.__len__() + __index).to_bytes(8, 'little')\n", tmpIdx, NameConv(f.Name))
				}
				//fmt.Fprintf(w, "    for (int i = 0; i < 8; i++) {\n")
				//fmt.Fprintf(w, "      dst[%d + i] = __tmp_%d[i];\n", s.DynamicFieldHeadOffsets[i], tmpIdx)
				//fmt.Fprintf(w, "    }\n")
				fmt.Fprintf(w, "    	dst[%d:%d] = __tmp_%d\n", s.DynamicFieldHeadOffsets[i], s.DynamicFieldHeadOffsets[i]+8, tmpIdx)
				tmpIdx++
				switch f.TypeInfo.FieldType {
				case ir.FieldType_STRUCT:
					//fmt.Fprintf(w, "    Uint8List __tmp_%d = %s.toBytes();\n", tmpIdx, NameConv(f.Name))
					//fmt.Fprintf(w, "    for (int i = 0; i < %s.lengthInBytes; i++) {\n", NameConv(f.Name))
					//fmt.Fprintf(w, "      dst[(__index + U64(i)).value.toInt()] = __tmp_%d[i];\n", tmpIdx)
					//fmt.Fprintf(w, "    }\n")
					fmt.Fprintf(w, "    	__tmp_%d = %s.toBytes()\n", tmpIdx, NameConv(f.Name))
					fmt.Fprintf(w, "    	dst[__index:__index+len(__tmp_%d)] = __tmp_%d\n", tmpIdx, tmpIdx)
				case ir.FieldType_BYTES:
					//fmt.Fprintf(w, "    for (int i = 0; i < %s.lengthInBytes; i++) {\n", NameConv(f.Name))
					//fmt.Fprintf(w, "      dst[(__index + U64(i)).value.toInt()] = %s[i];\n", NameConv(f.Name))
					//fmt.Fprintf(w, "    }\n")
					fmt.Fprintf(w, "    	dst[__index:__index+len(%s)] = %s\n", NameConv(f.Name), NameConv(f.Name))
				case ir.FieldType_STRING:
					//fmt.Fprintf(w, "    List<int> __tmp_%d = utf8.encode(%s);\n", tmpIdx, NameConv(f.Name))
					//tmpIdx++
					//fmt.Fprintf(w, "    Uint8List __tmp_%d = Uint8List.fromList(__tmp_%d);\n", tmpIdx, tmpIdx-1)
					//fmt.Fprintf(w, "    for (int i = 0; i < %s.length; i++) {\n", NameConv(f.Name))
					//fmt.Fprintf(w, "      dst[(__index + U64(i)).value.toInt()] = __tmp_%d[i];\n", tmpIdx)
					//fmt.Fprintf(w, "    }\n")
					fmt.Fprintf(w, "        __tmp_%d = %s.encode('utf-8')\n", tmpIdx, NameConv(f.Name))
					fmt.Fprintf(w, "    	dst[__index:__index+len(__tmp_%d)] = __tmp_%d\n", tmpIdx, tmpIdx)
				}

				if i != len(s.DynamicFields)-1 {
					switch f.TypeInfo.FieldType {
					case ir.FieldType_STRUCT:
						//fmt.Fprintf(w, "    __index = __index + U64(%s.lengthInBytes);\n", NameConv(f.Name))
						fmt.Fprintf(w, "    	__index = __index + %s.__len__()\n", NameConv(f.Name))
					case ir.FieldType_BYTES:
						//fmt.Fprintf(w, "	__index = __index + U64(%s.lengthInBytes);\n", NameConv(f.Name))
						fmt.Fprintf(w, "		__index = __index + len(%s)\n", NameConv(f.Name))
					case ir.FieldType_STRING:
						//fmt.Fprintf(w, "    __index = __index + U64(%s.length);\n", NameConv(f.Name))
						fmt.Fprintf(w, "    	__index = __index + len(%s)\n", NameConv(f.Name))
					}
				}
				tmpIdx++
			}
		}
		//fmt.Fprintf(w, "    return dst;\n")
		fmt.Fprintf(w, "    	return dst\n\n")
		//fmt.Fprintf(w, "  }\n\n")

		for _, f := range s.FixedFields {
			//fmt.Fprintf(w, "  %s get %s {\n", TypeConv(f.Type), NameConv(f.Name))
			fmt.Fprintf(w, "    @property\n")
			fmt.Fprintf(w, "    def %s(self) -> %s:\n", NameConv(f.Name), TypeConv(f.Type))
			switch f.TypeInfo.FieldType {
			case ir.FieldType_BOOL:
				fmt.Fprintf(w, "    	return s[%d] != 0;\n", f.Offset)
			case ir.FieldType_UINT:
				//fmt.Fprintf(w, "    U%d _value = U%d.fromBytes(vData.sublist(%d, %d));\n", f.TypeInfo.Size*8, f.TypeInfo.Size*8, f.Offset, f.Offset+f.TypeInfo.Size)
				//fmt.Fprintf(w, "    return _value;\n")
				fmt.Fprintf(w, "    	return %s.from_bytes(s[%d:%d], byteorder='little')\n", TypeConv(f.Type), f.Offset, f.Offset+f.TypeInfo.Size)
			case ir.FieldType_INT:
				fmt.Fprintf(w, "    	return %s.from_bytes(s[%d:%d], byteorder='little')\n", TypeConv(f.Type), f.Offset, f.Offset+f.TypeInfo.Size)
			case ir.FieldType_FLOAT:
				//fmt.Fprintf(w, "    F%d _value = F%d.fromBytes(vData.sublist(%d, %d));\n", f.TypeInfo.Size*8, f.TypeInfo.Size*8, f.Offset, f.Offset+f.TypeInfo.Size)
				//fmt.Fprintf(w, "    return _value;\n")
				switch f.TypeInfo.Size {
				case 4:
					fmt.Fprintf(w, "    	return struct.unpack('<f', s[%d:%d])[0]\n", f.Offset, f.Offset+f.TypeInfo.Size)
				case 8:
					fmt.Fprintf(w, "    	return struct.unpack('<d', s[%d:%d])[0]\n", f.Offset, f.Offset+f.TypeInfo.Size)
				}
			case ir.FieldType_STRUCT:
				//fmt.Fprintf(w, "    return %s.fromBytes(vData.sublist(%d, %d));\n", TypeConv(f.Type), f.Offset, f.Offset+f.TypeInfo.Size)
				fmt.Fprintf(w, "    	return %s.fromBytes(s[%d:%d])\n", TypeConv(f.Type), f.Offset, f.Offset+f.TypeInfo.Size)
			case ir.FieldType_ENUM:
				if f.TypeInfo.Size == 1 {
					//fmt.Fprintf(w, "    return %s.values[vData[%d]];\n", TypeConv(f.Type), f.Offset)
					fmt.Fprintf(w, "    return %s(s[%d])\n", TypeConv(f.Type), f.Offset)
				} else {
					//fmt.Fprintf(w, "    U%d _value = U%d.fromBytes(vData.sublist(%d, %d));\n", f.TypeInfo.Size*8, f.TypeInfo.Size*8, f.Offset, f.Offset+f.TypeInfo.Size)
					//fmt.Fprintf(w, "    return %s.values[_value.value];\n", TypeConv(f.Type))
					fmt.Fprintf(w, "    _value = %s.from_bytes(s[%d:%d], byteorder='little')\n", TypeConv(f.Type), f.Offset, f.Offset+f.TypeInfo.Size)
					fmt.Fprintf(w, "    return %s(_value)\n", TypeConv(f.Type))
				}
			}
			fmt.Fprintf(w, "\n")
		}

		for i, f := range s.DynamicFields {
			fmt.Fprintf(w, "  %s get %s {\n", TypeConv(f.Type), NameConv(f.Name))
			fmt.Fprintf(w, "    U64 __off0 = ")
			if i == 0 {
				fmt.Fprintf(w, "U64(%d);", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
			} else {
				fmt.Fprintf(w, "U64.fromBytes(vData.sublist(%d, %d));", s.DynamicFieldHeadOffsets[i]-8, s.DynamicFieldHeadOffsets[i])
			}
			fmt.Fprintf(w, "\n    U64 __off1 = ")
			fmt.Fprintf(w, "U64.fromBytes(vData.sublist(%d, %d));\n", s.DynamicFieldHeadOffsets[i], s.DynamicFieldHeadOffsets[i]+8)
			switch f.TypeInfo.FieldType {
			case ir.FieldType_STRUCT:
				fmt.Fprintf(w, "\n    return %s.fromBytes(vData.sublist(__off0.value.toInt(), __off1.value.toInt()));\n", TypeConv(f.Type))
			case ir.FieldType_STRING:
				fmt.Fprintf(w, "\n    return utf8.decode(vData.sublist(__off0.value.toInt(), __off1.value.toInt()));\n")
			case ir.FieldType_BYTES:
				fmt.Fprintf(w, "\n    return vData.sublist(__off0.value.toInt(), __off1.value.toInt());\n")
			}
			fmt.Fprintf(w, "\n")
		}

		fmt.Fprintf(w, "  bool vStructValidate() {\n")

		if s.IsFixed && len(s.DynamicFields) == 0 {
			fmt.Fprintf(w, "    return vData.lengthInBytes >= %d;\n", s.TotalFixedFieldSize)
		} else {
			fmt.Fprintf(w, "    if (vData.lengthInBytes < %d) {\n", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
			fmt.Fprintf(w, "      return false;\n")
			fmt.Fprintf(w, "    }\n")
			for i, f := range s.DynamicFieldHeadOffsets {
				_ = f
				fmt.Fprintf(w, "\n    U64 __off%d = ", i)
				if i == 0 {
					fmt.Fprintf(w, "U64(%d);", s.DynamicFieldHeadOffsets[len(s.DynamicFieldHeadOffsets)-1])
				} else {
					fmt.Fprintf(w, "U64.fromBytes(vData.sublist(%d, %d));", s.DynamicFieldHeadOffsets[i]-8, s.DynamicFieldHeadOffsets[i])
				}
			}
			fmt.Fprintf(w, "\n    U64 __off%d = U64(vData.lengthInBytes);", len(s.DynamicFieldHeadOffsets))
			var dynStructFields []*ir.Field
			for _, f := range s.DynamicFields {
				if f.TypeInfo.FieldType == ir.FieldType_STRUCT {
					dynStructFields = append(dynStructFields, f)
				}
			}
			if len(dynStructFields) == 0 {
				fmt.Fprintf(w, "\n    return ")
				for i, f := range s.DynamicFieldHeadOffsets {
					fmt.Fprintf(w, "__off%d <= __off%d ", i, i+1)
					if i != len(s.DynamicFieldHeadOffsets)-1 {
						fmt.Fprintf(w, "&& ")
					}
					_ = f
				}
				fmt.Fprintf(w, ";\n")
			} else {
				fmt.Fprintf(w, "\n    if (")
				for i, f := range s.DynamicFieldHeadOffsets {
					fmt.Fprintf(w, "__off%d <= __off%d ", i, i+1)
					if i != len(s.DynamicFieldHeadOffsets)-1 {
						fmt.Fprintf(w, "&& ")
					}
					_ = f
				}
				fmt.Fprintf(w, ") {\n")
				fmt.Fprintf(w, "      return ")
				for i, f := range dynStructFields {
					fmt.Fprintf(w, "%s.vStructValidate()", NameConv(f.Name))
					if i != len(dynStructFields)-1 {
						fmt.Fprintf(w, " && ")
					}
				}
				fmt.Fprintf(w, ";")
				fmt.Fprintf(w, "\n    }\n")
				fmt.Fprintf(w, "\n    return false;\n")
			}
		}
		fmt.Fprintf(w, "  }\n\n")
		fmt.Fprintf(w, "}\n\n")
	}
}
