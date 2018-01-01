package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func jumpSteps(jumps []int) int {
	var cur, steps int
	for {
		// From current position, obtain jump value
		// Increment value of current position
		// Move from current position by jump value places
		// Exit if outside of list
		if cur > len(jumps)-1 || cur < 0 {
			break
		}
		jump := jumps[cur]
		jumps[cur]++
		cur = cur + jump
		steps++
	}
	return steps
}

func jumpSteps2(jumps []int) int {
	var cur, steps int
	for {
		if cur > len(jumps)-1 || cur < 0 {
			break
		}
		jump := jumps[cur]
		if jump >= 3 {
			jumps[cur]--
		} else {
			jumps[cur]++
		}
		cur = cur + jump
		steps++
	}
	return steps
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	var jumps []int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		jump, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		jumps = append(jumps, jump)
	}
	// fmt.Println("Steps:", jumpSteps(jumps))
	// Don't run both - the underlying array has been modified!
	fmt.Println("Steps 2:", jumpSteps2(jumps))
}
