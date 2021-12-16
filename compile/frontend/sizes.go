package frontend

import "github.com/lemon-mint/vstruct/ast"

type sizeInfo struct {
	size      int
	isDynamic bool
}

var rawTypeSizeInfo = map[string]sizeInfo{
	"bool":    {1, false},
	"int8":    {1, false},
	"uint8":   {1, false},
	"int16":   {2, false},
	"uint16":  {2, false},
	"int32":   {4, false},
	"uint32":  {4, false},
	"int64":   {8, false},
	"uint64":  {8, false},
	"float32": {4, false},
	"float64": {8, false},
	"bytes":   {8, true},
	"string":  {8, true},
}

func (f *FrontEnd) getTypeSize(node *ast.Node) sizeInfo {
	var sizeInfo sizeInfo
	switch node.Type {
	case ast.NodeType_STRUCT:
		for _, s := range node.Struct.Fields {
			fInfo := f.getTypeSize(s.Type)
			if fInfo.isDynamic {
				sizeInfo.isDynamic = true
				sizeInfo.size = 0
				return sizeInfo
			}
			sizeInfo.size += fInfo.size
		}
	case ast.NodeType_ENUM:
		enumLen := len(node.Enum.Enums)
		switch {
		case enumLen <= 1<<8:
			sizeInfo.size = 1
		case enumLen <= 1<<16:
			sizeInfo.size = 2
		case enumLen <= 1<<32:
			sizeInfo.size = 4
		default:
			sizeInfo.size = 8
		}
	case ast.NodeType_ALIAS:
		sizeInfo = f.getTypeSize(node.Alias.Type)
	case ast.NodeType_RAWTYPE:
		tInfo, ok := rawTypeSizeInfo[node.RawType.Type]
		if !ok {
			sizeInfo.isDynamic = true
			sizeInfo.size = 0
			return sizeInfo
		}
		sizeInfo.size = tInfo.size
		sizeInfo.isDynamic = tInfo.isDynamic
	}
	return sizeInfo
}
