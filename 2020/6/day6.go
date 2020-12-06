package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1(scanner bufio.Scanner) {
	total1 := 0
	seen := make(map[rune]struct{})
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			total1 += len(seen)
			seen = make(map[rune]struct{})
			continue
		}
		for _, l := range t {
			seen[l] = struct{}{}
		}
	}
	total1 += len(seen)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}

	scanner := bufio.NewScanner(f)

	total1 := 0
	total2 := 0
	groupsz := 0
	seen := make(map[rune]int)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			for l := range seen {
				if seen[l] == groupsz {
					total2++
				}
			}
			total1 += len(seen)
			seen = make(map[rune]int)
			groupsz = 0
			continue
		}
		groupsz++
		for _, l := range t {
			seen[l]++
		}
	}
	total1 += len(seen)
	for l := range seen {
		if seen[l] == groupsz {
			total2++
		}
	}

	fmt.Println("Part 1:", total1)

	fmt.Println("Part 2:", total2)
}
