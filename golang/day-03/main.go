package main

import (
	"encoding/json"
)

type Toy struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Category string `json:"category"`
}

func (t Toy) String() string {
	toy, err := json.Marshal(t)
	if err != nil {
		panic("Can not marshal toy")
	}
	return string(toy)
}

func OrganizeInventory(inventory []Toy) map[string]map[string]int {
	organizedInventory := make(map[string]map[string]int)

	for _, toy := range inventory {
		_, exits := organizedInventory[toy.Category]
		if !exits {
			organizedInventory[toy.Category] = map[string]int{
				toy.Name: toy.Quantity,
			}
		} else {
			_, exists := organizedInventory[toy.Category][toy.Name]
			if !exists {
				organizedInventory[toy.Category][toy.Name] = toy.Quantity
			} else {
				organizedInventory[toy.Category][toy.Name] += toy.Quantity
			}
		}
	}

	return organizedInventory
}
