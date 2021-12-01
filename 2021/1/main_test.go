package main

import "testing"

var input = []int{
	199,
	200,
	208,
	210,
	200,
	207,
	240,
	269,
	260,
	263,
}

func TestCountIncrease(t *testing.T) {
	got := countIncrease(input)
	if got != 7 {
		t.Errorf("wanted 7; got %d", got)
	}
}

func TestSlidingWindow(t *testing.T) {
	windows := slidingWindow(input)
	got := countIncrease(windows)
	if got != 5 {
		t.Errorf("wanted 5; got %d", got)
	}
}
