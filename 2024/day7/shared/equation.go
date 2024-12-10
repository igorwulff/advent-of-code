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

	depth++

	if depth == len(e.Values) {
		return []int{value}
	}

	add := e.Calc(depth, value+e.Values[depth])
	mul := e.Calc(depth, value*e.Values[depth])

	results := append(add, mul...)

	if e.UseConcat {
		concat := e.Calc(depth, e.Concat(value, e.Values[depth]))
		results = append(results, concat...)
	}

	return results
}
