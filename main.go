package main

import (
	"os"
	"github.com/resurtm/adventofcode-2015/problem1"
	"github.com/resurtm/adventofcode-2015/problem2"
	"github.com/resurtm/adventofcode-2015/problem3"
	"github.com/resurtm/adventofcode-2015/problem4"
	"github.com/resurtm/adventofcode-2015/problem5"
	"github.com/resurtm/adventofcode-2015/problem6"
	"github.com/resurtm/adventofcode-2015/problem7"
)

func main() {
	if len(os.Args) != 2 {
		panic("must have exactly one argument")
	}
	switch os.Args[1] {
	case "--problem1":
		problem1.RunPartOne()
		problem1.RunPartTwo()
	case "--problem2":
		problem2.RunPartOne()
		problem2.RunPartTwo()
	case "--problem3":
		problem3.RunPartOne()
		problem3.RunPartTwo()
	case "--problem4":
		problem4.RunPartOne()
		problem4.RunPartTwo()
	case "--problem5":
		problem5.RunPartOne()
		problem5.RunPartTwo()
	case "--problem6":
		problem6.RunPartOne()
		problem6.RunPartTwo()
	case "--problem7":
		problem7.RunPartOne()
		problem7.RunPartTwo()
	default:
		panic("nothing to do")
	}
}
