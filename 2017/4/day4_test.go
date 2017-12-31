package main

import "testing"

func TestValidPassphrases(t *testing.T) {
	tests := []struct {
		words []string
		valid bool
	}{
		{[]string{"aa", "bb", "cc", "dd", "ee"}, true},
		{[]string{"aa", "bb", "cc", "dd", "aa"}, false},
		{[]string{"aa", "bb", "cc", "dd", "aaa"}, true},
	}

	for _, tc := range tests {
		valid := validPassphrase(tc.words)
		if valid != tc.valid {
			t.Errorf("%q, expected %t, got %t", tc.words, tc.valid, valid)
		}
	}
}

func TestAnagrams(t *testing.T) {
	tests := []struct {
		words []string
		valid bool
	}{
		{[]string{"abcde", "fghij"}, true},
		{[]string{"abcde", "xyz", "ecdab"}, false},
		{[]string{"a", "ab", "abc", "abd", "abf", "abj"}, true},
		{[]string{"iiii", "oiii", "ooii"}, true},
		{[]string{"oiii", "ioii", "iioi", "iiio"}, false},
	}

	for _, tc := range tests {
		valid := validAnagrams(tc.words)
		if valid != tc.valid {
			t.Errorf("%q, expected %t, got %t", tc.words, tc.valid, valid)
		}
	}
}
