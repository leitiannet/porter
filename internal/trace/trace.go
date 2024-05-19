package trace

import (
	"fmt"
)

func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}
