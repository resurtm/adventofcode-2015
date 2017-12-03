package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("must have exactly one argument")
	}
	switch os.Args[1] {
	default:
		panic("nothing to do")
	}
}
