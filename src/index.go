package main

import (
	"fmt"

	"closest"
	"math/rand"
)

// Simple test
// more test int closest/*_test.go
func main() {
	fmt.Printf("START\n")

	x := 3
	myList := []int{
		1, -1, 2, 1, 1, 3, 1, 1,
		4, 1, 5, 1, 3, 2, 1, 6, 0, 1, 3,
	}

	fmt.Printf("Mysort.CountByDefault: %d\n", Сlosest.CountByDefault)
	fmt.Printf("x: %d\n", x)

	fmt.Printf("myList: %#v\n\n", myList)

	find := Сlosest.Find(myList, x, 5)
	fmt.Printf("Find: %#v\n", find)

	myList = []int{}
	count := 10000000
	for count > 0 {
		a := rand.Intn(1000000)
		if rand.Intn(1000000) > 50000 {
			a = -1 * a
		}
		myList = append(myList, a)
		count--
	}

	find = Сlosest.FindHuge(myList, -111)
	fmt.Printf("FindHuge: %#v\n", find)
}
