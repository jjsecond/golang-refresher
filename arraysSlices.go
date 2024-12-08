package main

import "fmt"


func arraysSlices(){
	// have to have the types on both sides and the 3 is the length you expect
	// arrays must be fixed lengths
	// var ages [3]int = [3]int {20, 25 ,30}

	// shorthand
	var ages = [3]int{20, 25 ,30}

	names := [4]string{"Yoshi", "Mario", "Peach", "Bowser"}

	//change value in array
	names[1] = "Luigi"

	// print out the ages and the length of the array with the len func
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices
	// More like arrays in other languages and not fixed length
	// use arrays under the hood but more flexible

	// this is now a slice
	var scores = []int{100, 50, 60}

	// update value at index
	scores[2] = 25
	// append to a slice
	// this does not mutate the original array but instead returns a new variable, so have to overwrite the variable
	scores = append(scores, 60)

	fmt.Println(scores, len(scores))

	// slice ranges

	rangeOne := names[1:3]

	// go to from position two and go to end of the slice
	rangeTwo:= names[2:]

	// go from beginning of array/slice until the 3rd position
	rangeThree:= names[:3]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	// append to a range
	rangeOne = append(rangeOne, "Koopa")

	fmt.Println(rangeOne)
}