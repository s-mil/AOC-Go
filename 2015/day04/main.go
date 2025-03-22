package main

import (
	"context"
	"crypto/md5"
	_ "embed"
	"flag"
	"fmt"
	"math"
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

func hashFinder(ctx context.Context, input string, leadingZero int, result chan int) {

	for i := 0; i < math.MaxInt32; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			key := fmt.Sprintf("%s%d", input, i)
			hash := fmt.Sprintf("%x", md5.Sum([]byte(key)))
			if strings.HasPrefix(hash, strings.Repeat("0", leadingZero)) {
				result <- i
				return
			}
		}
	}

}

func part1(input string) int {
	leadingZero := 5
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go hashFinder(ctx,input, leadingZero, ch)

	result := <-ch
	cancel()
	return result

}

func part2(input string) int {
	leadingZero := 6
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go hashFinder(ctx,input, leadingZero, ch)

	result := <-ch
	cancel()
	return result
}
