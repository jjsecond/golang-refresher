package main

import "fmt"

func main() {
	age := 35
	name := "shaun"

	fmt.Print("hello, ")
	fmt.Print("my friend, ")
	fmt.Print("you! \n")

	fmt.Println("My age is",age, "and my name is", name)

	// formatted strings. Order matters and using a format specifier
	// %v means out put variable in default format, lots of different format specifiers
	fmt.Printf("My age is %v and my names is %v. \n", age, name)
	// puts quotes around them but should be used on strings. Will use a hash in place of the number. Would work with the age if it was a string
	fmt.Printf("My age is %q and my names is %q. \n", age, name)

	// get type of variable with format specifier  %T
	fmt.Printf("ages is of type %T \n", age)

	// using floats
	fmt.Printf("You scored %f point ! \n", 255.55)
	// can round them down
	fmt.Printf("You scored %0.1f point ! \n", 255.55)

	// Sprintf (save formatted strings)
	// returns formatted string and save it in a var
	var str = fmt.Sprintf("My age is %v and my names is %v. \n", age, name)
	fmt.Println("the saved string is:", str)


}