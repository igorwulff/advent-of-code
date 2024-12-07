package shared

import (
	"strconv"
)

type Equation struct {
	Result int
	Values []int
	Concat bool
}

func (e Equation) IsValid() bool {
	sums := e.Calc(0, e.Values[0])

	for _, sum := range sums {
		if sum == e.Result {
			return true
		}
	}

	return false
}

func (e Equation) Calc(depth, value int) []int {
	if value > e.Result {
		return []int{}
	}

	if depth == len(e.Values)-1 {
		return []int{value}
	}

	add := e.Calc(depth+1, value+e.Values[depth+1])
	mul := e.Calc(depth+1, value*e.Values[depth+1])

	if e.Concat {
		concatString := strconv.Itoa(value) + strconv.Itoa(e.Values[depth+1])
		concat, _ := strconv.Atoi(concatString)
		return append(append(add, mul...), e.Calc(depth+1, concat)...)
	}

	// Combine the results into a single slice
	return append(add, mul...)
}
