package main

import (
	"fmt"
)

// linking two files with package
// see greeting file

// they are automatically picked up by go but have to specify that we want to run that file aswell
// go run main.go greetings.go

// this can be picked up by greetings file if the package is main
var score = 99.5
// if score is in the main function then it is out of scope so can't be accessed in other files

func main(){
	sayHello("Mario")

	for _, v := range points {
		fmt.Println(v)
	}

	// this function gets the score from above and prints it, the function is defined in another file
	showScore()
}