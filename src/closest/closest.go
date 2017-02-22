package Сlosest

func Find(list []int, x int, count int) []int {

	if len(list) == 0 {
		return []int{}
	}

	// init local stack
	stack := initStack(list, x, count)

	// find better results
	for i := count; i < len(list); i++ {
		stack.push(list[i])

		// It's case when we have many elements which equal to x
		if stack.maxDif == 0 {
			return stack.result
		}
	}

	return stack.result
}

// likes Math.Abs(x - y)
func abs_dif(x, y int) int {

	dif := x - y

	if dif < 0 {
		return -dif
	}

	return dif
}
