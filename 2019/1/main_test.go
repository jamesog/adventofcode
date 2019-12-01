package main

import (
	"strconv"
	"testing"
)

func TestFuelForMass(t *testing.T) {
	tests := []struct {
		mass int
		want int
	}{
		{mass: 12, want: 2},
		{mass: 14, want: 2},
		{mass: 1969, want: 654},
		{mass: 100756, want: 33583},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.mass), func(t *testing.T) {
			got := fuel(tt.mass)
			if tt.want != got {
				t.Errorf("want %d; got %d", tt.want, got)
			}
		})
	}
}

func TestFuelForFuel(t *testing.T) {
	tests := []struct {
		mass int
		want int
	}{
		{mass: 14, want: 2},
		{mass: 1969, want: 966},
		{mass: 100756, want: 50346},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.mass), func(t *testing.T) {
			mf := fuel(tt.mass)
			got := mf + fuelForFuel(mf)
			if tt.want != got {
				t.Errorf("want %d; got %d", tt.want, got)
			}
		})
	}
}
