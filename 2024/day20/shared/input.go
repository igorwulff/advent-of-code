package shared

import (
	"strings"
)

func ParseInput(input string) *Grid {
	lines := strings.Split(input, "\n")

	g := Grid{
		Width:   len(lines[0]),
		Height:  len(lines),
		Cells:   make([]string, 0),
		EndTile: &Tile{},
		Walls:   []int{},
	}

	g.User = &User{
		Grid: &g,
		X:    0,
		Y:    0,
	}

	for y, line := range lines {
		for x, cell := range strings.Split(line, "") {
			g.Cells = append(g.Cells, cell)

			if cell == "#" {
				g.Walls = append(g.Walls, g.GetPos(x, y))
			}

			if cell == "E" {
				g.EndTile = &Tile{
					X: x,
					Y: y,
				}
			}

			if cell == "S" {
				g.User.X = x
				g.User.Y = y
			}
		}
	}

	return &g
}
