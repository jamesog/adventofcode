package main

import "testing"

func TestChecksum(t *testing.T) {
	tests := []struct {
		Input     []string
		C2        int
		C3        int
		WantCksum int
	}{
		{
			Input: []string{
				"abcdef",
				"bababc",
				"abbcde",
				"abcccd",
				"aabcdd",
				"abcdee",
				"ababab",
			},
			C2:        4,
			C3:        3,
			WantCksum: 4 * 3,
		},
	}

	for _, tt := range tests {
		res := checksum(tt.Input)
		if res != tt.WantCksum {
			t.Errorf("want checksum %d; got %d", tt.WantCksum, res)
		}
	}
}

func TestCommon(t *testing.T) {
	test := struct {
		Input []string
		Want  string
	}{
		Input: []string{
			"abcde",
			"fghij",
			"klmno",
			"pqrst",
			"fguij",
			"axcye",
			"wvxyz",
		},
		Want: "fgij",
	}

	res := common(test.Input)
	if res != test.Want {
		t.Errorf("want %q; got %q", test.Want, res)
	}
}
