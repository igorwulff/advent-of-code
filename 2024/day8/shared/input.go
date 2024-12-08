package shared

import (
	"strings"
)

func ParseInput(input string) Grid {
	lines := strings.Split(input, "\n")

	grid := Grid{
		Width:    len(lines[0]),
		Height:   len(lines),
		Antennas: make(map[string][]int),
	}

	for y, line := range lines {
		if line == "" {
			continue
		}

		row := strings.Split(line, "")

		for x, v := range row {
			if v == "." {
				continue
			}

			grid.Antennas[v] = append(grid.Antennas[v], grid.GetPos(x, y))
		}
	}

	return grid
}
