package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2025/day9/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	corners := shared.ParseInput(input)
	distances := shared.FindDistances(corners)

	return fmt.Sprint(distances[0].Size)
}
