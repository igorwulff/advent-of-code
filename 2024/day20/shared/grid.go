package shared

import (
	"slices"
)

var PicoSeconds = 100

type Grid struct {
	Width   int
	Height  int
	Cells   []string
	User    *User
	EndTile *Tile
	Walls   []int
}

type Tile struct {
	X int
	Y int
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

func (g *Grid) GetCell(x, y int) string {
	return g.Cells[g.GetPos(x, y)]
}

func (g *Grid) AssignCheatWall(pos int, prevPos int) {
	if prevPos != -1 {
		g.Cells[prevPos] = "#"
	}

	g.Cells[pos] = "."
}
func (g *Grid) Draw(path []int) {
	for i := 0; i < len(g.Cells); i++ {
		if i%g.Width == 0 {
			println()
		}

		if slices.Contains(path, i) {
			print("O")
		} else {
			print(g.Cells[i])
		}

		print(g.Cells[i])
	}
	println()
}
