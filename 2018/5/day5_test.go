package main

import "testing"

func TestPolymer(t *testing.T) {
	input := "dabAcCaCBAcCcaDA"
	want := "dabCBAcaDA"
	output := polymer(input)

	if want != output {
		t.Errorf("want %q; got %q", want, output)
	}
}

func TestShortestPolymer(t *testing.T) {
	input := "dabAcCaCBAcCcaDA"
	want := "daDA"
	output := shortestPolymer(input)

	if want != output {
		t.Errorf("want %q; got %q", want, output)
	}
}
