package shared

import (
	"strings"
)

type Grid struct {
	Width  int
	Height int
	Rows   [][]string
	Starts map[string][]int
}

func (g Grid) FindWords(word string) int {
	occ := 0

	for _, pos := range g.Starts[string(word[0])] {

		y := pos / g.Width
		x := pos % g.Width

		if g.FindWordInRow(word, x, y) {
			occ++
		}

		if g.FindWordInColumn(word, x, y) {
			occ++
		}

		occ += g.FindWordDiagonal(word, x, y)
	}

	return occ
}

func (g Grid) FindX() int {
	occ := 0

	return occ
}

func (g Grid) FindWordInRow(word string, x, y int) bool {
	if (x + len(word)) > g.Width {
		return false
	}

	value := strings.Join(g.Rows[y][x:x+len(word)], "")
	return value == word
}

func (g Grid) FindWordInColumn(word string, x, y int) bool {
	if (y + len(word)) > g.Height {
		return false
	}

	value := g.Rows[y][x] + g.Rows[y+1][x] + g.Rows[y+2][x] + g.Rows[y+3][x]
	return value == word
}

func (g Grid) FindWordDiagonal(word string, x, y int) int {
	occ := 0

	if (y + len(word) - 1) < g.Height {
		if (x + len(word) - 1) < g.Width {
			value := g.Rows[y][x] + g.Rows[y+1][x+1] + g.Rows[y+2][x+2] + g.Rows[y+3][x+3]
			if value == word {
				occ++
			}
		}

		if (x - len(word) + 1) >= 0 {
			value := g.Rows[y][x] + g.Rows[y+1][x-1] + g.Rows[y+2][x-2] + g.Rows[y+3][x-3]
			if value == word {
				occ++
			}
		}
	}

	return occ
}

// XMAS
// SAMX
