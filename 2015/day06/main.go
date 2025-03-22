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

type prefix int

type command struct{

}

func parseCommand(line string){

}

func turnOn(grid [][]bool, fromX int, fromY int, toX int, toY int){
	for i := fromX; i < toX; i++{
		for j := fromY; j <= toY; j++{
			grid[i][j] = true
		}
	}
	return
}

func part1(input string) int {
	grid := make([][]bool, 1000)
	for i := range grid {
		grid[i] = make([]bool,1000)
	}

	for _,command := range strings.Split(input, "/n") {
		switch{
		case strings.HasPrefix(command, "turn on"):

			
		case strings.HasPrefix(command, "toggle"):

			
		case strings.HasPrefix(command, "turn off"):

			
		default:
			
		}
	}


	return 0
}

func part2(input string) int {
	return 0
}
