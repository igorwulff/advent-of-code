package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day8/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	grid := shared.ParseInput(input)
	locations := make(map[int]struct{}, 0)

	for _, antenna := range grid.Antennas {
		for _, a1 := range antenna {
			for _, a2 := range antenna {
				if a1 == a2 {
					continue
				}

				x1 := grid.GetX(a1)
				y1 := grid.GetY(a1)
				x2 := grid.GetX(a2)
				y2 := grid.GetY(a2)

				mul := 0
				for {
					x := x1 + ((x1 - x2) * mul)
					y := y1 + ((y1 - y2) * mul)

					if !grid.InBounds(x, y) {
						break
					}

					locations[grid.GetPos(x, y)] = struct{}{}
					mul++
				}
			}
		}
	}

	return fmt.Sprint(len(locations))
}
