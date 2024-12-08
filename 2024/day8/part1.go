package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day8/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	grid := shared.ParseInput(input)
	locations := make(map[int]bool, 0)

	for _, antenna := range grid.Antennas {
		for _, l1 := range antenna {
			for _, l2 := range antenna {
				if l1 == l2 {
					continue
				}

				x1 := grid.GetX(l1)
				y1 := grid.GetY(l1)
				x2 := grid.GetX(l2)
				y2 := grid.GetY(l2)

				x := x1 + (x1 - x2)
				y := y1 + (y1 - y2)

				if grid.InBounds(x, y) {
					locations[grid.GetPos(x, y)] = true
				}
			}
		}
	}

	return fmt.Sprint(len(locations))
}
