package main

import (
	"fmt"
	"strings"
)

func getInitials(name string) (string, string) {
	s := strings.ToUpper(name)
	names := strings.Split(s, " ")

	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func multiplesReturn() {
	fn, sn := getInitials("mike lima")
	fmt.Println(fn, sn)

	fn2, sn2 := getInitials("mike")
	fmt.Println(fn2, sn2)
}
