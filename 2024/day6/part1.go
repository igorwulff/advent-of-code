package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day6/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	grid, guard := shared.ParseInput(input)

	for {
		if !guard.Move(grid) {
			break
		}
	}

	return fmt.Sprint(len(guard.GetVisited()))
}
