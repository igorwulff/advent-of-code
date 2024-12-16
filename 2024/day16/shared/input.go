package shared

import (
	"strings"
)

func ParseInput(input string) Grid {
	lines := strings.Split(input, "\n")

	g := Grid{
		Width:   len(lines[0]),
		Height:  len(lines),
		Cells:   make([]string, 0, len(lines[0])*len(lines)),
		EndTile: &EndTile{},
	}

	g.Reindeer = &Reindeer{
		Grid: &g,
	}

	for _, line := range lines {
		cells := strings.Split(line, "")

		for i, cell := range cells {
			if cell == "S" {
				x, y := g.GetXY(len(g.Cells) + i)
				g.Reindeer.X = x
				g.Reindeer.Y = y
				g.Reindeer.Dir = East
			}

			if cell == "E" {
				x, y := g.GetXY(len(g.Cells) + i)
				g.EndTile.X = x
				g.EndTile.Y = y
			}
		}

		g.Cells = append(g.Cells, cells...)
	}

	return g
}
