package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day4/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	grid := shared.ParseInput(input)

	return fmt.Sprint(grid.FindWords("XMAS") + grid.FindWords("SAMX"))
}
