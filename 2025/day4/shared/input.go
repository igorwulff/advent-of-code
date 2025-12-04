package shared

import (
	"strings"
)

type Grid struct {
	Width  int
	Height int
	Cells  []string
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

func (g *Grid) Mark(x, y int) {
	r := []rune(g.Cells[y])
	r[x] = '.'

	g.Cells[y] = string(r)
}

func (g *Grid) GetCell(x, y int) string {
	pos := g.Cells[y]
	val := string(pos[x])
	return val
}

func (g *Grid) Draw() {
	for y := 0; y < g.Height; y++ {
		println(g.Cells[y])

	}
	println()
}

func (g *Grid) CheckAdjacent(posX, posY int) int {
	count := 0

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			xx := posX + x
			yy := posY + y
			if (x == 0 && y == 0) || !g.InBounds(xx, yy) {
				continue
			}

			if g.GetCell(xx, yy) == "@" {
				count++
			}
		}
	}

	return count
}

func ParseInput(input string) *Grid {
	lines := strings.Split(input, "\n")

	g := Grid{
		Width:  len(lines[0]),
		Height: 0,
		Cells:  make([]string, 0),
	}

	for _, line := range lines {
		if line == "" {
			continue
		}
		g.Cells = append(g.Cells, line)
		g.Height++
	}

	return &g
}
