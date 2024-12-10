package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills fn
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// a receiver function
// by only allowing a bill object, can only be called by a bill object (b bill)
// so myBill.format()
// the b is still just a copy
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	//list items
	// -25 pushes the prices right ober to the left
	// it is why the colon is on the key rather than in the string
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	//add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)

	return fs
}

// update the tip
// on receiver function we only take a copy of the bill
func (b *bill) updateTip(tip float64) {
	// derefrences a struct
	// so updates the tip on the original value
	// (*b).tip = tip
	// but we don't have to can just b.tip = tip
	b.tip = tip
}

// add item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) saveBill() {
	// when we save data to a file we need to save it in byte slice format
	data := []byte(b.format())

	// last argument is permissions
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file")
}
