package csharp

import (
	"fmt"
	"io"
	"strings"

	"github.com/lemon-mint/vstruct/ir"
)

func Generate(w io.Writer, i *ir.IR, packageName string) error {
	var codedataBuf strings.Builder
	writeAliases(&codedataBuf, i)
	writeEnums(&codedataBuf, i)
	writeStructs(&codedataBuf, i)
	output := fmt.Sprintf(
		`namespace vstruct.%s
{
%s
}
`,
		packageName,
		codedataBuf.String(),
	)
	_, err := w.Write([]byte(output))
	return err
}
