package shared

type Grid struct {
	Width    int
	Height   int
	Antennas map[string][]int
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

func (g *Grid) GetPos(x, y int) int {
	return (g.Width * y) + x
}
