package frontend

import (
	"fmt"

	"github.com/lemon-mint/vstruct/ast"
	"github.com/lemon-mint/vstruct/ir"
)

type _struct struct {
	Name string

	s *ast.Struct

	isDynamic bool
	size      int

	filename string
	line     int
	col      int
}

type _enum struct {
	Name string

	e *ast.Enum

	optionslen int
	size       int

	filename string
	line     int
	col      int
}

type _alias struct {
	Name string

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

func (f *FrontEnd) Output() *ir.IR {
	return f.output
}

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
	for _, s := range f.aliases {
		f.output.Aliases = append(f.output.Aliases, &ir.Alias{
			Name:         s.Name,
			OriginalType: s.a.Type.Name,
		})
	}
	for _, s := range f.structs {
		t := &ir.Struct{
			Name:    s.Name,
			Size:    s.size,
			IsFixed: true,
		}
		var offset int
		var dynOffset int
		var last *ir.Field
		for _, Field := range s.s.Fields {
			ft := &ir.Field{
				Name: Field.Name,
			}
			ft.Type = Field.StrType
			fInfo := f.getTypeSize(Field.Type)
			ft.TypeInfo = fInfo
			if fInfo.IsDynamic {
				t.DynamicFieldHeadOffsets = append(t.DynamicFieldHeadOffsets, dynOffset)
				t.DynamicFields = append(t.DynamicFields, ft)
				dynOffset += 8
				t.IsFixed = false
			} else {
				ft.Offset = offset
				t.FixedFields = append(t.FixedFields, ft)
				offset += fInfo.Size
			}
			last = ft
		}
		_ = last
		t.DynamicFieldHeadOffsets = append(t.DynamicFieldHeadOffsets, dynOffset)
		t.TotalFixedFieldSize = offset
		t.DynamicHead = offset
		for i := range t.DynamicFieldHeadOffsets {
			t.DynamicFieldHeadOffsets[i] += t.DynamicHead
		}
		f.output.Structs = append(f.output.Structs, t)
	}

	for _, e := range f.enums {
		t := &ir.Enum{
			Name: e.Name,
			Size: e.size,
		}
		t.Options = e.e.Enums
		f.output.Enums = append(f.output.Enums, t)
	}
	return nil
}

func (f *FrontEnd) compileStruct(node *ast.Node) {
	sizeInfo := f.getTypeSize(node)
	s := &_struct{
		Name:      node.Name,
		s:         node.Struct,
		isDynamic: sizeInfo.IsDynamic,
		size:      sizeInfo.Size,
		filename:  node.File,
		line:      node.Token.Line,
		col:       node.Token.Col,
	}
	f.structs = append(f.structs, *s)
}

func (f *FrontEnd) compileEnum(node *ast.Node) {
	sizeInfo := f.getTypeSize(node)
	e := &_enum{
		Name:       node.Name,
		e:          node.Enum,
		optionslen: len(node.Enum.Enums),
		size:       sizeInfo.Size,
		filename:   node.File,
		line:       node.Token.Line,
		col:        node.Token.Col,
	}
	f.enums = append(f.enums, *e)
}

func (f *FrontEnd) compileAlias(node *ast.Node) {
	a := &_alias{
		Name:     node.Name,
		a:        node.Alias,
		filename: node.File,
		line:     node.Token.Line,
		col:      node.Token.Col,
	}
	f.aliases = append(f.aliases, *a)
}
