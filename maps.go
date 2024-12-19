package main

import "fmt"

func maps() {
	menu := map[string]float64{
		"soup":      1.99,
		"pye":       7.90,
		"salad":     6.00,
		"strogonof": 78.90,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])

	//looping maps

	for k, v := range menu {
		fmt.Println("key is", k, "value is", v)
	}

	//int as key type
	phonebooks := map[int]string{
		15997624783: "mike",
	}

	fmt.Println(phonebooks[15997624783])
}
