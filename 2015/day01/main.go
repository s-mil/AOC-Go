package main

import (
	"flag"
	"fmt"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int

	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	result := NotQuiteLisp(input, part)
	fmt.Println("Output: ", result)
}

func NotQuiteLisp(input string, part int) int {
	var level int = 0

	for index, item := range input {
		if item == '(' {
			level++
		} else {
			level--
		}
		if (part== 2 && level == -1){
			return index + 1
		}
	}

	return level
}
