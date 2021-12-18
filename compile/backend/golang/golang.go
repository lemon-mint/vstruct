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
	writeStructs(&codedataBuf, i)
	writeAliases(&codedataBuf, i)
	output := fmt.Sprintf(
		`package %s

import (
	"math"
)

var _ = math.Float32frombits
var _ = math.Float64frombits

%s
`,
		packageName,
		codedataBuf.String(),
	)
	fmted, err := format.Source([]byte(output))
	if err != nil {
		fmt.Println(output)
		return err
	}
	_, err = w.Write(fmted)
	return err
}
