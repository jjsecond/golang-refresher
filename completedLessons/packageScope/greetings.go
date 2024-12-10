// link it to main
package main

import "fmt"

// anything declared here will be accessible in the main function because the package is main
var points = []int{20, 90, 100, 45, 70}

func sayHello(name string) {
	fmt.Println("hello", name)
}

// referencing a var from the main.go and then this function is called in the main.go so files can share vars and be called in other places
func showScore() {
	fmt.Println("You scored this many points:", score)
}
