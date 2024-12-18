package main

import "fmt"

func loops() {
	x := 0

	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}
	fmt.Println("-------------------------")
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	names := []string{"mike", "cesar", "mae", "pai", "tio"}

	for i := 0; i < len(names); i++ {
		fmt.Println("name is:", names[i])
	}

	for index, value := range names {
		fmt.Println("index is:", index, "and name is:", value)
	}

	for _, value := range names {
		fmt.Println("name is:", value)
		value = "new string"
	}

	fmt.Println("names:", names)

}
