package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2025/day6/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	ws := shared.ParseInput(input)

	var sum int64 = 0
	for _, w := range *ws {
		sum += w.Calculate()
	}

	return fmt.Sprint(sum)
}
