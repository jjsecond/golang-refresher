package main

import (
	"fmt"
	"strings"
)

// returning multiple values
func getInitials(name string) (string, string) {
	// 1. capitalise the string
	s := strings.ToUpper(name)
	// 2. Slice to split the name
	names := strings.Split(s, " ")

	// 3. create another slice for the initials
	var initials []string
	for _, v := range names {
		// get first letter [:1]
		initials = append(initials, v[:1])
	}

	// return the initials
	if len(initials) > 1 {
		// this is how you return two strings
		return initials[0], initials[1]
	}

	// if there is no second value return an underscore
	return initials[0], "_"

}

func returnMultiplesValuesFn() {
	// can't store two values from the return
	// i:= getInitials("tifa lockhart")
	fn1, sn1 := getInitials("tifa lockhart")
	fn2, sn2 := getInitials("Cloud Strife")

	fmt.Println(fn1, sn1)
	fmt.Println(fn2, sn2)
}
