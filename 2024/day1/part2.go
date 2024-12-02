package main

import (
	"github.com/igorwulff/advent-of-code/2024/day1/shared"

	"fmt"
)

// Exported function to be called by the main application
func Part2(input string) string {
	left, right := shared.ParseInput(input)

	m := make(map[int]int)

	for _, l := range left {
		i := 0

		for _, r := range right {
			if l == r {
				i++
			}
		}

		m[l] = i
	}

	output := 0
	for k, v := range m {
		output += k * v
	}

	return fmt.Sprint(output)
}
