package main

import "fmt"

func loops() {
	// x :=0

	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	for i := 0; i < 5; i++ {
		fmt.Println("value of x is:", i)
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// like for in to get the index and value
	for index, value := range names {
		fmt.Printf("Position at index %v and the value is %v \n", index, value)
	}

	// if we don't want the index we can use _
	// value is a local copy of the variable
	for _, value := range names {
		fmt.Printf("The value is %v \n", value)

		// we cannot update the values like this
		value = "new string"
	}

	// this will print the original array, the previous loop will not reassign the value of the string
	fmt.Println(names)
}
