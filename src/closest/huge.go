package Сlosest

import (
	"errors"
	"sync"
)

const CountSingleGoroutine = 10000
const CountGoroutine = 8
const CountByDefault = 5

func FindHuge(list []int, x int, counts ...int) []int {

	count := len(list)

	countReturn, err := checkCounts(counts...)
	if err != nil {
		// TODO make errors processing
		return []int{}
	}
	// Too short list
	if count <= countReturn {
		return list
	}

	// list is not as long as we need
	if count <= CountSingleGoroutine {
		return Find(list, x, countReturn)
	}

	step := count / getCountGoroutine()

	return search(list, step, x, countReturn)
}

func getCountGoroutine() int {
	// TODO need to change: it depends on processor cores and memory
	return CountGoroutine
}

func checkCounts(counts ...int) (int, error) {
	// How many number we have to return
	count := CountByDefault
	if len(counts) > 0 {
		count = counts[0]
	}

	if count < 1 {
		return 0, errors.New("Bad returned count")
	}

	return count, nil
}

// search splits to goroutine by step
func search(list []int, step, x, countReturn int) []int {

	result_list := []int{}
	count := len(list)

	var globalLock sync.Mutex
	var wg sync.WaitGroup
	i := 0
	for i <= count {
		wg.Add(1)
		go func(list []int, x int, countReturn int) {
			out := Find(list, x, countReturn)

			globalLock.Lock()
			// Merge small result to common result list
			result_list = append(result_list, out...)
			globalLock.Unlock()

			wg.Done()
		}(list[i:min(i+step, count)], x, countReturn)
		i = i + step
	}
	wg.Wait()

	// Sort and find top of values which are the closest to x
	stack := newStack(x, countReturn)
	stack.result = result_list
	stack.Sort()
	return stack.getResult()
}

//
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
