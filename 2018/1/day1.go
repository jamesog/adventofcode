package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func changeFrequency(curFreq int, shift string) int {
	// Convert the string to an int
	// Atoi can deal with turning +1 into 1
	shiftInt, err := strconv.Atoi(shift)
	if err != nil {
		return 0
	}
	return curFreq + shiftInt

}

func partOne(input []string) int {
	freq := 0
	for _, shift := range input {
		freq = changeFrequency(freq, shift)
	}
	return freq
}

func partTwo(input []string) int {
	seenFreqs := make(map[int]int)
	freq := 0
	// Fixes the bug for the +1 -1 test - we were told it starts at 0 therefore it's already been "seen"
	seenFreqs[0] = 1
	for {
		for _, shift := range input {
			freq = changeFrequency(freq, shift)
			seenFreqs[freq]++
			if seenFreqs[freq] == 2 {
				return freq
			}
		}
	}
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
