package main

import "fmt"

func variables() {
	//string
	var name string = "Mario"
	var name2 = "Luigi"
	var name3 string

	fmt.Println(name, name2, name3)

	name3 = "Princess"

	fmt.Println(name, name2, name3)

	name4 := "Yosh"

	fmt.Println(name4)

	//ints

	var age int = 1
	var age2 = 2
	age3 := 2

	fmt.Println(age, age2, age3)

	//bits and memorry

	var bits int8 = 1
	var num3 uint = 4

	fmt.Println(bits, num3)

	//floats
	var floats float32 = 1000.00
	fmt.Println(floats)

}
