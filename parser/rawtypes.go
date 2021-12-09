package parser

import (
	"github.com/lemon-mint/vstruct/ast"
	"github.com/lemon-mint/vstruct/lexer"
)

var rawtypes []string = []string{
	"bool", "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "string", "bytes",
}

func isRawType(name string) bool {
	for _, rawtype := range rawtypes {
		if rawtype == name {
			return true
		}
	}
	return false
}

func getRawTypeNode(name string) *ast.Node {
	return &ast.Node{
		Type: ast.NodeType_RAWTYPE,
		Name: name,
		File: "builtin",
		Line: 0,
		Col:  0,
		Token: lexer.Token{
			Type:    lexer.TOKEN_RAWTYPE,
			Literal: name,
			File:    "builtin",
			Line:    0,
			Col:     0,
		},
		RawType: &ast.RawType{
			Type: name,
		},
	}
}
