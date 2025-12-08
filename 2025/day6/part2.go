package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2025/day6/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	ws := shared.ParseInput(input)

	var sum int64 = 0
	for _, w := range *ws {
		sum += w.CalculateRL()
	}

	return fmt.Sprint(sum)
}
