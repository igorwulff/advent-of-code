package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Multiply int = iota
	Add
)

type Equation struct {
	Result int
	Values []int
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

	// Combine the results into a single slice
	return append(addResults, mulResults...)
}

// Exported function to be called by the main application
func Part1(input string) string {
	equations := ParseInput(input)

	sum := 0
	for _, e := range equations {
		if e.IsValid() {
			sum += e.Result
		}
	}

	return fmt.Sprint(sum)
}

func ParseInput(input string) []Equation {
	equations := make([]Equation, 0)

	// Split input by new line
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		// Split line by space
		parts := strings.Split(line, " ")
		result, _ := strconv.Atoi(parts[0][:len(parts[0])-1])

		values := make([]int, 0)
		for i := 1; i < len(parts); i++ {
			v, _ := strconv.Atoi(parts[i])
			values = append(values, v)
		}

		equations = append(equations, Equation{
			Result: result,
			Values: values,
		})
	}

	return equations
}
