package problem4

import (
	"fmt"
	"crypto/md5"
)

type testCase struct {
	input  string
	result int
}

var testCasesPartOne = []testCase{
	{input: "abcdef", result: 609043},
	{input: "pqrstuv", result: 1048970},
	{input: "iwrupvqb", result: -1},
}

var testCasesPartTwo = []testCase{
	{input: "iwrupvqb", result: -1},
}

func RunPartOne() {
	fmt.Printf(">>> PART ONE <<<\n\n")
	for k, v := range testCasesPartOne {
		fmt.Printf("test case : %d\n", k)
		if len(v.input) > 250 {
			fmt.Printf("input     : [TOO LONG VALUE]\n")
		} else {
			fmt.Printf("input     : %s\n", v.input)
		}
		fmt.Printf("expected  : %d\n", v.result)
		fmt.Printf("result    : %d\n\n", v.solvePartOne())
	}
}

func RunPartTwo() {
	fmt.Printf(">>> PART TWO <<<\n\n")
	for k, v := range testCasesPartTwo {
		fmt.Printf("test case : %d\n", k)
		if len(v.input) > 250 {
			fmt.Printf("input     : [TOO LONG VALUE]\n")
		} else {
			fmt.Printf("input     : %s\n", v.input)
		}
		fmt.Printf("expected  : %d\n", v.result)
		fmt.Printf("result    : %d\n\n", v.solvePartTwo())
	}
}

func (tc *testCase) solvePartOne() int {
	i := 0
	for {
		x := fmt.Sprintf("%s%d", tc.input, i)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(x)))
		if hash[:5] == "00000" {
			break
		}
		i++
	}
	return i
}

func (tc *testCase) solvePartTwo() int {
	i := 0
	for {
		x := fmt.Sprintf("%s%d", tc.input, i)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(x)))
		if hash[:6] == "000000" {
			break
		}
		i++
	}
	return i
}
