package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func countIncrease(input []int) int {
	if len(input) == 0 {
		return 0
	}

	var increase int
	prev := input[0]

	for i, depth := range input {
		if i == 0 {
			continue
		}
		if depth > prev {
			increase++
		}
		prev = depth
	}

	return increase
}

func slidingWindow(input []int) []int {
	sums := make([]int, 0, len(input))
	for i := range input {
		if i < 2 {
			continue
		}
		sum := input[i-2] + input[i-1] + input[i]
		sums = append(sums, sum)
	}
	return sums
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}
	var depths []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, num)
	}

	fmt.Printf("Part 1: %d\n", countIncrease(depths))
	fmt.Printf("Part 2: %d\n", countIncrease(slidingWindow(depths)))
}
