package shared

import (
	"strconv"
	"strings"
)

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
