package main

import (
	"github.com/leitiannet/porter/hash"
	"github.com/leitiannet/porter/internal/trace"
)

func main() {
	trace.Printf("this is hasher using hash\n")
	hash.Hash("a", "b", "c")
}
