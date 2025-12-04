package shared

import (
	"strings"
)

type Grid struct {
	Width  int
	Height int
	Cells  [][]rune
}

func (g Grid) InBounds(x, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

func (g Grid) GetX(pos int) int {
	return pos % g.Width
}

func (g Grid) GetY(pos int) int {
	return pos / g.Width
}

func (g Grid) GetXY(pos int) (int, int) {
	return g.GetX(pos), g.GetY(pos)
}

func (g Grid) GetPos(x, y int) int {
	return (y * g.Width) + x
}

func (g *Grid) Mark(x, y int) {
	g.Cells[y][x] = '.'
}

func (g *Grid) GetCell(x, y int) rune {
	return g.Cells[y][x]
}

func (g *Grid) Draw() {
	for y := 0; y < g.Height; y++ {
		println(g.Cells[y])

	}
	println()
}

func (g *Grid) CheckAdjacent(posX, posY, max int) bool {
	for x := posX - 1; x <= posX+1; x++ {
		for y := posY - 1; y <= posY+1; y++ {
			if x == posX && y == posY {
				continue
			}

			if !g.InBounds(x, y) {
				continue
			}

			if g.GetCell(x, y) == '@' {
				max--
				if max == 0 {
					return false
				}
			}
		}
	}

	return true
}

func ParseInput(input string) *Grid {
	lines := strings.Split(input, "\n")

	g := Grid{
		Width:  len(lines[0]),
		Height: 0,
		Cells:  make([][]rune, 0),
	}

	for _, line := range lines {
		if line == "" {
			continue
		}
		g.Cells = append(g.Cells, []rune(line))
		g.Height++
	}

	return &g
}
