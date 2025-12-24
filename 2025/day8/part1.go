package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2025/day8/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	distances := shared.ParseInput(input)

	circuits, _ := shared.ConnectJunctions(distances, shared.ConnectCount)

	sum := 1
	for i := range 3 {
		sum *= len(circuits[i].Junctions)
	}

	return fmt.Sprint(sum)
}
