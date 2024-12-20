package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day20/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	grid := shared.ParseInput(input)

	path := grid.User.FindPath(grid.User.X, grid.User.Y, shared.North)

	sum := grid.User.FindCheats(path, 2)

	return fmt.Sprint(sum)
}
