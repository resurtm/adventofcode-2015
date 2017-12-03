package main

import (
	"os"
	"github.com/resurtm/adventofcode-2015/problem1"
)

func main() {
	if len(os.Args) != 2 {
		panic("must have exactly one argument")
	}
	switch os.Args[1] {
	case "--problem1":
		problem1.RunPartOne()
		problem1.RunPartTwo()
	default:
		panic("nothing to do")
	}
}
