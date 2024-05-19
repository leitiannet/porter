package main

import (
	"github.com/leitiannet/porter/auth"
	"github.com/leitiannet/porter/internal/trace"
)

func main() {
	trace.Printf("this is auther using auth\n")
	auth.Auth("a", "b", "c")
}
