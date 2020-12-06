package main

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

var input = []byte(`abc

a
b
c

ab
ac

a
a
a
a

b`)

func TestPart1(t *testing.T) {
	r := bytes.NewReader(input)
	scanner := bufio.NewScanner(r)
	total := 0
	group := 1
	seen := make(map[rune]struct{})
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			group++
			total += len(seen)
			seen = make(map[rune]struct{})
			continue
		}
		for _, l := range t {
			seen[l] = struct{}{}
		}
		fmt.Printf("[group %d] %s: %d, %v\n", group, t, len(seen), seen)
	}
	total += len(seen)
	if total != 11 {
		t.Errorf("want 11; got %d", total)
	}
}

func TestPart2(t *testing.T) {
	r := bytes.NewReader(input)
	scanner := bufio.NewScanner(r)
	total := 0
	group := 1
	groupsz := 0
	seen := make(map[rune]int)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			fmt.Printf("[group %d]: %d, %v\n", group, groupsz, seen)
			for l := range seen {
				if seen[l] == groupsz {
					total++
				}
			}
			group++
			groupsz = 0
			seen = make(map[rune]int)
			continue
		}
		groupsz++
		for _, l := range t {
			seen[l]++
		}
	}
	for l := range seen {
		if seen[l] == groupsz {
			total++
		}
	}
	if total != 6 {
		t.Errorf("want 6; got %d", total)
	}
}
