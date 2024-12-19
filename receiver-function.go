package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price must be a number", err)
			promptOpt(bill)
		}

		bill.addItem(name, p)

		fmt.Printf("item added - %v: $%0.2f\n", name, p)
	case "s":
		tip, _ := getInput("Enter tip amount: ", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The price must be a number", err)
			promptOpt(bill)
		}

		bill.updateTip(t)

		fmt.Println("tip added - ", t)
		promptOpt(bill)
	case "t":
		fmt.Println("You chose to save the bill", bill)
	default:
		fmt.Println("You chose not allowed")
		promptOpt(bill)
	}
}

func main() {
	myBill := createBill()
	promptOpt(myBill)

	fmt.Println(myBill)
}
