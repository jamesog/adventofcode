package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	opAdd      = 1
	opMultiply = 2
	opHalt     = 99
)

func loadProgram(program string) []int {
	ops := strings.Split(program, ",")
	ints := make([]int, len(ops))
	for i := range ops {
		op, err := strconv.Atoi(ops[i])
		if err != nil {
			log.Fatal(err)
		}
		ints[i] = op
	}
	return ints
}

func runProgram(program []int) []int {
	var pos int
RUN:
	for {
		op := program[pos]
		switch op {
		case opAdd:
			in1 := program[program[pos+1]]
			in2 := program[program[pos+2]]
			store := program[pos+3]
			// fmt.Printf("%d %d + %d = %d (st: %d)\n", pos, in1, in2, in1+in2, store)
			program[store] = in1 + in2
		case opMultiply:
			in1 := program[program[pos+1]]
			in2 := program[program[pos+2]]
			store := program[pos+3]
			// fmt.Printf("%d %d x %d = %d (st: %d)\n", pos, in1, in2, in1*in2, store)
			program[store] = in1 * in2
		case opHalt:
			// fmt.Printf("Halt: 0 = %d\n", program[0])
			break RUN
		}
		pos = pos + 4
	}
	return program
}

func brute(pgm []int) (noun, verb int) {
	for noun := 0; noun <= 99; noun++ {
		try := make([]int, len(pgm))
		for verb := 0; verb <= 99; verb++ {
			copy(try, pgm)
			try[1] = noun
			try[2] = verb
			out := runProgram(try)
			if out[0] == 19690720 {
				return noun, verb
			}
		}
	}
	return noun, verb
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	ints := loadProgram(scanner.Text())
	part1in := make([]int, len(ints))
	copy(part1in, ints)
	part1in[1] = 12
	part1in[2] = 2
	part1out := runProgram(part1in)
	fmt.Printf("Part 1 pos 0: %d\n", part1out[0])
	noun, verb := brute(ints)
	fmt.Printf("Part 2 noun: %d verb %d, answer = %d\n", noun, verb, 100*noun+verb)
}
