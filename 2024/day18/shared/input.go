package shared

import (
	"strconv"
	"strings"
)

func ParseInput(input string) Grid {
	lines := strings.Split(input, "\n")

	g := Grid{
		Width:  Width,
		Height: Height,
		Cells:  make([]string, 0, Width*Height),
		EndTile: &Tile{
			X: Width - 1,
			Y: Height - 1,
		},
		CorruptedBytes: []int{},
	}

	g.User = &User{
		Grid: &g,
		X:    0,
		Y:    0,
	}

	for i := 0; i < Width*Height; i++ {
		g.Cells = append(g.Cells, ".")
	}

	corruptedBytes := 0
	for _, line := range lines {
		coord := strings.Split(line, ",")
		x, _ := strconv.Atoi(coord[0])
		y, _ := strconv.Atoi(coord[1])

		if CorruptedBytes > corruptedBytes {
			corruptedBytes++
			g.Cells[g.GetPos(x, y)] = "#"
			continue
		}

		g.CorruptedBytes = append(g.CorruptedBytes, g.GetPos(x, y))
	}

	return g
}
