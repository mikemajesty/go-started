package main

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	bill := bill{
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
