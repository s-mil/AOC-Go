// influenced by https://github.com/alexchao26/advent-of-code-go/blob/main/scripts/skeleton/tmpls/main.go
package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
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

	if part == 1 {
		ans := part1(input)

		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func getFeet(input string) int {
	dimensions := strings.Split(input, "x")

	l, _ := strconv.Atoi(dimensions[0])
	w, _ := strconv.Atoi(dimensions[1])
	h, _ := strconv.Atoi(dimensions[2])

	var lw = l * w
	var wh = w * h
	var lh = l * h

	return 2*(lw+wh+lh) + min(lw, wh, lh)
}
func part1(input string) int {
	var presents = strings.Split(input, "\n")
	var sqrft int = 0
	for _, present := range presents {
		sqrft += getFeet(present)
	}

	return sqrft
}

func getRibbon(input string) int {
	dimensions := strings.Split(input, "x")

	l, _ := strconv.Atoi(dimensions[0])
	w, _ := strconv.Atoi(dimensions[1])
	h, _ := strconv.Atoi(dimensions[2])

	lengths := []int{l, w, h}
	slices.Sort(lengths)

	return (2 * lengths[0]) + (2 * lengths[1]) + (l * w * h)

}

func part2(input string) int {
	var presents = strings.Split(input, "\n")
	var ribbonFeet int = 0
	for _, present := range presents {
		ribbonFeet += getRibbon(present)
	}

	return ribbonFeet
}
