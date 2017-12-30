package main

import "testing"

func TestNext(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 1, 2, 2}, 3},
		{[]int{1, 1, 1, 1}, 4},
		{[]int{1, 2, 3, 4}, 0},
		{[]int{9, 1, 2, 1, 2, 1, 2, 9}, 9},
	}

	for _, tc := range tests {
		total := sumNext(tc.input)
		if total != tc.want {
			t.Errorf("wanted %d, got %d", tc.want, total)
		}
	}

}

func TestHalfway(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 1, 2}, 6},
		{[]int{1, 2, 2, 1}, 0},
		{[]int{1, 2, 3, 4, 2, 5}, 4},
		{[]int{1, 2, 3, 1, 2, 3}, 12},
		{[]int{1, 2, 1, 3, 1, 4, 1, 5}, 4},
	}

	for _, tc := range tests {
		total := sumHalfway(tc.input)
		if total != tc.want {
			t.Errorf("wanted %d, got %d", tc.want, total)
		}
	}
}
