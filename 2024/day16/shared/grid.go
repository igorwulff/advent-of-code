package shared

type Grid struct {
	Width    int
	Height   int
	Cells    []string
	Reindeer *Reindeer
	EndTile  *EndTile
}

type EndTile struct {
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
