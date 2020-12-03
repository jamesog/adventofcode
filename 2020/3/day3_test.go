package main

import (
	"bufio"
	"bytes"
	"testing"
)

var input = []byte(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`)

func TestTravel(t *testing.T) {
	var grid []string
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	tests := map[string]struct {
		right, down, want int
	}{
		"Right 3 down 1": {3, 1, 7},
		"Right 1 down 1": {1, 1, 2},
		"Right 5 down 1": {5, 1, 3},
		"Right 7 down 1": {7, 1, 4},
		"Right 1 down 2": {1, 2, 2},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			trees := travel(grid, tt.right, tt.down)
			if trees != tt.want {
				t.Errorf("want %d; got %d", tt.want, trees)
			}
		})
	}
}
