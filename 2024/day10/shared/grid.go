package shared

import "github.com/igorwulff/advent-of-code/utils"

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

func (g *Grid) FindPath(pos int, ends *map[int]int) {
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
				if _, ok := (*ends)[next]; ok {
					(*ends)[next]++
				} else {
					(*ends)[next] = 1
				}
			} else if g.Cells[next] == value+1 {
				g.FindPath(next, ends)
			}
		}
	}
}
