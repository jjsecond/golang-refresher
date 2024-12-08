package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	greeting:= "Hello there friends!"

	fmt.Println(strings.Contains(greeting, "Hello"))

	// strings method do not mutate original string
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))

	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))

	// get slice of three elements
	fmt.Println(strings.Split(greeting, " "))

	ages:=[]int{45, 20 , 35, 75, 60 ,50, 25}

	// mutates the ages
	sort.Ints(ages)

	fmt.Println(ages)

	// return the 
	// If the value is not found, it returns the index where the value would be inserted to maintain the sorted order.
	index:= sort.SearchInts(ages, 1)

	fmt.Println(index)

	names:= []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	//sort also mutates the original array and sorts it
	sort.Strings(names)
	fmt.Println(names)

	// should give us the position of bowser but if it does not exist it will give the position where it should be
	fmt.Println(sort.SearchStrings(names, "bowser"))
}