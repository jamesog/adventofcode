package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isValid(policy, pass string) bool {
	// x-y l
	pol := strings.SplitN(policy, " ", 2)
	times, l := strings.SplitN(pol[0], "-", 2), pol[1]

	min, _ := strconv.Atoi(times[0])
	max, _ := strconv.Atoi(times[1])

	cnt := strings.Count(pass, l)
	if cnt >= min && cnt <= max {
		return true
	}
	return false
}

func isValidOTCP(policy, pass string) bool {
	// x-y l
	pol := strings.SplitN(policy, " ", 2)
	times, l := strings.SplitN(pol[0], "-", 2), pol[1]

	pos1, _ := strconv.Atoi(times[0])
	pos2, _ := strconv.Atoi(times[1])

	// Accessing byte indexes isn't ideal, but the input appears to be entirely ASCII chars, so...
	if (string(pass[pos1-1]) == l || string(pass[pos2-1]) == l) &&
		(pass[pos1-1] != pass[pos2-1]) {
		return true
	}

	return false
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}

	var valid, validOTCP int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// x-y l: pass
		entry := strings.Split(scanner.Text(), ": ")
		if isValid(entry[0], entry[1]) {
			valid++
		}
		if isValidOTCP(entry[0], entry[1]) {
			validOTCP++
		}
	}

	fmt.Println("Part 1:", valid)
	fmt.Println("Part 1:", validOTCP)
}
