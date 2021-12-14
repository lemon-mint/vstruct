package frontend

import (
	"github.com/lemon-mint/vstruct/ast"
	"github.com/lemon-mint/vstruct/ir"
)

type FrontEnd struct {
	file *ast.File

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

func (f *FrontEnd) Compile() error {
	return nil
}
