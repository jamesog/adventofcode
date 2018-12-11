package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isLower(b byte) bool {
	if 'a' <= b && b <= 'z' {
		return true
	}
	return false
}

func isUpper(b byte) bool {
	if 'A' <= b && b <= 'Z' {
		return true
	}
	return false
}

func toLower(b byte) byte {
	if isLower(b) {
		return b
	}
	return b + ('a' - 'A')
}

func toUpper(b byte) byte {
	if isUpper(b) {
		return b
	}
	return b - ('a' - 'A')
}

func reacts(a, b byte) bool {
	if toLower(a) != toLower(b) {
		return false
	}
	return (isLower(a) && isUpper(b)) || (isUpper(a) && isLower(b))
}

// Loop over the input string.
// Compare current letter with next.
// If they react, add remainder string (next+i:) to output and recurse
//	If recursed, reset  break the loop
//	N.B. for a huge string ths might cause a huse chain of function calls / memory usage
// If no reaction, add current letter to output
func polymer(input string) string {
	// Use a strings.Builder to gradually build the output polymer
	var b strings.Builder

	for i := range input {
		// If we reached the last byte it can't have reacted so just append to the string builder and end the loop.
		if i == len(input)-1 {
			b.WriteByte(input[i])
			break
		}
		// Test if there is a reaction between the current and next bytes in the string,
		// if so write the remainder of the string (skipping the current and next bytes)
		// to the output and recurse the new string.
		if reacts(input[i], input[i+1]) {
			b.WriteString(input[i+2:])
			output := polymer(b.String())
			b.Reset()
			b.WriteString(output)
			break
		}
		b.WriteByte(input[i])
	}
	return b.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Only one line of input in this challenge
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Part one: %d\n", len(polymer(input)))
}
