package main

import (
	"fmt"
)

func maps() {
	// maps made of key value pairs, all keys and values in a single map must have some type

	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// loop through a map

	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	// ints as keytype

	phoneBook := map[int]string {
		239834398: "mario",
		939879778: "luigi",
		868726475: "peach",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[239834398])


	//update item in map

	phoneBook[239834398] = "bowser"

	fmt.Println(phoneBook)

	
}