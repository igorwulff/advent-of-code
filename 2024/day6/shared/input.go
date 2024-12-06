package shared

import (
	"strings"
)

func ParseInput(input string) (Grid, Guard) {
	lines := strings.Split(input, "\n")

	grid := Grid{
		Width:     len(lines[0]),
		Height:    len(lines),
		Obstacles: make([]int, 0),
	}

	guard := Guard{
		Path: make([]int, 0),
	}

	for y, line := range lines {
		if line == "" {
			continue
		}

		row := strings.Split(line, "")

		for x, cell := range row {
			if cell == "#" {
				grid.SetObstacle(x, y)
			}

			if cell == "^" {
				guard.X = x
				guard.Y = y
			}
		}
	}

	return grid, guard
}
