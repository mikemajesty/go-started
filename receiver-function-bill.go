package main

import "fmt"

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
		tip: 1.90,
	}

	return bill
}

//format bill

func (bill billFunction) fomat() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0

	for k, v := range bill.items {
		fs += fmt.Sprintf("%-10v ...$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-10v ...$%0.2f", "total:", total)

	return fs
}
