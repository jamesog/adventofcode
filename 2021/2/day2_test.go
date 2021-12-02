package main

import "testing"

func TestPosition(t *testing.T) {
	pos := position{}
	inputs := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	for _, input := range inputs {
		command(input, &pos)
	}
	if pos.horizontal != 15 {
		t.Errorf("want horizontal position 15; got %d", pos.horizontal)
	}
	if pos.depth != 10 {
		t.Errorf("want depth 10; got %d", pos.horizontal)
	}
}
