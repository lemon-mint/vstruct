package frontend

import (
	"github.com/lemon-mint/vstruct/ast"
	"github.com/lemon-mint/vstruct/ir"
)

type TypeInfo struct {
	size      int
	isDynamic bool

	FieldType ir.FieldType
}

var rawTypeSizeInfo = map[string]TypeInfo{
	"bool":    {1, false, ir.FieldType_BOOL},
	"int8":    {1, false, ir.FieldType_INT},
	"uint8":   {1, false, ir.FieldType_UINT},
	"int16":   {2, false, ir.FieldType_INT},
	"uint16":  {2, false, ir.FieldType_UINT},
	"int32":   {4, false, ir.FieldType_INT},
	"uint32":  {4, false, ir.FieldType_UINT},
	"int64":   {8, false, ir.FieldType_INT},
	"uint64":  {8, false, ir.FieldType_UINT},
	"float32": {4, false, ir.FieldType_FLOAT},
	"float64": {8, false, ir.FieldType_FLOAT},
	"bytes":   {8, true, ir.FieldType_BYTES},
	"string":  {8, true, ir.FieldType_STRING},
}

func (f *FrontEnd) getTypeSize(node *ast.Node) TypeInfo {
	var typeInfo TypeInfo
	switch node.Type {
	case ast.NodeType_STRUCT:
		typeInfo.FieldType = ir.FieldType_STRUCT
		for _, s := range node.Struct.Fields {
			fInfo := f.getTypeSize(s.Type)
			if fInfo.isDynamic {
				typeInfo.isDynamic = true
				typeInfo.size = 0
				return typeInfo
			}
			typeInfo.size += fInfo.size
		}
	case ast.NodeType_ENUM:
		typeInfo.FieldType = ir.FieldType_ENUM
		enumLen := len(node.Enum.Enums)
		switch {
		case enumLen <= 1<<8:
			typeInfo.size = 1
		case enumLen <= 1<<16:
			typeInfo.size = 2
		case enumLen <= 1<<32:
			typeInfo.size = 4
		default:
			typeInfo.size = 8
		}
	case ast.NodeType_ALIAS:
		typeInfo = f.getTypeSize(node.Alias.Type)
	case ast.NodeType_RAWTYPE:
		tInfo, ok := rawTypeSizeInfo[node.RawType.Type]
		typeInfo.FieldType = tInfo.FieldType
		if !ok {
			typeInfo.isDynamic = true
			typeInfo.size = 0
			return typeInfo
		}
		typeInfo.size = tInfo.size
		typeInfo.isDynamic = tInfo.isDynamic
	}
	return typeInfo
}
