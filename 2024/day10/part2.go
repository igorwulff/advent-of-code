package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day10/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	grid := shared.ParseInput(input)

	sum := 0
	for _, start := range grid.Starts {
		ends := make(map[int]int, 0)
		grid.FindPath(start, &ends)

		for _, v := range ends {
			sum += v
		}
	}

	return fmt.Sprint(sum)
}
