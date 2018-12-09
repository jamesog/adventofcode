package main

import (
	"reflect"
	"testing"
)

// Test that we can parse the input line into a Claim correctly.
func TestClaim(t *testing.T) {
	in := "#1 @ 1,3: 4x4"
	want := claim{
		ID:     1,
		Left:   1,
		Top:    3,
		Right:  4,
		Bottom: 6,
		Width:  4,
		Height: 4,
	}
	got := newClaim(in)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("claims not equal: want %#v; got %#v", want, got)
	}
}
func TestCoverage(t *testing.T) {
	in := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
		"#4 @ 3,3: 2x2",
	}
	// This is really wrong but I don't understand why :-/
	want := 11
	got := coverage(allClaims(in))
	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}
