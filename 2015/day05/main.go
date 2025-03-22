package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"sync"
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

func isStringNice(s string, part int) bool {

	switch {
	case part == 1:
		// Rule 1: Must contain at least 3 vowels
		vowelCount := 0
		for _, c := range s {
			if strings.ContainsRune("aeiou", c) {
				vowelCount++
			}
		}
		if vowelCount < 3 {
			return false
		}

		// Rule 2: Must contain at least one letter that appears twice in a row
		hasDouble := false
		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1] {
				hasDouble = true
				break
			}
		}
		if !hasDouble {
			return false
		}

		// Rule 3: Must not contain the strings "ab", "cd", "pq", or "xy"
		forbidden := []string{"ab", "cd", "pq", "xy"}
		for _, f := range forbidden {
			if strings.Contains(s, f) {
				return false
			}
		}

	case part == 2:
		hasPair := false
		for i := 0; i < len(s)-1; i++ {
			pair := s[i : i+2]
			if hasPair {
				break
			}
			for j := i + 2; j < len(s)-1; j++ {
				comPair := s[j : j+2]
				if comPair == pair {
					hasPair = true
					break
				}
			}
		}
		if !hasPair {
			return false
		}

		// rule 2: Contains repeated letter with one gap
		hasDouble := false
		for i := 0; i < len(s)-2; i++ {
			if s[i] == s[i+2] {
				hasDouble = true
				break
			}
		}

		return hasDouble
	}
	return true
}

func isNice(input string, result chan int, part int) {
	items := strings.Split(input, "\n")
	var wg sync.WaitGroup
	var mu sync.Mutex
	var count int
	wg.Add(len(items))

	for _, item := range items {
		go func(s string) {
			defer wg.Done()

			if isStringNice(s, part) {
				mu.Lock()
				count++
				mu.Unlock()
			}

		}(item)

	}

	wg.Wait()

	result <- count
	close(result)
}

func part1(input string) int {

	ch := make(chan int)

	go isNice(input, ch, 1)

	result := <-ch
	return result
}

func part2(input string) int {
	ch := make(chan int)

	go isNice(input, ch, 2)

	result := <-ch
	return result
}
