package frontend

import (
	"github.com/lemon-mint/vstruct/ast"
	"github.com/lemon-mint/vstruct/ir"
)

var rawTypeSizeInfo = map[string]ir.TypeInfo{
	"bool":    {Size: 1, IsDynamic: false, FieldType: ir.FieldType_BOOL},
	"int8":    {Size: 1, IsDynamic: false, FieldType: ir.FieldType_INT},
	"uint8":   {Size: 1, IsDynamic: false, FieldType: ir.FieldType_UINT},
	"int16":   {Size: 2, IsDynamic: false, FieldType: ir.FieldType_INT},
	"uint16":  {Size: 2, IsDynamic: false, FieldType: ir.FieldType_UINT},
	"int32":   {Size: 4, IsDynamic: false, FieldType: ir.FieldType_INT},
	"uint32":  {Size: 4, IsDynamic: false, FieldType: ir.FieldType_UINT},
	"int64":   {Size: 8, IsDynamic: false, FieldType: ir.FieldType_INT},
	"uint64":  {Size: 8, IsDynamic: false, FieldType: ir.FieldType_UINT},
	"float32": {Size: 4, IsDynamic: false, FieldType: ir.FieldType_FLOAT},
	"float64": {Size: 8, IsDynamic: false, FieldType: ir.FieldType_FLOAT},
	"string":  {Size: 0, IsDynamic: true, FieldType: ir.FieldType_STRING},
	"bytes":   {Size: 0, IsDynamic: true, FieldType: ir.FieldType_BYTES},
}

func (f *FrontEnd) getTypeSize(node *ast.Node) ir.TypeInfo {
	var typeInfo ir.TypeInfo
	switch node.Type {
	case ast.NodeType_STRUCT:
		typeInfo.FieldType = ir.FieldType_STRUCT
		for _, s := range node.Struct.Fields {
			fInfo := f.getTypeSize(s.Type)
			if fInfo.IsDynamic {
				typeInfo.IsDynamic = true
				typeInfo.Size = 0
				return typeInfo
			}
			typeInfo.Size += fInfo.Size
		}
	case ast.NodeType_ENUM:
		typeInfo.FieldType = ir.FieldType_ENUM
		enumLen := uint64(len(node.Enum.Enums))
		switch {
		case enumLen <= 1<<8:
			typeInfo.Size = 1
		case enumLen <= 1<<16:
			typeInfo.Size = 2
		case enumLen <= 1<<32:
			typeInfo.Size = 4
		default:
			typeInfo.Size = 8
		}
	case ast.NodeType_ALIAS:
		typeInfo = f.getTypeSize(node.Alias.Type)
	case ast.NodeType_RAWTYPE:
		tInfo, ok := rawTypeSizeInfo[node.RawType.Type]
		typeInfo.FieldType = tInfo.FieldType
		if !ok {
			typeInfo.IsDynamic = true
			typeInfo.Size = 0
			return typeInfo
		}
		typeInfo.Size = tInfo.Size
		typeInfo.IsDynamic = tInfo.IsDynamic
	}
	return typeInfo
}
