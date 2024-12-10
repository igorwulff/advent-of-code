package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/igorwulff/advent-of-code/2024/day10/shared"
)

// Exported function to be called by the main application
func Part1(input string) string {
	lines := strings.Split(input, "\n")

	grid := shared.Grid{
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

	sum := 0
	for _, start := range grid.Starts {
		ends := make(map[int]int, 0)
		grid.FindPath(start, &ends)
		sum += len(ends)
	}

	return fmt.Sprint(sum)
}
