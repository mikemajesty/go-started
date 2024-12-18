package main

import "fmt"

func printing() {
	//string
	var name string = "Hello, \n"

	fmt.Print(name)

	fmt.Println("skip")
	fmt.Println("line")

	age := 20
	myName := "Mike"

	fmt.Println("my age is", age, "and my name is", myName)

	fmt.Printf("my age is %v and my name is %v\n", age, myName)
	fmt.Printf("my age is %q and my name is %q\n", age, myName)
	fmt.Printf("age is of type %T\n", age)
	fmt.Printf("name is of type %T\n", name)
	fmt.Printf("you score %0.3f points\n", 19.89)

	var text = fmt.Sprintf("my age is %v and my name is %v\n", age, myName)
	fmt.Println("text", text)
}
