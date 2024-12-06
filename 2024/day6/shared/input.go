package shared

import (
	"strings"
)

func ParseInput(input string) (Grid, Guard) {
	lines := strings.Split(input, "\n")

	Width = len(lines[0])
	Height = len(lines)

	grid := Grid{}

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
				Obstacles.Store(grid.GetPos(x, y), true)
			}

			if cell == "^" {
				guard.SetPos(x, y, grid.GetPos(x, y))
			}
		}
	}

	return grid, guard
}
