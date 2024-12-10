package main

import (
	"fmt"
)

// this will not update anything
// func updateName(x string){
// 	x ="wedge"
// }

// this will not update anything
// the astrix means we are are accepting a pointer to what is stored at that memory location and devalue it,
// we are not passing the name in we are passing in a pointer to where the value is
// we can now update it
// updates like a wrapper type
func updateNamePointer(x *string) {
	*x = "wedge"
}

func main() {
	name := "tifa"

	// this gets us a pointer to the memory location
	fmt.Println("memory address of name is: ", &name)
	// returns memory address of name is:  0x14000010040

	// this stores the pointer in its own memory block
	m := &name
	fmt.Println("the memory address is: ", m)
	// memory address of name is:  0x14000010040 which is the same

	// when we use astreix on a pointer gets the value at that memory address which is tifa, it is a dereference value
	fmt.Println("value at memory address: ", *m)

	fmt.Println(name)
	// this updates the data in the original location
	updateNamePointer(m)
	fmt.Println(name)
}
