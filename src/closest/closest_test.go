package Сlosest

import (
	"fmt"
	"testing"
)

func check_point(t *testing.T, str string, list []int, val1, val2 int) {
	if val1 != val2 {
		t.Errorf(str, list, val1, val2)
	}
}

func Test_abcDiff(t *testing.T) {

	// t.Skip("no now")

	test_data := [][]int{
		[]int{1, 2, 1},
		[]int{10, 2, 8},
		[]int{-1, 2, 3},
		[]int{8, -2, 10},
	}

	for _, v := range test_data {
		if abcDiff(v[0], v[1]) != v[2] {
			t.Errorf("abcDiff error: %#v", v)
		}
	}

}

func Test_initStack(t *testing.T) {

	// t.Skip("no now")

	stack := initStack([]int{1, 2, 3, 4, 5}, 3, 3)

	check_point(t, "1-1. initStack %#v: %x %x", stack.result, stack.result[0], 3)
	check_point(t, "1-2. initStack %#v: %x %x", stack.result, stack.result[1], 2)
	check_point(t, "1-3. initStack %#v: %x %x", stack.result, stack.result[2], 1)
	check_point(t, "1-4. initStack %#v: %x %x", stack.result, stack.maxDif, 2) // 3 - 1
	check_point(t, "1-5. initStack %#v: %x %x", stack.result, stack.lastId, 2)
}

func Test_initStack_Negative(t *testing.T) {

	// t.Skip("no now")

	stack := initStack([]int{-3, -2, -1, -4, -5}, 3, 4)

	check_point(t, "1. initStack_Negative %#v: %x %x", stack.result, stack.result[0], -1)
	check_point(t, "2. initStack_Negative %#v: %x %x", stack.result, stack.result[1], -2)
	check_point(t, "3. initStack_Negative %#v: %x %x", stack.result, stack.result[2], -3)
	check_point(t, "3. initStack_Negative %#v: %x %x", stack.result, stack.result[3], -4)
	check_point(t, "4. initStack_Negative %#v: %x %x", stack.result, stack.maxDif, 7) // 3 - -4
	check_point(t, "5. initStack_Negative %#v: %x %x", stack.result, stack.lastId, 3)
}

func Test_initStack_Mix(t *testing.T) {
	// t.Skip("no now")

	stack := initStack([]int{-3, -2, 0, -5}, 3, 3)

	check_point(t, "1. initStack_Mix %#v: %x %x", stack.result, stack.result[0], 0)
	check_point(t, "2. initStack_Mix %#v: %x %x", stack.result, stack.result[1], -2)
	check_point(t, "3. initStack_Mix %#v: %x %x", stack.result, stack.result[2], -3)
	check_point(t, "4. initStack_Mix %#v: %x %x", stack.result, stack.maxDif, 6) // 3 - -3
	check_point(t, "5. initStack_Mix %#v: %x %x", stack.result, stack.lastId, 2)
}

func Test_push(t *testing.T) {

	// t.Skip("no now")

	// Pushs to middle
	stack := initStack([]int{-30, -20, -10, -40, -50}, 3, 4)
	stack.push(2)
	// check_point(t, "init_stack %#v: %x %x", list[0], 3, 0)
	check_point(t, "1-1. push %#v: %d %d", stack.result, stack.result[0], 2)
	check_point(t, "1-2. push %#v: %d %d", stack.result, stack.result[1], -10)
	check_point(t, "1-3. push %#v: %d %d", stack.result, stack.result[2], -20)
	check_point(t, "1-3. push %#v: %d %d", stack.result, stack.result[3], -30)

	stack.push(2)
	check_point(t, "2-1. push %#v: %d %d", stack.result, stack.result[0], 2)
	check_point(t, "2-2. push %#v: %d %d", stack.result, stack.result[1], 2)
	check_point(t, "2-3. push %#v: %d %d", stack.result, stack.result[2], -10)
	check_point(t, "2-3. push %#v: %d %d", stack.result, stack.result[3], -20)

	stack.push(2)
	check_point(t, "3-1. push %#v: %d %d", stack.result, stack.result[0], 2)
	check_point(t, "3-2. push %#v: %d %d", stack.result, stack.result[1], 2)
	check_point(t, "3-3. push %#v: %d %d", stack.result, stack.result[2], 2)
	check_point(t, "3-3. push %#v: %d %d", stack.result, stack.result[3], -10)

	stack.push(2)
	check_point(t, "4-1. push %#v: %d %d", stack.result, stack.result[0], 2)
	check_point(t, "4-2. push %#v: %d %d", stack.result, stack.result[1], 2)
	check_point(t, "4-3. push %#v: %d %d", stack.result, stack.result[2], 2)
	check_point(t, "4-3. push %#v: %d %d", stack.result, stack.result[3], 2)
}

func Test_Find(t *testing.T) {
	// t.Skip("no now")

	myList := []int{
		1, -1, 2, 1, 1, 3, 1, 1,
		4, 1, 5, 1, 1, 6, 0, 1, 3,
	}

	check_Find(t, myList, 3, 2)
}

func Test_Find2(t *testing.T) {
	// t.Skip("no now")

	myList := []int{
		1, -1, 2, 1, 1, 3, 1, 1,
		4, 1, 5, 1, 1, 6, 0, 1, 3,
	}

	check_Find(t, myList, 0, 1)
}

func Test_Find_zero(t *testing.T) {
	myList := []int{
		1, -1, 2, 0, 0, 3, 1, 1,
		4, 1, 5, 1, 0, 0, 0, 1, 3,
		1, -1, 2, 1, 1, 3, 1, 1,
		4, 1, 5, 1, 1, 6, 0, 1, 3,
	}

	check_Find(t, myList, 0, 0)
}

func check_Find(t *testing.T, myList []int, x, diff int) {
	find := Find(myList, x, 5)

	fmt.Printf("find: %#v\n", find)

	for i, v := range find {
		if abcDiff(v, x) > diff {
			t.Errorf("check_Find error. x: %d,\nmyList: %#v,\nfind: %#v,\ni: %d, x: %d", x, myList, find, i, x)
		}
	}
}
