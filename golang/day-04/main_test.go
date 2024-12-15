package main

import (
	"testing"
)

type Input struct {
	height   int
	ornament string
}

func TestOrganizeInventory(t *testing.T) {
	var tests = []struct {
		name     string
		input    Input
		expected string
	}{
		{
			"height: 5 ornament: *",
			Input{
				height:   5,
				ornament: "*",
			},
			"/*\n____*____\n___***___\n__*****__\n_*******_\n*********\n____#____\n____#____\n*/",
		},
		{
			"height: 5 ornament: *",
			Input{
				height:   3,
				ornament: "+",
			},
			"/*\n__+__\n_+++_\n+++++\n__#__\n__#__\n*/",
		},
		{
			"height: 5 ornament: *",
			Input{
				height:   6,
				ornament: "@",
			},
			"/*\n_____@_____\n____@@@____\n___@@@@@___\n__@@@@@@@__\n_@@@@@@@@@_\n@@@@@@@@@@@\n_____#_____\n_____#_____\n*/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateXmasTree(tt.input.height, tt.input.ornament)
			if got != tt.expected {
				t.Errorf("CreateXmasTree(%d, %s) = %s; want %s", tt.input.height, tt.input.ornament, got, tt.expected)
			}
		})
	}
}
