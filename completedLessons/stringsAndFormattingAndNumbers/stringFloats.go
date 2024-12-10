package main

import "fmt"

func stringFloats() {
	// different ways to define
	var nameOne string = "mario"
	var nameTwo = "Luigi"
	// can se it up later
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	// shorthand, cannot use shorthand outside a function
	nameFour := "yoshi"
	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory

	// outside of range of int8: -128 - 127
	// var numOne int8 = 128
	var numOne int8 = 127
	var numTwo int8 = -128

	// can't have negative number with unit
	// var numThree uint8 = -1
	var numThree uint8 = 255

	fmt.Println(numOne, numTwo, numThree)

	// Floats

	var scoreOne float32 = 25.69
	var scoreTwo float64 = 8882843879238429738479.7
	// defaults to float64
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
