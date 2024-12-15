package main

import "testing"

func TestCreateFrame(t *testing.T) {
	var tests = []struct {
		name     string
		input    []string
		expected string
	}{
		{"[midu, madeval, educavolpz]", []string{"midu", "madeval", "educavolpz"}, "**************\n* midu       *\n* madeval    *\n* educavolpz *\n**************\n"},
		{"[midu]", []string{"midu"}, "********\n* midu *\n********\n"},
		{"[a, bb, ccc]", []string{"a", "bb", "ccc"}, "*******\n* a   *\n* bb  *\n* ccc *\n*******\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateFrame(tt.input)
			if got != tt.expected {
				t.Errorf("CreateFrame(%s) = %s; want %s", tt.input, got, tt.expected)
			}
		})
	}
}
