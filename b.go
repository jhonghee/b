package b

import (
	"fmt"

	"github.com/jhonghee/d"
)

// Version returns the tagged version of the module
func Version() string {
	return fmt.Sprint("B v1.2", "->", d.Version())
}
