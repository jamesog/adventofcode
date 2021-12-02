package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Part 1: Calculate horizontal position and depth after all instructions.
//         Multiply final horizontal position by final depth.

type position struct {
	horizontal int
	depth      int
}

func (p *position) forward(move int) { p.horizontal += move }
func (p *position) up(move int)      { p.depth -= move }
func (p *position) down(move int)    { p.depth += move }

func command(input string, pos *position) {
	parts := strings.Split(input, " ")
	var cmd string
	var move int
	fmt.Sscanf(input, "%s %d", &cmd, &move)
	switch cmd {
	case "forward":
		pos.forward(move)
	case "up":
		pos.up(move)
	case "down":
		pos.down(move)
	default:
		log.Fatalf("unknown command: %s", parts[0])
	}
}

func main() {
	pos := position{}

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		command(scanner.Text(), &pos)
	}

	fmt.Printf("Part 1: %d\n", pos.horizontal*pos.depth)
}
