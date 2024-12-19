package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, read *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := read.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() billFunction {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Enter the bill name: ", reader)

	bill := newFuncionBill(name)
	fmt.Println("Created the bill -", bill.name)
	return bill
}

func promptOpt(bill billFunction) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose options (a - add item, s - save bill, t - add tip) ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Enter item name: ", reader)
		price, _ := getInput("Enter item price: ", reader)
		fmt.Printf(name, price)
	case "s":
		tip, _ := getInput("Enter tip amount: ", reader)
		fmt.Printf(tip)
	case "t":
		fmt.Printf("You chose %v\n", opt)
	default:
		fmt.Printf("You chose not allowed")
		promptOpt(bill)
	}
}

func main() {
	myBill := createBill()
	promptOpt(myBill)

	fmt.Println(myBill)
}
