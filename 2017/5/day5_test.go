package main

import "testing"

func TestJumps(t *testing.T) {
	test := struct {
		jumps []int
		steps int
	}{
		[]int{0, 3, 0, 1, -3}, 5,
	}

	steps := jumpSteps(test.jumps)
	if steps != test.steps {
		t.Errorf("expected %d, got %d", test.steps, steps)
	}
}

func TestJumps2(t *testing.T) {
	test := struct {
		jumps []int
		steps int
	}{
		[]int{0, 3, 0, 1, -3}, 10,
	}

	steps := jumpSteps2(test.jumps)
	if steps != test.steps {
		t.Errorf("expected %d, got %d", test.steps, steps)
	}
}
