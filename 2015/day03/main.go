package main

import (
	_ "embed"
	"flag"
	"fmt"
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

type Coordinate struct {
	x int
	y int
}

func visitHouse(houseMap map[Coordinate]int, position Coordinate) {
	houseMap[position] = houseMap[position] + 1

}

func part1(input string) int {

	var houseMap = make(map[Coordinate]int)
	var currentPos Coordinate = Coordinate{0, 0}
	houseMap[currentPos] = 1

	for _, direction := range input {
		switch {
		case direction == '<':
			currentPos = Coordinate{currentPos.x - 1, currentPos.y}
		case direction == '^':
			currentPos = Coordinate{currentPos.x, currentPos.y + 1}
		case direction == '>':
			currentPos = Coordinate{currentPos.x + 1, currentPos.y}
		case direction == 'v':
			currentPos = Coordinate{currentPos.x, currentPos.y - 1}
		}
		visitHouse(houseMap,
			currentPos)
	}

	return len(houseMap)
}

func part2(input string) int {
	var houseMap = make(map[Coordinate]int)

	var currentPos Coordinate = Coordinate{0, 0}
	var roboPos Coordinate = Coordinate{0,0}
	
	houseMap[currentPos] = 2

	for index, direction := range input {
		if index % 2 ==0 { switch {
		case direction == '<':
			currentPos = Coordinate{currentPos.x - 1, currentPos.y}
		case direction == '^':
			currentPos = Coordinate{currentPos.x, currentPos.y + 1}
		case direction == '>':
			currentPos = Coordinate{currentPos.x + 1, currentPos.y}
		case direction == 'v':
			currentPos = Coordinate{currentPos.x, currentPos.y - 1}
		}
		visitHouse(houseMap,
			currentPos)
		}else{
			switch {
			case direction == '<':
				roboPos = Coordinate{roboPos.x - 1, roboPos.y}
			case direction == '^':
				roboPos = Coordinate{roboPos.x, roboPos.y + 1}
			case direction == '>':
				roboPos = Coordinate{roboPos.x + 1, roboPos.y}
			case direction == 'v':
				roboPos = Coordinate{roboPos.x, roboPos.y - 1}
			}
			visitHouse(houseMap,
				roboPos)
		}
	}

	return len(houseMap)
}
