package shared

import "strconv"

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
	if depth == len(e.Values)-1 {
		return []int{value}
	}

	addResults := e.Calc(depth+1, value+e.Values[depth+1])
	mulResults := e.Calc(depth+1, value*e.Values[depth+1])

	if e.Concat {
		concat := strconv.Itoa(value) + strconv.Itoa(e.Values[depth+1])
		concatInt, _ := strconv.Atoi(concat)
		concatResults := e.Calc(depth+1, concatInt)
		return append(append(addResults, mulResults...), concatResults...)
	}

	// Combine the results into a single slice
	return append(addResults, mulResults...)
}
