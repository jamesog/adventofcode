package main

import "testing"

func TestChangeFrequency(t *testing.T) {
	tests := []struct {
		Input      []string
		WantOutput int
	}{
		{
			Input:      []string{"+1", "+1", "+1"},
			WantOutput: 3,
		},
		{
			Input:      []string{"+1", "+1", "-2"},
			WantOutput: 0,
		},
		{
			Input:      []string{"-1", "-2", "-3"},
			WantOutput: -6,
		},
	}

	for _, tt := range tests {
		freq := 0
		for _, f := range tt.Input {
			freq = changeFrequency(freq, f)
		}
		if freq != tt.WantOutput {
			t.Errorf("input %v: want %d; got %d", tt.Input, tt.WantOutput, freq)
		}
	}
}
