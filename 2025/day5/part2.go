package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2025/day5/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	ingr := shared.ParseInput(input)

	return fmt.Sprint(ingr.CountRanges(ingr.Fresh))
}
