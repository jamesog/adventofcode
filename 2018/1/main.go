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

func main() {
	freq := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		freq = changeFrequency(freq, scanner.Text())
	}
	fmt.Println(freq)
}
