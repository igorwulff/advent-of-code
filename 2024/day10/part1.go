package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/igorwulff/advent-of-code/utils"
)

type Grid struct {
	Width  int
	Height int
	Cells  []int
	Starts []int
}

func (g Grid) InBounds(x, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

func (g *Grid) GetX(pos int) int {
	return pos % g.Width
}

func (g *Grid) GetY(pos int) int {
	return pos / g.Width
}

func (g Grid) GetXY(pos int) (int, int) {
	return g.GetX(pos), g.GetY(pos)
}

func (g *Grid) GetPos(x, y int) int {
	return (g.Width * y) + x
}

func (g *Grid) FindPath(pos int, ends *map[int]struct{}) {
	x, y := g.GetXY(pos)
	value := g.Cells[pos]

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// Don't include diagonal values or center.
			if utils.AbsInt(i) == utils.AbsInt(j) {
				continue
			}

			xx := x + i
			yy := y + j

			if !g.InBounds(xx, yy) {
				continue
			}

			next := g.GetPos(xx, yy)
			if value == 8 && g.Cells[next] == 9 {
				(*ends)[next] = struct{}{}
			} else if g.Cells[next] == value+1 {
				g.FindPath(next, ends)
			}
		}
	}
}

// Exported function to be called by the main application
func Part1(input string) string {
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

	sum := 0
	for _, start := range grid.Starts {
		ends := make(map[int]struct{}, 0)
		grid.FindPath(start, &ends)
		sum += len(ends)
	}

	return fmt.Sprint(sum)
}
