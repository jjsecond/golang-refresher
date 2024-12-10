package main

import (
	"fmt"
)

func main() {
	myBill := newBill("John's bill")

	myBill.addItem("onion soup", 4.50)
	myBill.addItem("chicken sandwich", 8.25)
	myBill.addItem("coke", 1.50)
	myBill.addItem("cake", 3.99)

	myBill.updateTip(10)
	fmt.Println(myBill.format())

}
