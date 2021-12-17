package golang

import (
	"fmt"
	"go/format"
	"io"
	"strings"

	"github.com/lemon-mint/vstruct/ir"
)

func Generate(w io.Writer, i *ir.IR, packageName string) error {
	var codedataBuf strings.Builder
	writeEnums(&codedataBuf, i)
	output := fmt.Sprintf(
		`package %s

%s
`,
		packageName,
		codedataBuf.String(),
	)
	fmted, err := format.Source([]byte(output))
	if err != nil {
		return err
	}
	_, err = w.Write(fmted)
	return err
}
