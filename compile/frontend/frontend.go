package frontend

import (
	"fmt"

	"github.com/lemon-mint/vstruct/ast"
	"github.com/lemon-mint/vstruct/ir"
)

type _struct struct {
	s *ast.Struct

	isDynamic bool
	size      int

	filename string
	line     int
	col      int
}

type _enum struct {
	e *ast.Enum

	optionslen int
	size       int

	filename string
	line     int
	col      int
}

type _alias struct {
	a *ast.Alias

	filename string
	line     int
	col      int
}

type FrontEnd struct {
	file *ast.File

	structs []_struct
	enums   []_enum
	aliases []_alias

	output *ir.IR
}

func New(file *ast.File) *FrontEnd {
	f := &FrontEnd{
		file:   file,
		output: &ir.IR{},
	}
	f.output.FileName = file.Name
	return f
}

var ErrRawTypeOnRoot = fmt.Errorf("rawtype is not supported on root")

func (f *FrontEnd) Compile() error {
	for _, node := range f.file.Nodes {
		switch node.Type {
		case ast.NodeType_STRUCT:
			f.compileStruct(node)
		case ast.NodeType_ENUM:
			f.compileEnum(node)
		case ast.NodeType_ALIAS:
			f.compileAlias(node)
		case ast.NodeType_RAWTYPE:
			return ErrRawTypeOnRoot
		}
	}
	return nil
}

func (f *FrontEnd) compileStruct(node *ast.Node) {
	sizeInfo := f.getTypeSize(node)
	s := &_struct{
		s:         node.Struct,
		isDynamic: sizeInfo.isDynamic,
		size:      sizeInfo.size,
		filename:  node.File,
		line:      node.Token.Line,
		col:       node.Token.Col,
	}
	f.structs = append(f.structs, *s)
}

func (f *FrontEnd) compileEnum(node *ast.Node) {
	sizeInfo := f.getTypeSize(node)
	e := &_enum{
		e:          node.Enum,
		optionslen: len(node.Enum.Enums),
		size:       sizeInfo.size,
		filename:   node.File,
		line:       node.Token.Line,
		col:        node.Token.Col,
	}
	f.enums = append(f.enums, *e)
}

func (f *FrontEnd) compileAlias(node *ast.Node) {
	a := &_alias{
		a:        node.Alias,
		filename: node.File,
		line:     node.Token.Line,
		col:      node.Token.Col,
	}
	f.aliases = append(f.aliases, *a)
}
