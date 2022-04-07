package csharp

import "strings"

func NameConv(in string) string {
	return strings.Title(in)
}

func TypeConv(in string) string {
	switch in {
	case "uint8":
		return "byte"
	case "uint16":
		return "ushort"
	case "uint32":
		return "uint"
	case "uint64":
		return "ulong"
	case "int8":
		return "sbyte"
	case "int16":
		return "short"
	case "int32":
		return "int"
	case "int64":
		return "long"
	case "float32":
		return "float"
	case "float64":
		return "double"
	case "bool":
		return "bool"
	case "string":
		return "string"
	case "bytes":
		return "byte[]"
	default:
		return NameConv(in)
	}
}
