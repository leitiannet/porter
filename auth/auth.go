package auth

import (
	"github.com/leitiannet/porter/internal/trace"
)

func Auth(a ...interface{}) error {
	trace.Printf("Auth: %v\n", a)
	return nil
}
