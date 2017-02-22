package Сlosest

import (
	"math/rand"
	"testing"
)

func check_search_smoke(t *testing.T, myList []int, step, x, diff int) {
	res := search(myList, step, x, 5)

	for i, v := range res {
		if abs_dif(v, x) > diff {
			t.Errorf("Test_huge... error. x: %d,\nres: %#v,\ni: %d, x: %d,\ndiff: %d", x, res, i, x, diff)
		}
	}
}

func Test_search_neg(t *testing.T) {
	// // t.Skip("no now")

	x := -1
	diff := 1
	myList := []int{
		1, -1, 2, 1, 1, 3, 1, 1, 9,
		4, 1, 5, 1, 1, 6, 0, 1, 3, 9,
		4, 1, 5, 1, 1, 6, 0, 1, 3, 9,
		4, 1, 5, 1, 1, 6, 0, 1, 3,
		4, 1, 5, 1, 1, 6, 0, 1, 3,
		4, 1, 5, 1, 1, 6, 0, 1, 3, 10,
		4, 1, 5, 1, 1, 6, 0, 1, 3, 9,
	}

	check_search_smoke(t, myList, 10, x, diff)
}

func Test_search_smoke(t *testing.T) {
	// t.Skip("no now")

	x := 10
	diff := 1
	myList := []int{
		1, -1, 2, 1, 1, 3, 1, 1, 9,
		4, 1, 5, 1, 1, 6, 0, 1, 3, 9,
		4, 1, 5, 1, 1, 6, 0, 1, 3, 9,
		4, 1, 5, 1, 1, 6, 0, 1, 3,
		4, 1, 5, 1, 1, 6, 0, 1, 3,
		4, 1, 5, 1, 1, 6, 0, 1, 3, 10,
		4, 1, 5, 1, 1, 6, 0, 1, 3, 9,
	}

	check_search_smoke(t, myList, 10, x, diff)
}

func Test_search_negative(t *testing.T) {
	// // t.Skip("no now")

	x := -10
	diff := 1
	myList := []int{
		1, -1, 2, 1, 1, 3, 1, 1, 9,
		4, 1, 5, 1, 1, 6, 0, 1, 3, 9,
		4, 1, 5, 1, 1, 6, 0, 1, 3, -9,
		4, 1, 5, 1, 1, 6, 0, 1, 3,
		4, 1, 5, 1, 1, 6, 0, 1, 3,
		4, 1, 5, 1, 1, 6, 0, 1, 3, -10,
		4, 1, 5, 1, 1, 6, 0, 1, 3, -9,
		-1, -1, -2, -1, -1, 3, 1, 1, 9,
		-4, -1, -5, -1, -1, 6, 0, 1, 3, -9,
		-4, -1, -5, -1, -1, 6, 0, 1, 3, 9,
		-4, -1, -5, -1, 1, 6, 0, 1, 3,
		-4, -1, -5, -1, 1, 6, 0, 1, 3,
		-4, -1, -5, -1, 1, 6, 0, 1, 3, -10,
		-4, -1, -5, -1, 1, 6, 0, 1, 3, 9,
	}

	check_search_smoke(t, myList, 10, x, diff)
}

func Test_search_long(t *testing.T) {

	// // t.Skip("no now")

	x := 100
	diff := 1

	myList := []int{}

	t.Log("Start prepare data for test")
	count := 10000000
	for count > 0 {
		myList = append(myList, rand.Intn(100))
		count--
	}
	t.Log("Finish prepare data for test")

	check_search_smoke(t, myList, 1000000, x, diff)
}

func Test_FindHuge_long(t *testing.T) {

	// // t.Skip("no now")

	x := 100
	diff := 1

	myList := []int{}

	t.Log("Start prepare data for test")
	count := 10000000
	for count > 0 {
		myList = append(myList, rand.Intn(1000000))
		count--
	}
	t.Log("Finish prepare data for test")

	res := FindHuge(myList, x, 5)
	for i, v := range res {
		if abs_dif(v, x) > diff {
			t.Errorf("Test_huge_smoke error. x: %d,\nres: %#v,\ni: %d, x: %d,\ndiff: %d", x, res, i, x, diff)
		}
	}
}
