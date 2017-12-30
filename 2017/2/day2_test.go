package main

import "testing"

func TestDiff(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{5, 1, 9, 5}, 8},
		{[]int{7, 5, 3}, 4},
		{[]int{2, 4, 6, 8}, 6},
	}

	for _, tc := range tests {
		d := diff(tc.input)
		if d != tc.want {
			t.Errorf("expected %d, got %d", tc.want, d)
		}
	}
}

func TestDivisible(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{5, 9, 2, 8}, 4},
		{[]int{9, 4, 7, 3}, 3},
		{[]int{3, 8, 6, 5}, 2},
	}

	for _, tc := range tests {
		d := divisible(tc.input)
		if d != tc.want {
			t.Errorf("expected %d, got %d", tc.want, d)
		}
	}
}
