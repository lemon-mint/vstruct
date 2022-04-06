package typescript

import "strings"

func NameConv(in string) string {
	return strings.Title(in)
}

func TypeConv(in string) string {
	switch in {
	case "uint8":
		return "bigint"
	case "uint16":
		return "bigint"
	case "uint32":
		return "bigint"
	case "uint64":
		return "bigint"
	case "int8":
		return "bigint"
	case "int16":
		return "bigint"
	case "int32":
		return "bigint"
	case "int64":
		return "bigint"
	case "float32":
		return "number"
	case "float64":
		return "number"
	case "bool":
		return "bool"
	case "string":
		return "string"
	case "bytes":
		return "Uint8Array"
	default:
		return NameConv(in)
	}
}
