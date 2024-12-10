package shared

import (
	"strconv"
	"strings"
)

func ParseInput(input string) Grid {
	lines := strings.Split(input, "\n")

	grid := Grid{
		Width:  len(lines[0]),
		Height: len(lines),
		Cells:  make([]int, 0, len(lines[0])*len(lines)),
		Starts: make([]int, 0),
	}

	for y, line := range lines {
		if line == "" {
			continue
		}

		row := strings.Split(line, "")
		for x, v := range row {
			var value int
			if v == "." {
				value = -1
			} else {
				value, _ = strconv.Atoi(v)
			}

			grid.Cells = append(grid.Cells, value)

			if value == 0 {
				grid.Starts = append(grid.Starts, grid.GetPos(x, y))
			}
		}
	}

	return grid
}
