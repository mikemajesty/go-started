package main

import "fmt"

func main() {

	var ages [3]int = [3]int{2, 4, 6}
	var ages2 = [3]int{2, 4, 6}

	fmt.Println(ages, len(ages))
	fmt.Println(ages2, len(ages2))

	names := [4]string{"name1", "name2", "name3", "name4"}
	names[1] = "names mudado"
	fmt.Println(names, len(names))

	//slices
	var scores = []int{100, 440, 550}
	scores[1] = 25
	fmt.Println(scores)
	scores = append(scores, 666)
	fmt.Println(scores)

	//slices range

	range1 := names[1:3]
	range2 := names[2:]
	range3 := names[:3]
	fmt.Println(range1, range2, range3)
}
