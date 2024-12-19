package main

import "fmt"

func main() {
	myBill := newFuncionBill("maario's bill")
	myBill.updateTip(10)
	myBill.addItem("onion ring", 5.21)
	fmt.Println(myBill.fomat())
}
