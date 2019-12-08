package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInputs(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{name: "2", input: []int{1, 0, 0, 0, 99}, output: []int{2, 0, 0, 0, 99}},
		{name: "6", input: []int{2, 3, 0, 3, 99}, output: []int{2, 3, 0, 6, 99}},
		{name: "9801", input: []int{2, 4, 4, 5, 99, 0}, output: []int{2, 4, 4, 5, 99, 9801}},
		{name: "x", input: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, output: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pgm := runProgram(tt.input)
			if diff := cmp.Diff(tt.output, pgm); diff != "" {
				t.Fatalf("output mismatch (-want +got)\n%s", diff)
			}
		})
	}
}
