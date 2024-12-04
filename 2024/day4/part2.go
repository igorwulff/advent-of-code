package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day4/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	grid := shared.ParseInput(input, []string{"A"})

	return fmt.Sprint(grid.FindX())
}
