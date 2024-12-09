package main

import (
	"fmt"
)

/*
Similar to JS


Go is a "Pass by value language"
- makes a copy of a value when passed to a function

Group A: non-pointer values : strings, ints, floats, booleans, arrays, structs

Group B point wrapper values: slices, maps, functions
*/


// x is copy of the name variable and go just overwrites the variable's copy
func updateName(x string) string{
	x ="wedge"

	return x
}

func updateMenu(y map[string]float64){
	y["coffee"]= 2.99
}

func passByValue(){
	name:= "tifa"

	// need to overwrite the name as if we pass it to the function it will not update the original value but the copy
	name = updateName(name)

	fmt.Println(name)

	// group B types
	menu :=map[string]float64 {
		"pie": 5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)

	fmt.Println(menu)
}