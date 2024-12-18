package main

import (
	"fmt"
	"math"
)

func sayHello(name string) {
	fmt.Printf("Hello %v \n", name)
}

func cycleNames(names []string, f func(string)) {
	for _, v := range names {
		f(v)
	}
}

func cycleArea(r float64) float64 {
	return math.Pi * r * r
}

func functions() {
	sayHello("Mike")
	cycleNames([]string{"Mike", "Cesar", "Murilo"}, sayHello)
	a1 := cycleArea(10.5)
	a2 := cycleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f\n", a1, a2)

}
