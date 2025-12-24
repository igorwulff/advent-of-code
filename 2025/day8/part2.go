package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2025/day8/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	distances := shared.ParseInput(input)

	_, lastDistance := shared.ConnectJunctions(distances, len(distances))

	return fmt.Sprint(lastDistance.A.X * lastDistance.B.X)
}
