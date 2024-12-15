package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOrganizeInventory(t *testing.T) {
	var tests = []struct {
		Name     string
		input    []Toy
		expected map[string]map[string]int
	}{
		{
			"{ toys: { doll: 5, car: 5 }, sports: { ball: 2, racket: 4 } }",
			[]Toy{
				{Name: "doll", Quantity: 5, Category: "toys"},
				{Name: "car", Quantity: 3, Category: "toys"},
				{Name: "ball", Quantity: 2, Category: "sports"},
				{Name: "car", Quantity: 2, Category: "toys"},
				{Name: "racket", Quantity: 4, Category: "sports"},
			},
			map[string]map[string]int{
				"toys": {
					"doll": 5,
					"car":  5,
				},
				"sports": {
					"ball":   2,
					"racket": 4,
				},
			}},
		{
			"{ education: { book: 15 }, art: { paint: 3 } }",
			[]Toy{
				{Name: "book", Quantity: 10, Category: "education"},
				{Name: "book", Quantity: 5, Category: "education"},
				{Name: "paint", Quantity: 3, Category: "art"},
			},
			map[string]map[string]int{
				"education": {
					"book": 15,
				},
				"art": {
					"paint": 3,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := OrganizeInventory(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				gotStringify, _ := json.Marshal(got)
				expectedStringify, _ := json.Marshal(tt.expected)
				t.Errorf("OrganizeInventory(%v) = %s; want %s", tt.input, gotStringify, expectedStringify)
			}
		})
	}
}
