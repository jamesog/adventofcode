package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func checksum(input []string) int {
	var c2, c3 int
	for _, id := range input {
		chars := make(map[string]int, len(id))
		r := strings.NewReader(id)
		for {
			ch, _, err := r.ReadRune()
			if err == io.EOF {
				break
			}
			chars[string(ch)]++
		}
		var has2, has3 bool
		for _, count := range chars {
			switch count {
			case 2:
				has2 = true
			case 3:
				has3 = true
			}
		}
		if has2 {
			c2++
		}
		if has3 {
			c3++
		}
	}
	return c2 * c3
}

// Compare each value with all other values and return a difference in character counts
// The slow way is to compare i=N with i>N
// Maybe faster to sort first?

// Look at two strings, do a left/right comparison of characters
// If a[i] == b[i]: 👍 - store character in a string
// Else count failure
// Failure > 1 = 👎
// Otherwise return stored characters

func compare(left, right string) (string, bool) {
	var b strings.Builder
	var fail int
	for i := 0; i < len(left); i++ {
		if left[i] == right[i] {
			// Characters match, store them
			fmt.Fprint(&b, string(left[i]))
		} else {
			fail++
		}
	}

	if fail > 1 {
		return "", false
	}
	return b.String(), true
}

func common(input []string) string {
	for i, a := range input {
		for j, b := range input[i+1:] {
			if j == len(input)-1 {
				break
			}
			// fmt.Printf("a=%s b=%s\n", a, b)
			com, ok := compare(a, b)
			if !ok {
				// fmt.Println("fail")
				continue
			}
			return com
		}
	}
	return ""
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println("Part one:", checksum(input))
	fmt.Println("Part two:", common(input))
}
