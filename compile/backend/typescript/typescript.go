package typescript

import (
	"fmt"
	"io"
	"strings"

	"github.com/lemon-mint/vstruct/ir"
)

func Generate(w io.Writer, i *ir.IR, packageName string) error {
	var codedataBuf strings.Builder
	writeEnums(&codedataBuf, i)
	writeStructs(&codedataBuf, i)
	writeAliases(&codedataBuf, i)
	output := fmt.Sprintf(`
// Code generated by vstruct. DO NOT EDIT.
// Package Name: %s

%s
`,
		packageName,
		codedataBuf.String(),
	)
	_, err := w.Write([]byte(output))
	return err
}
