package shared

import (
	"slices"
	"strings"
)

func ParseInput(input string, values []string) Grid {
	lines := strings.Split(input, "\n")

	grid := Grid{
		Rows:   make([][]string, 0),
		Starts: make(map[string][]int),
		Height: len(lines),
		Width:  len(lines[0]),
	}

	for y, line := range lines {
		if line == "" {
			continue
		}

		row := strings.Split(line, "")

		for x, cell := range row {
			if !slices.Contains(values, cell) {
				continue
			}

			grid.Starts[cell] = append(grid.Starts[cell], (grid.Width*y)+x)
		}

		grid.Rows = append(grid.Rows, row)
	}

	return grid
}
