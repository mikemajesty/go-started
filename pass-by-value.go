package main

import "fmt"

func updateName(name string) string {
	name = "cesar"
	return name
}

func updateMenu(menu map[string]float64) {
	menu["soup"] = 2.00
}

func passByValue() {
	name := "mike"
	name = updateName(name)

	fmt.Println(name)

	menu := map[string]float64{
		"soup":      1.99,
		"pye":       7.90,
		"salad":     6.00,
		"strogonof": 78.90,
	}

	updateMenu(menu)

	fmt.Println(menu)
}
