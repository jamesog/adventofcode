package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func diff(input []int) int {
	var s, b int
	for i, n := range input {
		if i == 0 {
			s = n
		}
		switch {
		case n < s:
			s = n
		case n > b:
			b = n
		}
	}
	return b - s
}

func divisible(nums []int) int {
	var div [2]int
OUTER:
	for i, x := range nums {
		for j, y := range nums {
			// Don't compare against itself
			if i == j {
				continue
			}
			if math.Mod(float64(x), float64(y)) == 0 {
				div[0] = x
				div[1] = y
				continue OUTER
			}
		}
	}

	return div[0] / div[1]
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	var cksum1, cksum2 int

	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		var nums []int
		line := scanner.Text()
		toks := strings.Split(line, "\t")
		for _, t := range toks {
			n, err := strconv.Atoi(t)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, n)
		}
		cksum1 += diff(nums)
		cksum2 += divisible(nums)
	}

	fmt.Println("Checksum 1:", cksum1)
	fmt.Println("Checksum 2:", cksum2)
}
