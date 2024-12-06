package shared

import (
	"strings"
)

func ParseInput(input string) (Grid, Guard) {
	lines := strings.Split(input, "\n")

	grid := Grid{
		Width:     len(lines[0]),
		Height:    len(lines),
		Obstacles: make(map[int]bool, 0),
	}

	guard := Guard{
		Path:    make([]int, 0),
		Visited: make(map[int]Dir),
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
				guard.SetPos(grid, x, y)
			}
		}
	}

	return grid, guard
}
