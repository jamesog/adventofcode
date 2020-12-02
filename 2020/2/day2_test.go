package main

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		pol, pass        string
		valid, validOTCP bool
	}{
		{"1-3 a", "abcde", true, true},
		{"1-3 b", "cdefg", false, false},
		{"2-9 c", "ccccccccc", true, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s: %s", tt.pol, tt.pass), func(t *testing.T) {
			valid := isValid(tt.pol, tt.pass)
			if valid != tt.valid {
				t.Errorf("want %t; got %t", tt.valid, valid)
			}
			validOTCP := isValidOTCP(tt.pol, tt.pass)
			if validOTCP != tt.validOTCP {
				t.Errorf("want %t; got %t", tt.validOTCP, validOTCP)
			}
		})
	}
}
