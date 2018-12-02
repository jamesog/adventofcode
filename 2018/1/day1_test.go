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
		freq := partOne(tt.Input)
		if freq != tt.WantOutput {
			t.Errorf("input %v: want %d; got %d", tt.Input, tt.WantOutput, freq)
		}
	}
}

func TestChangeFrequencyDuplicate(t *testing.T) {
	tests := []struct {
		Input      []string
		WantOutput int
	}{
		{
			Input:      []string{"+1", "-1"},
			WantOutput: 0,
		},
		{
			Input:      []string{"+3", "+3", "+4", "-2", "-4"},
			WantOutput: 10,
		},
		{
			Input:      []string{"-6", "+3", "+8", "+5", "-6"},
			WantOutput: 5,
		},
		{
			Input:      []string{"+7", "+7", "-2", "-7", "-4"},
			WantOutput: 14,
		},
	}

	for _, tt := range tests {
		freq := partTwo(tt.Input)
		if freq != tt.WantOutput {
			t.Errorf("input %v: want %d; got %d", tt.Input, tt.WantOutput, freq)
		}
	}
}
