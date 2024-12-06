package main

import (
	"fmt"

	"github.com/igorwulff/advent-of-code/2024/day6/shared"
)

// Exported function to be called by the main application
func Part2(input string) string {
	grid, guard := shared.ParseInput(input)

	startx := guard.X
	starty := guard.Y

	stuck := 0

	for y := range grid.Width {
		for x := range grid.Height {
			if startx == x && starty == y {
				continue // Skip starting position
			}

			if grid.IsObstacle(x, y) {
				continue // Skip obstacles
			}

			grid.SetObstacle(x, y)

			for {
				if m, err := guard.Move(grid); !m {
					if err != nil {
						stuck++
					}
					break
				}
			}

			grid.RemoveObstacle(x, y)
			guard.Reset(grid, startx, starty)
		}
	}

	return fmt.Sprint(stuck)
}
