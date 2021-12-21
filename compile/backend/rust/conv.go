package rust

import (
	"strings"

	"github.com/lemon-mint/vstruct/ir"
)

func NameConv(in string) string {
	return strings.Title(in)
}

func TypeConv(in string) string {
	switch in {
	case "uint8":
		return "u8"
	case "uint16":
		return "u16"
	case "uint32":
		return "u32"
	case "uint64":
		return "u64"
	case "int8":
		return "i8"
	case "int16":
		return "i16"
	case "int32":
		return "i32"
	case "int64":
		return "i64"
	case "float32":
		return "f32"
	case "float64":
		return "f64"
	case "bool":
		return "bool"
	case "string":
		return "String"
	case "bytes":
		return "Vec<u8>"
	default:
		return NameConv(in)
	}
}

func NeedRef(in ir.FieldType) bool {
	switch in {
	case ir.FieldType_BOOL, ir.FieldType_UINT, ir.FieldType_INT, ir.FieldType_FLOAT, ir.FieldType_ENUM:
		return false
	default:
		return true
	}
}
