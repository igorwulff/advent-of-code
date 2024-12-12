package shared

import "strings"

func ParseInput(input string) Grid {
	lines := strings.Split(input, "\n")

	g := Grid{
		Width:   len(lines[0]),
		Height:  len(lines),
		Cells:   make([]string, 0, len(lines[0])*len(lines)),
		Regions: make([]*Region, len(lines[0])*len(lines)),
	}

	for _, line := range lines {
		g.Cells = append(g.Cells, strings.Split(line, "")...)
	}

	return g
}
