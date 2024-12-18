package main

import (
	"fmt"
	"sort"
	"strings"
)

func library() {
	var hello = "mike"
	fmt.Println(strings.Contains(hello, "ik"))

	fmt.Println(strings.ReplaceAll(hello, "ik", "IK"))

	fmt.Println(strings.Index(hello, "ik"))

	fmt.Println(strings.Split(hello, "i"))

	ages := []int{1, 2, 4, 5, 6, 7, 8, 99, 56, 66, 56, 33, 31}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 5)
	fmt.Println(index)

	names := []string{"mike", "cesar", "mae", "pai", "tio"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "tio"))

}
