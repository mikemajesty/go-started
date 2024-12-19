package main

import (
	"fmt"
	"os"
)

type billFunction struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newFuncionBill(name string) billFunction {
	bill := billFunction{
		name: name,
		items: map[string]float64{
			"soup":      1.99,
			"pye":       7.90,
			"salad":     6.00,
			"strogonof": 78.90,
		},
		tip: 0,
	}

	return bill
}

//format bill

func (bill *billFunction) fomat() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0

	for k, v := range bill.items {
		fs += fmt.Sprintf("%-10v ...$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-10v ...$%v\n\n", "tip:", bill.tip)

	fs += fmt.Sprintf("%-10v ...$%0.2f", "total:", total+bill.tip)

	return fs
}

func (bill *billFunction) updateTip(tip float64) {
	bill.tip = tip
}

func (bill *billFunction) addItem(name string, price float64) {
	bill.items[name] = price
}

func (bill *billFunction) save() {
	data := []byte(bill.fomat())
	err := os.WriteFile(""+bill.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}
