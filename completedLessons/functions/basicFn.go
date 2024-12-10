package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Good morning %v \n", name)
}

func sayBye(name string) {
	fmt.Printf("Say bye %v \n", name)
}

func cycleNames(name []string, f func(string)) {
	for _, value := range name {
		f(value)
	}
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func basicFn() {
	sayGreeting("mario")
	sayGreeting("Luigi")
	sayBye("mario")

	names := []string{"mario", "Luigi", "Cloud", "Tifa"}
	cycleNames(names, sayGreeting)
	cycleNames(names, sayBye)

	a1 := circleArea(10.34)
	a2 := circleArea(15)

	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)

}
