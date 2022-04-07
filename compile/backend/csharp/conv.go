package csharp

import "strings"

func NameConv(in string) string {
	return strings.Title(in)
}

func NumberConv(unsigned bool, bit int) string {
	switch unsigned {
	case true:
		switch bit {
		case 8:
			return "byte"
		case 16:
			return "ushort"
		case 32:
			return "uint"
		case 64:
			return "ulong"
		}
	case false:
		switch bit {
		case 8:
			return "sbyte"
		case 16:
			return "short"
		case 32:
			return "int"
		case 64:
			return "long"
		}
	}
	return "UInt64"
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

func AliasConv(primitive string) string {
	switch primitive {
	case "bool":
		return "System.Boolean"
	case "byte":
		return "System.Byte"
	case "sbyte":
		return "System.SByte"
	case "short":
		return "System.Int16"
	case "ushort":
		return "System.UInt16"
	case "int":
		return "System.Int32"
	case "uint":
		return "System.UInt32"
	case "long":
		return "System.Int64"
	case "ulong":
		return "System.UInt64"
	case "float":
		return "System.Single"
	case "double":
		return "System.Double"
	case "string":
		return "System.String"
	case "bytes":
		return "System.Byte[]"
	}
	return primitive
}
