package shared

import (
	"strconv"
)

type Equation struct {
	Result    int
	Values    []int
	UseConcat bool
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

func (e Equation) Concat(l, r int) int {
	concatString := strconv.Itoa(l) + strconv.Itoa(r)
	concat, _ := strconv.Atoi(concatString)
	return concat
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

	if e.UseConcat {
		return append(append(add, mul...), e.Calc(depth+1, e.Concat(value, e.Values[depth+1]))...)
	}

	// Combine the results into a single slice
	return append(add, mul...)
}
