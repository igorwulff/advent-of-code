package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2025/day4/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	g := shared.ParseInput(input)

	sum := 0
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			val := g.GetCell(x, y)
			if val == '.' {
				continue
			}

			if g.CheckAdjacent(x, y, 4) {
				sum++
			}
		}
	}

	return fmt.Sprint(sum)
}
