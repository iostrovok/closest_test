package Сlosest

import (
	"sort"
)

// myStack is internal structure which checks and saves result list
type myStack struct {
	result      []int // result
	x           int   // processing value
	maxDif      int   // internal var, == max distance from x to any from result list
	countReturn int   // how many number we have to return
	lastId      int   // internal var
}

func (this myStack) Len() int      { return len(this.result) }
func (this myStack) Swap(i, j int) { this.result[i], this.result[j] = this.result[j], this.result[i] }
func (this myStack) Less(i, j int) bool {
	return abcDiff(this.x, this.result[i]) < abcDiff(this.x, this.result[j])
}

func (this *myStack) Sort() {
	sort.Sort(this)
}
func (this *myStack) getResult() []int {
	if len(this.result) > this.countReturn {
		return this.result[:this.countReturn]
	}

	return this.result
}

func newStack(x, countReturn int) *myStack {
	return &myStack{
		result:      []int{},
		x:           x,
		countReturn: countReturn,
	}
}

/*
	initStack gets count first elment from list, puts to result and sotrs result
*/
func initStack(list []int, x, countReturn int) *myStack {

	if countReturn > len(list) {
		countReturn = len(list)
	}

	this := newStack(x, countReturn)
	this.result = make([]int, countReturn, countReturn)

	copy(this.result, list[:countReturn])
	sort.Sort(this)

	this.lastId = countReturn - 1
	this.maxDif = abcDiff(x, this.result[this.lastId])

	return this
}

// push It adds element to result list (if it's necessary) and sorts result list
func (this *myStack) push(val int) {

	dif := abcDiff(this.x, val)
	if dif >= this.maxDif {
		return
	}

	// put to last position in list
	this.result[this.lastId] = val

	// move to truth position
	// res must be sorted
	for i := this.lastId; i > 0; i-- {
		if dif >= abcDiff(this.x, this.result[i-1]) {
			break
		}
		this.result[i-1], this.result[i] = this.result[i], this.result[i-1]
	}

	this.maxDif = abcDiff(this.x, this.result[this.lastId])
}
