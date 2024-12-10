package main

import "fmt"

func booleansAndConditionals() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("Falling into else block")
	}

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			// what this says continue with the loop but break out of current iteration
			// so it will skip the print line when index is 1
			continue
		}

		if index > 2 {
			fmt.Println("breaking at position", index)
			// completely breaks the loop
			break
		}

		fmt.Printf("The value at pos %v is %v \n", index, value)
	}
}
