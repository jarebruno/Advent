package main

import "testing"

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestPrepareGifts(t *testing.T) {
	var tests = []struct {
		name     string
		input    []int
		expected []int
	}{
		{"[]", []int{}, []int{}},
		{"[3, 1, 2, 3, 4, 2, 5]", []int{3, 1, 2, 3, 4, 2, 5}, []int{1, 2, 3, 4, 5}},
		{"[6, 5, 5, 5, 5]", []int{6, 5, 5, 5, 5}, []int{5, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrepareGifts(tt.input)
			if !equal(got, tt.expected) {
				t.Errorf("PrepareGifts(%d) = %d; want %d", tt.input, got, tt.expected)
			}
		})
	}
}
