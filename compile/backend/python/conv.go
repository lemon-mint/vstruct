package python

func NameConv(in string) string {
	return in
}

func TypeConv(in string) string {
	switch in {
	case "uint8":
		return "int"
	case "uint16":
		return "int"
	case "uint32":
		return "int"
	case "uint64":
		return "int"
	case "int8":
		return "int"
	case "int16":
		return "int"
	case "int32":
		return "int"
	case "int64":
		return "int"
	case "float32":
		return "float"
	case "float64":
		return "float"
	case "bool":
		return "bool"
	case "string":
		return "str"
	case "bytes":
		return "bytearray"
	default:
		return NameConv(in)
	}
}
