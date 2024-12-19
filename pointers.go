package main

import "fmt"

func updateNewName(name *string) {
	*name = "cesar"
}

func pointers() {
	name := "mike"

	fmt.Println("memory address of name is: ", &name)

	m := &name

	fmt.Println("old name", name)
	updateNewName(m)
	fmt.Println("new name", *m)
}
