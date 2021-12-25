package dart

import "strings"

func NameConv(in string) string {
	return strings.Title(in)
}

func TypeConv(in string) string {
	switch in {
	case "uint8":
		return "U8"
	case "uint16":
		return "U16"
	case "uint32":
		return "U32"
	case "uint64":
		return "U64"
	case "int8":
		return "I8"
	case "int16":
		return "I16"
	case "int32":
		return "I32"
	case "int64":
		return "I64"
	case "float32":
		return "F32"
	case "float64":
		return "F64"
	case "bool":
		return "bool"
	case "string":
		return "String"
	case "bytes":
		return "Uint8List"
	default:
		return NameConv(in)
	}
}
