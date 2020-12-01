package main

import "testing"

func TestSum2(t *testing.T) {
	input := []int{1721,
		979,
		366,
		299,
		675,
		1456,
	}
	want1 := 1721
	want2 := 299

	num1, num2 := sum2(input)
	if num1 != want1 || num2 != want2 {
		t.Errorf("want (%d, %d); got (%d, %d)", want1, want2, num1, num2)
	}
}
func TestSum3(t *testing.T) {
	input := []int{1721,
		979,
		366,
		299,
		675,
		1456,
	}
	want1 := 979
	want2 := 366
	want3 := 675

	num1, num2, num3 := sum3(input)
	if num1 != want1 || num2 != want2 || num3 != want3 {
		t.Errorf("want (%d, %d, %d); got (%d, %d, %d)", want1, want2, want3, num1, num2, num3)
	}
}
